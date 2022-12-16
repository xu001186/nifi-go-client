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

type JvmSystemDiagnosticsSnapshotDto struct {
	FlowFileRepositoryStorageUsage *RepositoryUsageDto `json:"flowFileRepositoryStorageUsage,omitempty"`
	// Information about the Content Repository's usage
	ContentRepositoryStorageUsage []RepositoryUsageDto `json:"contentRepositoryStorageUsage,omitempty"`
	// Information about the Provenance Repository's usage
	ProvenanceRepositoryStorageUsage []RepositoryUsageDto `json:"provenanceRepositoryStorageUsage,omitempty"`
	// The maximum number of bytes that the JVM heap is configured to use for heap
	MaxHeapBytes int64 `json:"maxHeapBytes,omitempty"`
	// The maximum number of bytes that the JVM heap is configured to use, as a human-readable value
	MaxHeap string `json:"maxHeap,omitempty"`
	// Diagnostic information about the JVM's garbage collections
	GarbageCollectionDiagnostics []GarbageCollectionDiagnosticsDto `json:"garbageCollectionDiagnostics,omitempty"`
	// The number of CPU Cores available on the system
	CpuCores int32 `json:"cpuCores,omitempty"`
	// The 1-minute CPU Load Average
	CpuLoadAverage float64 `json:"cpuLoadAverage,omitempty"`
	// The number of bytes of RAM available on the system
	PhysicalMemoryBytes int64 `json:"physicalMemoryBytes,omitempty"`
	// The number of bytes of RAM available on the system as a human-readable value
	PhysicalMemory string `json:"physicalMemory,omitempty"`
	// The number of files that are open by the NiFi process
	OpenFileDescriptors int64 `json:"openFileDescriptors,omitempty"`
	// The maximum number of open file descriptors that are available to each process
	MaxOpenFileDescriptors int64 `json:"maxOpenFileDescriptors,omitempty"`
}
