package repository

import (
	"database/sql"
	"errors"
	"queueAppV2/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type AdminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{db : db}
}

func(ar *AdminRepository) GetAppointments(serviceID int) ([]models.AppointmentForAdmin, error) {
	result := make([]models.AppointmentForAdmin, 0)
	query := `SELECT a.AppointmentID, a.QueueNumber, u.FullName, u.PassportNumber, t.StartTime FROM Appointments AS a 
		INNER JOIN Users AS u ON a.UserID = u.UserID
		INNER JOIN TimeSlots AS t ON a.TimeSlotID = t.TimeSlotID
		WHERE t.ServiceID = $1
		ORDER BY t.StartTime ASC`
	rows, err := ar.db.Query(query, serviceID)
	if err != nil {
		return nil, errors.New("with db.Query " + err.Error())
	}
	for rows.Next() {
		var queue models.AppointmentForAdmin
		var startTime time.Time
		if err := rows.Scan(&queue.AppointmentID, &queue.QueueNumber, &queue.FullName, 
			&queue.PassportNumber, &startTime); err != nil {
				return nil, errors.New("with rows.Scan " + err.Error())
		}
		queue.StartTime = startTime.Format("15:04")
		result = append(result, queue)
	}
	if len(result) == 0{
		return nil, ErrEmptyAppointments
	}
	return result, nil
}

func(ar *AdminRepository) DeleteAppointment(queueID int) error {
	query := `DELETE FROM Appointments WHERE AppointmentID = $1`
	if _, err := ar.db.Exec(query, queueID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrAppointmentNotFound
		}
		return errors.New("with db.Exec " + err.Error())
	}
	return nil
}