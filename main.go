package main

import (
	"net/http"
	"fmt"
)

var handler http.HandlerFunc = func (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func main() {
	server := http.Server{
        Addr: "localhost:3000",
        Handler : handler,
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}