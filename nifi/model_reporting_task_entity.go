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

type ReportingTaskEntity struct {
	Revision *RevisionDto `json:"revision,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	Position *PositionDto `json:"position,omitempty"`
	Permissions *PermissionsDto `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	Component *ReportingTaskDto `json:"component,omitempty"`
	OperatePermissions *PermissionsDto `json:"operatePermissions,omitempty"`
	Status *ReportingTaskStatusDto `json:"status,omitempty"`
}
