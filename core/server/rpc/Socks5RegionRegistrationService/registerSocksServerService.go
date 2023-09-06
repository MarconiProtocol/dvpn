package Socks5RegionRegistrationService

import (
	"github.com/MarconiFoundation/dvpn/core/region"
	mlog "github.com/MarconiProtocol/log"
)

/*
	This is the ServerRegistration service, it allows machines to register the presence of a Socks5 server
	as well as the region it is located in. PeerNodes call this service running on a BootNode
*/
type Args struct {
	Ip     string
	Port   string
	Region string
}

type Response struct {
	Success bool
}

type ServerRegistration bool

func (s *ServerRegistration) RegisterSocksServer(args *Args, response *Response) error {

	mlog.GetLogger().Debug("Received RPC request to register ", args.Ip, ":", args.Port, " for region: ", args.Region)
	m := make(map[string]string)
	m["peerIp"] = args.Ip + ":" + args.Port
	m["region"] = args.Region
	region.RequestRegionResponseHandler(m)
	response.Success = true
	return nil
}
