// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplaceentitlementservice

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/client"
	"github.com/aws/aws-sdk-go-v2/aws/client/metadata"
	"github.com/aws/aws-sdk-go-v2/aws/request"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// MarketplaceEntitlementService provides the API operation methods for making requests to
// AWS Marketplace Entitlement Service. See this package's package overview docs
// for details on the service.
//
// MarketplaceEntitlementService methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type MarketplaceEntitlementService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "entitlement.marketplace" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName               // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the MarketplaceEntitlementService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MarketplaceEntitlementService client from just a session.
//     svc := marketplaceentitlementservice.New(mySession)
//
//     // Create a MarketplaceEntitlementService client with additional configuration
//     svc := marketplaceentitlementservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *MarketplaceEntitlementService {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *MarketplaceEntitlementService {
	if len(signingName) == 0 {
		signingName = "aws-marketplace"
	}
	svc := &MarketplaceEntitlementService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2017-01-11",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSMPEntitlementService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a MarketplaceEntitlementService operation and runs any
// custom request initialization.
func (c *MarketplaceEntitlementService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}