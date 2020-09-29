// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CancellationStatus string

// Enum values for CancellationStatus
const (
	CancellationStatusCancelling CancellationStatus = "CANCELLING"
	CancellationStatusCancelled  CancellationStatus = "CANCELLED"
)

type DeviceStatus string

// Enum values for DeviceStatus
const (
	DeviceStatusOnline  DeviceStatus = "ONLINE"
	DeviceStatusOffline DeviceStatus = "OFFLINE"
)

type DeviceType string

// Enum values for DeviceType
const (
	DeviceTypeQpu       DeviceType = "QPU"
	DeviceTypeSimulator DeviceType = "SIMULATOR"
)

type QuantumTaskStatus string

// Enum values for QuantumTaskStatus
const (
	QuantumTaskStatusCreated    QuantumTaskStatus = "CREATED"
	QuantumTaskStatusQueued     QuantumTaskStatus = "QUEUED"
	QuantumTaskStatusRunning    QuantumTaskStatus = "RUNNING"
	QuantumTaskStatusCompleted  QuantumTaskStatus = "COMPLETED"
	QuantumTaskStatusFailed     QuantumTaskStatus = "FAILED"
	QuantumTaskStatusCancelling QuantumTaskStatus = "CANCELLING"
	QuantumTaskStatusCancelled  QuantumTaskStatus = "CANCELLED"
)

type SearchQuantumTasksFilterOperator string

// Enum values for SearchQuantumTasksFilterOperator
const (
	SearchQuantumTasksFilterOperatorLt      SearchQuantumTasksFilterOperator = "LT"
	SearchQuantumTasksFilterOperatorLte     SearchQuantumTasksFilterOperator = "LTE"
	SearchQuantumTasksFilterOperatorEqual   SearchQuantumTasksFilterOperator = "EQUAL"
	SearchQuantumTasksFilterOperatorGt      SearchQuantumTasksFilterOperator = "GT"
	SearchQuantumTasksFilterOperatorGte     SearchQuantumTasksFilterOperator = "GTE"
	SearchQuantumTasksFilterOperatorBetween SearchQuantumTasksFilterOperator = "BETWEEN"
)