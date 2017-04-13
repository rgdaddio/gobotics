# Set the base image to use to Ubuntu
#FROM docker/whalesay:latest
#RUN apt-get -y update && apt-get install -y fortunes
#CMD /usr/games/fortune -a | cowsay

FROM ubuntu
RUN apt-get update && apt-get install -y  wget git make gcc
ADD . /usr/local/bin/
WORKDIR "/usr/local/bin/"
RUN make
CMD ./gobotics_server 
