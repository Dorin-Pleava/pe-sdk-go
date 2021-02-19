package app

import (
	"context"
	"fmt"

	"github.com/puppetlabs/pe-sdk-go/api/puppet-access/client/operations"
	"github.com/puppetlabs/pe-sdk-go/api/puppet-access/models"
	"github.com/puppetlabs/pe-sdk-go/token/filetoken"
)

// getToken queries the status endpoint of a puppetdb instance
func (puppetAccess *PuppetAccess) getToken() (string, error) {
	client, err := puppetAccess.Client.GetClient()
	if err != nil {
		return "", err
	}

	loginParameters := operations.NewLoginParamsWithContext(context.Background())
	body := models.LoginRequest{
		Login:    puppetAccess.Username,
		Password: puppetAccess.Password,
		Lifetime: puppetAccess.Lifetime,
		Label:    puppetAccess.Label,
	}
	loginParameters.SetBody(&body)
	response, err := client.Operations.Login(loginParameters)

	if err != nil {
		return "", err
	}

	if response.Payload.Token == "" {
		return "", fmt.Errorf("The response did not contain a token. Rerun with --debug to see full body")
	}

	return response.Payload.Token, nil
}

func (puppetAccess *PuppetAccess) saveToken(token string) error {
	filetoken := filetoken.NewFileToken(puppetAccess.TokenPath)
	return filetoken.Write(token)
}

// Login method
func (puppetAccess *PuppetAccess) Login(print bool) error {

	token, err := puppetAccess.getToken()
	if err != nil {
		return err
	}

	if print {
		fmt.Println(token)
		return nil
	}

	return puppetAccess.saveToken(token)
}
