package service

import (
	"errors"
	"fmt"
	"gobackend/model"
	"gobackend/utils"
	"math/rand"
	"time"
)

type UserService struct {
}

//注册函数
func (s *UserService) Register(username string, password string) (user model.User, err error) {

	//检测手机号码是否存在,
	userInstance := model.User{}
	userInstance.Mobile = username
	userInstance = userInstance.GetOne()

	//如果存在则返回提示已经注册
	if userInstance.Id > 0 {
		return userInstance, errors.New("该手机号已经注册")
	}

	//否则拼接插入数据
	userInstance.Mobile = username
	userInstance.Passwd = password
	userInstance.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	userInstance.Passwd = utils.MakePasswd(userInstance.Passwd, userInstance.Salt)
	userInstance.Createat = time.Now()
	//token 可以是一个随机数
	userInstance.Token = utils.MD5Encode(fmt.Sprintf("%08d", rand.Int31()))

	//插入 InserOne
	userInstance.CreateOne()

	return userInstance, nil
}

//登录函数
func (s *UserService) Login(
	username, //手机
	plainpwd string) (user model.User, err error) {

	//首先通过手机号查询用户
	userInstance := model.User{}
	userInstance.Mobile = username
	userInstance = userInstance.GetOne()
	//如果没有找到
	if userInstance.Id == 0 {
		return userInstance, errors.New("该用户不存在")
	}
	//查询到了比对密码
	if !utils.ValidatePasswd(plainpwd, userInstance.Salt, userInstance.Passwd) {
		return userInstance, errors.New("密码不正确")
	}
	//刷新token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := utils.MD5Encode(str)
	userInstance.UpdateByField("token", token)

	//返回数据
	return userInstance, nil
}
