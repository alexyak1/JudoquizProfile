package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	handleRequests()
}

func handleRequests() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8788"
	}
	fmt.Println(port)

	http.HandleFunc("/user", getUser)
	fmt.Println("UserService is running on :8788")
	http.ListenAndServe(port, nil)

}
func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Data")
}
