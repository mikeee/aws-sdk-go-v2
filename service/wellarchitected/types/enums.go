// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdditionalResourceType string

// Enum values for AdditionalResourceType
const (
	AdditionalResourceTypeHelpfulResource AdditionalResourceType = "HELPFUL_RESOURCE"
	AdditionalResourceTypeImprovementPlan AdditionalResourceType = "IMPROVEMENT_PLAN"
)

// Values returns all known values for AdditionalResourceType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdditionalResourceType) Values() []AdditionalResourceType {
	return []AdditionalResourceType{
		"HELPFUL_RESOURCE",
		"IMPROVEMENT_PLAN",
	}
}

type AnswerReason string

// Enum values for AnswerReason
const (
	AnswerReasonOutOfScope              AnswerReason = "OUT_OF_SCOPE"
	AnswerReasonBusinessPriorities      AnswerReason = "BUSINESS_PRIORITIES"
	AnswerReasonArchitectureConstraints AnswerReason = "ARCHITECTURE_CONSTRAINTS"
	AnswerReasonOther                   AnswerReason = "OTHER"
	AnswerReasonNone                    AnswerReason = "NONE"
)

// Values returns all known values for AnswerReason. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AnswerReason) Values() []AnswerReason {
	return []AnswerReason{
		"OUT_OF_SCOPE",
		"BUSINESS_PRIORITIES",
		"ARCHITECTURE_CONSTRAINTS",
		"OTHER",
		"NONE",
	}
}

type CheckFailureReason string

// Enum values for CheckFailureReason
const (
	CheckFailureReasonAssumeRoleError        CheckFailureReason = "ASSUME_ROLE_ERROR"
	CheckFailureReasonAccessDenied           CheckFailureReason = "ACCESS_DENIED"
	CheckFailureReasonUnknownError           CheckFailureReason = "UNKNOWN_ERROR"
	CheckFailureReasonPremiumSupportRequired CheckFailureReason = "PREMIUM_SUPPORT_REQUIRED"
)

// Values returns all known values for CheckFailureReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CheckFailureReason) Values() []CheckFailureReason {
	return []CheckFailureReason{
		"ASSUME_ROLE_ERROR",
		"ACCESS_DENIED",
		"UNKNOWN_ERROR",
		"PREMIUM_SUPPORT_REQUIRED",
	}
}

type CheckProvider string

// Enum values for CheckProvider
const (
	CheckProviderTrustedAdvisor CheckProvider = "TRUSTED_ADVISOR"
)

// Values returns all known values for CheckProvider. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CheckProvider) Values() []CheckProvider {
	return []CheckProvider{
		"TRUSTED_ADVISOR",
	}
}

type CheckStatus string

// Enum values for CheckStatus
const (
	CheckStatusOkay         CheckStatus = "OKAY"
	CheckStatusWarning      CheckStatus = "WARNING"
	CheckStatusError        CheckStatus = "ERROR"
	CheckStatusNotAvailable CheckStatus = "NOT_AVAILABLE"
	CheckStatusFetchFailed  CheckStatus = "FETCH_FAILED"
)

// Values returns all known values for CheckStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CheckStatus) Values() []CheckStatus {
	return []CheckStatus{
		"OKAY",
		"WARNING",
		"ERROR",
		"NOT_AVAILABLE",
		"FETCH_FAILED",
	}
}

type ChoiceReason string

// Enum values for ChoiceReason
const (
	ChoiceReasonOutOfScope              ChoiceReason = "OUT_OF_SCOPE"
	ChoiceReasonBusinessPriorities      ChoiceReason = "BUSINESS_PRIORITIES"
	ChoiceReasonArchitectureConstraints ChoiceReason = "ARCHITECTURE_CONSTRAINTS"
	ChoiceReasonOther                   ChoiceReason = "OTHER"
	ChoiceReasonNone                    ChoiceReason = "NONE"
)

// Values returns all known values for ChoiceReason. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChoiceReason) Values() []ChoiceReason {
	return []ChoiceReason{
		"OUT_OF_SCOPE",
		"BUSINESS_PRIORITIES",
		"ARCHITECTURE_CONSTRAINTS",
		"OTHER",
		"NONE",
	}
}

type ChoiceStatus string

// Enum values for ChoiceStatus
const (
	ChoiceStatusSelected      ChoiceStatus = "SELECTED"
	ChoiceStatusNotApplicable ChoiceStatus = "NOT_APPLICABLE"
	ChoiceStatusUnselected    ChoiceStatus = "UNSELECTED"
)

