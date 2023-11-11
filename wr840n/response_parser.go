package wr840n

import (
	"bufio"
	"strings"
	"strconv"
)

/*
	Status information Parser
*/
type StatusInfo struct {
	FirmwareVer, HardwareVer string
	Uptime int
}

type StatusType uint8

const (
	BasicInfo StatusType = iota
	NotRecognized
)

func ParseStatusInfoData(body string) StatusInfo {
	dat := StatusInfo{}

	scanner := bufio.NewScanner(strings.NewReader(body))

	for scanner.Scan() {
		line := scanner.Text()

		lineType := NotRecognized

		if line[0] == '[' {
			if line == "[0,0,0,0,0,0]2" {
				lineType = BasicInfo
			} else {
				scanner.Scan()
				continue
			}
		}
		if lineType == BasicInfo {
			scanner.Scan()			
			shouldSkip := false

			for !shouldSkip {
				subLine := scanner.Text()
				if subLine[0] == '[' {
					shouldSkip = true
					continue
				}

				delim := strings.IndexByte(subLine, '=')
				switch key := subLine[:delim]; key {
				case "softwareVersion":
					dat.FirmwareVer = subLine[delim+1:]
				case "hardwareVersion":
					dat.HardwareVer = subLine[delim+1:]
				case "upTime":
					up, _ := strconv.Atoi(subLine[delim+1:])
					dat.Uptime = up
				}
				scanner.Scan()
			}
			break
		}
	}

	return dat
}

/* 
	DHCP client list Parser 
*/
type DHCPClient struct {
	LeaseTimeRemaining, MacAddress, HostName, IPAddress string
}

// i need to find a better way to do this lol
func ParseDHCPData(body string) []DHCPClient {
	dat := make([]DHCPClient, 0)

	scanner := bufio.NewScanner(strings.NewReader(body))

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '[' {
			scanner.Scan()
		} else {
			cl := DHCPClient{}

			shouldSkip := false
			
			for !shouldSkip {
				subLine := scanner.Text()
				if subLine[0] == '[' {
					shouldSkip = true
					continue
				}
				delim := strings.IndexByte(subLine, '=')
				switch key := subLine[:delim]; key {
				case "leaseTimeRemaining":
					cl.LeaseTimeRemaining = subLine[delim+1:]
				case "MACAddress":
					cl.MacAddress = subLine[delim+1:]
				case "hostName":
					cl.HostName = subLine[delim+1:]
				case "IPAddress":
					cl.IPAddress = subLine[delim+1:]
				}
				scanner.Scan()
			}

			dat = append(dat, cl)


		}
	}
	return dat
}