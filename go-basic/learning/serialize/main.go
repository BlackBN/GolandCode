package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	config := map[string]string{"test": "{\"recollectBefore\":\"2022-10-27\",\"encryption\":false,\"logCollectorOutputType\":\"elasticsearch\",\"vminsert\":{\"endpoint\":\"http://172.19.39.61:8480/insert/0/prometheus/api/v1/write\"},\"blueking\":{\"endpoint\":\"http://172.19.39.61:8480/insert/0/prometheus/api/v1/write\"},\"opampserver\":{\"endpoint\":\"ws://172.19.39.61:4320/v1/opamp\"},\"kafka\":{\"version\":\"2.3.1\",\"brokers\":\"172.19.37.35:30994\",\"auth\":{\"enabled\":true,\"type\":\"sasl,tls,plain_text\",\"tls\":{\"insecure\":false,\"insecure_skip_verify\":true},\"sasl\":{\"mechanism\":\"PLAIN\",\"username\":\"testa\",\"password\":\"w+IxjnE38SivVIcuHaI41g==\"},\"plain_text\":{\"username\":\"testa\",\"password\":\"w+IxjnE38SivVIcuHaI41g==\"}}},\"elasticsearch\":{\"hosts\":\"172.19.39.61:9200\",\"urls\":\"172.19.39.61:9200\",\"username\":\"admin\",\"password\":\"w+IxjnE38SivVIcuHaI41g==\",\"protocol\":\"http\",\"tls\":{\"enabled\":false,\"skip_host_verify\":true}}}"}

	observeComponentConfig, err := json.Marshal(config)
	if err != nil {

		panic(err)
	}
	result, err := convertJsonToObserveComponentConfig("test", string(observeComponentConfig))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)
}

// ExtraComponentInfo ...
type ExtraComponentInfo struct {
	VMInsert    VMInsert    `yaml:"vminsert" json:"vminsert"`
	Blueking    Blueking    `yaml:"blueking" json:"blueking"`
	OpampServer OpampServer `yaml:"opampserver" json:"opampserver"`
}

// VMInsert ...
type VMInsert struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
}

// Blueking ...
type Blueking struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
}

// OpampServer ...
type OpampServer struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
}

type KafkaInfo struct {
	Kafka Kafka `yaml:"kafka" json:"kafka"`
}

type Kafka struct {
	Brokers     string    `yaml:"brokers" json:"brokers"`
	EnabledAuth bool      `yaml:"enabledAuth" json:"enabledAuth"`
	AuthType    string    `yaml:"authType" json:"authType"`
	Version     string    `yaml:"version" json:"version"`
	Auth        KafkaAuth `yaml:"auth" json:"auth"`
}

type KafkaAuth struct {
	TLS       KafkaAuthTLS  `yaml:"tls" json:"tls"`
	PlainText KafkaAuthText `yaml:"plain_text" json:"plain_text"`
	SASL      KafkaAuthSASL `yaml:"sasl" json:"sasl"`
}

type KafkaAuthTLS struct {
	Insecure           bool `yaml:"insecure" json:"insecure"`
	InsecureSkipVerify bool `yaml:"insecure_skip_verify" json:"insecure_skip_verify"`
}

type KafkaAuthText struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type KafkaAuthSASL struct {
	Mechanism string `yaml:"mechanism" json:"mechanism"`
	Username  string `yaml:"username" json:"username"`
	Password  string `yaml:"password" json:"password"`
}

type SpecialKafkaInfo struct {
	ClusterIDList []string `yaml:"clusterIDList" json:"clusterIDList"`
	Kafka         Kafka    `yaml:"kafka" json:"kafka"`
}

