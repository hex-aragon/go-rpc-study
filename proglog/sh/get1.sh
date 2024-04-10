#!/bin/bash

curl -X GET localhost:8080 -d '{"offset": 0}'
curl -X GET localhost:8080 -d '{"offset": 1}'
curl -X GET localhost:8080 -d '{"offset": 2}'
curl -X GET localhost:8080 -d '{"offset": 3}'
curl -X GET localhost:8080 -d '{"offset": 4}'
curl -X GET localhost:8080 -d '{"offset": 5}'
curl -X GET localhost:8080 -d '{"offset": 6}'
curl -X GET localhost:8080 -d '{"offset": 7}'
curl -X GET localhost:8080 -d '{"offset": 8}'
curl -X GET localhost:8080 -d '{"offset": 9}'
