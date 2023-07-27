// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a warm pool and its instances. For more information, see
// Warm pools for Amazon EC2 Auto Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeWarmPool(ctx context.Context, params *DescribeWarmPoolInput, optFns ...func(*Options)) (*DescribeWarmPoolOutput, error) {
	if params == nil {
		params = &DescribeWarmPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWarmPool", params, optFns, c.addOperationDescribeWarmPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWarmPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWarmPoolInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of instances to return with this call. The maximum value is
	// 50 .
	MaxRecords *int32

	// The token for the next set of instances to return. (You received this token
	// from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeWarmPoolOutput struct {

	// The instances that are currently in the warm pool.
	Instances []types.Instance

	// This string indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// The warm pool configuration details.
	WarmPoolConfiguration *types.WarmPoolConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWarmPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeWarmPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeWarmPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeWarmPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWarmPool(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeWarmPoolAPIClient is a client that implements the DescribeWarmPool
// operation.
type DescribeWarmPoolAPIClient interface {
	DescribeWarmPool(context.Context, *DescribeWarmPoolInput, ...func(*Options)) (*DescribeWarmPoolOutput, error)
}

var _ DescribeWarmPoolAPIClient = (*Client)(nil)

// DescribeWarmPoolPaginatorOptions is the paginator options for DescribeWarmPool
type DescribeWarmPoolPaginatorOptions struct {
	// The maximum number of instances to return with this call. The maximum value is
	// 50 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeWarmPoolPaginator is a paginator for DescribeWarmPool
type DescribeWarmPoolPaginator struct {
	options   DescribeWarmPoolPaginatorOptions
	client    DescribeWarmPoolAPIClient
	params    *DescribeWarmPoolInput
	nextToken *string
	firstPage bool
}

// NewDescribeWarmPoolPaginator returns a new DescribeWarmPoolPaginator
func NewDescribeWarmPoolPaginator(client DescribeWarmPoolAPIClient, params *DescribeWarmPoolInput, optFns ...func(*DescribeWarmPoolPaginatorOptions)) *DescribeWarmPoolPaginator {
	if params == nil {
		params = &DescribeWarmPoolInput{}
	}

	options := DescribeWarmPoolPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeWarmPoolPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeWarmPoolPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeWarmPool page.
func (p *DescribeWarmPoolPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeWarmPoolOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeWarmPool(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeWarmPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeWarmPool",
	}
}
