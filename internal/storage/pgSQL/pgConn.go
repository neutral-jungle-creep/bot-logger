package pgSQL

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

func NewPgConnect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), viper.GetString("linkToDB"))
	if err != nil {
		return conn, err
	}

	return conn, nil
}
