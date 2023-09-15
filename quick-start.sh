#!/bin/bash

docker-compose -f ./cert-app/database/docker-compose-postgres.yaml up &

./net.sh up createChannel -ca -c certapp -s couchdb

# Path is relative to test-network/network.sh
./net.sh deployCC -c certapp -ccn certificate-manager -ccp ../cert-app/chaincodes/certificate-manager -ccv 1 -ccl go