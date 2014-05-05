package models

import (
	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/utils"
)

// ElasticsearchConfig represents an Elasticsearch config.
type ElasticsearchConfig struct {
	BaseURL string `yaml:"base_url"`
}

// NewElasticsearchConfig parses a yaml file, generates a ElasticsearchConfig and returns it.
func NewElasticsearchConfig() (*ElasticsearchConfig, error) {
	config := &ElasticsearchConfig{}
	if err := utils.YamlUnmarshal(consts.ElasticsearchConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
