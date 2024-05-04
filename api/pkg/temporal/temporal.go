package temporal

import (
	"bosca.io/pkg/configuration"
	"context"
)
import "go.temporal.io/sdk/client"

func NewClient(ctx context.Context, cfg *configuration.ClientEndpoints) (client.Client, error) {
	c, err := client.DialContext(ctx, client.Options{
		HostPort: cfg.TemporalApiAddress,
	})
	return c, err
}
