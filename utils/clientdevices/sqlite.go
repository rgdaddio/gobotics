package clientdevices

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

//SqlLiteDevicesLib implements the ClientDevice interface
type SqlLiteDevices struct {
	db *sql.DB
}

func initDB(db *sql.DB) error {
	stmt, _ := db.Prepare("create table if not exists client_devices( " +
		" name text, platform text, mac_address text, ip_address varchar(15), stats_table text);")
	_, err := stmt.Exec()
	return err
}

type SqlLiteOptions struct {
	databaseName string
}

func NewSqlLiteDefaultOptions() SqlLiteOptions {
	return SqlLiteOptions{
		databaseName: "devices",
	}
}

// NewClient will initialize the Database and return the client
func NewSqlLiteClient(options SqlLiteOptions) (ClientDevices, error) {
	sldl := SqlLiteDevices{}

	db, err := sql.Open("sqlite3", options.databaseName)
	if err != nil {
		return nil, err
	}

	//db.SetMaxIdleConns(50)

	err = db.Ping() // make sure the database conn is alive
	if err != nil {
		return nil, err
	}

	err = initDB(db)
	if err != nil {
		return nil, err
	}

	sldl.db = db

	log.WithFields(log.Fields{"error": err, "db": db}).Debug("New SqlLiteDevicesLib db conn")
	return &sldl, err
}

// AddDevice - Insert into table
func (s *SqlLiteDevices) AddDevice(newDevice Device) error {
	stmt, err := s.db.Prepare("INSERT INTO client_devices( " +
		" name, platform, mac_address, ip_address " +
		" ) values(?,?,?,?)")

	if err != nil {
		return err
	}
	_, err2 := stmt.Exec(newDevice.Name, newDevice.Platform, newDevice.Mac, newDevice.Ip)
	if err2 != nil {
		return err
	}

	return nil
}

// UpdateDevice - update information on a specific device
func (s *SqlLiteDevices) UpdateDevice(device Device) error {

	stmt, err := s.db.Prepare("UPDATE client_devices SET  " +
		" mac_address = ?,  ip_address = ? " +
		" WHERE name = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(device.Mac, device.Ip, device.Name)
	if err != nil {
		return err
	}
	return nil
}

// FindDeviceByName - find a device by its nickname
func (s *SqlLiteDevices) FindDeviceByName(device_name string) (Device, error) {
	rows, err := s.db.Query("SELECT * from client_devices WHERE name = ?", device_name)
	defer rows.Close()

	if err != nil {
		return Device{}, err
	}

	device := Device{}

	if rows.Next() {
		rows.Scan(&device.Name, &device.Platform, &device.Mac, &device.Ip)
	}
	log.WithFields(log.Fields{"err": err, "device": device}).Debug("FindDeviceByName")

	return device, nil
}

// RemoveDeviceByName - remove device given its nickname
func (s *SqlLiteDevices) RemoveDeviceByName(device_name string) error {
	stmt, err := s.db.Prepare("DELETE FROM client_devices WHERE name = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(device_name)
	if err != nil {
		return err
	}
	return nil
}

// GetAllDevices - get all registered devices
func (s *SqlLiteDevices) GetAllDevices() (Devices, error) {
	devices := Devices{}

	rows, err := s.db.Query("SELECT name, platform, mac_address, ip_address from client_devices")
	defer rows.Close()

	if err != nil {
		return devices, err
	}

	for rows.Next() {
		device := Device{}

		rows.Scan(&device.Name, &device.Platform, &device.Mac, &device.Ip)
		devices = append(devices, device)
	}
	log.WithFields(log.Fields{"err": err, "devices": devices}).Debug("GetAllDevices")

	return devices, nil
}
