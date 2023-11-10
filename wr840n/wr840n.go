package wr840n

import (
	"fmt"

	"github.com/furtidev/tplinkctl/utils"
	"github.com/urfave/cli/v2"
)

const (
	dhcpClientURL = "http://192.168.0.1/cgi?5"
	dhcpClientPayload = "[LAN_HOST_ENTRY#0,0,0,0,0,0#0,0,0,0,0,0]0,4\r\nleaseTimeRemaining\r\nMACAddress\r\nhostName\r\nIPAddress\r\n"
)

var authToken string

func Setup(ctx *cli.Context) error {
	authToken = "Basic " + utils.EncodeBase64(ctx.String("user")+":"+ctx.String("pass"))
	return nil
}

func Clients(ctx *cli.Context) error {
	body, err := utils.MakeRequest(dhcpClientURL, dhcpClientPayload, authToken)
	if err != nil {
		return err
	}

	dat := ParseDHCPData(body)

	fmt.Printf("Currently %d clients are connected.\n", len(dat))

	for _, v := range dat {
		fmt.Printf("- %s\n    - %s\n    - %s\n", v.HostName, v.IPAddress, v.MacAddress)
	}
	return nil
}

func Status(ctx *cli.Context) error {
	fmt.Println("Status page not implemented yet.")
	return nil
}
