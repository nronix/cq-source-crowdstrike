package client

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/rs/zerolog"
)

type Client struct {
	logger   zerolog.Logger
	Spec     Spec
	Account  CrowdstrikeClient
	accounts *CsClients
}

func (c *Client) ID() string {
	return fmt.Sprintf("Crowdstrike:%s", c.Account.Name)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, spec *Spec, bk state.Client) (Client, error) {
	//falconClientID := os.Getenv("FALCON_CLIENT_ID")
	//falconSecret := os.Getenv("FALCON_SECRET")
	var clients = CsClients{}
	for _, acc := range spec.FALCON {
		logger.Info().Msg(fmt.Sprintf("Access key %w Secret %w cid %w", acc.ClientId, acc.ClientSecret, acc.MemberCid))
		if acc.MemberCid != "" {
			fc, err := falcon.NewClient(&falcon.ApiConfig{
				ClientId:     acc.ClientId,
				ClientSecret: acc.ClientSecret,
				MemberCID:    acc.MemberCid,
				Cloud:        falcon.Cloud(acc.ClientCloud),
				Debug:        false,
				Context:      ctx,
			})
			if err != nil {
				logger.Error().Msg(fmt.Sprintf("could not auth due to : %w", err))
			}
			clients.items = append(clients.items, CrowdstrikeClient{Name: acc.Name, CrowdStrike: fc, Backend: bk})
		} else {
			fc, err := falcon.NewClient(&falcon.ApiConfig{
				ClientId:     acc.ClientId,
				ClientSecret: acc.ClientSecret,

				Cloud:   falcon.Cloud(acc.ClientCloud),
				Debug:   false,
				Context: ctx,
			})
			if err != nil {
				logger.Error().Msg(fmt.Sprintf("No member CId , could not auth due to : %w", err))
			}
			clients.items = append(clients.items, CrowdstrikeClient{Name: acc.Name, CrowdStrike: fc, Backend: bk})
		}

	}

	return Client{
		logger:   logger,
		Spec:     *spec,
		accounts: &clients,
	}, nil
}

func (c *Client) WithAccount(account CrowdstrikeClient) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("account", account.Name).Logger()
	newC.Account = account
	return &newC
}
