FROM golang:1.4.2

WORKDIR /usr/bin
ADD badhttp /usr/bin/badhttp

CMD ./badhttp -port=$PORT
