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

type ControllerDto struct {
	// The id of the NiFi.
	Id string `json:"id,omitempty"`
	// The name of the NiFi.
	Name string `json:"name,omitempty"`
	// The comments for the NiFi.
	Comments string `json:"comments,omitempty"`
	// The number of running components in the NiFi.
	RunningCount int32 `json:"runningCount,omitempty"`
	// The number of stopped components in the NiFi.
	StoppedCount int32 `json:"stoppedCount,omitempty"`
	// The number of invalid components in the NiFi.
	InvalidCount int32 `json:"invalidCount,omitempty"`
	// The number of disabled components in the NiFi.
	DisabledCount int32 `json:"disabledCount,omitempty"`
	// The number of active remote ports contained in the NiFi.
	ActiveRemotePortCount int32 `json:"activeRemotePortCount,omitempty"`
	// The number of inactive remote ports contained in the NiFi.
	InactiveRemotePortCount int32 `json:"inactiveRemotePortCount,omitempty"`
	// The number of input ports contained in the NiFi.
	InputPortCount int32 `json:"inputPortCount,omitempty"`
	// The number of output ports in the NiFi.
	OutputPortCount int32 `json:"outputPortCount,omitempty"`
	// The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteListeningPort int32 `json:"remoteSiteListeningPort,omitempty"`
	// The HTTP(S) Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteHttpListeningPort int32 `json:"remoteSiteHttpListeningPort,omitempty"`
	// Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication).
	SiteToSiteSecure bool `json:"siteToSiteSecure,omitempty"`
	// If clustered, the id of the Cluster Manager, otherwise the id of the NiFi.
	InstanceId string `json:"instanceId,omitempty"`
	// The input ports available to send data to for the NiFi.
	InputPorts []PortDto `json:"inputPorts,omitempty"`
	// The output ports available to received data from the NiFi.
	OutputPorts []PortDto `json:"outputPorts,omitempty"`
}
