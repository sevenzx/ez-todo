package internal

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sevenzx/eztodo/global"
	"github.com/sevenzx/eztodo/model"
	"github.com/sevenzx/eztodo/util"
	"gorm.io/gorm"
)

type UserService struct{}

func (s UserService) Register(user *model.User) error {
	if !errors.Is(global.DB.Where("username = ?", user.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("user already registered")
	}
	// 没有找到记录所以可以创建
	user.UUID, _ = uuid.NewRandom()
	user.Password = util.BcryptHash(user.Password)
	err := global.DB.Create(&user).Error
	if err != nil {
		return errors.Wrap(err, "create user")
	} else {
		return nil
	}
}

func (s UserService) GetById(id uint) (*model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}
