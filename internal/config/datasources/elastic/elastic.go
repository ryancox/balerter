package elastic

import (
	"fmt"
	"strings"

	"github.com/balerter/balerter/internal/config/common"
)

// Elastic datasource config
type Elastic struct {
	// Name of the datasource
	Name string `json:"name" yaml:"name" hcl:"name,label"`
	// Host value
	Host string `json:"host" yaml:"host" hcl:"host"`
	// Port value
	Port int `json:"port" yaml:"port" hcl:"port"`
	// BasicAuth contains auth data, if needed
	BasicAuth *common.BasicAuth `json:"basicAuth" yaml:"basicAuth" hcl:"basicAuth,block"`
	// Timeout value
	Timeout int `json:"timeout" yaml:"timeout" hcl:"timeout,optional"`
	// Scheme specifies the url connection scheme of http or https
	Scheme string `json:"scheme" yaml:"scheme" hcl:"scheme"`
	// Sniff specifies whether or not to enable client sniffing
	Sniff string `json:"sniff" yaml:"sniff" hcl:"sniff"`
}

// Validate config
func (cfg Elastic) Validate() error {
	if strings.TrimSpace(cfg.Name) == "" {
		return fmt.Errorf("name must be not empty")
	}
	if cfg.Host == "" {
		return fmt.Errorf("host must be defined")
	}
	if cfg.Port == 0 {
		return fmt.Errorf("port must be defined")
	}
	if cfg.Timeout < 0 {
		return fmt.Errorf("timeout must be greater than 0")
	}
	if cfg.Scheme != "https" && cfg.Scheme != "http" {
		return fmt.Errorf("scheme must be either http or https")
	}
	if cfg.Sniff != "true" && cfg.Sniff != "false" {
		return fmt.Errorf("sniff must be either true or false")
	}

	return nil
}
