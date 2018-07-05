package web

import (
	"fmt"
	"github.com/chainHero/heroes-service/web/controllers"
	"net/http"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
    http.HandleFunc("/",app.LoginPageHandler)
    http.HandleFunc("/index.html",app.LoginPageHandler)
	http.HandleFunc("/home.html", app.HomeHandler)
	http.HandleFunc("/internal.html", app.LoginHandler)
	http.HandleFunc("/register.html", app.SignUpHandler)
	http.HandleFunc("/request.html", app.RequestHandler)
    http.HandleFunc("/history.html", app.HistoryHandler)
     http.HandleFunc("/logout.html", app.LogOutHandler)
	//http.HandleFunc("/home.html", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("GGMU")
		//http.Redirect(w, r, "/home.html", http.StatusTemporaryRedirect)
	//})

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}
