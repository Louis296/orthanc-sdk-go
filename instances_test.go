package orthanc_sdk_go

import (
	"testing"
)

func TestDownloadDicomFile(t *testing.T) {
	orthancClient := Client{Host: "27.17.30.150:20086"}
	req := DownloadDicomFileReq{Id: "7ffd725f-1feee9bb-276d522a-09ab4473-333f1a4a"}
	_, err := orthancClient.DownloadDicomFile(req)
	if err != nil {
		t.Fatalf("%v api error", "DownloadDicomFile")
	}
}
