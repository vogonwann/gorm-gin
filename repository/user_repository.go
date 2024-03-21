package repository

import "github.com/vogonwann/gorm-gin/model"

type UserRepository interface {
	Create(user model.User)
	Delete(user model.User)
	FindByUsername(userName string)
}
