package controller

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, `{"message": "Health Check Is Good!"}`)
}
