#!/bin/bash


# Define the server URL
URL="http://localhost:17000"

# Place the figure at (200, 200)
curl -X POST $URL -d "figure 200 200"

# Define the step size for each move
step_size= 100

# Move the figure to the right 5 times with a pause of 1 second between each move
for i in {1..5}
do
  move_distance=$((i * step_size))
  curl -X POST $URL -d "move $move_distance $move_distance"
  curl -X POST $URL -d "update"
  sleep 1
done

# Update the state to reflect all the moves
