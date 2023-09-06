package utils

import "testing"

func TestIsBootNodeTrue(t *testing.T) {
	ipAddress := "172.30.20.10"
	bootnodes := "172.30.20.10:8080"

	isBootnode := isIpAddressInSet(ipAddress, bootnodes)

	if isBootnode != true {
		t.Error("Error: expected true but got ", isBootnode)
	}

	t.Log("Success")
}

func TestIsBootNodeMultipleBootnodes(t *testing.T) {
	ipAddress := "172.30.20.10"
	bootnodes := "172.30.20.10:8080,172.30.20.11:8080"

	isBootnode := isIpAddressInSet(ipAddress, bootnodes)

	if isBootnode != true {
		t.Error("Error: expected true but got ", isBootnode)
	}

	t.Log("Success")
}

func TestIsBootNodeFalse(t *testing.T) {
	ipAddress := "172.30.20.11"
	bootnodes := "172.30.20.10:8080"

	isBootnode := isIpAddressInSet(ipAddress, bootnodes)

	if isBootnode != false {
		t.Error("Error: expected false but got ", isBootnode)
	}

	t.Log("Success")
}
