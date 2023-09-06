package configs

import (
	"io/ioutil"
	"os"
	"testing"
)

/**
This test checks to see if default values get set and can be retrieved if no config files are provided
*/
func TestNoConfigFile(t *testing.T) {

	// No config file will exist, we'll expect defaults to be set
	InitializeConfigs("test/")

	// After Initialization is complete, retrieve values
	logLevel := GetAppConfig().Log.Level
	logDir := GetAppConfig().Log.Dir
	bootNodeAddress := GetAppConfig().BootNode.Address
	bootNodeRPCPort := GetAppConfig().BootNode.RPC_PORT
	bootNodeTTL := GetAppConfig().BootNode.Server_Address_TTL_Minutes
	bootNodeRunSocks5 := GetAppConfig().BootNode.RUN_SOCKS_5
	announcementInterval := GetAppConfig().Announcement.Region_Announcement_Interval_Seconds

	// Checking to see if values match the defaults
	if logLevel != LOG_LEVEL {
		t.Error("Error: expected", LOG_LEVEL, "but got", logLevel)
	}

	if logDir != LOG_DIR {
		t.Error("Error: expected", LOG_DIR, "but got", logDir)
	}

	if bootNodeAddress != BOOTNODE_ADDRESS {
		t.Error("Error: expected", BOOTNODE_ADDRESS, "but got", bootNodeAddress)
	}

	if bootNodeRPCPort != BOOTNODE_RPC_PORT {
		t.Error("Error: expected", BOOTNODE_RPC_PORT, "but got", bootNodeRPCPort)
	}

	if bootNodeTTL != BOOTNODE_SERVER_ADDRESS_TTL_MINUTES {
		t.Error("Error: expected", BOOTNODE_SERVER_ADDRESS_TTL_MINUTES, "but got", bootNodeTTL)
	}

	if bootNodeRunSocks5 != BOOTNODE_RUN_SOCKS_5 {
		t.Error("Error: expected", BOOTNODE_RUN_SOCKS_5, "but got", bootNodeRunSocks5)
	}

	if announcementInterval != ANNOUNCEMENT_REGION_ANNOUNCEMENT_INTERVAL_SECONDS {
		t.Error("Error: expected", ANNOUNCEMENT_REGION_ANNOUNCEMENT_INTERVAL_SECONDS, "but got", announcementInterval)
	}

	t.Log("Success")
}

/**
This test will write a yaml file and see if the values get retrieved properly from the file
(overwriting defaults)
*/
func TestConfigFile(t *testing.T) {

	// Creating a test directory to store config files
	// Note: If you decide to change this value, make sure it is some empty directory/non-existent directory
	baseDir := "../../test"
	configPath := baseDir + CONFIG_PATH

	// Create the necessary directories
	os.MkdirAll(configPath, os.ModePerm)

	var yamlExample = []byte(`
LOG.Level: debug
LOG.Dir: /var/log/test
BOOTNODE.Address: 172.30.20.10
BOOTNODE.SERVER_ADDRESS_TTL_MINUTES: 4
BOOTNODE.RPC_PORT: 8000
BOOTNODE.RUN_SOCKS_5: true
Announcement.Region_Announcement_Interval_Seconds: 12
`)
	//fmt.Println("Writing to", configPath)
	err := ioutil.WriteFile(configPath+"/config.yml", yamlExample, 0644)

	if err != nil {
		t.Error("Error occurred while creating config yml file", err)
	}

	// Now initializing the config
	InitializeConfigs(baseDir)

	// Retrieving config values
	logLevel := GetAppConfig().Log.Level
	logDir := GetAppConfig().Log.Dir
	bootNodeAddress := GetAppConfig().BootNode.Address
	bootNodeRPCPort := GetAppConfig().BootNode.RPC_PORT
	bootNodeTTL := GetAppConfig().BootNode.Server_Address_TTL_Minutes
	bootNodeRunSocks5 := GetAppConfig().BootNode.RUN_SOCKS_5
	announcementInterval := GetAppConfig().Announcement.Region_Announcement_Interval_Seconds

	// Checking to see if values match the ones written to the config

	// Checking to see if values match the defaults
	if logLevel != "debug" {
		t.Error("Error: expected", "debug", "but got", logLevel)
	}

	if logDir != "/var/log/test" {
		t.Error("Error: expected", "/var/log/test", "but got", logDir)
	}

	if bootNodeAddress != "172.30.20.10" {
		t.Error("Error: expected", "172.30.20.10", "but got", bootNodeAddress)
	}

	if bootNodeRPCPort != "8000" {
		t.Error("Error: expected", 8000, "but got", bootNodeRPCPort)
	}

	if bootNodeTTL != 4 {
		t.Error("Error: expected", 4, "but got", bootNodeTTL)
	}

	if bootNodeRunSocks5 != true {
		t.Error("Error: expected", true, "but got", bootNodeRunSocks5)
	}

	if announcementInterval != 12 {
		t.Error("Error: expected", 12, "but got", announcementInterval)
	}

	// Cleaning up test directories
	err = os.RemoveAll(baseDir)

	if err != nil {
		t.Error("Error occurred while cleaning up test directories")
	}

	t.Log("Success")
}
