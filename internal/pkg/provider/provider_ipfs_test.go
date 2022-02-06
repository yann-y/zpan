package provider

import "testing"

func TestGetIPfsAddUrl(t *testing.T) {
	c := Config{
		Endpoint: "http://192.168.1.238.128:5001/",
	}
	ipfs, err := NewIpfsProvider(&c)
	if err != nil {
		t.Log(err)
	}
	url, _, _ := ipfs.SignedPutURL("", "", 0, true)
	t.Log(url)
}
