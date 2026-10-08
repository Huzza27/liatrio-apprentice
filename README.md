# Go API

A simple REST API built with Go and the Fiber framework.

## Technologies

- Go
- Fiber
- Docker
- GitHub Actions (To be implemented)

## Running Locally

Make sure Go is installed.

```bash
go mod download
go run .
```

The API will be available at `http://localhost:3000`.

## Running with Docker

Build the Docker image:

```bash
docker build -t liatrio-api .
```

Run the container:

```bash
docker run --rm -p 3000:3000 liatrio-api
```

Visit `http://localhost:3000` to see the API response.

## API Endpoints

### GET /

Returns a JSON response containing a message and the current Unix timestamp.

Example response:

```json
{
  "message": "My name is Sam",
  "timestamp": 1791417600
}
```
