package api

import (
	"encoding/json"
	"github.com/MarconiFoundation/dvpn/core/region"
	"github.com/gorilla/mux"
	"net/http"
)

/*
	Returns the addresses for some region
*/
func RegionHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	regionName := vars["region"]

	addresses := region.GetRegionHandler().GetAddressesForRegion(regionName)

	// Convert the response to JSON
	js, err := json.Marshal(addresses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

/*
	Returns all regions that the server has addresses for
*/
func GetRegions(w http.ResponseWriter, r *http.Request) {

	// For all available regions
	availableRegions := region.GetRegionHandler().GetAvailableRegions()
	// Convert the response to JSON
	js, err := json.Marshal(availableRegions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
