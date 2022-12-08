package cmn

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Connection *pgxpool.Pool
func init() {
	var errors error

	config, err := pgxpool.ParseConfig("postgres://postgres:123456789@localhost:5433/postgres")
	if err != nil {
		fmt.Println("err1:",err)
	}
	Connection, errors = pgxpool.ConnectConfig(context.Background(), config)

	if errors != nil {
		fmt.Println("err2:",errors)
	} else {
		fmt.Println("连接成功")
	}
}