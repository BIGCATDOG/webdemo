
package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)
type Reply struct {
	Uid int `json:"u_id"`
	UName string `json:"u_Name"`
}
func Echo(ws *websocket.Conn) {
	var err error
    fmt.Println( ws.LocalAddr() )
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)
        if err = websocket.JSON.Send(ws,Reply{Uid: 10,UName: "fuck"});err!=nil{
				fmt.Println("Can't send")
				break
		}
		//if err = websocket.Message.Send(ws, msg); err != nil {
		//	fmt.Println("Can't send")
		//	break
		//}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}