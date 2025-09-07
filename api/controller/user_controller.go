package controller

import (
	"encoding/json"
	"fmt"
	"game/api/request"
	"game/service/impl"
	"game/storage/mysql"
	"io"
	"net/http"
)

func Register(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, "Method not allowed")))
		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	user := request.UserRegisterRequest{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	mysqldb := mysql.NewDB()
	userService := service.NewUser(mysqldb)
	response, err := userService.Register(user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	res, err := json.Marshal(response)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
	}
	writer.Write(res)
}

func Login(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, "Method not allowed")))
		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	user := request.UserLoginRequest{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	mysqldb := mysql.NewDB()
	userService := service.NewUser(mysqldb)
	response, err := userService.Login(user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	res, err := json.Marshal(response)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
	}
	writer.Write(res)
}

func Profile(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, "Method not allowed")))
		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	user := request.UserProfileRequest{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	mysqldb := mysql.NewDB()
	userService := service.NewUser(mysqldb)
	response, err := userService.Profile(user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	res, err := json.Marshal(response)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	writer.Write(res)

}
