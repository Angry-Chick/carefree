package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Client struct {
	config_client.IConfigClient
}

type Config struct {
	constant.ClientConfig
}

func DefaultConfig(endpoint string, un string, pwd string) *Config {
	return &Config{
		constant.ClientConfig{
			Username:            un,
			Password:            pwd,
			NamespaceId:         "public",
			TimeoutMs:           5000,
			ListenInterval:      10000,
			Endpoint:            endpoint,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			RotateTime:          "1h",
			MaxAge:              3,
			LogLevel:            "debug",
		},
	}
}

func NewClient(cfg *Config) (*Client, error) {
	cli, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": cfg.ClientConfig,
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
