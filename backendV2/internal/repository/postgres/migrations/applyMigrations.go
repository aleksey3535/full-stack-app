package migrations

import "github.com/jmoiron/sqlx"

func ApplyMigrations(db *sqlx.DB) {
	usersQuery := `CREATE TABLE IF NOT EXISTS Users (
		UserID serial not null unique,
		FullName varchar(255) not null,
		PassportNumber varchar(50) not null unique
	);`
	servicesQuery := `CREATE TABLE IF NOT EXISTS Services (
		ServiceID integer not null unique,
		ServiceName varchar(255) not null,
		Alias varchar(10) not null unique,
		Cabinet integer
	);`
	timeSlotsQuery := `CREATE TABLE IF NOT EXISTS TimeSlots (
		TimeSlotID integer not null unique,
		ServiceID integer not null,
		QueuePosition integer not null,
		StartTime time not null,
		IsBusy boolean,
		FOREIGN KEY (ServiceID) REFERENCES Services(ServiceID)
	);`
	appointmentsQuery := `CREATE TABLE IF NOT EXISTS Appointments (
		AppointmentID serial not null unique,
		UserID integer not null,
		TimeSlotID integer not null unique,
		QueueNumber varchar(20) not null unique,
		FOREIGN KEY (UserID) REFERENCES Users(UserID),
		FOREIGN KEY (TimeSlotID) REFERENCES TimeSlots(TimeSlotID)
	);`
	db.MustExec(usersQuery)
	db.MustExec(servicesQuery)
	db.MustExec(timeSlotsQuery)
	db.MustExec(appointmentsQuery)
}

func CancelMigrations(db *sqlx.DB) {
	appointmentsQuery := `DROP TABLE IF EXISTS Appointments CASCADE`
	timeSlotsQuery := `DROP TABLE IF EXISTS TimeSlots CASCADE`
	servicesQuery := `DROP TABLE IF EXISTS Services CASCADE`
	usersQuery := `DROP TABLE IF EXISTS Users CASCADE`
	db.MustExec(appointmentsQuery)
	db.MustExec(servicesQuery)
	db.MustExec(timeSlotsQuery)
	db.MustExec(usersQuery)
}
