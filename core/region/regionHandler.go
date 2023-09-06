package region

import (
	mnet_ip "github.com/MarconiFoundation/dvpn/core/ip"
	"github.com/MarconiFoundation/dvpn/core/configs"
	"github.com/MarconiFoundation/dvpn/core/types/set"
	mlog "github.com/MarconiProtocol/log"
	"sync"
	"time"
	//"crypto/sha1"
	//"encoding/base64"
)

type RegionHandler struct {
	regionAddresses  *map[string]set.TTLSet
	regionAddressTTL time.Duration
}

var regionHandler *RegionHandler
var once sync.Once

func GetRegionHandler() *RegionHandler {
	once.Do(func() {
		regionHandler = &RegionHandler{}
		regionAddresses := make(map[string]set.TTLSet)
		regionHandler.regionAddresses = &regionAddresses
		regionHandler.regionAddressTTL = time.Duration(configs.GetAppConfig().BootNode.Server_Address_TTL_Minutes) * time.Minute
	})
	return regionHandler
}

func RequestRegionResponseHandler(args map[string]string) {
	// expect args to be in the form:
	ip := args["peerIp"]
	region := args["region"]

	if _, exists := (*GetRegionHandler().regionAddresses)[region]; !exists {
		// Create a new set
		(*GetRegionHandler().regionAddresses)[region] = *set.NewTTLSet(GetRegionHandler().regionAddressTTL)
	}

	// To avoid possible glitches caused by Boot Nodes broadcasting multiple regions
	mainIpAddress, err := mnet_ip.GetMainInterfaceIpAddress()
	if err != nil {
		mlog.GetLogger().Error("Error retrieving main interface IP Address")
		return
	}

	if ip == mainIpAddress {
		region = configs.GetMyRegion()
	}

	// Adding to the set
	set := (*GetRegionHandler().regionAddresses)[region]
	set.Add(ip)
}

func (rh *RegionHandler) GetAddressesForRegion(region string) []string {
	set := (*rh.regionAddresses)[region]
	mlog.GetLogger().Info("Received request to get regions for ", region)
	return set.GetValues()
}

/*
	This returns all regions that we have an address for
*/
func (rh *RegionHandler) GetAvailableRegions() []string {
	mlog.GetLogger().Info("Received request to get all available regions")
	availableRegions := make([]string, 0, len(*rh.regionAddresses))

	for region := range *rh.regionAddresses {

		set := (*rh.regionAddresses)[region]
		// Make sure we have at least one address present
		if set.Size() > 0 {
			availableRegions = append(availableRegions, region)
		}
	}

	return availableRegions
}

func (rh *RegionHandler) ClearSet(region string) {
	delete(*rh.regionAddresses, region)
}

// TODO: Do we want to even hash this?
func HashRegion(region string) (string, error) {
	// Return a 20 character hash
	/*regionBytes := []byte(region)
	hasher := sha1.New()
	hasher.Write(regionBytes)
	shaHash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	//return shaHash, nil*/
	return region, nil
}

