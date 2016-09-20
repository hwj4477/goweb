package goweb

import (
    
    "fmt"
    "net/http"
)

func Run(handler http.Handler, port string) {
    
    fmt.Println("Starting goweb port:" + port)
    http.ListenAndServe(":" + port, handler)
}
