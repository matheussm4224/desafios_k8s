package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)    // Aqui declaramos o path e o Handler
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil) // Aqui ficará a porta em que nosso Web server vai responder
}

func Greeting(s string) (string){
	return "<b>"+s+"</b>";
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, Greeting("Code.education Rocks!"))
}

