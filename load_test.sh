#!/bin/bash

# number of requests
num_requests=100

# url to test
url="http://localhost:1729/"

for i in $(seq 1 $num_requests); do
  curl -s -o /dev/null $url &
done

# wait for all background processes to finish
wait

echo "Completed $num_requests requests."
