// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of image scan findings for your account.
func (c *Client) ListImageScanFindings(ctx context.Context, params *ListImageScanFindingsInput, optFns ...func(*Options)) (*ListImageScanFindingsOutput, error) {
	if params == nil {
		params = &ListImageScanFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImageScanFindings", params, optFns, c.addOperationListImageScanFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImageScanFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImageScanFindingsInput struct {

	// An array of name value pairs that you can use to filter your results. You can
	// use the following filters to streamline results:
	//
	// * imageBuildVersionArn
	//
	// *
	// imagePipelineArn
	//
	// * vulnerabilityId
	//
	// * severity
	//
	// If you don't request a filter,
	// then all findings in your account are listed.
	Filters []types.ImageScanFindingsFilter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImageScanFindingsOutput struct {

	// The image scan findings for your account that meet your request filter criteria.
	Findings []types.ImageScanFinding

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service has'ot included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImageScanFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImageScanFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImageScanFindings{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImageScanFindings(options.Region), middleware.Before); err != nil {
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

// ListImageScanFindingsAPIClient is a client that implements the
// ListImageScanFindings operation.
type ListImageScanFindingsAPIClient interface {
	ListImageScanFindings(context.Context, *ListImageScanFindingsInput, ...func(*Options)) (*ListImageScanFindingsOutput, error)
}

var _ ListImageScanFindingsAPIClient = (*Client)(nil)

// ListImageScanFindingsPaginatorOptions is the paginator options for
// ListImageScanFindings
type ListImageScanFindingsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImageScanFindingsPaginator is a paginator for ListImageScanFindings
type ListImageScanFindingsPaginator struct {
	options   ListImageScanFindingsPaginatorOptions
	client    ListImageScanFindingsAPIClient
	params    *ListImageScanFindingsInput
	nextToken *string
	firstPage bool
}

// NewListImageScanFindingsPaginator returns a new ListImageScanFindingsPaginator
func NewListImageScanFindingsPaginator(client ListImageScanFindingsAPIClient, params *ListImageScanFindingsInput, optFns ...func(*ListImageScanFindingsPaginatorOptions)) *ListImageScanFindingsPaginator {
	if params == nil {
		params = &ListImageScanFindingsInput{}
	}

	options := ListImageScanFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImageScanFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImageScanFindingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImageScanFindings page.
func (p *ListImageScanFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImageScanFindingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListImageScanFindings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImageScanFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImageScanFindings",
	}
}
