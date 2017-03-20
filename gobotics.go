package main

import (
    "fmt"
    "time"
    "gobot.io/x/gobot"
    "gobot.io/x/gobot/drivers/gpio"
    "gobot.io/x/gobot/platforms/firmata"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)


func load_sql_file(db *sql.DB){
  stmt, _ := db.Prepare("create table if not exists client_devices( " +
      " name text, platform text, mac_address text, ip_address varchar(15) );" )
  _, err := stmt.Exec()
  if err != nil { panic(err) }
}

func start_hardware_interface(){
 firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
 //println(firmataAdaptor)
 led := gpio.NewLedDriver(firmataAdaptor, "13")
 //led.Toggle()

 work := func() {
     gobot.Every(1*time.Second, func() {
                led.Toggle()
                })
        }

     robot := gobot.NewRobot("bot",
             []gobot.Connection{firmataAdaptor},
			[]gobot.Device{led},
                work,
		)

        robot.Start()

}

func main() {
  db, _ := sql.Open("sqlite3", "./foo.db")

  load_sql_file(db)
  start_hardware_interface()
  fmt.Printf("%s", db)
  println("Welcome to gobotics")
}
