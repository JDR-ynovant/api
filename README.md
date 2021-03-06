# Candy Fight API

[![candy-fight api - ci](https://github.com/JDR-ynovant/api/actions/workflows/ci.yml/badge.svg)](https://github.com/JDR-ynovant/api/actions/workflows/ci.yml)
[![candy-fight api - cd](https://github.com/JDR-ynovant/api/actions/workflows/cd.yml/badge.svg)](https://github.com/JDR-ynovant/api/actions/workflows/cd.yml)

---

The application is written in `go` using [fiber](https://github.com/gofiber/fiber) and [mongo-driver](https://github.com/mongodb/mongo-go-driver)

## Development
For convenience, we use `air` to enable hot reload of the app. Install air, then just run it.
```bash
go get -u github.com/cosmtrek/air
air
```

## Run the application

### From sources
To build the API Server from source, `Golang >= 1.16` required. 
```bash
git clone https://github.com/JDR-ynovant/api.git

make build

candy-fight serve
```

### From Docker
You have to be authenticated to GitHub Container Registry ([ghcr.io](https://ghcr.io)) before
```bash
docker login ghcr.io -u <username> -p <password>

docker run -it -p 3000:3000 ghcr.io/JDR-ynovant/candy-fight-api:latest
```

## Configuration

// TODO

## API Documentation
Once the server is started, you will have access to the API Documentation at [`http://localhost:3000/swagger/`](http://localhost:3000/swagger/).
The doc is automatically generated from code annotations thanks to [swaggo/swag](https://github.com/swaggo/swag) and [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)