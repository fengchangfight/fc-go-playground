package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func main() {
    url := "http://192.168.38.94:8115/frame-rule"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"sentence":"I would like to book a ticket from Shanghai to Beijing","rule":"what","keylist":"origin,destination"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

