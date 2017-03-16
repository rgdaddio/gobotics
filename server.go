package main

/***

    TODO: How to properly daemonize. Apparently its discouraged to daemonize in golang. (threadings, coroutines, privalages related to fork)
            See:
                https://github.com/VividCortex/godaemon
                http://www.ryanday.net/2012/09/04/the-problem-with-a-golang-daemon/
                http://stackoverflow.com/questions/12486691/how-do-i-get-my-golang-web-server-to-run-in-the-background/
            Ideas:
                Supervisord
                upstart

    TODO: gunicorn in go or how do I mulitthread server in golang?
    TODO: Another webserver to talk to bots

***/

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"
)

type Bot struct {
    Name   string    `json:"name"`
    Uptime time.Time `json:"uptime"`
}

type Msg struct {
    Message   string    `json:"msg"`
    Timestamp time.Time `json:"timestamp"`
}

/***
    /list : list all running bots being managed

***/
func list(w http.ResponseWriter, req *http.Request) {
    type Bots []Bot
    bots := Bots{
        Bot{Name: "RasPI"},
        Bot{Name: "Arduino"},
    }

    json.NewEncoder(w).Encode(bots)
    log.Printf(req.Method)
    log.Printf(req.URL.Path)
    //fmt.Fprintf(w, "Hello LIST, %q", html.EscapeString(req.URL.Path))
}

/***
    /kill:  do a sys exit
***/
func die(w http.ResponseWriter, req *http.Request) {
    log.Printf(req.Method)
    log.Printf(req.URL.Path)
    msg := Msg{Message: "killing daemon...."}
    json.NewEncoder(w).Encode(msg)
    os.Exit(1)
}

func main() {
    log.SetOutput(os.Stdout)

    //TODO: Maybe use Gorilla Mux ot GIN? Docker uses mux
    mux := http.NewServeMux()
    mux.Handle("/list", http.HandlerFunc(list))
    mux.Handle("/die", http.HandlerFunc(die))

    log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
