FROM golang:1.16-alpine

RUN mkdir todoapp

WORKDIR /todoapp

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# RUN go get -d -v ./...
# RUN go install -v ./...
RUN ls
# RUN  go mod download
RUN go build -o /ajmalaseem
RUN ls
CMD ["/ajmalaseem"]