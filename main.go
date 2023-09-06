package main

import (
	"flag"
	"github.com/MarconiFoundation/dvpn/core/configs"
	httpServer "github.com/MarconiFoundation/dvpn/core/server/http"
	"github.com/MarconiFoundation/dvpn/core/server/rpc/Socks5RegionRegistrationService"
	"github.com/MarconiFoundation/dvpn/core/server/rpc/Socks5RegionRegistrationService/RegionAnnouncement"
	"github.com/MarconiFoundation/dvpn/core/server/socks5"
	"github.com/MarconiFoundation/dvpn/core/server/utils"
	mlog "github.com/MarconiProtocol/log"
	"os"
	"sync"
)

func main() {

	// Setting Command Line Argument options
	baseDir := flag.String("basedir", "/opt/marconi", "Base of directory tree that will "+
		"be used to store all configs and data files related to dvpn")
	location := flag.String("region", "US:CA", "Indicates the location/region of the server in the format <Country>:<State>")
	socks5Port := flag.String("socks5-port", "5657", "Port to run the socks5 server on")
	httpPort := flag.String("http-port", "8080", "Port to run the http server on")
	flag.Parse()

	// Setting Configurations
	configs.SetBaseDir(*baseDir)
	configs.InitializeConfigs(*baseDir)
	mlog.Init(configs.GetFullPath(configs.GetAppConfig().Log.Dir), configs.GetAppConfig().Log.Level)

	// Setting region
	configs.SetRegion(*location)

	rpcPort := configs.GetAppConfig().BootNode.RPC_PORT

	// Handling port conflicts
	if *socks5Port == *httpPort {
		mlog.GetLogger().Error("Cannot run Http and Socks5 Servers on the same port")
		os.Exit(1)
	}

	if rpcPort == *httpPort {
		mlog.GetLogger().Error("Cannot run Http and RPC Servers on the same port")
		os.Exit(1)
	}

	if *socks5Port == rpcPort {
		mlog.GetLogger().Error("Cannot run RPC and Socks5 Servers on the same port")
		os.Exit(1)
	}

	isBootNode := utils.IsBootNode()
	runningSocks5Server := false

	// Booting up Servers
	wg := &sync.WaitGroup{}

	// Initialize Socks5 server if we are a PeerNode or a BootNode and the Configuration states so
	if !isBootNode || configs.GetAppConfig().BootNode.RUN_SOCKS_5 {
		wg.Add(1)

		// Initialize Socks5 Server with some port
		socksServer, err := socks5.Initialize(*socks5Port)
		if err != nil {
			mlog.GetLogger().Error("Error occurred while trying to create Socks5 Server", err)
		} else {
			go socksServer.Start(wg)
			runningSocks5Server = true
		}
	}

	// Run some additional setup if the current node is a bootnode
	if isBootNode {

		wg.Add(2)

		// Initializing and Launching HTTP Server
		httpServer, err := httpServer.Initialize(*httpPort)
		if err != nil {
			mlog.GetLogger().Error("Error occurred while trying to create HTTP Server", err)
		} else {
			go httpServer.Start(wg)
		}

		rpcServer, err := Socks5RegionRegistrationService.Initialize(rpcPort)

		// Initializing and Launching RPC Server
		if err != nil {
			mlog.GetLogger().Error("Error occurred while trying to create Socks5RegionRegistration RPC Server", err)

		} else {
			go rpcServer.Start(wg)
		}

	}

	// Start announcing our Socks5 Server if we are running one
	if runningSocks5Server {
		ra := RegionAnnouncement.Initialize(*socks5Port)
		ra.StartRegionAnnouncement()
	}

	wg.Wait()
}
