// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a workload share. The owner of a workload can share it with other Amazon
// Web Services accounts and users in the same Amazon Web Services Region. Shared
// access to a workload is not removed until the workload invitation is deleted. If
// you share a workload with an organization or OU, all accounts in the
// organization or OU are granted access to the workload. For more information, see
// Sharing a workload
// (https://docs.aws.amazon.com/wellarchitected/latest/userguide/workloads-sharing.html)
// in the Well-Architected Tool User Guide.
func (c *Client) CreateWorkloadShare(ctx context.Context, params *CreateWorkloadShareInput, optFns ...func(*Options)) (*CreateWorkloadShareOutput, error) {
	if params == nil {
		params = &CreateWorkloadShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkloadShare", params, optFns, c.addOperationCreateWorkloadShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkloadShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Create Workload Share
type CreateWorkloadShareInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after the original request has completed successfully, the result of the
	// original request is returned. This token is listed as required, however, if you
	// do not specify it, the Amazon Web Services SDKs automatically generate one for
	// you. If you are not using the Amazon Web Services SDK or the CLI, you must
	// provide this token or the request will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// Permission granted on a workload share.
	//
	// This member is required.
	PermissionType types.PermissionType

	// The Amazon Web Services account ID, IAM role, organization ID, or organizational
	// unit (OU) ID with which the workload is shared.
	//
	// This member is required.
	SharedWith *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	//
	// This member is required.
	WorkloadId *string

	noSmithyDocumentSerde
}

// Input for Create Workload Share
type CreateWorkloadShareOutput struct {

	// The ID associated with the workload share.
	ShareId *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkloadShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkloadShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkloadShare{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateWorkloadShareMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorkloadShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkloadShare(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorkloadShare struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorkloadShare) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorkloadShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorkloadShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorkloadShareInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateWorkloadShareMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorkloadShare{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorkloadShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "CreateWorkloadShare",
	}
}
