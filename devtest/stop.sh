#!/bin/bash

# stop all containers
sudo docker stop $(sudo docker ps -aq)