package mapper

import (
	"assigment-final-project/domain/entity/users"
	"assigment-final-project/internal/repository/mysql/models"
)

func DomainUsersToModelsUsers(dataUsers *entity.Users) *models.UsersModels {
	return &models.UsersModels{
		UserId:    dataUsers.GetUserId(),
		Name:      dataUsers.GetName(),
		Username:  dataUsers.Username(),
		Password:  dataUsers.Password(),
		UserType:  dataUsers.UserType(),
		CreatedAt: dataUsers.CreatedAt(),
	}
}

func ModelsUsersToDomainUsers(modelUsers *models.UsersModels) *entity.Users {
	return entity.UserFromDb(&entity.DTOUsers{
		UserId:    modelUsers.UserId,
		Name:      modelUsers.Name,
		Username:  modelUsers.Username,
		Password:  modelUsers.Password,
		UserType:  modelUsers.UserType,
		CreatedAt: modelUsers.CreatedAt,
	})
}

func ToListDomainUser(dataUsers []*models.UsersModels) []*entity.Users {
	users := make([]*entity.Users, 0)
	for _, user := range dataUsers {
		usersDomain := ModelsUsersToDomainUsers(user)
		users = append(users, usersDomain)
	}
	return users
}
