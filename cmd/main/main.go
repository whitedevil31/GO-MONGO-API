package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/whitedevil31/go-mongo-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStudentRoutes(r)
	http.Handle("/", r)
	viper.SetConfigFile(".env")
	value, _ := viper.Get("PORT").(string)
	host := "0.0.0.0" + ":" + value
	fmt.Println("APP RUNNING ON " + host)
	http.ListenAndServe(host, r)

}
