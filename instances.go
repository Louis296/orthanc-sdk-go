package orthanc_sdk_go

import (
	"fmt"
	"github.com/louis296/orthanc-sdk-go/api"
)

type DownloadDicomFileReq struct {
	Id string
}

func (c *Client) DownloadDicomFile(req DownloadDicomFileReq) ([]byte, error) {
	url := fmt.Sprintf(api.URL["DownloadDicomFile"], req.Id)
	resp, err := api.DoGET(c.Host+url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
