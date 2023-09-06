package RegionAnnouncement

import (
	"fmt"
	"github.com/MarconiFoundation/dvpn/core/configs"
	"github.com/MarconiFoundation/dvpn/core/server/rpc/Socks5RegionRegistrationService"
	mnet_ip "github.com/MarconiFoundation/dvpn/core/ip"
	mlog "github.com/MarconiProtocol/log"
	"net/rpc"
	"time"
)

/*
	This interface is responsible for announcing our server's ip address and port that the socks5 server is running on
	Additionally, we also broadcast the region that the machine is located in. This gets sent to the BootNode, which then
	keeps track of all regions and the machines in them
*/

type RegionAnnouncement struct {
	region                   string
	socks5ServerAddress      string
	socks5ServerPort         string
	interval                 int
	client                   *rpc.Client
	announcementStarted      bool
	regionAnnouncementSignal *chan bool
}

func Initialize(socks5Port string) *RegionAnnouncement {
	ra := &RegionAnnouncement{}

	// Initializing values used in the announcement
	mainIpAddr, err := mnet_ip.GetMainInterfaceIpAddress()
	ra.socks5ServerAddress = mainIpAddr
	ra.socks5ServerPort = socks5Port

	// Initializing announcement interval (how often we announce)
	ra.interval = configs.GetAppConfig().Announcement.Region_Announcement_Interval_Seconds

	signal := make(chan bool)
	ra.regionAnnouncementSignal = &signal
	ra.announcementStarted = false

	ra.region = configs.GetMyRegion()

	// Initialize rpc client to communicate with the rpc server running on the BootNode
	client, err := rpc.DialHTTP("tcp", configs.GetAppConfig().BootNode.Address+":"+configs.GetAppConfig().BootNode.RPC_PORT)
	if err != nil {
		mlog.GetLogger().Fatal("dialing:", err)
	}

	ra.client = client

	return ra
}

func (ra *RegionAnnouncement) AnnounceRegion() {

	// Makes the call to the Socks5RegionRegistrationService
	mlog.GetLogger().Debug("Announcing")
	// Args: ip, port, region
	args := &Socks5RegionRegistrationService.Args{ra.socks5ServerAddress, ra.socks5ServerPort, ra.region}
	response := &Socks5RegionRegistrationService.Response{}

	ra.client.Call("ServerRegistration.RegisterSocksServer", args, &response)

}

func (ra *RegionAnnouncement) StartRegionAnnouncement() {
	// Make calls on an interval
	if ra.announcementStarted {
		mlog.GetLogger().Debug("RegionRouteAnnouncement already started, no-op")
		return
	}

	ra.announcementStarted = true
	mlog.GetLogger().Info(fmt.Sprintf("Starting Region Announcement - announcing %s:%s for region %s", ra.socks5ServerAddress, ra.socks5ServerPort, ra.region))

	for {
		select {
		case <-*ra.regionAnnouncementSignal:
			return
		default:
			// Announcement
			ra.AnnounceRegion()
			// Sleeping
			time.Sleep(time.Duration(ra.interval) * time.Second)
		}
	}
}

func (ra *RegionAnnouncement) StopRegionAnnouncement() {
	if !ra.announcementStarted {
		mlog.GetLogger().Debug("RegionAnnouncement is not started, no-op")
		return
	}
	ra.announcementStarted = false
	*ra.regionAnnouncementSignal <- true
}
