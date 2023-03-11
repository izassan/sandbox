package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func connectWebsocket(ws *websocket.Conn){
    defer ws.Close()

    err := websocket.Message.Send(ws, "hello websocket client!")
    if err != nil{
        log.Fatal(err)
    }

    for {
        // recieve message
        msg := ""
        err = websocket.Message.Receive(ws, &msg)
        if err != nil{
            log.Fatal(err)
        }

        // return message
        if msg == "exit"{
            break
        }
        err = websocket.Message.Send(ws, fmt.Sprintf(`server>> %q`, msg))
        if err != nil{
            log.Fatal(err)
        }
    }
    err = websocket.Message.Send(ws, "ByeBye")
    if err != nil{
        log.Fatal(err)
    }
}

func testMiddleware(w http.ResponseWriter, r *http.Request){
    fmt.Println("connection test middleware")
    websocket.Handler(connectWebsocket).ServeHTTP(w, r)
}

func main() {
    websocketHandler := websocket.Handler(connectWebsocket)

    http.Handle("/ws", websocketHandler)
    http.HandleFunc("/ws/middle", testMiddleware)

    fmt.Println("Server Listening on ':31888'")
    err := http.ListenAndServe(":31888", nil)
    if err != nil{
        panic(err)
    }
}
