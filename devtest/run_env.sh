#!/bin/bash

if [[ $EUID -ne 0 ]]
then
    echo This should be run as root.
    exit 1
fi

RANDOM=$$
first_ip="172.30.20.10"
num_nodes=3
network_name="marconi_dvpn_net"
expose_port="24800-65535"
domain_name="p1.dev-docker.prd.vm.tc"

regions=("US:CA" "US:OH" "CAD:ON")

# get the os version
OS_VERS=$(awk -F= '/^ID/&&!/^ID_LIKE/{ OS=$2 } /^VERSION_ID/{ print OS $2 }' /etc/os-release | \
  sed 's/\"//g')

# assume that this is always going to be a class C network
IFS=. read ip1 ip2 ip3 ip4 <<< "$first_ip"
subnet_octets=$ip1.$ip2.$ip3
offset_octet=$ip4

create_network() {
  # Create network for local test cluster
  docker network rm $network_name
  docker network create --subnet=$subnet_octets.0/24 $network_name
}

run_node() {
    node_name=n$1
    last_octet="$(( $offset_octet + $1 ))"
    node_type="$2"
    len=${#regions[@]}
    num=$(($RANDOM%len))

    region=${regions[$num]}

    if [ $OS_VERS == "centos7" ]; then
        TERMINAL_EMULATOR=xterm
    else
        TERMINAL_EMULATOR=xterm
    fi

    docker stop -t 0 $node_name
    docker rm $node_name
    CMD="docker run \
        -it \
        --privileged \
        --net $network_name \
        --ip $subnet_octets.$last_octet \
        --name $node_name \
        --hostname $node_name.$domain_name \
        --expose=$expose_port/udp \
        --expose=$expose_port \
        dvpn_$node_type /bin/bash -c './start_dvpn_$node_type.sh $region &> /opt/marconi/var/log/marconi/startup.log & /bin/bash'"
    $TERMINAL_EMULATOR -e sh -c "$CMD" &>/dev/null &
}

run_dvpn_bootnode() {
    # start a terminal running bootnode
    run_node 0 bootnode
}

run_dvpn_peer() {
    # start terminals running peers
    run_node $1 peer
}

create_network
run_dvpn_bootnode
for i in $(seq 1 $num_nodes)
do
    run_dvpn_peer $i
done
