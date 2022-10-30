package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/roll", rollerHandler)
	http.HandleFunc("/minRoll", rollerMinHandler)
	http.HandleFunc("/addRoll", rollerAddHandler)
	log.Fatal(http.ListenAndServe(":9001", nil))
	return;
}
