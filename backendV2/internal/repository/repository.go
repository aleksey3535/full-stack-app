package repository

import (
	"errors"
	"queueAppV2/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryI interface {
	GetFreeSlots(serviceID int) ([]models.Slot, error)
	RegisterNewUser(slotID int, user models.User) (models.Appointment, error)
}

type AdminRepositoryI interface {
	GetAppointments(serviceID int) ([]models.AppointmentForAdmin, error)
	DeleteAppointment(queueID int) error
}

var (
	ErrServiceNotFound = errors.New("service not found")
	ErrEmptyFreeSlots = errors.New("free slots is empty")
	ErrBusySlot = errors.New("this slot is busy")
	ErrAlreadyRegistered = errors.New("this user already in queue")
	ErrEmptyAppointments = errors.New("no appointments found")
	ErrAppointmentNotFound = errors.New("appointment not found")
)

type Repository struct {
	UserRepositoryI
	AdminRepositoryI
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepositoryI: NewUserRepository(db),
		AdminRepositoryI:NewAdminRepository(db),
	}
}