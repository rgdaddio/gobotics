package main

import (
    "time"
    "gobot.io/x/gobot"
    "gobot.io/x/gobot/drivers/gpio"
    "gobot.io/x/gobot/platforms/firmata"
)



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

  start_hardware_interface()
  println("Welcome to gobotics")
  send_req("http://localhost:8080/list")
  send_req("http://localhost:8080/die")
}
