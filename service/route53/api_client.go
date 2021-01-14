// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	route53cust "github.com/aws/aws-sdk-go-v2/service/route53/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"strings"
	"time"
)

const ServiceID = "Route 53"
const ServiceAPIVersion = "2013-04-01"

// Client provides the API client to make operations call for Amazon Route 53.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer aws.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

// WithAPIOptions returns a functional option for setting the Client's APIOptions
// option.
func WithAPIOptions(optFns ...func(*middleware.Stack) error) func(*Options) {
	return func(o *Options) {
		o.APIOptions = append(o.APIOptions, optFns...)
	}
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}
func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	ctx = middleware.ClearStackValues(ctx)
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:        cfg.Region,
		HTTPClient:    cfg.HTTPClient,
		Credentials:   cfg.Credentials,
		APIOptions:    cfg.APIOptions,
		Logger:        cfg.Logger,
		ClientLogMode: cfg.ClientLogMode,
	}
	resolveAWSRetryerProvider(cfg, &opts)
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = awshttp.NewBuildableClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSRetryerProvider(cfg aws.Config, o *Options) {
	if cfg.Retryer == nil {
		return
	}
	o.Retryer = cfg.Retryer()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) error {
	mw := v4.NewSignHTTPRequestMiddleware(v4.SignHTTPRequestMiddlewareOptions{
		CredentialsProvider: o.Credentials,
		Signer:              o.HTTPSignerV4,
		LogSigning:          o.ClientLogMode.IsSigning(),
	})
	return stack.Finalize.Add(mw, middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time, optFns ...func(*v4.SignerOptions)) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = newDefaultV4Signer(*o)
}

func newDefaultV4Signer(o Options) *v4.Signer {
	return v4.NewSigner(func(so *v4.SignerOptions) {
		so.Logger = o.Logger
		so.LogSigning = o.ClientLogMode.IsSigning()
	})
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer:          o.Retryer,
		LogRetryAttempts: o.ClientLogMode.IsRetries(),
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) error {
	return awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return awshttp.AddResponseErrorMiddleware(stack)
}

func addSanitizeURLMiddleware(stack *middleware.Stack) error {
	return route53cust.AddSanitizeURLMiddleware(stack, route53cust.AddSanitizeURLMiddlewareOptions{SanitizeHostedZoneIDInput: sanitizeHostedZoneIDInput})
}

// Check for and split apart Route53 resource IDs, setting only the last piece.
// This allows the output of one operation e.g. foo/1234 to be used as input in
// another operation (e.g. it expects just '1234')
func sanitizeHostedZoneIDInput(input interface{}) error {
	switch i := input.(type) {
	case *ActivateKeySigningKeyInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *AssociateVPCWithHostedZoneInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *ChangeResourceRecordSetsInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *CreateHostedZoneInput:
		if i.DelegationSetId != nil {
			idx := strings.LastIndex(*i.DelegationSetId, `/`)
			v := (*i.DelegationSetId)[idx+1:]
			i.DelegationSetId = &v
		}

	case *CreateKeySigningKeyInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *CreateQueryLoggingConfigInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *CreateReusableDelegationSetInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *CreateTrafficPolicyInstanceInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *CreateVPCAssociationAuthorizationInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *DeactivateKeySigningKeyInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *DeleteHostedZoneInput:
		if i.Id != nil {
			idx := strings.LastIndex(*i.Id, `/`)
			v := (*i.Id)[idx+1:]
			i.Id = &v
		}

	case *DeleteKeySigningKeyInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *DeleteReusableDelegationSetInput:
		if i.Id != nil {
			idx := strings.LastIndex(*i.Id, `/`)
			v := (*i.Id)[idx+1:]
			i.Id = &v
		}

	case *DeleteVPCAssociationAuthorizationInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *DisableHostedZoneDNSSECInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *DisassociateVPCFromHostedZoneInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *EnableHostedZoneDNSSECInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *GetChangeInput:
		if i.Id != nil {
			idx := strings.LastIndex(*i.Id, `/`)
			v := (*i.Id)[idx+1:]
			i.Id = &v
		}

	case *GetDNSSECInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *GetHostedZoneInput:
		if i.Id != nil {
			idx := strings.LastIndex(*i.Id, `/`)
			v := (*i.Id)[idx+1:]
			i.Id = &v
		}

	case *GetHostedZoneLimitInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *GetReusableDelegationSetInput:
		if i.Id != nil {
			idx := strings.LastIndex(*i.Id, `/`)
			v := (*i.Id)[idx+1:]
			i.Id = &v
		}

	case *GetReusableDelegationSetLimitInput:
		if i.DelegationSetId != nil {
			idx := strings.LastIndex(*i.DelegationSetId, `/`)
			v := (*i.DelegationSetId)[idx+1:]
			i.DelegationSetId = &v
		}

	case *ListHostedZonesInput:
		if i.DelegationSetId != nil {
			idx := strings.LastIndex(*i.DelegationSetId, `/`)
			v := (*i.DelegationSetId)[idx+1:]
			i.DelegationSetId = &v
		}

	case *ListHostedZonesByNameInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *ListQueryLoggingConfigsInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *ListResourceRecordSetsInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *ListTrafficPolicyInstancesInput:
		if i.HostedZoneIdMarker != nil {
			idx := strings.LastIndex(*i.HostedZoneIdMarker, `/`)
			v := (*i.HostedZoneIdMarker)[idx+1:]
			i.HostedZoneIdMarker = &v
		}

	case *ListTrafficPolicyInstancesByHostedZoneInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *ListTrafficPolicyInstancesByPolicyInput:
		if i.HostedZoneIdMarker != nil {
			idx := strings.LastIndex(*i.HostedZoneIdMarker, `/`)
			v := (*i.HostedZoneIdMarker)[idx+1:]
			i.HostedZoneIdMarker = &v
		}

	case *ListVPCAssociationAuthorizationsInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *TestDNSAnswerInput:
		if i.HostedZoneId != nil {
			idx := strings.LastIndex(*i.HostedZoneId, `/`)
			v := (*i.HostedZoneId)[idx+1:]
			i.HostedZoneId = &v
		}

	case *UpdateHostedZoneCommentInput:
		if i.Id != nil {
			idx := strings.LastIndex(*i.Id, `/`)
			v := (*i.Id)[idx+1:]
			i.Id = &v
		}

	default:
		break
	}
	return nil
}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}
