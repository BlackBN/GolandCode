package main

import (
	"GolandCode/bn/bn-orm/database-sql/log"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
)

func main() {
	// funcMap := template.FuncMap{
	// 	"strupper": upper,
	// }
	// t := template.New("test1")
	// tmpl, err := t.Funcs(funcMap).Parse(`{{strupper .}}`)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = tmpl.Execute(os.Stdout, "go programming")

	TestToYaml()
}

func upper(str string) string {
	return strings.ToUpper(str)
}

type otel struct {
	KafkaInfo  *KafkaInfo `yaml:"KafkaInfo,omitempty" json:"KafkaInfo,omitempty"`
	KafkaInfoW *KafkaInfo `yaml:"KafkaInfo2,omitempty" json:"KafkaInfo2,omitempty"`
}

type KafkaInfo struct {
	Kafka Kafka `yaml:"kafka" json:"kafka"`
}
type Kafka struct {
	Brokers     string   `yaml:"brokers" json:"brokers"`
	Server      []string `yaml:"server" json:"server"`
	EnabledAuth bool     `yaml:"enabledAuth" json:"enabledAuth"`
	Auth        Auth     `yaml:"auth" json:"auth"`
}

type Auth struct {
	Tls        Tls        `yaml:"tls" json:"tls"`
	Sasl       Sasl       `yaml:"sasl" json:"sasl"`
	Plain_text Plain_text `yaml:"plain_text" json:"plain_text"`
}

type Tls struct {
	Insecure             bool `yaml:"insecure" json:"insecure"`
	Insecure_skip_verify bool `yaml:"insecure_skip_verify" json:"insecure_skip_verify"`
}

type Sasl struct {
	Mechanism string `yaml:"mechanism" json:"mechanism"`
	Username  string `yaml:"username" json:"username"`
	Password  string `yaml:"password" json:"password"`
}

type Plain_text struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type Obser struct {
}

func TestToYaml() {

	kafkaInfo := KafkaInfo{
		Kafka: Kafka{
			Brokers:     "127.0.0.1,123.3.3.3,134.3.4.5",
			Server:      []string{"127.0.0.1", "123.4.4.4"},
			EnabledAuth: true,
			Auth: Auth{
				Tls: Tls{
					Insecure:             false,
					Insecure_skip_verify: true,
				},
				Sasl: Sasl{
					Mechanism: "mechanism",
					Username:  "admin",
					Password:  "password",
				},
			},
		},
	}

	otelOrigin := otel{
		KafkaInfo: &kafkaInfo,
	}

	datas, err := json.Marshal(otelOrigin)
	if err != nil {

	}
	fmt.Println(string(datas))
	otel := &otel{}
	err = yaml.Unmarshal([]byte(datas), otel)
	if err != nil {
		log.Error(err, "failed to unmarshal log collector conf")
	}
	changePaasword(&otel.KafkaInfo.Kafka)
	fmt.Printf(renderHelmOverrides("opentelemetry-collector", opentelemetryCollectorOverrides, otel))
}

func changePaasword(kafka *Kafka) {
	kafka.Auth.Sasl.Password = "changePassword"
}

const (
	metallbProvider = "metallb"
	volclbProvider  = "volcengine"
)

var opentelemetryCollectorOverrides = `
{{ if . }}
{{ if .KafkaInfoW }}
{{ toYaml .KafkaInfoW }}
{{ end }}

{{ toYaml .KafkaInfo }}
{{- end }}
`

func renderHelmOverrides(name, text string, data interface{}) string {
	tpl := template.New(name)
	// add func
	tpl.Funcs(map[string]interface{}{
		"isMetalLB": isMetalLB,
		"isVolcLB":  isVolcLB,
		"toYaml":    toYAML,
	})
	tpl.Funcs(sprig.TxtFuncMap())

	_, err := tpl.Parse(text)
	if err != nil {
		log.Errorf("render %v err: %v", name, err)
	}

	buff := &bytes.Buffer{}
	err = tpl.Execute(buff, data)
	if err != nil {
		log.Errorf("render %v Execute err: %v", name, err)
	}

	//log.Infof("addon %v rendor result %v, text: %v, data: %+v", name, buff.String(), text, data)
	return buff.String()
}

func isMetalLB(p string) bool {
	return p == metallbProvider
}

func isVolcLB(p string) bool {
	return p == volclbProvider || p == ""
}

func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}
