package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

var format string;

func main() {
    port := os.Getenv("PORT")
    if(port == "") {
        port = "8080"
    }

    dat, err := ioutil.ReadFile("format.txt")
    if(err != nil) {
        panic(err)
    }
    format = string(dat)
    fmt.Println("Format: " + format)

    fmt.Println("Start listen port " + port)
    
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":" + port, nil)
    fmt.Println("Terminating")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, format, r.URL.Path[1:])
}