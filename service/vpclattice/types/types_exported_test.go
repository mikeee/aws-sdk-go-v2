// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
)

func ExampleHeaderMatchType_outputUsage() {
	var union types.HeaderMatchType
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.HeaderMatchTypeMemberContains:
		_ = v.Value // Value is string

	case *types.HeaderMatchTypeMemberExact:
		_ = v.Value // Value is string

	case *types.HeaderMatchTypeMemberPrefix:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string
var _ *string

func ExampleMatcher_outputUsage() {
	var union types.Matcher
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MatcherMemberHttpCode:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExamplePathMatchType_outputUsage() {
	var union types.PathMatchType
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PathMatchTypeMemberExact:
		_ = v.Value // Value is string

	case *types.PathMatchTypeMemberPrefix:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExampleRuleAction_outputUsage() {
	var union types.RuleAction
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RuleActionMemberFixedResponse:
		_ = v.Value // Value is types.FixedResponseAction

	case *types.RuleActionMemberForward:
		_ = v.Value // Value is types.ForwardAction

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ForwardAction
var _ *types.FixedResponseAction

func ExampleRuleMatch_outputUsage() {
	var union types.RuleMatch
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RuleMatchMemberHttpMatch:
		_ = v.Value // Value is types.HttpMatch

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.HttpMatch
