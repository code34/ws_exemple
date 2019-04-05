package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    type User struct {
        Firstname string `json:"firstname"`
        Lastname string `json:"lastname"`
    }

    type httpbin struct {
        Origin string `json:"origin"`
        Headers map[string]string `json:"headers"`
        Myuser map[string]string `json:"json"`
    }

    u := User{Firstname: "doy", Lastname: "nic"}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(u)
    res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
    newhttpbin := httpbin{}
    json.NewDecoder(res.Body).Decode(&newhttpbin)
    fmt.Println(newhttpbin.Myuser)
}