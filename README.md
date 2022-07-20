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
### If running baremetal:
- git
- curl
- Go v1.18 (recommended), v1.13 (supported best-effort)
- Postgres loaded with data

### If running Docker container:
- Docker
- Docker Compose
- curl

## How to Deploy
### Via Baremetal
TODO

### Via Docker
TODO

## How to Use REST API

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
