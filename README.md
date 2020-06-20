# gobotics
Imagine a robotics solution that provides:
*  robotics apis and libraries middleware
*  policy interfaces to robotic devices
*  DB services, secure communication layers, and policy administration.
*  clean interfaces with command and control of robotic communities and swarms through
*  access and control via commandline tool
*  access and control via a webportal
*  can be used as director level software services to control bots from shop floor, through the supply chain, all the way to customer delivery and billing.

Gobotics is a skunkworkds proof of concept project to help us learn more about golang/web development while keeping the points above in mind.

Currently it is composed of the following:
* gobotics_server: REST API server (golang)
* gobotics_client: commandline client (golang)

Pending:
* Frontend WebApp: TODO React/Redux web application
* actual robotics apis and libraries? (unsure what this will actually be.)


#TODO
* go mod
* serve react app
* build compiled react app

* ssl
* jwt cookie

* better go coding habbits
* logging
* Front end <- -> backend

The server will run mostly on ssl. To generate dev certs use:
$ openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 3000 -nodes