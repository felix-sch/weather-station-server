# weather-station-server

## Compile and Run with Docker

- `$ git clone --recursive https://github.com/felix-sch/weather-station-server.git`
- setup submodule https://github.com/felix-sch/weather-station-app 
- `$ docker build -t ws-server .`
- `$ docker run --publish 6060:8080 --name ws-server --rm ws-server`
- open [localhost:6060/monkeys](localhost:6060/monkeys) in browser
