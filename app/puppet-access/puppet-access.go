package app

import (
	api "github.com/puppetlabs/pe-sdk-go/api/puppet-access"
)

// PuppetAccess interface
type PuppetAccess struct {
	Username  string
	Password  string
	Lifetime  string
	URL       string
	Label     string
	Cacert    string
	TokenPath string
	Client    api.Client
}

// NewWithConfig creates a puppet code application with configuration
func NewWithConfig(username, password, lifetime, url, label, cacert, tokenPath string) *PuppetAccess {
	return &PuppetAccess{
		Username:  username,
		Password:  password,
		Lifetime:  lifetime,
		URL:       url,
		Label:     label,
		Cacert:    cacert,
		TokenPath: tokenPath,
		Client:    api.NewClient(username, password, lifetime, url, label, cacert),
	}
}

// NewWithMinimalConfig creates a puppet code application with minimal configuration
func NewWithMinimalConfig(username, password, url, cacert, tokenPath string) *PuppetAccess {
	return &PuppetAccess{
		Username:  username,
		Password:  password,
		URL:       url,
		Cacert:    cacert,
		TokenPath: tokenPath,
		Client:    api.NewClient(username, password, "", url, "", cacert),
	}
}

// New creates an unconfigured puppet-db application
func New() *PuppetAccess {
	return &PuppetAccess{
		Client: api.NewClient("", "", "", "", "", ""),
	}
}
