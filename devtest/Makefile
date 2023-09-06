OS_VER = $(shell awk -F= '/^ID/&&!/^ID_LIKE/{ OS=$$2 } /^VERSION_ID/{ print OS $$2 }' /etc/os-release | sed 's/\"//g')

.PHONY: all build pack pack_data build_dvpn docker phony_explicit 

all: pack 

pack: build pack_dvpn pack_data

build_dvpn:
		cd .. && ./build.sh

pack_data:
		rm -rf dvpn_bootnode/data
		mkdir -p dvpn_bootnode/data
		cp -R data/* dvpn_bootnode/data

		rm -rf dvpn_peer/data
		mkdir -p dvpn_peer/data
		cp -R data/* dvpn_peer/data

pack_dvpn: build_dvpn
		mkdir -p dvpn_bootnode/bin
		mkdir -p dvpn_peer/bin
		cp ../out/dvpn dvpn_peer/bin
		cp ../out/dvpn dvpn_bootnode/bin

docker: dvpn_peer dvpn_bootnode

dvpn_%: phony_explicit 
		docker build -t $@ -f $@/dockerfile_$(OS_VER) $@