type observeComponentConfig struct {
	Elasticsearch          Elasticsearch      `yaml:"elasticsearch" json:"elasticsearch"`
	Kafka                  Kafka              `yaml:"kafka" json:"kafka"`
	SpecialKafkaInfos      []SpecialKafkaInfo `yaml:"specialKafkaInfos" json:"specialKafkaInfos"`
	Vminsert               VMInsert           `yaml:"vminsert" json:"vminsert"`
	Blueking               Blueking           `yaml:"blueking" json:"blueking"`
	OpampServer            OpampServer        `yaml:"opampserver" json:"opampserver"`
	LogCollectorOutputType string             `yaml:"logCollectorOutputType" json:"logCollectorOutputType"`
	Encryption             bool               `yaml:"encryption" json:"encryption"`
	RecollectBefore        string             `yaml:"recollectBefore" json:"recollectBefore"`
}

type Elasticsearch struct {
	Hosts    string `yaml:"hosts" json:"hosts"`
	Urls     string `yaml:"urls" json:"urls"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Protocol string `yaml:"protocol" json:"protocol"`
	TLS      ESTLS  `yaml:"tls" json:"tls"`
}
type ESTLS struct {
	Enabled        bool `yaml:"enabled" json:"enabled"`
	SkipHostVerify bool `yaml:"skip_host_verify" json:"skip_host_verify"`
}

type observeComponentExtraInfo struct {
	Biz     string `yaml:"Biz" json:"Biz"`
	EnvType string `yaml:"EnvType" json:"EnvType"`
	Region  string `yaml:"Region" json:"Region"`
	Vpc     string `yaml:"Vpc" json:"Vpc"`
}

func convertJsonToObserveComponentExtraInfo(config string) (*observeComponentExtraInfo, error) {
	info := observeComponentExtraInfo{}
	if config != "" {
		if err := json.Unmarshal([]byte(config), &info); err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("unmarshal config %s to observeComponentExtraInfo is error", config))
		}
	}
	return &info, nil
}

func convertJsonToObserveComponentConfig(region, config string) (*observeComponentConfig, error) {
	configMapData := make(map[string]string)
	if config != "" {
		if err := json.Unmarshal([]byte(config), &configMapData); err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("unmarshal config %s to map[string]string is error", config))
		}
	}
	//fmt.Printf("configMapData : %v\n", configMapData)
	observeComponentConfigJson, exist := configMapData[region]
	//fmt.Printf("observeComponentConfigJson : %s\n", observeComponentConfigJson)

	if !exist || observeComponentConfigJson == "" {
		return nil, fmt.Errorf("this region %s has not observe component info", region)
	}
	info := &observeComponentConfig{
		Elasticsearch: Elasticsearch{
			Hosts:    "127.0.0.1:9200",
			Urls:     "127.0.0.1:9200",
			Username: "changeme",
			Password: "changeme",
			Protocol: "http",
			TLS: ESTLS{
				Enabled:        false,
				SkipHostVerify: true,
			},
		},
		Kafka: Kafka{
			Brokers:     "127.0.0.1:9092,127.0.0.1:9093",
			Version:     "2.3.1",
			EnabledAuth: false,
			AuthType:    "sasl",
			Auth: KafkaAuth{

				TLS: KafkaAuthTLS{
					Insecure:           false,
					InsecureSkipVerify: true,
				},
				PlainText: KafkaAuthText{
					Username: "changeme",
					Password: "changeme",
				},
				SASL: KafkaAuthSASL{
					Mechanism: "PLAIN",
					Username:  "changeme",
					Password:  "changeme",
				},
			},
		},
		Vminsert: VMInsert{
			Endpoint: "changeme",
		},
		Blueking: Blueking{
			Endpoint: "changeme",
		},
		OpampServer: OpampServer{
			Endpoint: "ws://ip:4320/v1/opamp",
		},
		LogCollectorOutputType: "kafka",
		Encryption:             false,
		RecollectBefore:        "2022-10-27",
	}
	if err := json.Unmarshal([]byte(observeComponentConfigJson), info); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("unmarshal config %s to observeComponentConfig is error", observeComponentConfigJson))
	}
	//fmt.Printf("observeComponentConfig : %+v\n", info)
	return info, nil
}
