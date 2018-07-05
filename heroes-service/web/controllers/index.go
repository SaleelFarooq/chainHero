package controllers

import (
	"fmt"
	//"encoding/json"
	"net/http"
	"github.com/gorilla/securecookie"
)

import "database/sql"
import "log"
import _ "github.com/go-sql-driver/mysql"


func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func (app *Application)LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		http.Redirect(w, r, "/home.html", 302)
	} else {
		
		renderTemplate1(w, r, "index.html", nil)
	}
}

func (app *Application) GHomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HAla")
	userName := getUserName(r)
	fmt.Println(userName)

    //if userName != "" {
	renderTemplate(w, r, "home.html", nil)
    //}else {
		
		//renderTemplate(w, r, "index.html", nil)
	//}
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))



func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func Authenticate(username string,password string,typ string) string {
	db, err := sql.Open("mysql","root:saleel@tcp(127.0.0.1:3306)/account")
    if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM credentials where username=? and password=? and type=?",username,password,typ)
	if err != nil {
		panic(err.Error())
	}
	if(results.Next()){
		return "VALID"
	}
	return "INVALID"

}

// voter login handler

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request){ 

  
  fmt.Println("Continue 2")
  name := r.PostFormValue("username")
  pwd := r.PostFormValue("password")
  typ := r.PostFormValue("type")
   redirectTarget := "/"
    
    a:=Authenticate(name,pwd,typ)
    fmt.Println(a)
    

    if(a=="VALID"){
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //w.WriteHeader(http.StatusAccepted)
    fmt.Println("Logged In")
    setSession(name, w)
			
    http.Redirect(w, r, "/home.html", 302)
    
 //http.Redirect(w, r, redirectTarget, 302)
    }else{
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Println("Incorrect credentials")
    msg:="Didn't Logged In"
    fmt.Fprintf(w,"%s\n",msg)
    }
   
 fmt.Println(redirectTarget)


}




func Register(username string,password string,typ string) string{
str:=""
	db, err := sql.Open("mysql", "root:saleel@tcp(127.0.0.1:3306)/account")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    insert, err := db.Query("INSERT INTO credentials VALUES (?,?,?)",username,password,typ)
    
    if err != nil {
        panic(err.Error())
        str ="NOT ADDED"
    }else{
     str = "ADDED"
}
    

defer insert.Close()
 return str   
}

func (app *Application) SignUpHandler(w http.ResponseWriter,r *http.Request){

name := r.PostFormValue("username")
  pwd1 := r.PostFormValue("password1")
  pwd2 := r.PostFormValue("password2")
  typ :=  r.PostFormValue("type")
    //fmt.Fprintf(w, "Hello, %s! .. you are going to register %s --- %s", name,pwd1,pwd2)
//msg := "2 password don't match"
   if pwd1==pwd2{
     a := Register(name,pwd1,typ)
    fmt.Println(a)
    if(a=="ADDED"){
	setSession(name, w)
	fmt.Println("ADDED SUCCESSSFULLY")		
    http.Redirect(w, r, "/home.html", 302)
}else{
fmt.Println("FAILURE DURING SIGNING UP")
http.Redirect(w, r, "/index.html", 302) 
}
}else{
	fmt.Println("2 PASSWORDS DONT MATCH")
http.Redirect(w, r, "/index.html", 302) 
}


}
// logout handler

func (app *Application) LogOutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}
