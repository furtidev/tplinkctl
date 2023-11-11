package wr840n

import (
	"fmt"

	"github.com/furtidev/tplinkctl/utils"
	"github.com/urfave/cli/v2"
)

const (
	statusPageURL = "http://192.168.0.1/cgi?1&1&1&5&5&5&5&5&5&5"
	statusPagePayload = "[SYS_MODE#0,0,0,0,0,0#0,0,0,0,0,0]0,1\r\nmode\r\n[IGD#0,0,0,0,0,0#0,0,0,0,0,0]1,1\r\nLANDeviceNumberOfEntries\r\n[IGD_DEV_INFO#0,0,0,0,0,0#0,0,0,0,0,0]2,3\r\nsoftwareVersion\r\nhardwareVersion\r\nupTime\r\n[WAN_COMMON_INTF_CFG#0,0,0,0,0,0#0,0,0,0,0,0]3,1\r\nWANAccessType\r\n[WAN_IP_CONN#0,0,0,0,0,0#0,0,0,0,0,0]4,0\r\n[WAN_PPP_CONN#0,0,0,0,0,0#0,0,0,0,0,0]5,0\r\n[WAN_L2TP_CONN#0,0,0,0,0,0#0,0,0,0,0,0]6,0\r\n[WAN_PPTP_CONN#0,0,0,0,0,0#0,0,0,0,0,0]7,0\r\n[L2_BRIDGING_ENTRY#0,0,0,0,0,0#0,0,0,0,0,0]8,1\r\nbridgeName\r\n[LAN_WLAN#0,0,0,0,0,0#0,0,0,0,0,0]9,12\r\nstatus\r\nSSID\r\nBSSID\r\nchannel\r\nautoChannelEnable\r\nstandard\r\nbeaconType\r\nbasicEncryptionModes\r\nX_TP_Bandwidth\r\npossibleDataTransmitRates\r\nWPAAuthenticationMode\r\nIEEE11iAuthenticationMode\r\n"

	dhcpClientURL = "http://192.168.0.1/cgi?5"
	dhcpClientPayload = "[LAN_HOST_ENTRY#0,0,0,0,0,0#0,0,0,0,0,0]0,4\r\nleaseTimeRemaining\r\nMACAddress\r\nhostName\r\nIPAddress\r\n"
)

var authToken string

func Setup(ctx *cli.Context) error {
	authToken = "Basic " + utils.EncodeBase64(ctx.String("user")+":"+ctx.String("pass"))
	return nil
}

func Status(ctx *cli.Context) error {
	body, err := MakeRequest(statusPageURL, statusPagePayload, authToken)
	if err != nil {
		return err
	}

	dat := ParseStatusInfoData(body)

	fmt.Printf("Firmware version: %s\nHardware version: %s\nUptime: %.2f hours\n", dat.FirmwareVer, dat.HardwareVer, float32(dat.Uptime/3600))

	return nil
}

func Clients(ctx *cli.Context) error {
	body, err := MakeRequest(dhcpClientURL, dhcpClientPayload, authToken)
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
