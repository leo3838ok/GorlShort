# GorlShort
![Go](https://github.com/leozhantw/GorlShort/workflows/Go/badge.svg?branch=master)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/leozhantw/GorlShort/blob/master/LICENSE)

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
