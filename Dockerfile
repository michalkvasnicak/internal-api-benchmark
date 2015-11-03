FROM ubuntu:14.04
RUN apt-get install software-properties-common -y
RUN add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) main universe"
RUN apt-get update && apt-get install curl build-essential git pkg-config -y
RUN curl -O https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
RUN curl -O -L https://github.com/nanomsg/nanomsg/releases/download/0.7-beta/nanomsg-0.7-beta.tar.gz
RUN curl -O http://download.zeromq.org/zeromq-4.1.3.tar.gz
RUN tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
RUN tar -xzf zeromq-4.1.3.tar.gz
RUN tar -xzf nanomsg-0.7-beta.tar.gz
RUN cd zeromq-4.1.3 && ./configure --without-libsodium && make && make install && ldconfig
RUN cd nanomsg-0.7-beta && ./configure && make && make install && ldconfig
#RUN mkdir -p ~/go; echo "export GOPATH=$HOME/go" >> ~/.bashrc
#RUN echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc
RUN mkdir -p ~/go
ENV GOPATH $HOME/go
ENV PATH $PATH:$HOME/go/bin:/usr/local/go/bin
RUN go get github.com/michalkvasnicak/internal-api-benchmark