package provider

import (
	"net/http"
	"net/url"

	shell "github.com/ipfs/go-ipfs-api"
)

// IpfsProvider ipfs对等存储
type IpfsProvider struct {
	client   *shell.Shell
	endpoint string
}

func NewIpfsProvider(conf *Config) (Provider, error) {
	return newIpfsProvider(conf)
}

func newIpfsProvider(conf *Config) (*IpfsProvider, error) {
	client := shell.NewShell(conf.Endpoint)
	return &IpfsProvider{
		client:   client,
		endpoint: conf.Endpoint,
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
	u, err := url.Parse(p.endpoint)
	if err != nil {
		return "", nil, err
	}
	u.Path = "/api/v0/add"
	parm := url.Values{}
	parm.Add("pin", "true")
	parm.Add("wrap-with-directory", "false")
	parm.Add("stream-channels", "true")
	parm.Add("progress", "false")
	u.RawQuery = parm.Encode()
	return u.String(), nil, err
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
