package main

import (
	"game/api/controller"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", controller.HealthCheckHandler)
	mux.HandleFunc("/users/register", controller.Register)
	http.ListenAndServe(":8080", mux)

	//mySqlDb := mysql.NewDB()
	//phoneNumber := "09120000000"
	//isUnique, _ := mySqlDb.IsPhoneNumberUnique(phoneNumber)
	//fmt.Println("before persist", isUnique)
	//user, _ := mySqlDb.Save(entity.User{Username: "admin", PhoneNumber: phoneNumber})
	//fmt.Println(user.Id)
	//isUnique, _ = mySqlDb.IsPhoneNumberUnique(phoneNumber)
	//fmt.Println("after persist", isUnique)

}
