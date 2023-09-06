package mnet_ip

import (
  "github.com/MarconiFoundation/dvpn/core/cmd"
  "errors"
  "fmt"
  "strings"
)

/*
  Returns the main network interface ip address
*/
func GetMainInterfaceIpAddress() (string, error) {
  cmdSuite, err := msys_cmd.GetSuite()
  if err != nil {
    return "", errors.New(fmt.Sprintf("Failed to get cmd suite: %s", err))
  }
  mainInterfaceIp, err := cmdSuite.GetMainInterfaceIpAddress()
  if err != nil {
    return "", errors.New(fmt.Sprintf("cmdSuite.GetMainInterfaceIpAddress() failed with error: %s", err))
  }
  return strings.TrimSpace(mainInterfaceIp), nil
}

/*
  Returns the node's gateway ip address
*/
func GetOwnGatewayIpAddress() (string, error) {
  cmdSuite, err := msys_cmd.GetSuite()
  if err != nil {
    return "", errors.New(fmt.Sprintf("Failed to get cmd suite: %s", err))
  }
  gateway, err := cmdSuite.GetOwnGatewayIpAddress()
  if err != nil {
    return "", errors.New(fmt.Sprintf("cmdSuite.GetOwnGatewayIpAddress() failed with error: %s", err))
  }
  return gateway, nil
}