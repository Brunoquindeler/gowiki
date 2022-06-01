package main

import (
	"log"
	"net/http"
)

const HOST = ":8080"

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Printf("server started on: %s", HOST)
	log.Fatal(http.ListenAndServe(HOST, nil))
}
