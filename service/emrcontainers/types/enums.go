// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ContainerProviderType string

// Enum values for ContainerProviderType
const (
	ContainerProviderTypeEks ContainerProviderType = "EKS"
)

// Values returns all known values for ContainerProviderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerProviderType) Values() []ContainerProviderType {
	return []ContainerProviderType{
		"EKS",
	}
}

type EndpointState string

// Enum values for EndpointState
const (
	EndpointStateCreating             EndpointState = "CREATING"
	EndpointStateActive               EndpointState = "ACTIVE"
	EndpointStateTerminating          EndpointState = "TERMINATING"
	EndpointStateTerminated           EndpointState = "TERMINATED"
	EndpointStateTerminatedWithErrors EndpointState = "TERMINATED_WITH_ERRORS"
)

// Values returns all known values for EndpointState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EndpointState) Values() []EndpointState {
	return []EndpointState{
		"CREATING",
		"ACTIVE",
		"TERMINATING",
		"TERMINATED",
		"TERMINATED_WITH_ERRORS",
	}
}

type FailureReason string

// Enum values for FailureReason
const (
	FailureReasonInternalError      FailureReason = "INTERNAL_ERROR"
	FailureReasonUserError          FailureReason = "USER_ERROR"
	FailureReasonValidationError    FailureReason = "VALIDATION_ERROR"
	FailureReasonClusterUnavailable FailureReason = "CLUSTER_UNAVAILABLE"
)

// Values returns all known values for FailureReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailureReason) Values() []FailureReason {
	return []FailureReason{
		"INTERNAL_ERROR",
		"USER_ERROR",
		"VALIDATION_ERROR",
		"CLUSTER_UNAVAILABLE",
	}
}

type JobRunState string

// Enum values for JobRunState
const (
	JobRunStatePending       JobRunState = "PENDING"
	JobRunStateSubmitted     JobRunState = "SUBMITTED"
	JobRunStateRunning       JobRunState = "RUNNING"
	JobRunStateFailed        JobRunState = "FAILED"
	JobRunStateCancelled     JobRunState = "CANCELLED"
	JobRunStateCancelPending JobRunState = "CANCEL_PENDING"
	JobRunStateCompleted     JobRunState = "COMPLETED"
)

// Values returns all known values for JobRunState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobRunState) Values() []JobRunState {
	return []JobRunState{
		"PENDING",
		"SUBMITTED",
		"RUNNING",
		"FAILED",
		"CANCELLED",
		"CANCEL_PENDING",
		"COMPLETED",
	}
}

type PersistentAppUI string

// Enum values for PersistentAppUI
const (
	PersistentAppUIEnabled  PersistentAppUI = "ENABLED"
	PersistentAppUIDisabled PersistentAppUI = "DISABLED"
)

// Values returns all known values for PersistentAppUI. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PersistentAppUI) Values() []PersistentAppUI {
	return []PersistentAppUI{
		"ENABLED",
		"DISABLED",
	}
}

type TemplateParameterDataType string

// Enum values for TemplateParameterDataType
const (
	TemplateParameterDataTypeNumber TemplateParameterDataType = "NUMBER"
	TemplateParameterDataTypeString TemplateParameterDataType = "STRING"
)

// Values returns all known values for TemplateParameterDataType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TemplateParameterDataType) Values() []TemplateParameterDataType {
	return []TemplateParameterDataType{
		"NUMBER",
		"STRING",
	}
}

type VirtualClusterState string

// Enum values for VirtualClusterState
const (
	VirtualClusterStateRunning     VirtualClusterState = "RUNNING"
	VirtualClusterStateTerminating VirtualClusterState = "TERMINATING"
	VirtualClusterStateTerminated  VirtualClusterState = "TERMINATED"
	VirtualClusterStateArrested    VirtualClusterState = "ARRESTED"
)

// Values returns all known values for VirtualClusterState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VirtualClusterState) Values() []VirtualClusterState {
	return []VirtualClusterState{
		"RUNNING",
		"TERMINATING",
		"TERMINATED",
		"ARRESTED",
	}
}
