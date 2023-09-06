package httpServer

import (
	"fmt"
	"github.com/MarconiFoundation/dvpn/core/server/http/api"
	"github.com/MarconiFoundation/dvpn/core/ip"
	mlog "github.com/MarconiProtocol/log"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type HttpServer struct {
	port   string
	ip     string
	router *mux.Router
}

func Initialize(port string) (*HttpServer, error) {

	httpServer := HttpServer{}
	httpServer.port = port

	mainIpAddr, err := mnet_ip.GetMainInterfaceIpAddress()
	httpServer.ip = mainIpAddr

	// Building the router
	router := mux.NewRouter()

	// Adding 'routes'
	router.HandleFunc("/region/{region}", api.RegionHandler)
	router.HandleFunc("/regions", api.GetRegions)

	httpServer.router = router

	if err != nil {
		return nil, err
	}

	return &httpServer, nil
}

func (hs *HttpServer) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	mlog.GetLogger().Info(fmt.Sprintf("Starting http server on %s:%s", hs.ip, hs.port))
	http.ListenAndServe(hs.ip+":"+hs.port, hs.router)
}
