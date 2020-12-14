package bridge

type Bridge struct {
	IpAddr string
	ApiKey string
}

func GetBridge(ipAddr string, apiKey string) Bridge {
	return Bridge{IpAddr: ipAddr, ApiKey: apiKey}
}
