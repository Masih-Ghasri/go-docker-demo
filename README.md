# go-docker-demo

A tiny demo HTTP server written in Go — intended as a minimal example for learning how to Dockerize a Go app. The server exposes two simple endpoints and listens on port `8070`. The main source is `main.go`.

## What this project does

* Starts an HTTP server on port 8070.
* `GET /` responds with `Hello World "<path>"`.
* `GET /hi` responds with `Hello "<path>"`.
* Minimal, dependency-light Go app intended for experimenting with Dockerizing and basic HTTP routing.

## Files

* `main.go` — the Go program that runs the HTTP server.
* `go.mod` — module file (not shown here).
* `Dockerfile` — Docker container description (you uploaded one).
* `README.md` — this file (replaced/expanded).

## Requirements

* Go (1.18+ recommended) installed locally, for running or building without Docker.
* Docker (for building and running the container).

## Run locally (no Docker)

From project root:

```bash
# download deps (if needed)
go mod tidy

# run directly
go run main.go
```

Then open `http://localhost:8070/` and `http://localhost:8070/hi` in your browser or use curl:

```bash
curl http://localhost:8070/
curl http://localhost:8070/hi
```

Behavior: the server echoes the requested URL path in the greeting; it listens on `:8070`.

## Build and run with Go (binary)

```bash
go build -o go-docker-demo ./main.go
./go-docker-demo
# or (portable)
PORT=8070 ./go-docker-demo
```

## Docker — build & run

Assuming your `Dockerfile` is in the project root, a common pattern:

```bash
# build image (replace tag as you like)
docker build -t go-docker-demo:latest .

# run container, publish port 8070
docker run --rm -p 8070:8070 go-docker-demo:latest
```

Then visit `http://localhost:8070/` and `http://localhost:8070/hi`.

If you prefer a multi-stage, small image, your Dockerfile might already use `scratch` or `alpine` + `go build` — the commands above still apply.

## Example requests & responses

```bash
$ curl http://localhost:8070/
Hello World "/"

$ curl http://localhost:8070/hi
Hello "/hi"
```

(The server escapes and prints the request path before responding.)

## Notes & suggestions

* This is intentionally minimal — great for practicing Docker builds, port mapping, and simple health checks.
* Add graceful shutdown handling (context + signal) if you plan to run this in production.
* Add a `/health` endpoint if you want Kubernetes/Docker health probes.
* Consider adding unit tests around handlers and a Makefile for common tasks.

## License

Pick a license you prefer (MIT is common for demos). Add a `LICENSE` file if you want to share this publicly.

If you want, I can:
* produce a polished `Dockerfile` (small multi-stage) and include it here, or
* add a `Docker Compose` example, or
* add graceful shutdown and a `/health` endpoint to `main.go`.

Tell me which of the above you'd like and I'll add it directly.