// Values returns all known values for ChoiceStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChoiceStatus) Values() []ChoiceStatus {
	return []ChoiceStatus{
		"SELECTED",
		"NOT_APPLICABLE",
		"UNSELECTED",
	}
}

type DifferenceStatus string

// Enum values for DifferenceStatus
const (
	DifferenceStatusUpdated DifferenceStatus = "UPDATED"
	DifferenceStatusNew     DifferenceStatus = "NEW"
	DifferenceStatusDeleted DifferenceStatus = "DELETED"
)

// Values returns all known values for DifferenceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DifferenceStatus) Values() []DifferenceStatus {
	return []DifferenceStatus{
		"UPDATED",
		"NEW",
		"DELETED",
	}
}

type ImportLensStatus string

// Enum values for ImportLensStatus
const (
	ImportLensStatusInProgress ImportLensStatus = "IN_PROGRESS"
	ImportLensStatusComplete   ImportLensStatus = "COMPLETE"
	ImportLensStatusError      ImportLensStatus = "ERROR"
)

// Values returns all known values for ImportLensStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportLensStatus) Values() []ImportLensStatus {
	return []ImportLensStatus{
		"IN_PROGRESS",
		"COMPLETE",
		"ERROR",
	}
}

type LensStatus string

// Enum values for LensStatus
const (
	LensStatusCurrent    LensStatus = "CURRENT"
	LensStatusNotCurrent LensStatus = "NOT_CURRENT"
	LensStatusDeprecated LensStatus = "DEPRECATED"
	LensStatusDeleted    LensStatus = "DELETED"
	LensStatusUnshared   LensStatus = "UNSHARED"
)

// Values returns all known values for LensStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LensStatus) Values() []LensStatus {
	return []LensStatus{
		"CURRENT",
		"NOT_CURRENT",
		"DEPRECATED",
		"DELETED",
		"UNSHARED",
	}
}

type LensStatusType string

// Enum values for LensStatusType
const (
	LensStatusTypeAll       LensStatusType = "ALL"
	LensStatusTypeDraft     LensStatusType = "DRAFT"
	LensStatusTypePublished LensStatusType = "PUBLISHED"
)

// Values returns all known values for LensStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LensStatusType) Values() []LensStatusType {
	return []LensStatusType{
		"ALL",
		"DRAFT",
		"PUBLISHED",
	}
}

type LensType string

// Enum values for LensType
const (
	LensTypeAwsOfficial  LensType = "AWS_OFFICIAL"
	LensTypeCustomShared LensType = "CUSTOM_SHARED"
	LensTypeCustomSelf   LensType = "CUSTOM_SELF"
)

// Values returns all known values for LensType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LensType) Values() []LensType {
	return []LensType{
		"AWS_OFFICIAL",
		"CUSTOM_SHARED",
		"CUSTOM_SELF",
	}
}

type MetricType string

// Enum values for MetricType
const (
	MetricTypeWorkload MetricType = "WORKLOAD"
)

// Values returns all known values for MetricType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MetricType) Values() []MetricType {
	return []MetricType{
		"WORKLOAD",
	}
}

type NotificationType string

// Enum values for NotificationType
const (
	NotificationTypeLensVersionUpgraded   NotificationType = "LENS_VERSION_UPGRADED"
	NotificationTypeLensVersionDeprecated NotificationType = "LENS_VERSION_DEPRECATED"
)

// Values returns all known values for NotificationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationType) Values() []NotificationType {
	return []NotificationType{
		"LENS_VERSION_UPGRADED",
		"LENS_VERSION_DEPRECATED",
	}
}

type OrganizationSharingStatus string

// Enum values for OrganizationSharingStatus
const (
	OrganizationSharingStatusEnabled  OrganizationSharingStatus = "ENABLED"
	OrganizationSharingStatusDisabled OrganizationSharingStatus = "DISABLED"
)

// Values returns all known values for OrganizationSharingStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationSharingStatus) Values() []OrganizationSharingStatus {
	return []OrganizationSharingStatus{
		"ENABLED",
		"DISABLED",
	}
}

type PermissionType string

// Enum values for PermissionType
const (
	PermissionTypeReadonly    PermissionType = "READONLY"
	PermissionTypeContributor PermissionType = "CONTRIBUTOR"
)

// Values returns all known values for PermissionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PermissionType) Values() []PermissionType {
	return []PermissionType{
		"READONLY",
		"CONTRIBUTOR",
	}
}

type ReportFormat string

// Enum values for ReportFormat
const (
	ReportFormatPdf  ReportFormat = "PDF"
	ReportFormatJson ReportFormat = "JSON"
)

