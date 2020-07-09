# GorlShort
![Go](https://github.com/leo3838ok/GorlShort/workflows/Go/badge.svg?branch=master)

`GorlShort`is a URL shortener written with Golang、MySQL and Redis.

It is used to store and get the ranking of specific content or condition.

## API
  - [POST] `/v1/create` It accepts POST form data with a parameter **url** and returns json response with short URL.
  - [GET] `/{code}` If the code is valid and found in the database, request will be redirected to original URL.

## Development
Deploy on localhost
```shell script
$ make dev
```
