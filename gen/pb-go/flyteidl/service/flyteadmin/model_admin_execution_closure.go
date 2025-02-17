/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

import (
	"time"
)

type AdminExecutionClosure struct {
	// Output URI in the case of a successful execution.
	Outputs *AdminLiteralMapBlob `json:"outputs,omitempty"`
	// Error information in the case of a failed execution.
	Error_ *CoreExecutionError `json:"error,omitempty"`
	// In the case of a user-specified abort, this will pass along the user-supplied cause.
	AbortCause string `json:"abort_cause,omitempty"`
	// In the case of a user-specified abort, this will pass along the user and their supplied cause.
	AbortMetadata *AdminAbortMetadata `json:"abort_metadata,omitempty"`
	// Raw output data produced by this execution.
	OutputData *CoreLiteralMap `json:"output_data,omitempty"`
	ComputedInputs *CoreLiteralMap `json:"computed_inputs,omitempty"`
	// Most recent recorded phase for the execution.
	Phase *CoreWorkflowExecutionPhase `json:"phase,omitempty"`
	// Reported time at which the execution began running.
	StartedAt time.Time `json:"started_at,omitempty"`
	// The amount of time the execution spent running.
	Duration string `json:"duration,omitempty"`
	// Reported time at which the execution was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Reported time at which the execution was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The notification settings to use after merging the CreateExecutionRequest and the launch plan notification settings. An execution launched with notifications will always prefer that definition to notifications defined statically in a launch plan.
	Notifications []AdminNotification `json:"notifications,omitempty"`
	// Identifies the workflow definition for this execution.
	WorkflowId *CoreIdentifier `json:"workflow_id,omitempty"`
}
