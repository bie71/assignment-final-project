package categories_test

import (
	entity "assigment-final-project/domain/entity/categories"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository "assigment-final-project/internal/repository/mysql"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	ctx            = context.Background()
	db             = mysql_connection.InitMysqlDB()
	repoCategories = repository.NewCategoryRepoImpl(db)
)

func TestInsertCategory(t *testing.T) {
	err := repoCategories.InsertCategory(ctx, entity.NewCategories(&entity.DTOCategories{
		CategoryId: "123",
		Name:       "Console",
	}))

	assert.NoError(t, err)
}

func TestInsertListCategory(t *testing.T) {
	listCategory := make([]*entity.Categories, 0)

	for i := 0; i < 5; i++ {
		catgory := entity.NewCategories(&entity.DTOCategories{
			CategoryId: "123" + strconv.Itoa(i),
			Name:       "Console",
		})
		listCategory = append(listCategory, catgory)
	}

	err := repoCategories.InsertListCategory(ctx, listCategory)
	fmt.Println(err)
	assert.NoError(t, err)
}

func TestFindCategoryById(t *testing.T) {
	category, err := repoCategories.FindCategory(ctx, "123")
	assert.NoError(t, err)
	assert.NotEmpty(t, category)
	assert.Equal(t, "123", category.CategoryId())
	assert.Equal(t, "Console", category.CategoryName())
}

func TestGetCategories(t *testing.T) {
	categories, err := repoCategories.GetCategories(ctx, 0, 5)
	assert.NoError(t, err)
	assert.NotEmpty(t, categories)
}

func TestDeleteCategory(t *testing.T) {
	err := repoCategories.DeleteCategory(ctx, "123")
	assert.NoError(t, err)
}
