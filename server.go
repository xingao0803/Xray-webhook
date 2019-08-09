package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Get the Alert!")

    fmt.Println("Raw Request: ", request)
    fmt.Println()

    fmt.Println("method: ", request.Method)
    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        fmt.Printf("read body err, %v\n", err)
        return
    }
    println("json: ", string(body))

}

func main() {
	http.HandleFunc("/xray/", handler)
	http.ListenAndServe(":9999", nil)
}
