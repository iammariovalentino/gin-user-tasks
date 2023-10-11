package service

import (
	"fmt"
	"gin-user-tasks/src/app/oauth_client/handler"
	taskHandler "gin-user-tasks/src/app/tasks/handler"
	taskQuery "gin-user-tasks/src/app/tasks/query"
	taskUsecase "gin-user-tasks/src/app/tasks/usecase"
	userHandler "gin-user-tasks/src/app/users/handler"
	userQuery "gin-user-tasks/src/app/users/query"
	userUseCase "gin-user-tasks/src/app/users/usecase"
	"gin-user-tasks/src/pkg/config"
	"gin-user-tasks/src/pkg/middleware"
	"gin-user-tasks/src/platform/db"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/generates"
	"github.com/go-oauth2/oauth2/manage"
	"github.com/go-oauth2/oauth2/models"
	"github.com/go-oauth2/oauth2/server"
	"github.com/go-oauth2/oauth2/store"
	"golang.org/x/sync/errgroup"
)

var defaultAppPort = 8080
var defaultOauthPort = 9096

type Service struct {
	userService userUseCase.UserUsecase
	taskService taskUsecase.TaskUsecase
}

func New() *Service {
	conn := db.NewMySQL(config.Env).GetConn()
	dbInst := db.ConnectGorm(conn)

	userQuery := userQuery.NewUserQuery(dbInst)
	taskQuery := taskQuery.NewTaskQuery(dbInst)

	userUsecase := userUseCase.NewUserUsecase(userQuery)
	taskUsecase := taskUsecase.NewTaskUsecase(taskQuery)

	return &Service{
		userService: userUsecase,
		taskService: taskUsecase,
	}
}

func (s *Service) buildAppHandler() http.Handler {

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := store.NewClientStore()
	clientStore.Set(config.Env.OauthConfig.ClientID, &models.Client{
		ID:     config.Env.OauthConfig.ClientID,
		Secret: config.Env.OauthConfig.ClientSecret,
		Domain: "http://localhost:8080",
	})

	manager.MapClientStorage(clientStore)
	srv := server.NewServer(server.NewConfig(), manager)

	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.POST("/token", func(c *gin.Context) {
		srv.HandleTokenRequest(c.Writer, c.Request)
	})
	g.POST("/authorize", func(c *gin.Context) {
		srv.HandleTokenRequest(c.Writer, c.Request)
	})

	userHandler := userHandler.NewUserHandler(s.userService)
	userGroup := g.Group("/users")
	userHandler.Mount(userGroup)

	taskHandler := taskHandler.NewTaskHandler(s.taskService)
	taskGroup := g.Group("/tasks")
	taskGroup.Use(middleware.HandleTokenVerify(middleware.Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			ctx.AbortWithError(500, err)
		},
		TokenKey: "http://localhost:9096/oauth/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
		OauthServer: srv,
	}))

	taskHandler.Mount(taskGroup)

	return g
}

func (s *Service) buildOauthHandler() http.Handler {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	oauthHandler := handler.NewOauthClientHandler(config.Env.OauthConfig)
	oauthGroup := g.Group("/oauth")
	oauthHandler.Mount(oauthGroup)

	return g
}

func (s *Service) Run() {
	appPort := config.Env.AppPort
	if appPort == 0 {
		appPort = defaultAppPort
	}

	oauthPort := config.Env.OauthPort
	if oauthPort == 0 {
		oauthPort = defaultAppPort
	}

	appHandler := s.buildAppHandler()
	oauthHandler := s.buildOauthHandler()

	appServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", appPort),
		Handler:      appHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	oauthServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", oauthPort),
		Handler:      oauthHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var g errgroup.Group

	g.Go(func() error {
		return appServer.ListenAndServe()
	})

	g.Go(func() error {
		return oauthServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
