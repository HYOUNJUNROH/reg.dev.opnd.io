package config

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var Version = struct {
	AppName     string    `yaml:"app_name" json:"app_name"`
	GitCommit   string    `yaml:"git_commit" json:"git_commit"`
	GitBranch   string    `yaml:"git_branch" json:"git_branch"`
	GitTag      string    `yaml:"git_tag" json:"git_tag"`
	GitVersion  string    `yaml:"git_version" json:"git_version"`
	GitDateTime time.Time `yaml:"git_datetime" json:"git_datetime"`
}{}

func GetVersion(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, Version, "  ")
}
