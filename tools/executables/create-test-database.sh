#!/bin/bash

cd src/utils/tests/docker
docker compose rm -sf
docker compose up --build -d