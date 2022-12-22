package helper

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

type Rows struct {
	TotalRows int `dbq:"total_rows"'`
}

func CountTotalRows(ctx context.Context, db *sql.DB, tableName string) *Rows {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT COUNT(*) AS total_rows FROM %s`, tableName)
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: Rows{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	q := dbq.MustQ(ctx, db, stmt, opts)
	return q.(*Rows)
}
