# Binquiz Server App

A go based application to play Binquiz. The game itself is multiplayer and will be moderated by a quizmaster and some players in the teams. Stay tuned in  for the first event.

---
## Table of contents
- Technology
  - External Comms
  - Package structure
  - Installation
  - Launch application
  - Test application
  - Environment Variables
  - Building Image
  - Dependencies
- Author
---
## Technology
### External Comms
- MongoDB client for storing entitites and relationships.
- Web sockets networking for seamless quizzing experience with every player getting an update at the same time.
- Jwt Token is used for authenticating any REST api requests. A new token is sent out to a user when they login.
---

### Package structure
```
core
├───actions
├───comms
├───db
├───errors
├───models
├───services
├───tests
└───utils
```
---

### Installation

- The instructions for installation are the following. Please run them inorder.
```go
go get -d -v ./...
go mod download
go mod vendor
go mod verify
```
---

### Launch application
Start the service as a go application by running 
```go
go run main.go
```
---

### Test application
Test the go application with coverage by running 
```go
go test -v -cover ./...
```
---

### Environment Variables
- `MONGO_CLIENT_ID` = <MONGO_CLIENT_ID>
- `PORT` = <SERVER_PORT>
- `GAME_DATABASE` = <DATABASE_NAME>
- `SSL_KEY` = <PATH_TO_KEY_FILE>
- `SSL_CERT` = <PATH_TO_CERT_FILE>
---
### Building Image
The images reside in [DockerHub](https://hub.docker.com/repository/docker/atulanand206/binquiz-server). To build a docker image, run the following command in the repository's root directory.
```
docker build . -t atulanand206/binquiz-server:vx.x
```

Once the image is built, push it out to DockerHub so that it can be consumed by the production server.
```
docker push atulanand206/binquiz-server:vx.x
```

---
### Dependencies   

- ##### Core
  - [godotenv](github.com/joho/godotenv)
  - [go-nanoid](github.com/matoous/go-nanoid/v2)
  - [pointer](github.com/xorcare/pointer)
  - [uuid](github.com/google/uuid)

- ##### Comms
  - [go-network](github.com/atulanand206/go-network)
  - [jwt-go](github.com/dgrijalva/jwt-go/v4)
  - [websocket](github.com/gorilla/websocket)

- ##### Data mgmt
  - [go-mongo](github.com/atulanand206/go-mongo)

- ##### Testing 
  - [testify](github.com/stretchr/testify)

---

## Author
- [Atul Anand](https://github.com/atulanand206)