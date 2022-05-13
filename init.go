package orthanc_sdk_go

func InitClient(host string) *Client {
	return &Client{Host: host}
}
