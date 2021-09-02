package services

import (
	//"fmt"
	"ormalta/problem-3.2/config"
	"ormalta/problem-3.2/models"
	"ormalta/problem-3.2/middlewares"
)


func GetUsers() ([]models.User, error) {
	var users []models.User

	res := config.Db.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func GetUserById(targetId int) (models.User, int, error) {
	var user models.User

	res := config.Db.Find(&user, targetId)

	if res.Error != nil {
		return models.User{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.User{}, 0, nil
	}

	return user, 1, nil
}

func AddUser(newUser *models.User) (models.User, error) {
	res := config.Db.Create(newUser)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return *newUser, nil
}

func EditUser(newData models.User, targetId int) (models.User, int, error) {
	targetUser := models.User{}
	res := config.Db.Find(&targetUser, targetId)

	if res.Error != nil {
		return models.User{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.User{}, 0, nil
	}

	res = config.Db.Model(&targetUser).Updates(newData)

	if res.Error != nil {
		return models.User{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.User{}, 0, nil
	}

	return targetUser, 1, nil
}

func DeleteUser(targetId int) (models.User, int, error) {	
	targetUser := models.User{}
	res := config.Db.Find(&targetUser, targetId)

	if res.Error != nil {
		return models.User{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.User{}, 0, nil
	}
	
	deleted := targetUser

	res = config.Db.Unscoped().Delete(&targetUser)

	if res.Error != nil {
		return models.User{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.User{}, 0, nil
	}

	return deleted, 1, nil

}

func LoginUser(user *models.User) (string ,error) {
	res := config.Db.Where("email = ? AND password = ?", user.Email, user.Password).First(user)

	if res.Error != nil {
		return "", res.Error
	}

	token, err := middlewares.CreateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil

}