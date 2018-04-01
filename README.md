# Basic Redis Usage on Golang
Simple Prevention Multiple Request from Frontend with Redis

### Getting Started
**redissimulation** is one of simple example on how to utilize Redis technology into Golang webapp usage. This example utilizes redis to prevent multiple requests from frontend in certain period of time (in this case, 10 seconds).

### Pre-requisites
- Go 1.8.1 or Above (https://golang.org/dl/)
- Redis Server 3.0.6 or Above (https://redis.io/download)

### Installing
1. Make sure you have fulfilled all prerequisites above
2. Clone repo to your local computer
```
git clone https://github.com/wicoady1/redissimulation.git
```
3. Build and Run
(This code example for UNIX Command, Windows and other Linux Distro may use different commands)
```
go build && ./redissimulation
```
4. Go to browser and access link below
```
localhost:5555
```

### Authors
- Kennard Wicoady
