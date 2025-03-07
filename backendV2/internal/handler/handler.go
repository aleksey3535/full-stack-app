package handler

import (
	"log/slog"
	"net/http"
	"queueAppV2/internal/config"
	"queueAppV2/internal/middleware"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

type UserHandlerI interface {
	FreeSlotsHandler(w http.ResponseWriter, r *http.Request)
	RegisterHandler(w http.ResponseWriter, r *http.Request)
}

type AdminHandlerI interface {
	StatusHandler(w http.ResponseWriter, r *http.Request)
	DoneHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)

}

type RepositoryI interface {
	UserRepositoryI
	AdminRepositoryI
}

type Handler struct {
	mw *middleware.MiddleWare
	UserHandlerI
	AdminHandlerI
}

func New(mw *middleware.MiddleWare, log *slog.Logger, repo RepositoryI, cfg *config.Config) *Handler {
	return &Handler{
		mw:           mw,
		UserHandlerI: NewUserHandler(log, repo),
		AdminHandlerI: NewAdminHandler(log, repo, cfg),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	mux := mux.NewRouter()
	mux.Use(h.mw.UseHeaders)
	log.Info("[GET] /api/user/{id}/status - возвращает свободные записи на услугу с id = {id}",
	"output", "{data:[{id, startTime}]}")
	mux.HandleFunc("/api/user/{id}/status", h.UserHandlerI.FreeSlotsHandler).Methods(http.MethodGet)
	log.Info("[POST] /api/user/register/{id} - регистрирует пользователя на временной слот с id = {id}.", 
	"input", "{fullName, passportNumber}", 
	"output", "{queueNumber, fullName, passportNumber, startTime, Cabinet}" )
	mux.HandleFunc("/api/user/register/{id}", h.UserHandlerI.RegisterHandler).Methods(http.MethodPost)
	mux.HandleFunc("/api/user/register/{id}", func(w http.ResponseWriter, r *http.Request){w.WriteHeader(http.StatusOK)}).Methods(http.MethodOptions)
	log.Info("[GET] /api/admin/{id}/status - возвращает текущий статус очереди для оператора услуги {id}.",
	"output", "{data:[{id, queueNumber, fullName, passportNumber, startTime}]}")
	mux.HandleFunc("/api/admin/{id}/status", h.AdminHandlerI.StatusHandler).Methods(http.MethodGet)
	log.Info("[GET] /api/admin/{id}/done/{queueID} - завершает запись с id={queueID} для оператора услуги {id}.")
	mux.HandleFunc("/api/admin/{id}/done/{queueID}", h.AdminHandlerI.DoneHandler).Methods(http.MethodGet)
	log.Info("[POST] /api/admin/login - авторизация для оператора.", "input", "{login, password}", "output", "200 - success, 401 - permission denied")
	mux.HandleFunc("/api/admin/login", h.AdminHandlerI.LoginHandler).Methods(http.MethodPost)
	mux.HandleFunc("/api/admin/login", func(w http.ResponseWriter, r *http.Request){w.WriteHeader(http.StatusOK)}).Methods(http.MethodOptions)
	return mux
	
}
