package criteria_test

import (
	entity "assigment-final-project/domain/entity/criteria"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository "assigment-final-project/internal/repository/mysql"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ctx          = context.Background()
	db           = mysql_connection.InitMysqlDB()
	repoCriteria = repository.NewCriteriaRepoImpl(db)
)

func TestInsertCriteria(t *testing.T) {
	err := repoCriteria.InsertCriteria(ctx, entity.NewCriteria(&entity.DTOCriteria{
		Name: "Game",
	}))

	fmt.Println(err)
	assert.NoError(t, err)
}
func TestGetCriteria(t *testing.T) {
	result, err := repoCriteria.GetCriteria(ctx)
	fmt.Println(err)
	for _, criteria := range result {
		fmt.Println(criteria)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
