# Candy Fight API

---

The application is written in `go` using [Fiber]() and [mongo-driver]()

## Run the application

### From sources
To build the API Server from source, `Golang >= 1.16` required. 
```bash
git clone https://github.com/JDR-ynovant/api.git

go mod download
go build -o candy-fight main.go
chmod +x candy-fight

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