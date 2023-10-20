FROM golang:1.21.1-alpine3.18

# Gelistirme Ortami Self Issued Certificate
COPY certs/minica.pem /usr/local/share/ca-certificates/minica.crt
RUN update-ca-certificates

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/golang-gin-boilerplate ./cmd

EXPOSE 8989

CMD [ "/bin/golang-gin-boilerplate" ]
