package socks5

import (
	"fmt"
	"github.com/MarconiFoundation/dvpn/core/ip"
	mlog "github.com/MarconiProtocol/log"
	"github.com/armon/go-socks5"
	"sync"
)

type SockServer struct {
	Server *socks5.Server
	ip     string
	port   string
}

func Initialize(port string) (*SockServer, error) {
	sockServer := SockServer{}

	mainIpAddr, err := mnet_ip.GetMainInterfaceIpAddress()
	sockServer.ip = mainIpAddr

	sockServer.port = port

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		return nil, err
	}
	sockServer.Server = server
	return &sockServer, nil
}

func (s *SockServer) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	mlog.GetLogger().Info(fmt.Sprintf("Starting socks5 server on %s:%s", s.ip, s.port))

	err := s.Server.ListenAndServe("tcp", s.ip+":"+s.port)
	if err != nil {
		mlog.GetLogger().Error(fmt.Sprintf("Socks Server failed! %s", err.Error()))
	}
}
