package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/LockettPundt/GOserver/routes"
	utils "github.com/LockettPundt/GOserver/utils"
	"github.com/slayer/autorestart"
)

func main() {
	autorestart.StartWatcher()
	fileServer := http.FileServer(http.Dir("./static"))

	dbName := utils.GoDotEnvVariable("DB_NAME")
	db_URI := utils.GoDotEnvVariable("DB_URI")

	fmt.Println(dbName)
	fmt.Println(db_URI)

	http.Handle("/", fileServer)
	http.HandleFunc("/form", routes.FormHandler)
	http.HandleFunc("/hello", routes.HelloHandler)

	port := ":5055"
	fmt.Printf(" ✔ Server starting on PORT %v ✔  \n", port)
	if err := http.ListenAndServe(":5055", nil); err != nil {
		log.Fatal(err)
	}
}
