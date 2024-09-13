package client

import (
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/crowdstrike/gofalcon/falcon/client"
)

type CrowdstrikeClient struct {
	Name        string
	CrowdStrike *client.CrowdStrikeAPISpecification
	Backend     state.Client
}
type CsClients struct {
	items []CrowdstrikeClient
}
