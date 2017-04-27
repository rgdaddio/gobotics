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
  //  Uptime time.Time `json:"uptime"`
}

type Devices []Device

type Msg struct {
    Message   string    `json:"msg"`
    Timestamp time.Time `json:"timestamp"`
}


/**

TODO: Move db funcs to a file of its own
**/

func init_db(db *sql.DB){
  stmt, _ := db.Prepare("create table if not exists client_devices( " +
      " name text, platform text, mac_address text, ip_address varchar(15), stats_table text);" )
  _, err := stmt.Exec()
  if err != nil { panic(err) }
}

func add_client_device(db *sql.DB, new_device Device){
    fmt.Println(new_device)

    stmt, err := db.Prepare("INSERT INTO client_devices( " +
                          " name, platform, mac_address, ip_address " +
                          " ) values(?,?,?,?)")
  if err != nil { fmt.Println("HI"); panic(err) }
  _, err = stmt.Exec(new_device.Name, new_device.Platform, new_device.Mac, new_device.Ip)
  if err != nil { panic(err) }
}

func update_client_device(db *sql.DB, device Device){
    fmt.Println(device)
    fmt.Println(device.Mac)


    stmt, err := db.Prepare("UPDATE client_devices SET  " +
                          " mac_address = ?,  ip_address = ? " +
                          " WHERE name = ?")
  if err != nil { fmt.Println("HI"); panic(err) }
  res, err := stmt.Exec(device.Mac, device.Ip, device.Name)
  if err != nil { panic(err) }
    affect, _ := res.RowsAffected()
    log.Println(affect)
}

func find_client_device(db *sql.DB, device_name string) Device{
    rows, err := db.Query("SELECT * from client_devices WHERE name = ?", device_name)
    if err != nil { log.Println("HI"); log.Fatal(err) }
    defer rows.Close()
    
    device := Device{}
    var name string 
    var platform string 
    var mac_address string 
    var ip_address string 
    if rows.Next() {
        rows.Scan(&name, &platform, &mac_address, &ip_address)
        device = Device{
            Name: name,
            Platform: platform,
            Mac: mac_address,
            Ip: ip_address,
          }
        log.Println(device)
    }
    return device
}

func get_client_devices(db *sql.DB) Devices{
    rows, err := db.Query("SELECT * from client_devices")
    if err != nil { log.Println("HI"); log.Fatal(err) }
    defer rows.Close()
   
    devices := Devices{} 
    for rows.Next() {
        device := Device{}
        var name string 
        var platform string 
        var mac_address string 
        var ip_address string 

        rows.Scan(&name, &platform, &mac_address, &ip_address)
        device = Device{
            Name: name,
            Platform: platform,
            Mac: mac_address,
            Ip: ip_address,
          }
        devices = append(devices, device)
    }
    return devices
}

func remove_client_device(db *sql.DB, device_name string) int64 {
    stmt, err := db.Prepare("DELETE FROM client_devices WHERE name = ?")
    if err != nil { panic(err) }
    res, err := stmt.Exec(device_name)
    if err != nil { panic(err) }
    affect, _ := res.RowsAffected()
    return affect
}


/***
    URI: /client/devices
    paths: 
        GET:
            responses:
                200:
                    description: list of all devices being managed
***/
func devices(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "GET":
            // List information on all devices
            devices := get_client_devices(db)
            json.NewEncoder(w).Encode(devices)

        default:
            // Give an error message.
            msg := Msg{Message: "Only GET supported for this api"}
            json.NewEncoder(w).Encode(msg)
    }
}

/***
    URI: /client/die 
     do a sys exit
***/
func die(w http.ResponseWriter, req *http.Request) {
    log.Printf(req.Method)
    log.Printf(req.URL.Path)
    msg := Msg{Message: "killing daemon...."}
    json.NewEncoder(w).Encode(msg)
    os.Exit(1)
}

/***
    URI: /client/device
    paths: 
        GET:
            query parameters:
                device: name of device to get info for
            responses:
                200:
                    description: return information on device
	POST:
            parameters:
                name string : name of device
                mac_address string : mac_address of device
                ip_address  string: ip of device
                platform string: platform
            responses:
                200:
                    description: device entry created
	PUT:
            parameters:
                name string : name of device
                mac_address string : mac_address of device
                ip_address  string: ip of device
                platform string: platform
            responses:
                200:
                    description: device entry updated based on information in body
	DELETE:
            query parameters:
                device: name of device to get info for
            responses:
                200:
                    description: device removed 
***/
func device(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "GET":
            // List information on a specific device
            log.Println(req.RequestURI)
            url_par, _ := url.Parse(req.RequestURI)
            qmap,  _ := url.ParseQuery(url_par.RawQuery)
            ret := find_client_device(db, qmap["device"][0])
            json.NewEncoder(w).Encode(ret)
        case "POST":
            // Add a new device.
            new_device := Device{}
            decoder := json.NewDecoder(req.Body)
            decoder.Decode(&new_device)
            log.Println(new_device)
            add_client_device(db, new_device)

        case "PUT":
            // Update an existing record.
            new_device := Device{}
            decoder := json.NewDecoder(req.Body)
            decoder.Decode(&new_device)
            log.Println(new_device)
            if (Device{}) == new_device {
                msg := Msg{Message: "Please give information for update"}
                json.NewEncoder(w).Encode(msg)
            }

            if new_device.Name == ""{
                msg := Msg{Message: "You must specify name when trying to update device"}
                json.NewEncoder(w).Encode(msg)
            }

            if new_device.Platform == "" && new_device.Mac == "" && new_device.Ip == ""{
                msg := Msg{Message: "Please give information for update"}
                json.NewEncoder(w).Encode(msg)
            }

            update_client_device(db, new_device)

        case "DELETE":
            // Remove the record.
            url_par, _ := url.Parse(req.RequestURI)
            qmap,  _ := url.ParseQuery(url_par.RawQuery)
            ret := remove_client_device(db, qmap["device"][0])
            if ret > 0 {
                msg := Msg{Message: "Device Removed"}
                json.NewEncoder(w).Encode(msg)
            } else{
                msg := Msg{Message: "Device name not found"}
                json.NewEncoder(w).Encode(msg)
            }
            json.NewEncoder(w).Encode(ret)

        default:
            // Give an error message.
            log.Println("Unknown Method")
    }
}

// TODO: this is scketches what we want if we ever want to start
// dynamically allocating REST APIs  
func dynamically_add_api( server *http.ServeMux, path string, handler_func func(http.ResponseWriter,*http.Request)){
    server.Handle(path, http.HandlerFunc(handler_func))
}

func main() {
    log.SetOutput(os.Stdout)

    var err error
    db, err = sql.Open("sqlite3", "./foo.db")
    //db.SetMaxIdleConns(50)
    fmt.Printf("%s\n", db)

    err = db.Ping() // make sure the database conn is alive
    if err != nil {
        log.Fatalf("Error on opening database connection: %s", err.Error())
    }
    init_db(db)

    //TODO: Maybe use Gorilla Mux ot GIN? Docker uses mux

    client_api_server := http.NewServeMux()
    client_api_server.Handle("/client/die", http.HandlerFunc(die))
    client_api_server.Handle("/client/device", http.HandlerFunc(device))
    client_api_server.Handle("/client/devices", http.HandlerFunc(devices))


    client_api_server.Handle("/security/scanDevicePort", http.HandlerFunc(scan_port))

    log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", client_api_server))
    // Debugging purposes
    //log.Fatal(http.ListenAndServe(":8080", client_api_server))

}
