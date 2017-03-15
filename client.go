package main

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func send_req(url string) string {
    fmt.Printf("%s\n", url)
    var ret string = "error"
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    req, err := client.Get(url)

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer req.Body.Close()
        contents, err := ioutil.ReadAll(req.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Fprintf(os.Stdout, "Response\n")
        fmt.Printf("%s\n", string(contents))
        ret = string(contents)
    }
    return ret
}

func main() {
    send_req("http://localhost:8080/list")
    send_req("http://localhost:8080/die")
}
