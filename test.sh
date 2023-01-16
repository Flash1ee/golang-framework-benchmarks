#!/bin/bash
repeats=3
rm -rf out-ab
mkdir out-ab
for (( i = 0; i < repeats; i++ )); do
  echo $((i+1))
  docker-compose -f docker-compose.yml up -d
  sleep 10
  make run-ab-get num=$((i+1)) dir=./out-ab
  make run-ab-post num=$((i+1)) dir=./out-ab
  docker-compose down
done
