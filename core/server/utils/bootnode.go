package utils

import (
	"github.com/MarconiFoundation/dvpn/core/configs"
	httpServer "github.com/MarconiFoundation/dvpn/core/server/http"
	mnet_ip "github.com/MarconiFoundation/dvpn/core/ip"
	mlog "github.com/MarconiProtocol/log"
	"strings"
	"sync"
)

/*
	Wrapper around isIpAddressInSet
	Compares the main interface ip address to the bootnodes in the config
	This is used to determine if this server is in a fact a bootnode
*/
func IsBootNode() bool {
	ipAddress, err := mnet_ip.GetMainInterfaceIpAddress()

	if err != nil {
		mlog.GetLogger().Error("Error retrieving main interface IP Address", err)
		return false
	}

	bootnode := configs.GetAppConfig().BootNode.Address
	return bootnode == ipAddress
}

/*
	Checks if an ip address is present in a 'set'
	a set is a string of addresses in the form <address>:<port> and are separated by commas
	(this is how they are stored in the config currently.)
*/
func isIpAddressInSet(ipAddress string, addresses string) bool {

	//Bootnodes are separated by comas and in the form <address>:<port>
	bootnodeAddresses := strings.Split(addresses, ",")

	for _, bootnodeAddress := range bootnodeAddresses {
		address := strings.Split(bootnodeAddress, ":")[0]
		if address == ipAddress {
			return true
		}
	}
	return false
}

func BootNodeSetup(httpPort string, wg *sync.WaitGroup) {
	// Initialize Http Server with some port
	httpServer, err := httpServer.Initialize(httpPort)
	if err != nil {
		mlog.GetLogger().Error("Error occurred while trying to create HTTP Server", err)
	} else {
		go httpServer.Start(wg)
	}
}
