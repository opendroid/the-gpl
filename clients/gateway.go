package clients

import (
	"github.com/opendroid/the-gpl/clients/df"
)

// Gateway all external API calls made by the-gpl
type Gateway struct {
	DialogFlowES df.Bot
}

// NewGateway returns a new instance of Gateway
func NewGateway(df df.Bot) Gateway {
	return Gateway{DialogFlowES: df}
}
