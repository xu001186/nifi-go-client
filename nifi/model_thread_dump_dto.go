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

type ThreadDumpDto struct {
	// The ID of the node in the cluster
	NodeId string `json:"nodeId,omitempty"`
	// The address of the node in the cluster
	NodeAddress string `json:"nodeAddress,omitempty"`
	// The port the node is listening for API requests.
	ApiPort int32 `json:"apiPort,omitempty"`
	// The stack trace for the thread
	StackTrace string `json:"stackTrace,omitempty"`
	// The name of the thread
	ThreadName string `json:"threadName,omitempty"`
	// The number of milliseconds that the thread has been executing in the Processor
	ThreadActiveMillis int64 `json:"threadActiveMillis,omitempty"`
	// Indicates whether or not the user has requested that the task be terminated. If this is true, it may indicate that the thread is in a state where it will continue running indefinitely without returning.
	TaskTerminated bool `json:"taskTerminated,omitempty"`
}