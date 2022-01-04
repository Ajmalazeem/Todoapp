FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . ./

RUN go build -o /file-k.exe 
COPY db.env ./
RUN ls

EXPOSE 8000

CMD [ "/file-k.exe" ]