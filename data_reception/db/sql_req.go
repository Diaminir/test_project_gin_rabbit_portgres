package db

import (
	"context"
	"gin_db/dto"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	db *pgx.Conn
}

func NewDbLogin() (*Postgres, error) {
	connFig := "postgres://qwert:12345@postgres:5432/cars"
	conn, err := pgx.Connect(context.Background(), connFig)
	if err != nil {
		return &Postgres{}, nil
	}
	return &Postgres{
		db: conn,
	}, nil
}

func (pg *Postgres) DbClose() {
	pg.db.Close(context.Background())
}

func (pg *Postgres) DbSearch(desCar dto.DesiredCarDTO) (dto.CarInfoBD, error) {
	data := dto.CarInfoBD{}
	Sql := "SELECT * FROM view_marks_options WHERE name=$1 AND model=$2 AND engine=$3 AND generation=$4"
	err := pg.db.QueryRow(context.Background(), Sql, desCar.Title, desCar.Model, desCar.Engine, desCar.Genetation).Scan(&data.Name, &data.Model, &data.Engine, &data.Genetation, &data.Price)
	if err != nil {
		return dto.CarInfoBD{}, err
	}
	return data, nil
}
