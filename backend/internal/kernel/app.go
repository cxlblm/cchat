package kernel

import (
	"cchart/internal/http"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log/slog"
	httpS "net/http"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	Config *Config
	Engine *gin.Engine
	Logger *slog.Logger
	DB     *gorm.DB
	Cache  *redis.Client
	Router *http.Router
	Srv    *httpS.Server
}

func (s *Application) Run() {

	s.Logger.Info("server starting")
	s.Router.Register()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := s.Srv.ListenAndServe(); err != nil && !errors.Is(err, httpS.ErrServerClosed) {
			s.Logger.Error("http srv run error" + err.Error())
		}
	}()
	if s.Config.Nsq.Use {
		go func() {}()
	}

	<-ctx.Done()

	s.Stop(stop)
}

func (s *Application) Stop(stop context.CancelFunc) {
	stop()
	s.Logger.Info("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Srv.Shutdown(ctx); err != nil {
		s.Logger.Error("Server forced to shutdown: " + err.Error())
	}

	s.Logger.Info("srv exiting")
}
