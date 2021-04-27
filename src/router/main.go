package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	"context"
	"github.com/go-redis/redis/v8"
)

func getMessage(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {

	http.SetCookie(resp, &http.Cookie{Name: "UserName", Value: "alex", MaxAge: 3000, Expires: time.Now().Add(5 * time.Minute)})
	http.SetCookie(resp, &http.Cookie{Name: "SessionId", Value: "2292929", MaxAge: 3000, Expires: time.Now().Add(5 * time.Minute)})
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("get message from server post"))

}
func addMessage(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {
	cookie, err := req.Cookie("userName")
	if err == nil {
		fmt.Println(cookie.Value)
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Added message from server post"))
}
func update(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("updated message from server post"))
}
func delete(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("delete message from server post"))
}
func redirect(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {
   fmt.Println( req.FormValue("redirecturl") )
   // resp.WriteHeader(http.StatusTemporaryRedirect)
   http.Redirect(resp,req,req.FormValue("redirecturl"),http.StatusTemporaryRedirect)
    //resp.Write([]byte("delete message from server post"))
}

func DownloadFile(resp http.ResponseWriter, req *http.Request, param httprouter.Params)  {
	//file,err:=os.Open(param.ByName("filename"))
	//if err!=nil{
	//	resp.WriteHeader(http.StatusNotFound)
	//	resp.Write([]byte(err.Error()))
	//	return
	//}
	http.ServeFile(resp,req,param.ByName("filename"))
	//buf := make([]byte,1)
	//resp.Header().Add("Content-Type", "application/octet-stream")
	//resp.Header().Add("Content-Disposition", "attachment; filename=\""+"fuckdufji"+"\"")
	//for ;;{
	//	size,err:=file.Read(buf)
	//	if err!=nil{
	//		if err == io.EOF{
	//			resp.Write(buf[:size])
	//
	//			return
	//		} else{
	//
	//		}
	//	}else{
	//		resp.Write(buf[:size])
	//	}
	//}

}
func main() {
	ExampleClient()
	router := httprouter.New()
	router.GET("/getMessage", getMessage)
    router.GET("/redirect", redirect)
	router.POST("/addMessage", addMessage)
	router.PUT("/updateMessage", update)
	router.DELETE("/delete", delete)
	router.GET("/download/:filename",DownloadFile)
	if err := http.ListenAndServe("localhost:7070", router); err != nil {
		fmt.Println("server internal exception!")
	}
}

var ctx = context.Background()

type people struct {
	UserId int `redis:"userId`
	UserName string `redis:userName`
	Gender bool `redis:gender`
}
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	err1 := rdb.HMSet(ctx,"alex","gender",false,"userId",5,"userName","hhhh").Err()
	if err1 != nil {
		panic(err1)
	}
	var p1 people

	 rdb.HMGet(ctx,"alex","gender","userId")

	err2 := rdb.HGetAll(ctx, "alex").Scan(&p1)
	if err != nil {
		panic(err2)
	}

	rdb.LPush(ctx,"list1","hhh",false,"tueue")

	//rdb.HMSet(ctx,"people.alex",)
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	rdb.ZAdd(ctx,"userMoneys",&redis.Z{Score: 100,Member: "alex"},&redis.Z{Score: 22,Member: "joe"})
	//vals, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
	//	Keys: []string{"zset1", "zset2"},
	//	Weights: []float64{2.0, 3.0},
	//}).Result()
	//println( vals )
	res ,_:= rdb.ZPopMax(ctx,"userMoneys",1).Result()
	res1:= rdb.ZRange(ctx,"userMoneys",-1,33)
	println(res,res1)

	rdb.Pipeline()
	// Output: key value
	// key2 does not exist
}
