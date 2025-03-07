package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"queueAppV2/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)


type UserRepository struct {
	db *sqlx.DB
}


func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func(ur *UserRepository) GetFreeSlots(serviceID int) ([]models.Slot, error) {
	var startTime time.Time
	result := make([]models.Slot, 0)
	query := `SELECT TimeSlotID, StartTime from TimeSlots WHERE ServiceID = $1 AND IsBusy = $2`
	rows, err := ur.db.Query(query, serviceID, false)
	if err != nil {
		return nil, errors.New("with db.Query " + err.Error())
	}
	for rows.Next() {
		var slot models.Slot
		if err := rows.Scan(&slot.TimeSlotID, &startTime); err != nil {
			return nil, errors.New("with rows.Scan " + err.Error())
		}
		slot.StartTime = startTime.Format("15:04")
		result = append(result, slot)
	}
	return result, nil
}

func(ur *UserRepository) RegisterNewUser(slotID int, user models.User) (models.Appointment, error) {
	serviceID,queueNumber, cabinet, startTime, err := getData(ur.db, slotID)
	if err != nil {
		return models.Appointment{}, err
	}
	// начало транзакции
	tx, err := ur.db.Begin()
	if err != nil {
		tx.Rollback()
		return models.Appointment{}, errors.New("with db.Begin " + err.Error())
	}
	userID, err := createUser(tx, user) 
	if err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}
	if err = checkRegistration(tx, userID, serviceID); err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}
	if err = createAppointment(tx, userID, slotID, queueNumber); err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}
	if err = updateTimeSlot(tx, slotID); err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return models.Appointment{}, errors.New("with tx.Commit " + err.Error())
	}
	return models.Appointment{
		QueueNumber: queueNumber,
		FullName: user.FullName,
		PassportNumber: user.PassportNumber,
		StartTime: startTime,
		Cabinet: cabinet,
	}, nil
}

func checkRegistration(tx *sql.Tx, userID, serviceID int) error {
	queryLastServiceID := `SELECT t.ServiceID 
		FROM Appointments AS a INNER JOIN TimeSlots AS t ON a.TimeSlotID = t.TimeSlotID
		WHERE a.UserID = $1`
	rows, err  := tx.Query(queryLastServiceID, userID)
	if err != nil {
		return errors.New("with checkRegistration tx.Query " + err.Error())
	}
	for rows.Next() {
		var lastServiceID int
		if err = rows.Scan(&lastServiceID); err != nil {
			return errors.New("with checkRegistration rows.Scan " + err.Error())
		}
		if serviceID == lastServiceID {
			return ErrAlreadyRegistered
		}
	}
	return nil
}


func updateTimeSlot(tx *sql.Tx, slotID int) error {
	query := `UPDATE TimeSlots SET IsBusy = $1 WHERE TimeSlotID = $2`
	if _, err := tx.Exec(query, true, slotID); err != nil {
		return errors.New("with updateTimeSlot" + err.Error())
	}
	return nil
}

func createAppointment(tx *sql.Tx, userID, slotID int, queueNumber string) error {
	query := `INSERT INTO Appointments(UserID, TimeSlotID, QueueNumber) VALUES ($1, $2, $3)`
	if _,err := tx.Exec(query, userID, slotID, queueNumber); err != nil {
		return ErrBusySlot
	}
	return nil
}

func createUser(tx *sql.Tx, user models.User) (int, error) {
	var userID int
	getQuery := `SELECT UserID from Users WHERE PassportNumber = $1`
	row := tx.QueryRow(getQuery, user.PassportNumber)
	if err := row.Scan(&userID); err == nil {
		return userID, nil
	} else {
		if !errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("with createUser getQuery " + err.Error())
		}
	}

	createQuery := `INSERT INTO Users(FullName, PassportNumber) VALUES ($1, $2) RETURNING UserID`
	row = tx.QueryRow(createQuery, user.FullName, user.PassportNumber)
	if err := row.Scan(&userID); err != nil {
		return 0, errors.New("with create user " + err.Error())
	}
	return userID, nil
}

// returning serviceID, queueNumber, cabinet, startTime and error
func getData(db *sqlx.DB, slotID int) (int,string, int, string, error) {
	var serviceID, queuePosition, cabinet int
	var serviceAlias string
	var startTime time.Time
	queryServiceID := `SELECT ServiceID, QueuePosition, StartTime from TimeSlots WHERE TimeSlotID = $1`
	row := db.QueryRow(queryServiceID, slotID)
	if err := row.Scan(&serviceID, &queuePosition, &startTime); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", 0, "", ErrServiceNotFound
		}
		return 0, "", 0, "", errors.New("with queryServiceID row.Scan" + err.Error())
	}
	query := `SELECT Alias, Cabinet from Services WHERE ServiceID = $1`
	row = db.QueryRow(query,serviceID)
	if err := row.Scan(&serviceAlias, &cabinet); err != nil {
		return 0, "", 0, "", errors.New("with query row.Scan" + err.Error())
	}
	queueNumber := serviceAlias + fmt.Sprint(queuePosition)
	startTimeFormatted := startTime.Format("15:04")
	return serviceID, queueNumber, cabinet, startTimeFormatted, nil
}
