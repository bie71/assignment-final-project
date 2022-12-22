package test

import (
	"assigment-final-project/helper"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"context"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	rows := helper.CountTotalRows(context.Background(), mysql_connection.InitMysqlDB(), "users")
	fmt.Println(rows)
}
