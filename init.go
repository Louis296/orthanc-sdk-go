package orthanc_sdk_go

import "github.com/louis296/orthanc-sdk-go/api"

func InitClient(host string) *api.Client {
	return &api.Client{Host: host}
}
