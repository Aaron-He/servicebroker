# Service Broker POC ![Build Status](https://travis-ci.org/cncf/servicebroker.svg?branch=master)

This repo has the PoC code the CNCF Service Broker WG.

All PRs must be signed with a DCO.

## Building

To build everything just run: `make` and that should leave you with a
`service_controller` executable in the `k8s/service_controller/` directory
along with a Docker image called`service_controller`.

## Running

`docker run -ti service_controller` should bring up a Service Controller
listengin on port 10000.

## Testing

TBD
