// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches an active index. Use this API to search your documents using query. The
// Query API enables to do faceted search and to filter results based on document
// attributes. It also enables you to provide user context that Amazon Kendra uses
// to enforce document access control in the search results. Amazon Kendra searches
// your index for text content and question and answer (FAQ) content. By default
// the response contains three types of results.
//
// * Relevant passages
//
// * Matching
// FAQs
//
// * Relevant documents
//
// You can specify that the query return only one type
// of result using the QueryResultTypeFilter parameter. Each query returns the 100
// most relevant results.
func (c *Client) Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error) {
	if params == nil {
		params = &QueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Query", params, optFns, c.addOperationQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryInput struct {

	// The identifier of the index to search. The identifier is returned in the
	// response from the CreateIndex API.
	//
	// This member is required.
	IndexId *string

	// Enables filtered searches based on document attributes. You can only provide one
	// attribute filter; however, the AndAllFilters, NotFilter, and OrAllFilters
	// parameters contain a list of other filters. The AttributeFilter parameter
	// enables you to create a set of filtering rules that a document must satisfy to
	// be included in the query results.
	AttributeFilter *types.AttributeFilter

	// Overrides relevance tuning configurations of fields or attributes set at the
	// index level. If you use this API to override the relevance tuning configured at
	// the index level, but there is no relevance tuning configured at the index level,
	// then Amazon Kendra does not apply any relevance tuning. If there is relevance
	// tuning configured at the index level, but you do not use this API to override
	// any relevance tuning in the index, then Amazon Kendra uses the relevance tuning
	// that is configured at the index level. If there is relevance tuning configured
	// for fields at the index level, but you use this API to override only some of
	// these fields, then for the fields you did not override, the importance is set to
	// 1.
	DocumentRelevanceOverrideConfigurations []types.DocumentRelevanceConfiguration

	// An array of documents attributes. Amazon Kendra returns a count for each
	// attribute key specified. This helps your users narrow their search.
	Facets []types.Facet

	// Query results are returned in pages the size of the PageSize parameter. By
	// default, Amazon Kendra returns the first page of results. Use this parameter to
	// get result pages after the first one.
	PageNumber *int32

	// Sets the number of results that are returned in each page of results. The
	// default page size is 10. The maximum number of results returned is 100. If you
	// ask for more than 100 results, only 100 are returned.
	PageSize *int32

	// Sets the type of query. Only results for the specified query type are returned.
	QueryResultTypeFilter types.QueryResultType

	// The input query text for the search. Amazon Kendra truncates queries at 30 token
	// words, which excludes punctuation and stop words. Truncation still applies if
	// you use Boolean or more advanced, complex queries.
	QueryText *string

	// An array of document attributes to include in the response. You can limit the
	// response to include certain document attributes. By default all document
	// attributes are included in the response.
	RequestedDocumentAttributes []string

	// Provides information that determines how the results of the query are sorted.
	// You can set the field that Amazon Kendra should sort the results on, and specify
	// whether the results should be sorted in ascending or descending order. In the
	// case of ties in sorting the results, the results are sorted by relevance. If you
	// don't provide sorting configuration, the results are sorted by the relevance
	// that Amazon Kendra determines for the result.
	SortingConfiguration *types.SortingConfiguration

	// Enables suggested spell corrections for queries.
	SpellCorrectionConfiguration *types.SpellCorrectionConfiguration

	// The user context token or user and group information.
	UserContext *types.UserContext

	// Provides an identifier for a specific user. The VisitorId should be a unique
	// identifier, such as a GUID. Don't use personally identifiable information, such
	// as the user's email address, as the VisitorId.
	VisitorId *string

	noSmithyDocumentSerde
}

type QueryOutput struct {

	// Contains the facet results. A FacetResult contains the counts for each attribute
	// key that was specified in the Facets input parameter.
	FacetResults []types.FacetResult

	// The list of featured result items. Featured results are displayed at the top of
	// the search results page, placed above all other results for certain queries. If
	// there's an exact match of a query, then certain documents are featured in the
	// search results.
	FeaturedResultsItems []types.FeaturedResultsItem

	// The identifier for the search. You use QueryId to identify the search when using
	// the feedback API.
	QueryId *string

	// The results of the search.
	ResultItems []types.QueryResultItem

	// A list of information related to suggested spell corrections for a query.
	SpellCorrectedQueries []types.SpellCorrectedQuery

	// The total number of items found by the search; however, you can only retrieve up
	// to 100 items. For example, if the search found 192 items, you can only retrieve
	// the first 100 of the items.
	TotalNumberOfResults *int32

	// A list of warning codes and their messages on problems with your query. Amazon
	// Kendra currently only supports one type of warning, which is a warning on
	// invalid syntax used in the query. For examples of invalid query syntax, see
	// Searching with advanced query syntax
	// (https://docs.aws.amazon.com/kendra/latest/dg/searching-example.html#searching-index-query-syntax).
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQuery{}, middleware.After)
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
	if err = addOpQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "Query",
	}
}
