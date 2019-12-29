package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

func main() {
    port := os.Getenv("PORT")
    if(port == "") {
        port = "8080"
    }

    dat, err := ioutil.ReadFile("format.txt")
    if(err != nil) {
        panic(err)
    }
    fmt.Println("Format: " + string(dat))

    fmt.Println("Start listen port " + port)
    
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":" + port, nil)
    fmt.Println("Terminating")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}