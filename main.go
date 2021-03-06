package main

import (
	"log"
	"net/http"
	"chat"
	//"os"
)

func main() {

	//port := os.Getenv("PORT")

	port:="8080"
	if port == "" {
		//log.Fatal("$PORT must be set")
		port="8080"
	}
	log.Println("PUERTO FINAL:"+port)
	log.SetFlags(log.Lshortfile)

	// websocket server
	server := chat.NewServer("/entry")
	go server.Listen()
	// static files
	http.Handle("/", http.FileServer(http.Dir("src/github.com/Alejosd/chatSd/webroot")))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
