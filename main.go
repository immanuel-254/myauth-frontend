package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/immanuel-254/myauth-frontend/pages"
)

func main() {
	// Run the app
	component := pages.Base(pages.Home())

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./style.css")
	})
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./script.js")
	})
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/daily-login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			date := r.FormValue("dailyDate")

			fmt.Fprintf(w, "%s", date)
		}
	})

	http.Handle("/", templ.Handler(component))
	http.Handle("/profile", templ.Handler(pages.Base(pages.ProfilePage())))
	http.Handle("/login", templ.Handler(pages.Base(pages.LoginPage())))
	http.Handle("/logout", templ.Handler(pages.Base(pages.LogoutPage())))
	http.Handle("/change-password", templ.Handler(pages.Base(pages.ChangePasswordPage())))
	http.Handle("/reset-password", templ.Handler(pages.Base(pages.ResetPasswordPage())))

	fmt.Println("Listening on :5000")
	http.ListenAndServe(":5000", nil)
}
