package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	log.Println("home endpoint hit")
	w.Write([]byte("Hello from Snippetbox"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Snippetbox"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Println("err: ", err)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Display a specific Snippetbox"))
	fmt.Fprint(w, "Display a specific snippet with ID %id...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/about", about)
	mux.HandleFunc("/health", health)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux) // host:port
	log.Fatal("log Fatal: ", err)
}
