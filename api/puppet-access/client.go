package api

import "github.com/puppetlabs/pe-sdk-go/api/puppet-access/client"

//Client is interface to the api client
type Client interface {
	GetClient() (*client.PuppetAccess, error)
}
