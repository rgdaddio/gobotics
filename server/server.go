package server

import (
    "log"
    "fmt"
    "net/http"
    "os"
    "database/sql"
)

var db *sql.DB // global variable to share it between main and the HTTP handler
                //represents a connection pool, not a single connection



/**

TODO: Move db funcs to a file of its own
**/

func init_db(db *sql.DB){
  stmt, _ := db.Prepare("create table if not exists client_devices( " +
      " name text, platform text, mac_address text, ip_address varchar(15), stats_table text);" )
  _, err := stmt.Exec()
  if err != nil { panic(err) }
}


func Serve() {
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

    s := &http.Server{
        Addr: ":8080",
        Handler: client_api_server,
    }

    //log.Fatal(s.ListenAndServeTLS("cert.pem", "key.pem"))
    // Debugging purposes
    log.Fatal(s.ListenAndServe())

}
