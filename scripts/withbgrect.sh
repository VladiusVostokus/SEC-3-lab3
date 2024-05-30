#!/bin/bash

curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "bgrect 100 100 200 200"
curl -X POST http://localhost:17000 -d "figure 200 200"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "update"

