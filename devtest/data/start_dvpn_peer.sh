#!/bin/bash

# add n0 hostname to hosts file
echo "172.30.20.10 n0.p1.dev-docker.prd.vm.tc" >> /etc/hosts

/opt/marconi/bin/dvpn --region=${1}