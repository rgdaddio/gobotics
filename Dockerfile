# Set the base image to use to Ubuntu
#FROM docker/whalesay:latest
#RUN apt-get -y update && apt-get install -y fortunes
#CMD /usr/games/fortune -a | cowsay

FROM ubuntu
RUN apt-get update && apt-get install -y  wget git make gcc
ADD . /usr/local/bin/


# SHOULD BE DELETED EVENTUALLY right now im using so the docker contrainer doesnt die
RUN echo 'ping localhost &' > /bootstrap.sh
RUN echo 'sleep infinity' >> /bootstrap.sh
RUN chmod +x /bootstrap.sh

CMD /bootstrap.sh
