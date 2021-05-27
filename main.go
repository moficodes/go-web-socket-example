package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/moficodes/go-web-socket-example/config"

	"github.com/moficodes/go-web-socket-example/repository"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	db := config.InitDB()
	config.CreateRedisClient()
	wsServer := NewWebsocketServer(&repository.RoomRepository{DB: db}, &repository.UserRepository{DB: db})
	go wsServer.Run()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWS(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)
	log.Println("starting server on port", *addr)
	log.Fatalln(http.ListenAndServe(*addr, mux))
}
