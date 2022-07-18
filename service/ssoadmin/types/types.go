// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// These are Amazon Web Services SSO identity store attributes that you can
// configure for use in attributes-based access control (ABAC). You can create
// permissions policies that determine who can access your Amazon Web Services
// resources based upon the configured attribute values. When you enable ABAC and
// specify AccessControlAttributes, Amazon Web Services SSO passes the attribute
// values of the authenticated user into IAM for use in policy evaluation.
type AccessControlAttribute struct {

	// The name of the attribute associated with your identities in your identity
	// source. This is used to map a specified attribute in your identity source with
	// an attribute in Amazon Web Services SSO.
	//
	// This member is required.
	Key *string

	// The value used for mapping a specified attribute to an identity source.
	//
	// This member is required.
	Value *AccessControlAttributeValue

	noSmithyDocumentSerde
}

// The value used for mapping a specified attribute to an identity source. For more
// information, see Attribute mappings
// (https://docs.aws.amazon.com/singlesignon/latest/userguide/attributemappingsconcept.html)
// in the Amazon Web Services Single Sign-On User Guide.
type AccessControlAttributeValue struct {

	// The identity source to use when mapping a specified attribute to Amazon Web
	// Services SSO.
	//
	// This member is required.
	Source []string

	noSmithyDocumentSerde
}

// The assignment that indicates a principal's limited access to a specified Amazon
// Web Services account with a specified permission set. The term principal here
// refers to a user or group that is defined in Amazon Web Services SSO.
type AccountAssignment struct {

	// The identifier of the Amazon Web Services account.
	AccountId *string

	// The ARN of the permission set. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and Amazon Web Services Service Namespaces in the Amazon
	// Web Services General Reference.
	PermissionSetArn *string

	// An identifier for an object in Amazon Web Services SSO, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in Amazon Web Services SSO, see the Amazon
	// Web Services SSO Identity Store API Reference.
	PrincipalId *string

	// The entity type for which the assignment will be created.
	PrincipalType PrincipalType

	noSmithyDocumentSerde
}

// The status of the creation or deletion operation of an assignment that a
// principal needs to access an account.
type AccountAssignmentOperationStatus struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The message that contains an error or exception in case of an operation failure.
	FailureReason *string

	// The ARN of the permission set. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and Amazon Web Services Service Namespaces in the Amazon
	// Web Services General Reference.
	PermissionSetArn *string

	// An identifier for an object in Amazon Web Services SSO, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in Amazon Web Services SSO, see the Amazon
	// Web Services SSO Identity Store API Reference.
	PrincipalId *string

	// The entity type for which the assignment will be created.
	PrincipalType PrincipalType

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	// TargetID is an Amazon Web Services account identifier, typically a 10-12 digit
	// string (For example, 123456789012).
	TargetId *string

	// The entity type for which the assignment will be created.
	TargetType TargetType

	noSmithyDocumentSerde
}

// Provides information about the AccountAssignment creation request.
type AccountAssignmentOperationStatusMetadata struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	noSmithyDocumentSerde
}

// A structure that stores the details of the Amazon Web Services managed IAM
// policy.
type AttachedManagedPolicy struct {

	// The ARN of the Amazon Web Services managed IAM policy. For more information
	// about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service
	// Namespaces in the Amazon Web Services General Reference.
	Arn *string

	// The name of the Amazon Web Services managed IAM policy.
	Name *string

	noSmithyDocumentSerde
}

// Specifies the name and path of the IAM customer managed policy. You must have an
// IAM policy that matches the name and path in each Amazon Web Services account
// where you want to deploy your permission set.
type CustomerManagedPolicyReference struct {

	// The name of the policy document.
	//
	// This member is required.
	Name *string

	// The path for the policy. The default is /. For more information, see Friendly
	// names and paths
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names)
	// in the Identity and Access Management user guide.
	Path *string

	noSmithyDocumentSerde
}

// Specifies the attributes to add to your attribute-based access control (ABAC)
// configuration.
type InstanceAccessControlAttributeConfiguration struct {

	// Lists the attributes that are configured for ABAC in the specified Amazon Web
	// Services SSO instance.
	//
	// This member is required.
	AccessControlAttributes []AccessControlAttribute

	noSmithyDocumentSerde
}

// Provides information about the SSO instance.
type InstanceMetadata struct {

	// The identifier of the identity store that is connected to the SSO instance.
	IdentityStoreId *string

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services
	// Service Namespaces in the Amazon Web Services General Reference.
	InstanceArn *string

	noSmithyDocumentSerde
}

// Filters he operation status list based on the passed attribute value.
type OperationStatusFilter struct {

	// Filters the list operations result based on the status attribute.
	Status StatusValues

	noSmithyDocumentSerde
}

// Specifies the configuration of the Amazon Web Services managed or customer
// managed policy that you want to set as a permissions boundary. Specify either
// CustomerManagedPolicyReference to use the name and path of a customer managed
// policy, or ManagedPolicyArn to use the ARN of an Amazon Web Services managed IAM
// policy. A permissions boundary represents the maximum permissions that any
// policy can grant your role. For more information, see Permissions boundaries for
// IAM entities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
// in the Identity and Access Management User Guide. Policies used as permissions
// boundaries do not provide permissions. You must also attach an IAM policy to the
// role. To learn how the effective permissions for a role are evaluated, see IAM
// JSON policy evaluation logic
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
// in the Identity and Access Management User Guide.
type PermissionsBoundary struct {

	// Specifies the name and path of the IAM customer managed policy. You must have an
	// IAM policy that matches the name and path in each Amazon Web Services account
	// where you want to deploy your permission set.
	CustomerManagedPolicyReference *CustomerManagedPolicyReference

	// The Amazon Web Services managed policy ARN that you want to attach to a
	// permission set as a permissions boundary.
	ManagedPolicyArn *string

	noSmithyDocumentSerde
}

// An entity that contains IAM policies.
type PermissionSet struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The description of the PermissionSet.
	Description *string

	// The name of the permission set.
	Name *string

	// The ARN of the permission set. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and Amazon Web Services Service Namespaces in the Amazon
	// Web Services General Reference.
	PermissionSetArn *string

	// Used to redirect users within the application during the federation
	// authentication process.
	RelayState *string

	// The length of time that the application user sessions are valid for in the
	// ISO-8601 standard.
	SessionDuration *string

	noSmithyDocumentSerde
}

// A structure that is used to provide the status of the provisioning operation for
// a specified permission set.
type PermissionSetProvisioningStatus struct {

	// The identifier of the Amazon Web Services account from which to list the
	// assignments.
	AccountId *string

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The message that contains an error or exception in case of an operation failure.
	FailureReason *string

	// The ARN of the permission set that is being provisioned. For more information
	// about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service
	// Namespaces in the Amazon Web Services General Reference.
	PermissionSetArn *string

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	noSmithyDocumentSerde
}

// Provides information about the permission set provisioning status.
type PermissionSetProvisioningStatusMetadata struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	noSmithyDocumentSerde
}

// A set of key-value pairs that are used to manage the resource. Tags can only be
// applied to permission sets and cannot be applied to corresponding roles that
// Amazon Web Services SSO creates in Amazon Web Services accounts.
type Tag struct {

	// The key for the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
