package configs

type Configuration struct {
	Log struct {
		Level string // indicates the log level ('info', 'debug' etc.)
		Dir   string // Directory in which to store logs in
	}

	// Configurations specific to the BootNode
	BootNode struct {
		Address                    string // Indicates the address of the bootnode used for discovery by peer nodes
		Server_Address_TTL_Minutes int    // Indicates how long the bootnode will keep a server address for a region without receiving an announcement
		RPC_PORT                   string // Indicates which port to run the rpc client on
		RUN_SOCKS_5                bool   // Indicates if the BootNode should run a Socks5 server
	}

	Announcement struct {
		Region_Announcement_Interval_Seconds int // Indicates how often the machine's region + socks5 server address gets announced to the bootnode
	}
}
