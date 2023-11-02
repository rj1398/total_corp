FROM golang:1.21.1

WORKDIR /total_corp

COPY . .

Run go get -d -v ./...
Run go install -v ./...

Expose 8080
cmd ["total_corp"]