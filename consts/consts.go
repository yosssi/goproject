// Package consts provides constants.
package consts

// consts definition
const (
	YmlExtension            = ".yml"
	ConfigPath              = "./config/"
	LoggerConfigPath        = ConfigPath + "logger" + YmlExtension
	ElasticsearchConfigPath = ConfigPath + "elasticsearch" + YmlExtension
	GitHubConfigPath        = ConfigPath + "github" + YmlExtension
	QueryUpdate             = "q=language:go&sort=updated&order=asc&page=%d"
	IndexGoProject          = "goproject"
	TypeGitHubRepository    = "github_repository"
)
