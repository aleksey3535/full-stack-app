package local

import (
	"log/slog"
	"net/http"
	"os"
	"queueAppV2/internal/config"
	"queueAppV2/internal/handler"
	"queueAppV2/internal/middleware"
	"queueAppV2/internal/repository"
	"queueAppV2/internal/repository/postgres"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron/v3"
)

type App struct {
	cfg *config.Config
	handler *handler.Handler
	db *sqlx.DB
}

func New() *App {
	cfg := config.MustLoad()
	log := setupLogger(cfg)
	db := postgres.MustCreate(cfg)
	repository := repository.New(db)
	mw := middleware.New(log)
	h := handler.New(mw, log, repository, cfg)

	return &App{
		handler: h,
		cfg: cfg,
		db: db,
	}
}

func(a *App) Run() error {
	port := ":" + a.cfg.Port
	log.Info("Server starting at port" + port)
	go restartDb(a.db)
	return http.ListenAndServe(port, a.handler.InitRoutes())
}

func restartDb(db *sqlx.DB) {
	c := cron.New()
	runTask := func() {
		queries := []string {
			"DELETE FROM Appointments",
			"ALTER TABLE Appointments AUTO_INCREMENT = 1",
			"DELETE FROM Users",
			"ALTER TABLE Users AUTO_INCREMENT = 1",
			"UPDATE TimeSlots SET IsBusy = false",
		}
		for _, query := range queries {
			db.MustExec(query)
		}
		log.Info("Database cleared")
	}
	if _, err := c.AddFunc("0 0 * * *", runTask); err != nil {
		panic(err)
	}
	c.Start()
	select {}
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger
	switch cfg.Env {
	case "local":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelError,
		}))
	}
	return log
}