package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conns []*websocket.Conn //把所有的conn都存起來

func handle(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	conns = append(conns, conn) //把所有的conn都存起來
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("Are you saying:"+string(p)+"?")) //把所有的conn都存起來
		}

		fmt.Println(m, string(p)) //string(p)才會印出文字

	}
	defer conn.Close()
	log.Println("conn close")
}

func main() {

	http.HandleFunc("/", handle)

	http.ListenAndServe(":8080", nil)

}
