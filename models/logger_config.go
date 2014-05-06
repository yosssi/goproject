package models

import (
	"github.com/yosssi/gologger"
	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/utils"
)

// LoggerConfig represents a logger config.
type LoggerConfig struct {
	Name              string `yaml:"name"`
	Level             string `yaml:"level"`
	File              string `yaml:"file"`
	OutputFileColored bool   `yaml:"output_file_colored"`
}

// Logger generates a Logger and returns it.
func (l *LoggerConfig) Logger() *gologger.Logger {
	return gologger.NewLogger(l.Name, l.Level, l.File, l.OutputFileColored)
}

// NewLoggerConfig parses a yaml file, generates a LoggerConfig and returns it.
func NewLoggerConfig() (*LoggerConfig, error) {
	config := &LoggerConfig{}
	if err := utils.YamlUnmarshal(consts.LoggerConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
