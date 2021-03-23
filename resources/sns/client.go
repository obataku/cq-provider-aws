package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log  hclog.Logger
	accountID string
	region    string
	svc       *sns.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       sns.NewFromConfig(awsConfig, func(options *sns.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(ctx context.Context, resource string, config interface{}) error {
	switch resource {
	case "subscriptions":
		return c.subscriptions(ctx, config)
	case "topics":
		return c.topics(ctx, config)
	default:
		return fmt.Errorf("unsupported resource sns.%s", resource)
	}
}
