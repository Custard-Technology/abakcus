## Abakcus: Small Business Tools

### Development

The backend lives in the `backend` directory and is a Go module.
To run the API locally you need a MongoDB instance and the following
environment variables (see `backend/.env` for an example):

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB=abakcus
```

Start the service with:

```sh
cd backend
go run ./cmd/api
```

Connection attempts are logged on startup and the process will exit if the
MongoDB handshake fails.

