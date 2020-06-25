package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
)

//TODO add license
//TODO Godocs?
//TODO unit tests
var db *sql.DB // global variable to share it between main and the HTTP handler
//represents a connection pool, not a single connection

func init_db(db *sql.DB) {
	stmt, _ := db.Prepare("create table if not exists client_devices( " +
		" name text, platform text, mac_address text, ip_address varchar(15), stats_table text);")
	_, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func Serve() {
	log.SetOutput(os.Stdout)

	var err error
	db, err = sql.Open("sqlite3", "./foo.db")
	//db.SetMaxIdleConns(50)
	fmt.Printf("%#v\n", db)

	err = db.Ping() // make sure the database conn is alive
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
	init_db(db)

	//TODO: Maybe use Gorilla Mux ot GIN? Docker uses mux

	client_api_server := http.NewServeMux()

	//TODO add discovery

	// React App
	// TODO if this were a real app: serve with NGINX?
	path := "/Users/ssikdar1/go/src/github.com/rgdaddio/gobotics/gobotics-frontend/build/"
	log.Fatal(http.ListenAndServe(":8080", staticFileHandler(http.FileServer(http.Dir(path)))))

	client_api_server.Handle("/client/die", http.HandlerFunc(die))
	client_api_server.Handle("/client/device", http.HandlerFunc(device))
	client_api_server.Handle("/client/devices", http.HandlerFunc(devices))
	client_api_server.Handle("/health", http.HandlerFunc(healthcheck))

	httpMux := reflect.ValueOf(client_api_server).Elem()
	finList := httpMux.FieldByIndex([]int{1})
	fmt.Println(finList)

	s := &http.Server{
		Addr:    ":8080",
		Handler: client_api_server,
	}

	//log.Fatal(s.ListenAndServeTLS("cert.pem", "key.pem"))
	// Debugging purposes
	log.Fatal(s.ListenAndServe())

}
