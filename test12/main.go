package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("test")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

	//select {}
}



