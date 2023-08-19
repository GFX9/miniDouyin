package pg

import (
	"fmt"
	"miniDouyin/biz/model/miniDouyin/api"
)

// 处理登录请求
// 并填充response结构体
func DBUserLogin(request *api.UserLoginRequest, response *api.UserLoginResponse) {
	user := DBUserFromLoginRequest(request)

	if user.QueryUser() {
		// user存在
		fmt.Printf("user = %+v\n", user)

		// 校验密码
		if user.Passwd != request.Password {
			response.StatusCode = 1
			response.StatusMsg = "Password wrong!"
			return
		}
		response.StatusCode = 0
		response.UserID = int64(user.ID)
		response.Token = user.Username + user.Passwd
		response.StatusMsg = "Login successfully!"
		return
	}
	response.StatusCode = 2
	response.StatusMsg = "User not exist!"
}

// 处理注册请求
// 并填充response结构体
func DBUserRegister(request *api.UserRegisterRequest, response *api.UserRegisterResponse) {
	user := DBUserFromRegisterRequest(request)

	if user.QueryUser() {
		// user存在
		response.StatusCode = 1
		response.StatusMsg = "User already exists!"
		return
	}

	if user.insert() {
		response.StatusCode = 0
		response.UserID = int64(user.ID)
		response.Token = user.Token
		response.StatusMsg = "Register succesffully!"
	} else {
		response.StatusCode = 2
		response.StatusMsg = "Register failed!"
	}
}

// 获取User信息
func DBGetUserinfo(request *api.UserRequest, response *api.UserResponse) {
	user, err := DBGetUser(request)
	if err != nil {
		// 没有找到用户或token失败
		response.StatusCode = 1
		response.StatusMsg = err.Error()
	}
	// 填充结构体
	response.User = user.ToApiUser()
	response.StatusCode = 0
	response.StatusMsg = "Get user information successfully!"
}
