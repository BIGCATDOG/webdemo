package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)
const (
	CreatUserTable = "CREATE TABLE IF NOT EXISTS userinfo( uid INT )"
	CreateDataBase = "CREATE DATABASE userdetails; "
)
func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello web server") // 这个写入到 w 的是输出到客户端的

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		token := r.Form.Get("token")
		if token != "" {
			// 验证 token 的合法性
		} else {
			// 不存在 token 报错
		}

		// 请求的是登录数据，那么执行登录的逻辑判断
		if len(r.Form["username"][0]) == 0 || len(r.Form["password"][0]) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid username or password")

			return
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("Interesting:", r.Form["interest"])
        fmt.Fprintln(w,"hello login success!")
	}

}

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main() {
	initDB()
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload",upload)
	v := url.Values{}
	v.Add("hh", "fuck")
	fmt.Println(v.Get("hh"))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Server init is failed!")
	}
}

func initDB()  {
	db ,err := sql.Open("mysql","root:Ndalex123@tcp(localhost:3306)/")
	if err!=nil{
		fmt.Println(err)
	}
	smt ,err:= db.Prepare(CreateDataBase)
	if err!=nil{
		fmt.Println(err)
	}
	smt.Exec()
	if err!=nil{
		fmt.Println(err)
	}
	_,err1:= db.Exec("USE `userinfo` ")
	if err1!=nil{
		fmt.Println(err)
	}

	if err!=nil{
		fmt.Println(err)
	}
	smt1 ,err:= db.Prepare(CreatUserTable)
	if err!=nil{
		fmt.Println(err)
	}
	smt1.Exec()
	if err!=nil{
		fmt.Println(err)
	}

}