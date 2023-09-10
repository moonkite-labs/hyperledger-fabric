#!/bin/bash

./net.sh down
docker-compose -f ./cert-app/database/docker-compose-postgres.yaml down --volumes