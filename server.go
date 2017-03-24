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


***/

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "os"
    "time"
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

var db *sql.DB // global variable to share it between main and the HTTP handler
                //represents a connection pool, not a single connection


type Device struct {
    Name   string    `json:"name"`
    Platform string `json:"platform"`
    Mac string `json:"mac_address"`
    Ip string `json:"ip_address"`
    Uptime time.Time `json:"uptime"`
}

type Msg struct {
    Message   string    `json:"msg"`
    Timestamp time.Time `json:"timestamp"`
}


/**

TODO: Move db funcs to a file of its own
**/
func add_client_device(db *sql.DB, new_device Device){
    fmt.Println(new_device)

    stmt, err := db.Prepare("INSERT INTO client_devices( " +
                          " name, platform, mac_address, ip_address " +
                          " ) values(?,?,?,?)")
  if err != nil { fmt.Println("HI"); panic(err) }
  _, err = stmt.Exec(new_device.Name, new_device.Platform, new_device.Mac, new_device.Ip)
  if err != nil { panic(err) }
}

func find_client_device(device_name string) Device{
    log.Println(device_name)
    rows, err := db.Query("SELECT * from client_devices WHERE name = ?", device_name)
    if err != nil { log.Println("HI"); log.Fatal(err) }
    defer rows.Close()
    
    device := Device{}
    if rows.Next() {
        rows.Scan(&device)
        log.Println(device)
    }
    return device
}

/***
    /list : list all running bots being managed

***/
func list(w http.ResponseWriter, req *http.Request) {
    type Bots []Device
    bots := Bots{
        Device{Name: "RasPI"},
        Device{Name: "Arduino"},
    }

    json.NewEncoder(w).Encode(bots)
    log.Printf(req.Method)
    log.Printf(req.URL.Path)
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

func device(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "GET":
            // List information on a specific device
            log.Println(req.RequestURI)
            url_par, _ := url.Parse(req.RequestURI)
            qmap,  _ := url.ParseQuery(url_par.RawQuery)
            log.Println(qmap["device"][0])
            ret := find_client_device(qmap["device"][0])
            log.Println(ret)
            json.NewEncoder(w).Encode(ret)
        case "POST":
            // Add a new device.
            new_device := Device{}
            decoder := json.NewDecoder(req.Body)
            decoder.Decode(&new_device)
            add_client_device(db, new_device)

        case "PUT":
            // Update an existing record.
            log.Println("PUT not yet implemented")

        case "DELETE":
            // Remove the record.
            log.Println("DELETE not yet implemented")

        default:
            // Give an error message.
            log.Println("Unknown Error")
    }
}

func main() {
    log.SetOutput(os.Stdout)

    var err error
    db, err = sql.Open("sqlite3", "./foo.db")
    //db.SetMaxIdleConns(50)
    fmt.Printf("%s", db)

    err = db.Ping() // make sure the database conn is alive
    if err != nil {
        log.Fatalf("Error on opening database connection: %s", err.Error())
    }

    //TODO: Maybe use Gorilla Mux ot GIN? Docker uses mux

    client_api_server := http.NewServeMux()
    client_api_server.Handle("/client/list", http.HandlerFunc(list))
    client_api_server.Handle("/client/die", http.HandlerFunc(die))
    client_api_server.Handle("/client/device", http.HandlerFunc(device))
    // log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", client_api_server))
    // Debugging purposes
    log.Fatal(http.ListenAndServe(":8080", client_api_server))

}
