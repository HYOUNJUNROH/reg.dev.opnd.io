/*
Copyright © 2022 Openerd <jsmoon@openerd.com>
*/
package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"git.dev.opnd.io/gc/backend-admin/cmd"
	"git.dev.opnd.io/gc/backend-admin/pkg/config"
)

var (
	AppName     = "backend"
	GitCommit   = "development"
	GitBranch   = "development"
	GitTag      = "development"
	GitVersion  = "development"
	GitDateTime = "0001-01-01T00:00:00+00:00"
)

func main() {
	fmt.Printf("%v : %v\n", AppName, GitVersion)

	config.Version.AppName = AppName
	config.Version.GitCommit = GitCommit
	config.Version.GitBranch = GitBranch
	config.Version.GitTag = GitTag
	config.Version.GitVersion = GitVersion
	var err error
	config.Version.GitDateTime, err = validateGitDateTime(GitDateTime)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}

func validateGitDateTime(v string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, GitDateTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid GitDateTime: %s", err.Error())
	}
	if !config.IsDevelopment() && t.IsZero() {
		return time.Time{}, errors.New("invalid GitDateTime: zero value")
	}
	return t, nil
}
