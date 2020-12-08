package config

import (
	"net/url"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Client struct {
	CC config_client.IConfigClient
	NC naming_client.INamingClient
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
	// 创建服务配置客户端
	cc, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  cfg.cc,
		"serverConfigs": cfg.sc,
	})
	if err != nil {
		return nil, err
	}

	// TODO(ljy): 服务发现和服务配置分离
	nc, err := clients.CreateNamingClient(map[string]interface{}{
		"clientConfig":  cfg.cc,
		"serverConfigs": cfg.sc,
	})
	if err != nil {
		return nil, err
	}
	return &Client{CC: cc, NC: nc}, err
}

func (cli *Client) GetConfig(dataID string, group string) (string, error) {
	return cli.CC.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
}

func (cli *Client) GetService(service string, group string, cluster []string) (model.Service, error) {
	s, err := cli.NC.GetService(vo.GetServiceParam{
		ServiceName: service,
		Clusters:    cluster, // default value is DEFAULT
		GroupName:   group,   // default value is DEFAULT_GROUP
	})
	if err != nil {
		return model.Service{}, err
	}
	return s, err
}

func (cli *Client) RegisterInstance(param vo.RegisterInstanceParam) (bool, error) {
	return cli.NC.RegisterInstance(param)
}
