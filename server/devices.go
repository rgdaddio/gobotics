package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Device struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Mac      string `json:"mac_address"`
	Ip       string `json:"ip_address"`
	//  Uptime time.Time `json:"uptime"`
}

type Devices []Device

// struct methods?
func addClientDevice(db *sql.DB, new_device Device) {
	stmt, err := db.Prepare("INSERT INTO client_devices( " +
		" name, platform, mac_address, ip_address " +
		" ) values(?,?,?,?)")
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec(new_device.Name, new_device.Platform, new_device.Mac, new_device.Ip)
	if err2 != nil {
		panic(err)
	}
}

func updateClientDevice(db *sql.DB, device Device) {

	stmt, err := db.Prepare("UPDATE client_devices SET  " +
		" mac_address = ?,  ip_address = ? " +
		" WHERE name = ?")
	if err != nil {
		fmt.Println("HI")
		panic(err)
	}
	res, err := stmt.Exec(device.Mac, device.Ip, device.Name)
	if err != nil {
		panic(err)
	}
	affect, _ := res.RowsAffected()
	log.Println(affect)
}

func findClientDevice(db *sql.DB, device_name string) (Device, error) {
	rows, err := db.Query("SELECT * from client_devices WHERE name = ?", device_name)
	if err != nil {
		return Device{}, err
	}
	defer rows.Close()

	device := Device{}
	var name string
	var platform string
	var mac_address string
	var ip_address string
	if rows.Next() {
		rows.Scan(&name, &platform, &mac_address, &ip_address)
		device = Device{
			Name:     name,
			Platform: platform,
			Mac:      mac_address,
			Ip:       ip_address,
		}
		log.Println(device)
	}
	return device, err
}

func getClientDevices(db *sql.DB) Devices {
	rows, err := db.Query("SELECT name, platform, mac_address, ip_address from client_devices")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	devices := Devices{}
	for rows.Next() {
		device := Device{}
		var name string
		var platform string
		var mac_address string
		var ip_address string

		rows.Scan(&name, &platform, &mac_address, &ip_address)
		log.Println(rows)
		log.Println(name)
		device = Device{
			Name:     name,
			Platform: platform,
			Mac:      mac_address,
			Ip:       ip_address,
		}
		devices = append(devices, device)
	}
	return devices
}

func remove_client_device(db *sql.DB, device_name string) int64 {
	stmt, err := db.Prepare("DELETE FROM client_devices WHERE name = ?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(device_name)
	if err != nil {
		panic(err)
	}
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
func DevicesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// List information on all devices
		devices := getClientDevices(db)
		json.NewEncoder(w).Encode(devices)

	default:
		// Give an error message.
		msg := ServerMsg{Message: "HTTP Method not supported"}
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
	msg := ServerMsg{Message: "killing daemon...."}
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
func DeviceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// List information on a specific device
		log.Println(req.RequestURI)
		url_par, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(url_par.RawQuery)
		// if no device throw 404
		ret, err := findClientDevice(db, qmap["device"][0])
		if err != nil {

		}

		json.NewEncoder(w).Encode(ret)
	case "POST":
		// Add a new device.
		new_device := Device{}
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&new_device)
		addClientDevice(db, new_device)

	case "PUT":
		// Update an existing record.
		new_device := Device{}
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&new_device)
		if (Device{}) == new_device {
			msg := ServerMsg{Message: "Please give information for update"}
			json.NewEncoder(w).Encode(msg)
		}

		if new_device.Name == "" {
			msg := ServerMsg{Message: "You must specify name when trying to update device"}
			json.NewEncoder(w).Encode(msg)
		}

		if new_device.Platform == "" && new_device.Mac == "" && new_device.Ip == "" {
			msg := ServerMsg{Message: "Please give information for update"}
			json.NewEncoder(w).Encode(msg)
		}

		updateClientDevice(db, new_device)

	case "DELETE":
		// Remove the record.
		url_par, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(url_par.RawQuery)
		ret := remove_client_device(db, qmap["device"][0])
		if ret > 0 {
			msg := ServerMsg{Message: "Device Removed"}
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := ServerMsg{Message: "Device name not found"}
			json.NewEncoder(w).Encode(msg)
		}
		json.NewEncoder(w).Encode(ret)

	default:
		// Give an error message.
		log.Println("Unknown Method")
	}
}
