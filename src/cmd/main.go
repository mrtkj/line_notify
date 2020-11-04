package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/line_notify/src/client/line"
	"github.com/line_notify/src/config"
	"github.com/line_notify/src/controller"
	"github.com/line_notify/src/db"
	"github.com/line_notify/src/repository"
	"github.com/line_notify/src/service"
	"go.uber.org/dig"
)

type server struct {
	config   *config.Config
	db       *db.DBContext
	schedule *controller.ScheduleController
}

func NewServer(
	config *config.Config,
	db *db.DBContext,
	schedule *controller.ScheduleController,
) *server {
	return &server{
		config:   config,
		db:       db,
		schedule: schedule,
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func main() {
	port := os.Getenv("PORT")

	container := Register()
	err := container.Invoke(
		func(server *server) {
			gin.SetMode("release")
			r := gin.New()
			// repository.SetTransactional(server.db)

			g := r.Group("/api")
			g.Use(CheckHeader())
			{
				server.schedule.Endpoint(g)
			}

			srv := &http.Server{
				Addr:    ":" + port,
				Handler: r,
			}
			log.Printf("Starting server at Port %s", port)

			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					panic(err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Print("Shutdown server ...")
		})

	if err != nil {
		panic(err)
	}
}

func Register() *dig.Container {
	container := dig.New()
	provide(container, config.NewConfig)

	provide(container, func(cfg *config.Config) db.Config { return cfg.GetDBConfig() })
	provide(container, func(cfg *config.Config) line.Config { return cfg.GetLINEConfig() })
	provide(container, db.NewDB)
	provide(container, line.NewClient)
	provide(container, repository.NewSchedulesRepository)
	provide(container, repository.NewUsersRepository)
	provide(container, service.NewScheduleService)
	provide(container, controller.NewScheduleController)
	provide(container, NewServer)
	return container
}

func provide(container *dig.Container, constructor interface{}) {
	if err := container.Provide(constructor); err != nil {
		panic(err)
	}
}

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Request.Header
		a := h.Get("Authorization")
		at := "Bearer " + os.Getenv("ACCESS_TOKEN")

		if a != at {
			c.JSON(http.StatusUnauthorized, "access_token is invalid.")
			c.Abort()
		}
		c.Next()
	}
}
