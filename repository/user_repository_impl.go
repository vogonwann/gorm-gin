package repository

import (
	"errors"

	"github.com/vogonwann/gorm-gin/helper"
	"github.com/vogonwann/gorm-gin/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (t *UserRepositoryImpl) Create(user model.User) {
	result := t.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (t *UserRepositoryImpl) Delete(user model.User) {
	var toDelete model.User
	result := t.Db.Where("id = ?", user.Id).Delete(&toDelete)
	helper.ErrorPanic(result.Error)
}

func (t *UserRepositoryImpl) FindByUsername(userName string) (model.User, error) {
	var user model.User
	result := t.Db.Where("userName = ?", userName).First(&user)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("User not found!")
	}
}
