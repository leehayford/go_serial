package main

import (
	"log"
	"net/http"

	"github.com/leehayford/go_serial/pkg/ser"
)

func main() {

	/*SERIAL STUFF*/
	log.Println("\tSetting up serial port...")
	ser.RegisterRoutes()

	/*FRONT END STUFF - SVELTE*/
	static := http.FileServer(http.Dir("./web/public"))
	http.Handle("/", static)

	log.Println("\tStarting server, listening on port 8002")
	http.ListenAndServe("127.0.0.1:8002", nil)
	// http.ListenAndServe("0.0.0.0:8002", nil)

}
