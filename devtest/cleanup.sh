#!/bin/bash

# stop all containers
bash stop.sh

# clean up docker images
sudo docker system prune
sudo docker image rm dvpn_peer dvpn_bootnode