package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"os"
)

type ConnectionFactory struct {
	Connect *pgx.Conn
}

func (connectionFactory *ConnectionFactory) getConnect() (*pgx.Conn, error) {
	if connectionFactory.Connect != nil {
		return connectionFactory.Connect, nil
	}
	db_url := "host=" + viper.Get("database.host").(string) + " "
	db_url += "port=" + viper.Get("database.port").(string) + " "
	db_url += "user=" + viper.Get("database.user").(string) + " "
	db_url += "password=" + viper.Get("database.password").(string) + " "
	db_url += "dbname=" + viper.Get("database.dbname").(string) + " "
	db_url += "sslmode=disable"
	conn, err := pgx.Connect(context.Background(), os.Getenv(db_url))
	if err != nil {
		connectionFactory.Connect = conn
		defer conn.Close(context.Background())
		return conn, nil
	}
	return nil, err
}
