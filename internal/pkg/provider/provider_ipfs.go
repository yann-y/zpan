package provider

import (
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
)

// OBSProvider 华为云
type IpfsProvider struct {
	client *shell.Shell
}

func NewIpfsProvider(conf *Config) (Provider, error) {
	return newIpfsProvider(conf)
}

func newIpfsProvider(conf *Config) (*IpfsProvider, error) {
	client := shell.NewShell(conf.Endpoint)
	return &IpfsProvider{
		client: client,
	}, nil
}

func (p *IpfsProvider) SetupCORS() error {
	return nil
}

// 查看文件权限
func (p *IpfsProvider) Head(object string) (*Object, error) {

	return &Object{
		Key:  "",
		ETag: "aws.StringValue(hOut.ETag)",
		Type: "aws.StringValue(hOut.ContentType)",
	}, nil
}

// List returns the remote objects
func (p *IpfsProvider) List(prefix string) ([]Object, error) {

	return nil, nil
}

func (p *IpfsProvider) Move(object, newObject string) error {
	return nil
}

func (p *IpfsProvider) SignedPutURL(key, filetype string, filesize int64, public bool) (string, http.Header, error) {

	return "", nil, nil
}

func (p *IpfsProvider) SignedGetURL(key, filename string) (string, error) {

	return "", nil
}

func (p *IpfsProvider) PublicURL(key string) string {

	return ""
}

func (p *IpfsProvider) ObjectDelete(key string) error {
	return nil
}

func (p *IpfsProvider) ObjectsDelete(objectKeys []string) error {
	return nil
}
