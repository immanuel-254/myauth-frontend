package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/immanuel-254/myauth-frontend/components"
	"github.com/immanuel-254/myauth-frontend/pages"
)

func main() {
	// First create the index.html
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = pages.Base(components.NavBar(), pages.LoginPage()).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

	// Run the app
	component := pages.Base(components.NavBar())

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./style.css")
	})
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.Handle("/", templ.Handler(component))
	http.Handle("/login", templ.Handler(pages.Base(pages.LoginPage())))

	fmt.Println("Listening on :5000")
	http.ListenAndServe(":5000", nil)
}

// <nav aria-label="Main" data-orientation="horizontal" dir="ltr" class="relative z-10 flex max-w-max flex-1 items-center justify-center my-1">
// <div style="position:relative"><ul data-orientation="horizontal" class="group flex flex-1 list-none items-center justify-center space-x-1" dir="ltr"><li class="text-sm px-2"><a href="#">Docs</a></li><li class="text-sm px-2"><a href="#">Contact</a></li><li class="text-sm px-2"><a href="#">About</a></li></ul></div><div class="absolute left-0 top-full flex justify-center"></div></nav>
