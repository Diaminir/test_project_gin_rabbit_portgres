FROM golang:1.25

WORKDIR /data_reception

COPY go.mod go.sum ./
RUN go mod download

COPY . .


CMD [ "go", "run", "main.go" ]