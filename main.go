package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello")

		data, err := ioutil.ReadAll(r.Body)

		log.Println(string(data))

		if err != nil {
			http.Error(rw, "Error occurred", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Invalid"))
			return
		}

		fmt.Fprintf(rw, "Hello")

	})

	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Bye")

		data, _ := ioutil.ReadAll(r.Body)
		log.Println(string(data))

		fmt.Fprintf(rw, "Bye")

	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening")
}
