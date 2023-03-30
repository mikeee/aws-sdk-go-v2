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

// List the Packages that are associated with an Image Build Version, as determined
// by Amazon Web Services Systems Manager Inventory at build time.
func (c *Client) ListImagePackages(ctx context.Context, params *ListImagePackagesInput, optFns ...func(*Options)) (*ListImagePackagesOutput, error) {
	if params == nil {
		params = &ListImagePackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImagePackages", params, optFns, c.addOperationListImagePackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImagePackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagePackagesInput struct {

	// Filter results for the ListImagePackages request by the Image Build Version ARN
	//
	// This member is required.
	ImageBuildVersionArn *string

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImagePackagesOutput struct {

	// The list of Image Packages returned in the response.
	ImagePackageList []types.ImagePackage

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

func (c *Client) addOperationListImagePackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImagePackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImagePackages{}, middleware.After)
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
	if err = addOpListImagePackagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImagePackages(options.Region), middleware.Before); err != nil {
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

// ListImagePackagesAPIClient is a client that implements the ListImagePackages
// operation.
type ListImagePackagesAPIClient interface {
	ListImagePackages(context.Context, *ListImagePackagesInput, ...func(*Options)) (*ListImagePackagesOutput, error)
}

var _ ListImagePackagesAPIClient = (*Client)(nil)

// ListImagePackagesPaginatorOptions is the paginator options for ListImagePackages
type ListImagePackagesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImagePackagesPaginator is a paginator for ListImagePackages
type ListImagePackagesPaginator struct {
	options   ListImagePackagesPaginatorOptions
	client    ListImagePackagesAPIClient
	params    *ListImagePackagesInput
	nextToken *string
	firstPage bool
}

// NewListImagePackagesPaginator returns a new ListImagePackagesPaginator
func NewListImagePackagesPaginator(client ListImagePackagesAPIClient, params *ListImagePackagesInput, optFns ...func(*ListImagePackagesPaginatorOptions)) *ListImagePackagesPaginator {
	if params == nil {
		params = &ListImagePackagesInput{}
	}

	options := ListImagePackagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImagePackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImagePackagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImagePackages page.
func (p *ListImagePackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImagePackagesOutput, error) {
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

	result, err := p.client.ListImagePackages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImagePackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImagePackages",
	}
}