// Values returns all known values for ReportFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportFormat) Values() []ReportFormat {
	return []ReportFormat{
		"PDF",
		"JSON",
	}
}

type Risk string

// Enum values for Risk
const (
	RiskUnanswered    Risk = "UNANSWERED"
	RiskHigh          Risk = "HIGH"
	RiskMedium        Risk = "MEDIUM"
	RiskNone          Risk = "NONE"
	RiskNotApplicable Risk = "NOT_APPLICABLE"
)

// Values returns all known values for Risk. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Risk) Values() []Risk {
	return []Risk{
		"UNANSWERED",
		"HIGH",
		"MEDIUM",
		"NONE",
		"NOT_APPLICABLE",
	}
}

type ShareInvitationAction string

// Enum values for ShareInvitationAction
const (
	ShareInvitationActionAccept ShareInvitationAction = "ACCEPT"
	ShareInvitationActionReject ShareInvitationAction = "REJECT"
)

// Values returns all known values for ShareInvitationAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShareInvitationAction) Values() []ShareInvitationAction {
	return []ShareInvitationAction{
		"ACCEPT",
		"REJECT",
	}
}

type ShareResourceType string

// Enum values for ShareResourceType
const (
	ShareResourceTypeWorkload ShareResourceType = "WORKLOAD"
	ShareResourceTypeLens     ShareResourceType = "LENS"
)

// Values returns all known values for ShareResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShareResourceType) Values() []ShareResourceType {
	return []ShareResourceType{
		"WORKLOAD",
		"LENS",
	}
}

type ShareStatus string

// Enum values for ShareStatus
const (
	ShareStatusAccepted    ShareStatus = "ACCEPTED"
	ShareStatusRejected    ShareStatus = "REJECTED"
	ShareStatusPending     ShareStatus = "PENDING"
	ShareStatusRevoked     ShareStatus = "REVOKED"
	ShareStatusExpired     ShareStatus = "EXPIRED"
	ShareStatusAssociating ShareStatus = "ASSOCIATING"
	ShareStatusAssociated  ShareStatus = "ASSOCIATED"
	ShareStatusFailed      ShareStatus = "FAILED"
)

// Values returns all known values for ShareStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ShareStatus) Values() []ShareStatus {
	return []ShareStatus{
		"ACCEPTED",
		"REJECTED",
		"PENDING",
		"REVOKED",
		"EXPIRED",
		"ASSOCIATING",
		"ASSOCIATED",
		"FAILED",
	}
}

type TrustedAdvisorIntegrationStatus string

// Enum values for TrustedAdvisorIntegrationStatus
const (
	TrustedAdvisorIntegrationStatusEnabled  TrustedAdvisorIntegrationStatus = "ENABLED"
	TrustedAdvisorIntegrationStatusDisabled TrustedAdvisorIntegrationStatus = "DISABLED"
)

// Values returns all known values for TrustedAdvisorIntegrationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TrustedAdvisorIntegrationStatus) Values() []TrustedAdvisorIntegrationStatus {
	return []TrustedAdvisorIntegrationStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}

type WorkloadEnvironment string

// Enum values for WorkloadEnvironment
const (
	WorkloadEnvironmentProduction    WorkloadEnvironment = "PRODUCTION"
	WorkloadEnvironmentPreproduction WorkloadEnvironment = "PREPRODUCTION"
)

// Values returns all known values for WorkloadEnvironment. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkloadEnvironment) Values() []WorkloadEnvironment {
	return []WorkloadEnvironment{
		"PRODUCTION",
		"PREPRODUCTION",
	}
}

type WorkloadImprovementStatus string

// Enum values for WorkloadImprovementStatus
const (
	WorkloadImprovementStatusNotApplicable    WorkloadImprovementStatus = "NOT_APPLICABLE"
	WorkloadImprovementStatusNotStarted       WorkloadImprovementStatus = "NOT_STARTED"
	WorkloadImprovementStatusInProgress       WorkloadImprovementStatus = "IN_PROGRESS"
	WorkloadImprovementStatusComplete         WorkloadImprovementStatus = "COMPLETE"
	WorkloadImprovementStatusRiskAcknowledged WorkloadImprovementStatus = "RISK_ACKNOWLEDGED"
)

// Values returns all known values for WorkloadImprovementStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkloadImprovementStatus) Values() []WorkloadImprovementStatus {
	return []WorkloadImprovementStatus{
		"NOT_APPLICABLE",
		"NOT_STARTED",
		"IN_PROGRESS",
		"COMPLETE",
		"RISK_ACKNOWLEDGED",
	}
}
