#!/bin/bash

BASE_URL="http://localhost:8080/api/v1/parties"

kill_process() {
  local port=$1
  local pid=$(lsof -ti:$port)
  echo "Killing any process on port $port..."

  if [ -n "$pid" ]; then
    kill -9 $pid
    echo
  fi
}

kill_process 8080

echo "Starting Docker..."
docker compose up -d

sleep 3  # wait for deps to start up

echo "Running Go-Party-Service..."
go run main.go & # run it in the background
GO_PID=$! # grab process ID to kill it later

sleep 3  # wait for server to start up

echo "Testing GET /api/v1/parties/{id} endpoint"
echo -n "Enter Party ID: "
read partyId

curl -X GET "${BASE_URL}/${partyId}" | jq .

# cleanup
echo -e "\nStopping Go server"
kill $GO_PID
kill_process 8080
docker compose down