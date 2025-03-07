package migrations

import "github.com/jmoiron/sqlx"

func InitServices(db *sqlx.DB) {
		queryList := []struct {
		ServiceID   int
		ServiceName string
		Alias       string
		Cabinet     int
	}{
		{
			ServiceID:   1,
			ServiceName: "Продление визы",
			Alias:       "В",
			Cabinet:     512,
		},
		{
			ServiceID:   2,
			ServiceName: "Продление миграционного учета",
			Alias:       "М",
			Cabinet:     510,
		},
		{
			ServiceID:   3,
			ServiceName: "Подача документов на приглашение",
			Alias:       "П",
			Cabinet:     519,
		},
	}
	for _, query := range queryList {
		db.Exec("INSERT INTO Services(ServiceID, ServiceName, Alias, Cabinet) VALUES($1, $2, $3, $4)",
			query.ServiceID, query.ServiceName, query.Alias, query.Cabinet)
	}
}