<div align="center">
    <h2 align="center" style="border-bottom: none">Drone Flight Logs</h2>
    <hr/>
</div>

[![CI](https://github.com/rpparas/flight_log/actions/workflows/ci.yml/badge.svg)](https://github.com/rpparas/flight_log/actions/workflows/ci.yml)


## Definitions
robot = drone
flight = ...
TODO

## System Requirements
### Basics
- git    for downloading this repo
- curl   for fetching resources
- jq     for parsing results

### If running baremetal:
- Go v1.18 (recommended), v1.13 (supported best-effort)
- Postgres loaded with data
- pgadmin4

### If running Docker container:
- Docker
- Docker Compose
- curl

## How to Deploy
### Via Baremetal
TODO

### Via Docker
TODO

## How to use REST API
Create a drone
```bash
curl --location --request POST 'http://localhost:8000/api/v1/robots' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "Alpha3"
}'
```

Record a flight using drone/robot id:
```
curl --location --request POST 'http://localhost:8000/api/v1/flights' \
--header 'Content-Type: application/json' \
--data-raw '{
    "robotId": "e570b6c0-9bb0-47c9-a358-b984ed402406",
    "startTime": "2022-08-16T00:00:00+00:00"
    
}
'
```

Retrieving flight logs
```
curl --location --request GET 'http://localhost:8000/api/v1/flights?generation=1&from=2022-07-01T00:00:00Z&to=2023-07-03T00:00:00Z&maxDurationMins=16'
```
<hr>

## Feature Status
- Resource Creation
  - [x] Create Robot (single via JSON)
  - [x] Create Flight (single via JSON)
  - [x] Create Flights (bulk via CSV)
- Resource Retrieval
  - [x] Get Robot (using ID)
  - [x] Get Flight (using ID)
  - [x] Get Flights (using generation, date range, duration, combo)
- Testing
  - [x] integration tests
  - [x] test automation (CI)
  - [ ] code coverage
- Data Validation
  - [x] URL params in GET requests
  - [ ] payload in POST requests
- Documentation
  - [ ] Docker builds
  - [ ] setting up DB
  - [ ] CLI: running bare-metal
  - [ ] CLI: running via Docker Compose
  - [ ] Web UI
- Security
  - [ ] API Auth Token
  - [ ] request rate limiting
  - [ ] restricting IP ranges 
- Scalability
  - [ ] streaming when loading CSV
  - [ ] limiting / paginating results
- Data Storage
  - [ ] optimized schema
  - [ ] caching results
- Other
  - [ ] Support for gRPC


## Known Limitations
TODO
