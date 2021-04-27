package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type Reply struct {
	Uid   int    `json:"u_id"`
	UName string `json:"u_Name"`
}

func Echo(ws *websocket.Conn) {
	var err error
	fmt.Println(ws.LocalAddr())
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)
		if err = websocket.JSON.Send(ws, Reply{Uid: 10, UName: "fuck"}); err != nil {
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
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test_DB?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	insertDataToMysql(db)
	rows,err1:=db.Query("SELECT userId,userName FROM user_info WHERE userId BETWEEN 0 AND 1000 LIMIT 5")
	if err1!=nil{

	}
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next(){
		var value string
		var userId int
        rows.Scan(&userId,&value)
		fmt.Printf("userId is %d userName is %s",userId,value)

	}
	defer db.Close()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func insertDataToMysql(db *sql.DB) {
	res, err := db.Exec("INSERT INTO user_info(userId,userName,nick_name) VALUES (100,'hhhh','fuck67')")
	if err != nil {

	}
	fmt.Print(res)
}
