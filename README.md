# weather-station-server

## Compile and Run with Docker

- `$ docker build -t ws-server .`
- `$ docker run --publish 6060:8080 --name ws-server --rm ws-server`
- open [localhost:6060/monkeys](localhost:6060/monkeys) in browser
