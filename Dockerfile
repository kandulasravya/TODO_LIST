FROM golang:latest
 
WORKDIR /app
 
COPY go.mod ./
#COPY go.sum ./
 
RUN apt update
RUN apt install wget -y
RUN ["apt-get", "-y", "install", "vim"]
RUN apt install build-essential -y
 
RUN go mod download
 
COPY . .
 
RUN go build -o main .
#CMD ["/todo-list"]
ENTRYPOINT ["./main"]
