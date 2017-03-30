package main
import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "strconv"
)

type PortScan struct {
    Name   string    `json:"name"`
    Ip string `json:"ip_address"`
    Port int `json:"port"`
    Closed bool `json:"tcp_port_closed"`
}

/***
    URI: /security/scanDevicePort
    paths: 
        GET:
            query parameters:
                device: name of device to port scan
		portNo: port number
***/
func scan_port(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "GET":
            // List information on a specific device
            log.Println(req.RequestURI)
            url_par, _ := url.Parse(req.RequestURI)
            qmap,  _ := url.ParseQuery(url_par.RawQuery)
            log.Println(qmap)
            ret := find_client_device(db, qmap["device"][0])
            port_num, _ := strconv.Atoi(qmap["portNo"][0])
            ret2 :=port_scan(ret.Ip, port_num)
            scan := PortScan{
                Name: qmap["device"][0],
                Ip: ret.Ip,
                Port: port_num,
                Closed: ret2,
            }
 
            json.NewEncoder(w).Encode(scan)

        default:
            // Give an error message.
            msg := Msg{Message: "Only GET supported for this api"}
            json.NewEncoder(w).Encode(msg)
    }
}
