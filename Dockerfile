FROM golang:1.18-bullseye

WORKDIR /wotoplatform

COPY . /wotoplatform/
RUN go mod download

RUN go build -o wp-server

EXPOSE 443

CMD [ "./wp-server" ]