package main

// Importing libraries
import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/WEBAPP1/handler"
)

// Calling the functions
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/signup", handler.Signup)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	fmt.Println("Starting the server at :8080")
	http.ListenAndServe(":8080", nil)
}
