package configs

import (
	"fmt"
	mlog "github.com/MarconiProtocol/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"sync"
)

// Defaults for config values
const (
	LOG_DIR                                           = "/var/log/marconi"
	LOG_LEVEL                                         = "info"
	BOOTNODE_ADDRESS                                  = "<Not Configured>"
	BOOTNODE_SERVER_ADDRESS_TTL_MINUTES               = 2
	BOOTNODE_RPC_PORT                                 = "5000"
	BOOTNODE_RUN_SOCKS_5                              = false
	ANNOUNCEMENT_REGION_ANNOUNCEMENT_INTERVAL_SECONDS = 10
)

const CONFIG_PATH = "/etc/dvpn"
const CONFIG_EXT = "yml"

var config *Configuration
var configLock sync.RWMutex

var configViperInst *viper.Viper

func InitializeConfigs(baseDir string) {
	InitializeDHTConfig(baseDir)
}

func InitializeDHTConfig(baseDir string) {
	configViperInst = viper.New()

	configName := "config"

	configViperInst.SetConfigName(configName)
	configViperInst.SetConfigType(CONFIG_EXT)
	configViperInst.AddConfigPath(baseDir + CONFIG_PATH)

	// Set Default Values for the DHT config
	setConfigDefaults()

	// Read in the DHT configuration file
	readAndLoadConfig()

	// triggered on initial read and every time DHTConfig is modified
	configViperInst.OnConfigChange(func(event fsnotify.Event) {
		// Read in config file
		readAndLoadConfig()
	})

	// Watch the Configuration file for any changes
	configViperInst.WatchConfig()
}

func setConfigDefaults() {
	/*
		Log Related Configurations
	*/
	configViperInst.SetDefault("log.level", LOG_LEVEL)
	configViperInst.SetDefault("log.dir", LOG_DIR)

	/*
		BootNode Related Configurations
	*/
	configViperInst.SetDefault("bootnode.address", BOOTNODE_ADDRESS)
	configViperInst.SetDefault("bootnode.server_address_ttl_minutes", BOOTNODE_SERVER_ADDRESS_TTL_MINUTES)
	configViperInst.SetDefault("bootnode.rpc_port", BOOTNODE_RPC_PORT)
	configViperInst.SetDefault("bootnode.run_socks5", BOOTNODE_RUN_SOCKS_5)

	/*
		Announcement Related Configurations
	*/
	configViperInst.SetDefault("announcement.region_announcement_interval_seconds", ANNOUNCEMENT_REGION_ANNOUNCEMENT_INTERVAL_SECONDS)
}

func readAndLoadConfig() {
	if err := configViperInst.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			// pass if appConfig file does not exist
		default:
			mlog.GetLogger().Error(fmt.Sprintf("Error reading app config file, %s\n", err.Error()))
			os.Exit(1)
		}
	}
	// parse appConfig file into their corresponding struct
	LoadConfig()
}

func GetAppConfig() *Configuration {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func LoadConfig() {
	newConfiguration := new(Configuration)
	if err := configViperInst.Unmarshal(newConfiguration); err != nil {
		fmt.Printf("unable to decode application config file into struct, %v", err)
	}
	configLock.Lock()
	config = newConfiguration
	configLock.Unlock()
}
