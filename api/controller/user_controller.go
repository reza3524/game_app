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
		fmt.Fprintf(writer, "Method not allowed")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	user := request.RegisterUserRequest{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	mysqldb := mysql.NewDB()
	userService := service.NewUser(mysqldb)
	_, err = userService.Register(user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	writer.Write([]byte(`saved!`))
}

func Login(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, "Method not allowed")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
	}
	user := request.LoginUserDto{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	mysqldb := mysql.NewDB()
	userService := service.NewUser(mysqldb)
	err = userService.Login(user)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	writer.Write([]byte(`LoginSuccess!`))
}
