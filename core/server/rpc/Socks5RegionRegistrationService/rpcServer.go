package Socks5RegionRegistrationService

import (
	"fmt"
	mlog "github.com/MarconiFoundation/log"
	"github.com/MarconiFoundation/dvpn/core/ip"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

type RPCServer struct {
	port string
	ip   string
}

func Initialize(port string) (*RPCServer, error) {
	rpcServer := RPCServer{}
	rpcServer.port = port
	mainIpAddr, err := mnet_ip.GetMainInterfaceIpAddress()
	rpcServer.ip = mainIpAddr

	if err != nil {
		return nil, err
	}
	return &rpcServer, nil
}

func (r *RPCServer) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	mlog.GetLogger().Info(fmt.Sprintf("Starting RPC service on %s:%s", r.ip, r.port))

	serviceRegistration := new(ServerRegistration)
	rpc.Register(serviceRegistration)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", r.ip+":"+r.port)

	if e != nil {
		mlog.GetLogger().Fatal("Failed to start rpc service", e)
	}

	go http.Serve(l, nil)
}
