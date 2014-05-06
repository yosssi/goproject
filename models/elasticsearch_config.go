package models

import (
	"github.com/yosssi/goelasticsearch"
	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/utils"
)

// ElasticsearchConfig represents an Elasticsearch config.
type ElasticsearchConfig struct {
	BaseURL string `yaml:"base_url"`
}

// ElasticsearchClient generates an Elasticsearch client and returns it.
func (e *ElasticsearchConfig) ElasticsearchClient() *goelasticsearch.Client {
	return goelasticsearch.NewClient(e.BaseURL)
}

// NewElasticsearchConfig parses a yaml file, generates a ElasticsearchConfig and returns it.
func NewElasticsearchConfig() (*ElasticsearchConfig, error) {
	config := &ElasticsearchConfig{}
	if err := utils.YamlUnmarshal(consts.ElasticsearchConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
