# gobotics
Golang robotics apis and libraries<br>

Gobotics is middleware that provides policy interfaces to robotic devices.<br><br>

Gobotics provides DB services, secure communication layers, and policy administration.<br> 
Providing clean interfaces with command and control of robotic communities and swarms through<br>
easy to implement APIs.<br><br>

Gobotics can be used as director level software services to control bots from shop floor, through<br>
the supply chain, all the way to customer delivery and billing.<br><br>

Get it:<br>
git clone https://github.com/rgdaddio/gobotics<br><br>

This has been tested on Ubuntu 14.04 installs and runs out-of-the-box just type make<br><br>

The server will run mostly on ssl. To generate dev certs use:
$ openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 3000 -nodes

#TODO
* better go coding habbits
* logging
* Front end <- -> backend 
* jwt cookie
