# gobotics

Imagine a robotics solution that provides:
*  robotics apis and libraries middleware
*  policy interfaces to robotic devices
*  DB services, secure communication layers, and policy administration.
*  clean interfaces with command and control of robotic communities and swarms through
*  access and control via commandline tool & webportal 
*  can be used as director level software services to control bots from shop floor, through the supply chain, all the way to customer delivery and billing.

Gobotics is a skunkworkds proof of concept project to help us learn more about golang/web development. It tries to focus on the IoT/device space.

Currently it is composed of the following:
* [`gobotics_server`](https://github.com/rgdaddio/gobotics/tree/master/server): REST API server (golang) - A simple CRUD REST API to add,get,update, & delete information on devices. 
* `gobotics_client`: a commandline client to the `gobotics_server` (golang)
* `gobotics-frontend`: React Frontend web application

Pending:
* Frontend WebApp: TODO React/Redux web application
* actual robotics apis and libraries? (unsure what this will actually be.)


# TODO

* client library

* ssl
* jwt cookie

* better go coding habbits
* logging
* Front end <- -> backend 

* Swagger
* Protobuff
* Cassandra Client
* Have AddDevice return a uuid, that people can use to query later! ( will help with Cassandra etc)
Natural key? Surrogate key - UUID?
Partition key, clustering column?
Get all devices could be problematically large ( paginate, skip if using casandra?)

The server will run mostly on ssl. To generate dev certs use:
$ openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 3000 -nodes

# License
MIT LICENSE - enjoy!
