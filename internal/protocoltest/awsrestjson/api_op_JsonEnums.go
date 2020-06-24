// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes enums as top level properties, in lists, sets, and maps.
func (c *Client) JsonEnums(ctx context.Context, params *JsonEnumsInput, optFns ...func(*Options)) (*JsonEnumsOutput, error) {
	stack := middleware.NewStack("JsonEnums", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "JsonEnums",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	stack.Serialize.Add(&awsRestjson1_serializeOpJsonEnums{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonEnums{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "JsonEnums",
			Err:           err,
		}
	}
	out := result.(*JsonEnumsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonEnumsInput struct {
	FooEnum1    types.FooEnum
	FooEnum2    types.FooEnum
	FooEnum3    types.FooEnum
	FooEnumList []types.FooEnum
	FooEnumMap  map[string]types.FooEnum
	FooEnumSet  []types.FooEnum
}

type JsonEnumsOutput struct {
	FooEnum1    types.FooEnum
	FooEnum2    types.FooEnum
	FooEnum3    types.FooEnum
	FooEnumList []types.FooEnum
	FooEnumMap  map[string]types.FooEnum
	FooEnumSet  []types.FooEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
