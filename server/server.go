package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rgdaddio/gobotics/utils/clientdevices"
	log "github.com/sirupsen/logrus"
)

// ServerMsg - default return message for all endpoints
type ServerMsg struct {
	Message string `json:"message"`
}

type Server struct {
	DeivcesClient clientdevices.ClientDevices
}

/***
    URI: /client/die
     do a sys exit
***/
func (s *Server) die(w http.ResponseWriter, req *http.Request) {
	log.Printf(req.Method)
	log.Printf(req.URL.Path)
	msg := ServerMsg{Message: "killing daemon...."}
	json.NewEncoder(w).Encode(msg)
	os.Exit(1)
}

func Serve() {

	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting HTTP Server")

	s := Server{}

	sqlliteOptions := clientdevices.NewSqlLiteDefaultOptions()
	devicesClient, err := clientdevices.NewSqlLiteClient(sqlliteOptions)
	if err != nil {
		log.Fatal("Error instantiating devicesClient")
	}
	s.DeivcesClient = devicesClient

	client_api_server := http.NewServeMux()

	//TODO add discovery

	// React App
	path := "/Users/ssikdar1/go/src/github.com/rgdaddio/gobotics/gobotics-frontend/build/"

	client_api_server.Handle("/", s.staticFileHandler(http.FileServer(http.Dir(path))))

	client_api_server.Handle("/die", http.HandlerFunc(s.die))
	client_api_server.Handle("/client/device", http.HandlerFunc(s.DeviceHandler))
	client_api_server.Handle("/client/devices", http.HandlerFunc(s.DevicesHandler))
	client_api_server.Handle("/healthcheck", http.HandlerFunc(s.HealthCheckHandler))

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: client_api_server,
	}

	//log.Fatal(s.ListenAndServeTLS("cert.pem", "key.pem"))
	// Debugging purposes
	log.Fatal(httpServer.ListenAndServe())

}
