# Decentralized VPN

<p align="center"> 
<img src="https://user-images.githubusercontent.com/13502814/69590012-ecc9fb00-0fa2-11ea-9047-6d3e3d1a70d9.png">
</p>

## Running 

Getting Dependencies 

`./get_deps.sh`

To Build 

`./build.sh`

To Run

`./out/dvpn`

Tests

`./test.sh`

`go test <specific test file>`

### Command Line Arguments

dvpn can be run with the following command line arguments:


**basedir** [default: /opt/marconi] Base of directory tree that will be used to store all configs and data files related to dv

**region** [default: "US:CA"] Indicates the location/region of the server

**socks5-port** [default: 5657] Port to run the socks5 server on

**http-port** [default: 8080] Port to run the http server on (only applicable to bootnodes)

## API

Nodes running in 'Bootnode mode' also have an API exposed on whatever port is indicated by the httpPort command line argument. 
In order to run a node in "bootnode" mode, simply include the node's address in the BootNode.Address configuration. 

### Region

#### Get all Ip addresses for a Region:

Returns all ip addresses for a specific region. (All Machines broadcasting that region)

`/region/{region}` 

**Parameters**

`region` The region to query


#### Get all Available Regions:

Returns all available regions, in which each region has at least 1 IP address associated with it.

`/regions`

**Parameters**

None


## Dev Test

There is a simple testing framework that lives under the devtest directory. 
Configurations for this can be edited by editing `devtest/data/config.yml`. 

To build binaries and docker images run:

`./setup_env.sh`

To launch containers: 

`sudo ./run_env.sh`

**Note:** Running the above will launch 4 containers

**n0:** BootNode (Has Http Server, RPC Server, Socks5 Server (optional) running) 

**n1, n2, n3:** PeerNodes (Has socks5Server running)

## Configuration

_Parameters_

```
	Log.Level                                             string   indicates the log level ('info', 'debug' etc.)
	Log.Dir                                               string   Directory in which to store logs in
	BootNode.Address                                      string   Indicates the address of the bootnode used for discovery by peer nodes
	BootNode.RPC_Port                                     string   Indicates which port to run the rpc client on
	Bootnode.Server_Address_TTL_Minutes                   int      Indicates how long the bootnode will keep a server address for a region without receiving an announcement
	Bootnode.Run_Socks_5                                  bool     Indicates if the BootNode should run a Socks5 server
	Announcement.Region_Announcement_Interval_Seconds     int      Indicates how often the machine's region + socks5 server address gets announced to the bootnode                     
```

_Example Configuration:_ 

```
LOG.dir: /var/log/marconi
LOG.Level: debug
BootNode.Address: "172.30.20.10"
BootNode.Server_Address_TTL_Minutes: 1
BootNode.RPC_PORT: "5000"
BootNode.Run_Socks_5: false
Announcement.Region_Announcement_Interval_Seconds: 5
```

