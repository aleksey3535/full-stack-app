package migrations

import "github.com/jmoiron/sqlx"

func InitTimeSlots(db *sqlx.DB) {
		queryList := []struct {
		TimeSlotID int
		ServiceID int
		QueueNumber int 
		StartTime string
	}{
		{
			TimeSlotID: 1,
			ServiceID: 1,
			QueueNumber: 1,
			StartTime: "14:00:00",
		},
		{
			TimeSlotID: 2,
			ServiceID: 1,
			QueueNumber: 2,
			StartTime: "14:15:00",
		},
		{
			TimeSlotID: 3,
			ServiceID: 1,
			QueueNumber: 3,
			StartTime: "14:30:00",
		},
		{
			TimeSlotID: 4,
			ServiceID: 1,
			QueueNumber: 4,
			StartTime: "14:45:00",
		},
		{
			TimeSlotID: 5,
			ServiceID: 1,
			QueueNumber: 5,
			StartTime: "15:00:00",
		},
		{
			TimeSlotID: 6,
			ServiceID: 1,
			QueueNumber: 6,
			StartTime: "15:15:00",
		},
		{
			TimeSlotID: 7,
			ServiceID: 1,
			QueueNumber: 7,
			StartTime: "15:30:00",
		},
		{
			TimeSlotID: 8,
			ServiceID: 1,
			QueueNumber: 8,
			StartTime: "15:45:00",
		},
		{
			TimeSlotID: 9,
			ServiceID: 1,
			QueueNumber: 9,
			StartTime: "16:00:00",
		},
		{
			TimeSlotID: 10,
			ServiceID: 1,
			QueueNumber: 10,
			StartTime: "16:15:00",
		},
		{
			TimeSlotID: 11,
			ServiceID: 2,
			QueueNumber: 1,
			StartTime: "14:00:00",
		},
		{
			TimeSlotID: 12,
			ServiceID: 2,
			QueueNumber: 2,
			StartTime: "14:15:00",
		},
		{
			TimeSlotID: 13,
			ServiceID: 2,
			QueueNumber: 3,
			StartTime: "14:30:00",
		},
		{
			TimeSlotID: 14,
			ServiceID: 2,
			QueueNumber: 4,
			StartTime: "14:45:00",
		},
		{
			TimeSlotID: 15,
			ServiceID: 2,
			QueueNumber: 5,
			StartTime: "15:00:00",
		},
		{
			TimeSlotID: 16,
			ServiceID: 2,
			QueueNumber: 6,
			StartTime: "15:15:00",
		},
		{
			TimeSlotID: 17,
			ServiceID: 2,
			QueueNumber: 7,
			StartTime: "15:30:00",
		},
		{
			TimeSlotID: 18,
			ServiceID: 2,
			QueueNumber: 8,
			StartTime: "15:45:00",
		},
		{
			TimeSlotID: 19,
			ServiceID: 2,
			QueueNumber: 9,
			StartTime: "16:00:00",
		},
		{
			TimeSlotID: 20,
			ServiceID: 2,
			QueueNumber: 10,
			StartTime: "16:15:00",
		},
		{
			TimeSlotID: 21,
			ServiceID: 3,
			QueueNumber: 1,
			StartTime: "14:00:00",
		},
		{
			TimeSlotID: 22,
			ServiceID: 3,
			QueueNumber: 2,
			StartTime: "14:10:00",
		},
		{
			TimeSlotID: 23,
			ServiceID: 3,
			QueueNumber: 3,
			StartTime: "14:20:00",
		},
		{
			TimeSlotID: 24,
			ServiceID: 3,
			QueueNumber: 4,
			StartTime: "14:30:00",
		},
		{
			TimeSlotID: 25,
			ServiceID: 3,
			QueueNumber: 5,
			StartTime: "14:40:00",
		},
		{
			TimeSlotID: 26,
			ServiceID: 3,
			QueueNumber: 6,
			StartTime: "14:50:00",
		},
		{
			TimeSlotID: 27,
			ServiceID: 3,
			QueueNumber: 7,
			StartTime: "15:00:00",
		},
		{
			TimeSlotID: 28,
			ServiceID: 3,
			QueueNumber: 8,
			StartTime: "15:10:00",
		},
		{
			TimeSlotID: 29,
			ServiceID: 3,
			QueueNumber: 9,
			StartTime: "15:20:00",
		},
		{
			TimeSlotID: 30,
			ServiceID: 3,
			QueueNumber: 10,
			StartTime: "15:30:00",
		},
		{
			TimeSlotID: 31,
			ServiceID: 3,
			QueueNumber: 11,
			StartTime: "15:40:00",
		},
		{
			TimeSlotID: 32,
			ServiceID: 3,
			QueueNumber: 12,
			StartTime: "15:50:00",
		},
		{
			TimeSlotID: 33,
			ServiceID: 3,
			QueueNumber: 13,
			StartTime: "16:00:00",
		},
		{
			TimeSlotID: 34,
			ServiceID: 3,
			QueueNumber: 14,
			StartTime: "16:10:00",
		},
		{
			TimeSlotID: 35,
			ServiceID: 3,
			QueueNumber: 15,
			StartTime: "16:20:00",
		},
		{
			TimeSlotID: 36,
			ServiceID: 3,
			QueueNumber: 16,
			StartTime: "16:30:00",
		},
		{
			TimeSlotID: 37,
			ServiceID: 3,
			QueueNumber: 17,
			StartTime: "16:40:00",
		},
		{
			TimeSlotID: 38,
			ServiceID: 3,
			QueueNumber: 18,
			StartTime: "16:50:00",
		},
	}
	for _, query := range queryList {
		db.Exec("INSERT INTO TimeSlots(TimeSlotID, ServiceID, QueuePosition,StartTime, IsBusy) VALUES ($1, $2, $3, $4, $5)",
		query.TimeSlotID,query.ServiceID, query.QueueNumber,query.StartTime, false)
	}

}