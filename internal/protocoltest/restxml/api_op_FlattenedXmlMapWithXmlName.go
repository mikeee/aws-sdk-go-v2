// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Flattened maps with @xmlName
func (c *Client) FlattenedXmlMapWithXmlName(ctx context.Context, params *FlattenedXmlMapWithXmlNameInput, optFns ...func(*Options)) (*FlattenedXmlMapWithXmlNameOutput, error) {
	stack := middleware.NewStack("FlattenedXmlMapWithXmlName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpFlattenedXmlMapWithXmlNameMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opFlattenedXmlMapWithXmlName(options.Region), middleware.Before)

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
			OperationName: "FlattenedXmlMapWithXmlName",
			Err:           err,
		}
	}
	out := result.(*FlattenedXmlMapWithXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FlattenedXmlMapWithXmlNameInput struct {
	MyMap map[string]*string
}

type FlattenedXmlMapWithXmlNameOutput struct {
	MyMap map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpFlattenedXmlMapWithXmlNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpFlattenedXmlMapWithXmlName{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpFlattenedXmlMapWithXmlName{}, middleware.After)
}

func newServiceMetadataMiddleware_opFlattenedXmlMapWithXmlName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "FlattenedXmlMapWithXmlName",
	}
}
