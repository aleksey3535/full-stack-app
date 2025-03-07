package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"queueAppV2/internal/config"
	"queueAppV2/internal/models"
	"queueAppV2/internal/repository"

	"github.com/gorilla/mux"
)

type AdminRepositoryI interface {
	GetAppointments(serviceID int) ([]models.AppointmentForAdmin, error)
	DeleteAppointment(queueID int) error
}

type AdminHandler struct {
	log  *slog.Logger
	repo AdminRepositoryI
	cfg  *config.Config
}

func NewAdminHandler(log *slog.Logger, repo AdminRepositoryI, cfg *config.Config) *AdminHandler {
	return &AdminHandler{log: log, repo: repo, cfg : cfg}
}

type DataForStatus struct {
	Data []models.AppointmentForAdmin `json:"data"`
}

func (ah *AdminHandler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	const op = "handler.adminHandler.StatusHandler"
	log := ah.log.With("op", op)
	serviceIDStr := mux.Vars(r)["id"]
	serviceID, errorMessage, ok := validateServiceID(serviceIDStr)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
	}
	data, err := ah.repo.GetAppointments(serviceID)
	if err != nil {
		if errors.Is(err, repository.ErrEmptyAppointments) {
			message := ErrorWrapper(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write(message)
			return
		}
		InternalErrorHandler(w)
		log.Error("occurred with repo.getAppointments " + err.Error())
		return
	}
	dataMessage := DataForStatus{Data: data}
	message, err := json.Marshal(dataMessage)
	if err != nil {
		InternalErrorHandler(w)
		log.Error("occurred with marshalling json " + err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

type SuccessMessage struct {
	Message string `json:"message"`
}

func (ah *AdminHandler) DoneHandler(w http.ResponseWriter, r *http.Request) {
	const op = "handler.adminHandler.DoneHandler"
	log := ah.log.With("op", op)
	idStr := mux.Vars(r)["queueID"]
	id, errorMessage, ok := validateQueueID(idStr)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage)
	}
	if err := ah.repo.DeleteAppointment(id); err != nil {
		if errors.Is(err, repository.ErrEmptyAppointments) {
			message := ErrorWrapper(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write(message)
			return
		}
		InternalErrorHandler(w)
		log.Error("occurred with repo.DeleteAppointment " + err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	message, _ := json.Marshal(SuccessMessage{Message: "success"})
	w.Write(message)
}

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (ah *AdminHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	const op = "handler.adminHandler.LoginHandler"
	log := ah.log.With("op", op)
	var data LoginData
	body, err := io.ReadAll(r.Body)
	if err != nil {
		InternalErrorHandler(w)
		log.Error("occurred with reading body " + err.Error())
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		InvalidDataHandler(w)
		log.Error("occurred with unmarshalling json " + err.Error())
		return
	}
	if validateCredentials(ah.cfg, data) {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
