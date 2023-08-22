// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new kdb cluster.
func (c *Client) CreateKxCluster(ctx context.Context, params *CreateKxClusterInput, optFns ...func(*Options)) (*CreateKxClusterOutput, error) {
	if params == nil {
		params = &CreateKxClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKxCluster", params, optFns, c.addOperationCreateKxClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKxClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKxClusterInput struct {

	// The number of availability zones you want to assign per cluster. This can be
	// one of the following
	//   - SINGLE – Assigns one availability zone per cluster.
	//   - MULTI – Assigns all the availability zones per cluster.
	//
	// This member is required.
	AzMode types.KxAzMode

	// A structure for the metadata of a cluster. It includes information like the
	// CPUs needed, memory of instances, and number of instances.
	//
	// This member is required.
	CapacityConfiguration *types.CapacityConfiguration

	// A unique name for the cluster that you want to create.
	//
	// This member is required.
	ClusterName *string

	// Specifies the type of KDB database that is being created. The following types
	// are available:
	//   - HDB – A Historical Database. The data is only accessible with read-only
	//   permissions from one of the FinSpace managed kdb databases mounted to the
	//   cluster.
	//   - RDB – A Realtime Database. This type of database captures all the data from
	//   a ticker plant and stores it in memory until the end of day, after which it
	//   writes all of its data to a disk and reloads the HDB. This cluster type requires
	//   local storage for temporary storage of data during the savedown process. If you
	//   specify this field in your request, you must provide the
	//   savedownStorageConfiguration parameter.
	//   - GATEWAY – A gateway cluster allows you to access data across processes in
	//   kdb systems. It allows you to create your own routing logic using the
	//   initialization scripts and custom code. This type of cluster does not require a
	//   writable local storage.
	//
	// This member is required.
	ClusterType types.KxClusterType

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	// The version of FinSpace managed kdb to run.
	//
	// This member is required.
	ReleaseLabel *string

	// The configuration based on which FinSpace will scale in or scale out nodes in
	// your cluster.
	AutoScalingConfiguration *types.AutoScalingConfiguration

	// The availability zone identifiers for the requested regions.
	AvailabilityZoneId *string

	// The configurations for a read only cache storage associated with a cluster.
	// This cache will be stored as an FSx Lustre that reads from the S3 store.
	CacheStorageConfigurations []types.KxCacheStorageConfiguration

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	// A description of the cluster.
	ClusterDescription *string

	// The details of the custom code that you want to use inside a cluster when
	// analyzing a data. It consists of the S3 source bucket, location, S3 object
	// version, and the relative path from where the custom code is loaded into the
	// cluster.
	Code *types.CodeConfiguration

	// Defines the key-value pairs to make them available inside the cluster.
	CommandLineArguments []types.KxCommandLineArgument

	// A list of databases that will be available for querying.
	Databases []types.KxDatabaseConfiguration

	// An IAM role that defines a set of permissions associated with a cluster. These
	// permissions are assumed when a cluster attempts to access another cluster.
	ExecutionRole *string

	// Specifies a Q program that will be run at launch of a cluster. It is a relative
	// path within .zip file that contains the custom code, which will be loaded on the
	// cluster. It must include the file name itself. For example, somedir/init.q .
	InitializationScript *string

	// The size and type of the temporary storage that is used to hold data during the
	// savedown process. This parameter is required when you choose clusterType as
	// RDB. All the data written to this storage space is lost when the cluster node is
	// restarted.
	SavedownStorageConfiguration *types.KxSavedownStorageConfiguration

	// A list of key-value pairs to label the cluster. You can add up to 50 tags to a
	// cluster.
	Tags map[string]string

	// Configuration details about the network where the Privatelink endpoint of the
	// cluster resides.
	VpcConfiguration *types.VpcConfiguration

	noSmithyDocumentSerde
}

type CreateKxClusterOutput struct {

	// The configuration based on which FinSpace will scale in or scale out nodes in
	// your cluster.
	AutoScalingConfiguration *types.AutoScalingConfiguration

	// The availability zone identifiers for the requested regions.
	AvailabilityZoneId *string

	// The number of availability zones you want to assign per cluster. This can be
	// one of the following
	//   - SINGLE – Assigns one availability zone per cluster.
	//   - MULTI – Assigns all the availability zones per cluster.
	AzMode types.KxAzMode

	// The configurations for a read only cache storage associated with a cluster.
	// This cache will be stored as an FSx Lustre that reads from the S3 store.
	CacheStorageConfigurations []types.KxCacheStorageConfiguration

	// A structure for the metadata of a cluster. It includes information like the
	// CPUs needed, memory of instances, and number of instances.
	CapacityConfiguration *types.CapacityConfiguration

	// A description of the cluster.
	ClusterDescription *string

	// A unique name for the cluster.
	ClusterName *string

	// Specifies the type of KDB database that is being created. The following types
	// are available:
	//   - HDB – A Historical Database. The data is only accessible with read-only
	//   permissions from one of the FinSpace managed kdb databases mounted to the
	//   cluster.
	//   - RDB – A Realtime Database. This type of database captures all the data from
	//   a ticker plant and stores it in memory until the end of day, after which it
	//   writes all of its data to a disk and reloads the HDB. This cluster type requires
	//   local storage for temporary storage of data during the savedown process. If you
	//   specify this field in your request, you must provide the
	//   savedownStorageConfiguration parameter.
	//   - GATEWAY – A gateway cluster allows you to access data across processes in
	//   kdb systems. It allows you to create your own routing logic using the
	//   initialization scripts and custom code. This type of cluster does not require a
	//   writable local storage.
	ClusterType types.KxClusterType

	// The details of the custom code that you want to use inside a cluster when
	// analyzing a data. It consists of the S3 source bucket, location, S3 object
	// version, and the relative path from where the custom code is loaded into the
	// cluster.
	Code *types.CodeConfiguration

	// Defines the key-value pairs to make them available inside the cluster.
	CommandLineArguments []types.KxCommandLineArgument

	// The timestamp at which the cluster was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	// A list of databases that will be available for querying.
	Databases []types.KxDatabaseConfiguration

	// A unique identifier for the kdb environment.
	EnvironmentId *string

	// An IAM role that defines a set of permissions associated with a cluster. These
	// permissions are assumed when a cluster attempts to access another cluster.
	ExecutionRole *string

	// Specifies a Q program that will be run at launch of a cluster. It is a relative
	// path within .zip file that contains the custom code, which will be loaded on the
	// cluster. It must include the file name itself. For example, somedir/init.q .
	InitializationScript *string

	// The last time that the cluster was modified. The value is determined as epoch
	// time in milliseconds. For example, the value for Monday, November 1, 2021
	// 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	// A version of the FinSpace managed kdb to run.
	ReleaseLabel *string

	// The size and type of the temporary storage that is used to hold data during the
	// savedown process. This parameter is required when you choose clusterType as
	// RDB. All the data written to this storage space is lost when the cluster node is
	// restarted.
	SavedownStorageConfiguration *types.KxSavedownStorageConfiguration

	// The status of cluster creation.
	//   - PENDING – The cluster is pending creation.
	//   - CREATING – The cluster creation process is in progress.
	//   - CREATE_FAILED – The cluster creation process has failed.
	//   - RUNNING – The cluster creation process is running.
	//   - UPDATING – The cluster is in the process of being updated.
	//   - DELETING – The cluster is in the process of being deleted.
	//   - DELETED – The cluster has been deleted.
	//   - DELETE_FAILED – The cluster failed to delete.
	Status types.KxClusterStatus

	// The error message when a failed state occurs.
	StatusReason *string

	// Configuration details about the network where the Privatelink endpoint of the
	// cluster resides.
	VpcConfiguration *types.VpcConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKxClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKxCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKxCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateKxClusterResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateKxClusterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKxClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKxCluster(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateKxCluster struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKxCluster) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKxCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKxClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKxClusterInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateKxClusterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKxCluster{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKxCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace",
		OperationName: "CreateKxCluster",
	}
}

type opCreateKxClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateKxClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateKxClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "finspace"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "finspace"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("finspace")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateKxClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateKxClusterResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
