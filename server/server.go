package server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// ServerMsg - default return message for all endpoints
type ServerMsg struct {
	Message string `json:"message"`
}

func Serve() {

	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting HTTP Server")

	//TODO: instantiate SqlLiteDevicesLib
	var err error

	client_api_server := http.NewServeMux()

	//TODO add discovery

	// React App
	path := "/Users/ssikdar1/go/src/github.com/rgdaddio/gobotics/gobotics-frontend/build/"

	client_api_server.Handle("/", staticFileHandler(http.FileServer(http.Dir(path))))

	client_api_server.Handle("/client/die", http.HandlerFunc(die))
	client_api_server.Handle("/client/device", http.HandlerFunc(DeviceHandler))
	client_api_server.Handle("/client/devices", http.HandlerFunc(DevicesHandler))
	client_api_server.Handle("/healthcheck", http.HandlerFunc(HealthCheckHandler))

	s := &http.Server{
		Addr:    ":8080",
		Handler: client_api_server,
	}

	//log.Fatal(s.ListenAndServeTLS("cert.pem", "key.pem"))
	// Debugging purposes
	log.Fatal(s.ListenAndServe())

}
