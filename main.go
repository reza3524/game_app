package main

import (
	"fmt"
	"game/api/controller"
	"net/http"
)

func main() {
	port := 8080

	mux := http.NewServeMux()
	mux.HandleFunc("/health", controller.HealthCheckHandler)
	mux.HandleFunc("/users/register", controller.Register)
	mux.HandleFunc("/users/login", controller.Login)
	mux.HandleFunc("/users/profile/", controller.Profile)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("server is up and running on port %d:\n", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
	//mySqlDb := mysql.NewDB()
	//phoneNumber := "09120000000"
	//isUnique, _ := mySqlDb.IsPhoneNumberUnique(phoneNumber)
	//fmt.Println("before persist", isUnique)
	//user, _ := mySqlDb.Save(entity.User{Username: "admin", PhoneNumber: phoneNumber})
	//fmt.Println(user.Id)
	//isUnique, _ = mySqlDb.IsPhoneNumberUnique(phoneNumber)
	//fmt.Println("after persist", isUnique)

}
