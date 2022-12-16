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

type PortStatusDto struct {
	// The id of the port.
	Id string `json:"id,omitempty"`
	// The id of the parent process group of the port.
	GroupId string `json:"groupId,omitempty"`
	// The name of the port.
	Name string `json:"name,omitempty"`
	// Whether the port has incoming or outgoing connections to a remote NiFi.
	Transmitting bool `json:"transmitting,omitempty"`
	// The run status of the port.
	RunStatus string `json:"runStatus,omitempty"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
	AggregateSnapshot *PortStatusSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []NodePortStatusSnapshotDto `json:"nodeSnapshots,omitempty"`
}
