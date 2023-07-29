package main

import (
    "fmt"
    "log"
    "net/http"
)


func main(){

    fileServer := http.FileServer(http.Dir('./static/'))
    http.Handle("/",fileServer)
    http.HandleFunc("/form",formHandler)
    http.HandleFunc("/hello",helloHandler)

    fmt.Printf("Server starting at 8080 port")

    if err:

}


