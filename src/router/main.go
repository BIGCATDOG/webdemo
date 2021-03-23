package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
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
