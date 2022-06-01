package main

import (
	"log"
	"net/http"
)

const HOST = ":8080"

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Printf("server started on: %s", HOST)
	log.Fatal(http.ListenAndServe(HOST, nil))
}
