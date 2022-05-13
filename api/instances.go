package api

import (
	"fmt"
	"github.com/louis296/orthanc-sdk-go/model"
)

func (c *Client) DownloadDicomFile(req model.DownloadDicomFileReq) ([]byte, error) {
	url := fmt.Sprintf(URL["DownloadDicomFile"], req.Id)
	resp, err := doGET(c.Host+url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
