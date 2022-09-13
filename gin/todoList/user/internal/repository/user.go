package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/service"

)

type User struct {
	UserID         uint      `gorm:"primarykey"` // 主键
	UserName       string    `gorm:"unique"` // 用户名(唯一索引)
	NickName       string // 昵称
	PasswordDigest string // 密码(密文)
}


const (
	PassWordCost = 12 // 密码加密难度
)

// CheckUserExist 查询用户是否存在
func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error;
		err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowUserInfo 获取用户详情
func (user *User) ShowUserInfo (req *service.UserRequest)(err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

// BuildUser 视图返回 序列化user
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserID),
		NickName: item.NickName,
		UserName: item.UserName,
	}
	return &userModel
}
// UserCreate 添加用户
func (u *User) UserCreate(req *service.UserRequest) (user User, err error)  {
	if exist := u.CheckUserExist(req); exist {
		return User{},errors.New("用户名已存在")
	}

	user = User{
		UserName:       req.UserName,
		NickName:       req.NickName,
	}
	if err = user.SetPassword(req.Password);err != nil {
		return user,err
	}
	return user,DB.Model(&User{}).Create(&user).Error
}


// SetPassword 加密密码
func (u *User) SetPassword (password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 检验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}