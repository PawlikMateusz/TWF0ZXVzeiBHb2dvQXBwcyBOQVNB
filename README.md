# url-collector

This repository is part of the recruitment task. The task was to create a micro-service for an awesome space visualization tool. url-collector is responsible for collecting URLs from external NASA API.

## Features

* context cancellation
* limit max concurrent request
* request params validation
* docker multi-stage build for image size optimization
* configuration via env variables


## Quick Start Guide

* Prerequisites
    * Installed Docker on your machine

1. Build a docker image

```sh
docker build -t url-collector . 
```

2. Run the application

```sh
docker run -p 8080:8080 \
    --env PORT=8080 \
    --env API_KEY=DEMO_KEY \
    --env CONCURRENT_REQUESTS=5 \
    url-collector
```

3. Send request

```sh
# using curl 
curl "0.0.0.0:8080/pictures?start_date=2021-03-09&end_date=2021-03-11"
# using httpie (better response formatting)
# install httpie `sudo apt install httpie`
http GET "0.0.0.0:8080/pictures?start_date=2021-03-09&end_date=2021-03-11"
```

## Running unit tests

```sh
go test ./... 
```

## To do

* logging
* tests

## Place for discussion

* Change to another image provider, just implement ImageProvider interface and switch to a new provider
* We can easily reuse code between microservices, for example pkg/client/PooledHTTPClient 