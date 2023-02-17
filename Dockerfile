FROM golang:1.19.0-alpine

WORKDIR /app

COPY . /app

RUN go mod download
RUN cd cmd/webhook && go build -o /workflows-manager

EXPOSE 8080

CMD [ "/workflows-manager" ]