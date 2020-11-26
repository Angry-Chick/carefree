package config

import (
	"net/url"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Client struct {
	config_client.IConfigClient
}

type Config struct {
	sc []constant.ServerConfig
	cc constant.ClientConfig
}

func DefaultConfig(endpoint string, un string, pwd string) (*Config, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}
	return &Config{
		sc: []constant.ServerConfig{{
			IpAddr: u.Hostname(),
			Port:   uint64(port),
			Scheme: "http",
		}},
		cc: constant.ClientConfig{
			Username:            un,
			Password:            pwd,
			NamespaceId:         "public",
			TimeoutMs:           5000,
			ListenInterval:      10000,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			RotateTime:          "1h",
			MaxAge:              3,
			LogLevel:            "debug",
		},
	}, nil
}

func NewClient(cfg *Config) (*Client, error) {
	cli, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  cfg.cc,
		"serverConfigs": cfg.sc,
	})
	if err != nil {
		return nil, err
	}
	return &Client{cli}, err
}

func (cli *Client) GetConfig(dataID string, group string) (string, error) {
	return cli.IConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
}
