package main_test

import (
	"testing"

	"github.com/GMcD/api-semaphore/module/env"
)

func TestGetUser(t *testing.T) {

	var v env.Variables

	_, err := v.GetUser()
	if err != nil {
		t.Errorf("Failed to get user from GitHub : %v", err)
	}
}

func TestGetRepo(t *testing.T) {

	var v env.Variables

	_, err := v.GetRepo(env.GH_repo)
	if err != nil {
		t.Errorf("Failed to get repository from GitHub : %v", err)
	}
}

func TestGetGithubEnvironment(t *testing.T) {

	var v env.Variables

	vars, err := v.GetGithubEnvironment(env.GH_repo)
	if err != nil {
		t.Errorf("Failed to get action variables/secrets from GitHub : %v", err)
	}
	if len(*vars) != env.GH_EnvLen {
		t.Errorf("Wrong count of action variables/secrets : %v", err)
	}
}

func TestGetOrganizations(t *testing.T) {

	var v env.Variables

	v.GetOrganizations()
}
