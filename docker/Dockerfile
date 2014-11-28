FROM ubuntu
RUN apt-get update
RUN apt-get install -y libreadline-dev libconfig-dev libssl-dev lua5.2 liblua5.2-dev libevent-dev
RUN apt-get install -y lua-lgi-dev lua5.2-socket
RUN apt-get install -y git
RUN apt-get install -y build-essential
RUN apt-get install -y golang
RUN git clone --recursive https://github.com/vysheng/tg.git
WORKDIR /tg
RUN ./configure
RUN make
VOLUME /app
CMD bash
