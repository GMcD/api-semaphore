package env

import (
	"context"
	"fmt"

	"github.com/google/go-github/v61/github"
)

// Move to secret store
const GH_token = "ghp_***"

// Move to input?
const GH_user = "GMcD"
const GH_repo = "api-semaphore"
const GH_perpage = 100
const GH_EnvLen = 11

type Variables struct {
}

type Secrets struct {
}

func (*Variables) GetUser() (*github.User, error) {

	client := github.NewClient(nil).WithAuthToken(GH_token)

	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}
	if *user.Login != GH_user {
		return nil, fmt.Errorf("Unexpected user %v", user.Login)
	}
	return user, nil
}

func (v *Variables) GetRepo(name string) (*github.Repository, error) {

	client := github.NewClient(nil).WithAuthToken(GH_token)

	user, err1 := v.GetUser()
	if err1 != nil {
		return nil, err1
	}
	repo, _, err2 := client.Repositories.Get(context.Background(), *user.Login, name)
	if err2 != nil {
		return nil, err2
	}
	return repo, nil
}

func (v *Variables) GetGithubEnvironment(name string) (*map[string]string, error) {

	client := github.NewClient(nil).WithAuthToken(GH_token)

	user, err1 := v.GetUser()
	if err1 != nil {
		return nil, err1
	}

	// Fit list of Variables to one page, unless the list goes too long..
	vars, _, err2 := client.Actions.ListRepoVariables(context.Background(), *user.Login, name, &github.ListOptions{PerPage: GH_perpage})
	if err2 != nil {
		return nil, err2
	}
	// Fit list of Secret names to one page, unless the list goes too long..
	secrets, _, err3 := client.Actions.ListRepoSecrets(context.Background(), *user.Login, name, &github.ListOptions{PerPage: GH_perpage})
	if err3 != nil {
		return nil, err3
	}

	var env = make(map[string]string)

	for _, variable := range vars.Variables {
		env[variable.Name] = variable.Value
	}
	// Do not get Secret Values back from API
	for _, secret := range secrets.Secrets {
		env[secret.Name] = "***"
	}

	return &env, nil
}

func (*Variables) GetOrganizations() ([]*github.Organization, error) {

	client := github.NewClient(nil).WithAuthToken(GH_token)

	orgs, _, err := client.Organizations.List(context.Background(), GH_user, nil)
	if err != nil {
		err = fmt.Errorf("Cannot connect to Github as %v: %v", GH_user, err)
		return nil, err
	}

	fmt.Println(orgs)

	return orgs, nil

}
