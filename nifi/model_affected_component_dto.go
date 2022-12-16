/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.17.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

type AffectedComponentDto struct {
	// The UUID of the Process Group that this component is in
	ProcessGroupId string `json:"processGroupId,omitempty"`
	// The UUID of this component
	Id string `json:"id,omitempty"`
	// The type of this component
	ReferenceType string `json:"referenceType,omitempty"`
	// The name of this component.
	Name string `json:"name,omitempty"`
	// The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state.
	State string `json:"state,omitempty"`
	// The number of active threads for the referencing component.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The validation errors for the component.
	ValidationErrors []string `json:"validationErrors,omitempty"`
}
