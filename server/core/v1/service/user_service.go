package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sevenzx/eztodo/global"
	"github.com/sevenzx/eztodo/model"
	"github.com/sevenzx/eztodo/util"
	"gorm.io/gorm"
)

type userService struct{}

// Register 注册
func (s *userService) Register(user *model.User) error {
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

// Login 用户登录
func (s *userService) Login(username string, password string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ?", username).First(&user).Error
	if err == nil {
		if ok := util.BcryptCheck(password, user.Password); !ok {
			return nil, errors.New("incorrect password")
		} else {
			return &user, nil
		}
	} else {
		return nil, errors.New("no such username")
	}
}

// GetUserByUuid 通过uuid获取用户信息
func (s *userService) GetUserByUuid(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := global.DB.Where("uuid = ?", id).First(&user).Error
	return &user, err
}
