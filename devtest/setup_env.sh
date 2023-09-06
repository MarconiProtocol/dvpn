set -e

# build the docker image for bootnode and peer
echo -e "[ Building the docker images for bootnode and peer ]\n"
make
sudo make -j docker
