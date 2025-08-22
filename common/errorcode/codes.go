// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package errorcode

func init() {
	// Register all error codes
	registerInfraCodes()
	registerFrontendCodes()
	registerHistoryCodes()
	registerMatchingCodes()
	registerWorkerCodes()
}

func registerInfraCodes() {
	// Persistence (1100-1199)
	Register(1101, ComponentInfra, "Database connection failed")
	Register(1102, ComponentInfra, "Transaction commit failed")
	Register(1103, ComponentInfra, "Persistence timeout")
	Register(1104, ComponentInfra, "Nexus endpoint create/update failed")
	Register(1105, ComponentInfra, "Nexus endpoint get failed")
	Register(1106, ComponentInfra, "Next page token serialization failed")
	Register(1107, ComponentInfra, "Mutable state retrieval failed")
	Register(1108, ComponentInfra, "History branch trim failed")
	Register(1109, ComponentInfra, "Nexus endpoints table version increment failed")
	Register(1110, ComponentInfra, "Nexus endpoint delete operation failed")
	Register(1111, ComponentInfra, "Nexus endpoint delete rows affected error")
	Register(1112, ComponentInfra, "Service startup failed")
	Register(1113, ComponentInfra, "Service shutdown failed")
	Register(1114, ComponentInfra, "Cluster metadata configuration error")
	Register(1115, ComponentInfra, "Service lifecycle hook failed")

	// Clustering (1200-1299)
	Register(1201, ComponentInfra, "Cluster membership change")
	Register(1202, ComponentInfra, "Ring membership failed")
	Register(1203, ComponentInfra, "Ring refresh failed")
	Register(1204, ComponentInfra, "Cluster metadata refresh failed")
	Register(1205, ComponentInfra, "Ring refresh on scheduled event failed")
	Register(1206, ComponentInfra, "Listener notification channel full")
	Register(1207, ComponentInfra, "Ring refresh by request failed")
	Register(1208, ComponentInfra, "Periodic ring refresh failed")
}

func registerFrontendCodes() {
	// Request Validation (2000-2099)
	Register(2001, ComponentFrontend, "Invalid workflow ID provided")
	Register(2002, ComponentFrontend, "Invalid namespace name")
	Register(2003, ComponentFrontend, "Missing required request parameter")
	Register(2004, ComponentFrontend, "Invalid request format")

	// Authentication/Authorization (2100-2199)
	Register(2101, ComponentFrontend, "Authentication required")
	Register(2102, ComponentFrontend, "Invalid authentication token")
	Register(2103, ComponentFrontend, "Insufficient permissions")
	Register(2104, ComponentFrontend, "Token expired")

	// Rate Limiting (2200-2299)
	Register(2201, ComponentFrontend, "Request rate limit exceeded")
	Register(2202, ComponentFrontend, "Namespace rate limit exceeded")
	Register(2203, ComponentFrontend, "User rate limit exceeded")

	// Namespace Operations (2300-2399)
	Register(2301, ComponentFrontend, "Namespace not found")
	Register(2302, ComponentFrontend, "Namespace already exists")
	Register(2303, ComponentFrontend, "Namespace operation failed")

	// Worker and Task Queue Operations (2400-2499)
	Register(2401, ComponentFrontend, "Failed to record worker heartbeat")
	Register(2402, ComponentFrontend, "Unable to call matching service")
	Register(2403, ComponentFrontend, "Worker task reachability check failed")
	Register(2404, ComponentFrontend, "Schedule decoding failed")
	Register(2405, ComponentFrontend, "Nexus task queue poll failed")

	// Admin Operations (2500-2599)
	Register(2501, ComponentFrontend, "Failed to get search attributes")
	Register(2502, ComponentFrontend, "Failed to close replication messages server")
	Register(2503, ComponentFrontend, "Failed to describe history host")

	// Schedule Operations (2600-2699)
	Register(2601, ComponentFrontend, "Schedule memo encoding failed")

	// Nexus Operations (2700-2799)
	Register(2701, ComponentFrontend, "Nexus operation panic captured")
	Register(2702, ComponentFrontend, "Nexus payload size exceeds limit")
	Register(2703, ComponentFrontend, "Nexus link URL parsing failed")
	Register(2704, ComponentFrontend, "Nexus forwarded start operation request failed")
	Register(2705, ComponentFrontend, "Nexus forwarded cancel operation request failed")
	Register(2706, ComponentFrontend, "Nexus HTTP client creation failed")
	Register(2707, ComponentFrontend, "Nexus service base URL construction failed")
	Register(2708, ComponentFrontend, "Nexus failure marshaling failed")
	Register(2709, ComponentFrontend, "Nexus response body write failed")
	Register(2710, ComponentFrontend, "Nexus invalid URL provided")
	Register(2711, ComponentFrontend, "Nexus invalid namespace name")
	Register(2712, ComponentFrontend, "Nexus claims retrieval failed")
	Register(2713, ComponentFrontend, "Nexus invalid endpoint ID")
	Register(2714, ComponentFrontend, "Nexus namespace lookup failed")
	Register(2715, ComponentFrontend, "OpenAPI spec reader initialization failed")
	Register(2716, ComponentFrontend, "OpenAPI spec send failed")
	Register(2717, ComponentFrontend, "OpenAPI spec checksum verification failed")
	Register(2718, ComponentFrontend, "Nexus endpoints persistence listing failed")
	Register(2719, ComponentFrontend, "Nexus endpoint client generic error")

	// Warnings (2900-2999)
	Register(2901, ComponentFrontend, "Unspecified task queue kind")
}

func registerHistoryCodes() {
	// Workflow Execution (3000-3099)
	Register(3001, ComponentHistory, "Workflow execution not found")
	Register(3002, ComponentHistory, "Workflow in invalid state")
	Register(3003, ComponentHistory, "Workflow execution already exists")
	Register(3004, ComponentHistory, "Workflow task processing failed")

	// Activity Management (3100-3199)
	Register(3101, ComponentHistory, "Activity not found")
	Register(3102, ComponentHistory, "Activity execution failed")
	Register(3103, ComponentHistory, "Activity timeout exceeded")

	// State Management (3300-3399)
	Register(3301, ComponentHistory, "Shard ownership lost")
	Register(3302, ComponentHistory, "Current branch changed")
	Register(3303, ComponentHistory, "Condition check failed")

	// Transfer Queue Operations (3400-3499)
	Register(3401, ComponentHistory, "Transfer request cancel failed")
	Register(3402, ComponentHistory, "Transfer signal workflow failed")
	Register(3403, ComponentHistory, "Transfer child workflow start failed")
	Register(3404, ComponentHistory, "Transfer auto-reset workflow failed")

	// Workflow State Management (3500-3599)
	Register(3501, ComponentHistory, "Mutable state dirty transaction error")
	Register(3502, ComponentHistory, "Database data inconsistency detected")
	Register(3503, ComponentHistory, "Sync versioned transition task missing")

	// Queue Processing (3600-3699)
	Register(3601, ComponentHistory, "Task processor not registered")
	Register(3602, ComponentHistory, "Engine retrieval failed")
	Register(3603, ComponentHistory, "DLQ replication tasks fetch failed")

	// Archival Operations (3700-3799)
	Register(3701, ComponentHistory, "Visibility URI parsing failed")
	Register(3702, ComponentHistory, "History URI parsing failed")

	// Queue Executable Operations (3800-3899)
	Register(3801, ComponentHistory, "Task payload serialization for OTEL span failed")
	Register(3802, ComponentHistory, "Queue executable panic captured")
}

func registerMatchingCodes() {
	// Task Queue Management (4000-4099)
	Register(4001, ComponentMatching, "Task queue not found")
	Register(4002, ComponentMatching, "Task queue partition failed")
	Register(4003, ComponentMatching, "Task forwarding failed")
	Register(4004, ComponentMatching, "Task dropped due to non-retryable errors")
	Register(4005, ComponentMatching, "Deployment registration channel unlocked error")
	Register(4006, ComponentMatching, "Partition force load failed")

	// Worker Management (4100-4199)
	Register(4101, ComponentMatching, "No available workers")
	Register(4102, ComponentMatching, "Worker polling failed")
	Register(4103, ComponentMatching, "Worker assignment failed")

	// Persistence Operations (4200-4299)
	Register(4201, ComponentMatching, "Task store operation failure")

	// User Data Management (4300-4399)
	Register(4301, ComponentMatching, "User data fetch from parent failed")
	Register(4302, ComponentMatching, "User data version mismatch")
	Register(4303, ComponentMatching, "Replication task publish failed")
	Register(4304, ComponentMatching, "User data update function failed")
	Register(4305, ComponentMatching, "User data push to matching node failed")
	Register(4306, ComponentMatching, "User data version request exceeded known version")

	// Nexus Endpoint Operations (4400-4499)
	Register(4401, ComponentMatching, "Namespace lookup by ID failed")
	Register(4402, ComponentMatching, "Nexus endpoint creation failed")
	Register(4403, ComponentMatching, "Nexus endpoint update failed")
	Register(4404, ComponentMatching, "Nexus endpoint deletion failed")
	Register(4405, ComponentMatching, "Nexus endpoints ownership check failed")
	Register(4406, ComponentMatching, "Matching node not Nexus endpoints table owner")

	// Deployment Operations (4500-4599)
	Register(4501, ComponentMatching, "Deployment version registration error")
	Register(4502, ComponentMatching, "Deployment version wait timeout")
}

func registerWorkerCodes() {
	// System Workers (5000-5099)
	Register(5001, ComponentWorker, "System worker startup failed")
	Register(5002, ComponentWorker, "Background task processing failed")
	Register(5003, ComponentWorker, "Worker shutdown timeout")
	Register(5004, ComponentWorker, "Membership listener unregister failed")
	Register(5005, ComponentWorker, "Existing timer found in error handler")
	Register(5006, ComponentWorker, "SDK worker failed to start out of retries")
	Register(5007, ComponentWorker, "Failed to look up worker hosts")
	Register(5008, ComponentWorker, "SDK worker non-retryable error")

	// Archival (5100-5199)
	Register(5101, ComponentWorker, "History archival failed")
	Register(5102, ComponentWorker, "Visibility archival failed")

	// Scanner Operations (5200-5299)
	Register(5201, ComponentWorker, "Scanner heartbeat details recovery failed")

	// Replication Operations (5300-5399)
	Register(5301, ComponentWorker, "Replication tasks fetch failed")
	Register(5302, ComponentWorker, "Replication tasks apply failed")
	Register(5303, ComponentWorker, "Replication tasks DLQ put failed")
	Register(5304, ComponentWorker, "Namespace replication task processing failed")
	Register(5305, ComponentWorker, "Task queue user data replication task processing failed")

	// Namespace Operations (5400-5499)
	Register(5401, ComponentWorker, "Delete namespace child workflow error")

	// Batcher Operations (5500-5599)
	Register(5501, ComponentWorker, "Batch operation completion failed")
	Register(5502, ComponentWorker, "Batch operation namespace mismatch")
	Register(5503, ComponentWorker, "Batch reset options deserialization failed")
	Register(5504, ComponentWorker, "Batch post-reset operation deserialization failed")
	Register(5505, ComponentWorker, "Batch heartbeat recovery failed")
	Register(5506, ComponentWorker, "Batch workflow count estimation failed")
	Register(5507, ComponentWorker, "Batch operation task processing failed")
	Register(5508, ComponentWorker, "Batch workflow history reverse read failed")
	Register(5509, ComponentWorker, "Batch workflow history read failed")

	// Worker Deployment Operations (5600-5699)
	Register(5601, ComponentWorker, "Worker deployment query handler setup failed")
	Register(5602, ComponentWorker, "Worker deployment workflow lock acquisition failed")
	Register(5603, ComponentWorker, "Worker deployment update canceled before start")
	Register(5604, ComponentWorker, "Worker deployment version check failed")
	Register(5605, ComponentWorker, "Worker deployment version workflow execution failed")
	Register(5606, ComponentWorker, "Worker deployment version registration failed")
	Register(5607, ComponentWorker, "Worker deployment version polling failed")
	Register(5608, ComponentWorker, "Worker deployment version validation failed")
	Register(5609, ComponentWorker, "Worker deployment version timeout")

	// Scheduler Operations (5700-5799)
	Register(5701, ComponentWorker, "Scheduler workflow error")
	Register(5702, ComponentWorker, "Scheduler action execution failed")
	Register(5703, ComponentWorker, "Scheduler policy validation failed")
	Register(5704, ComponentWorker, "Scheduler trigger evaluation failed")
	Register(5705, ComponentWorker, "Scheduler state transition failed")

	// Scanner Operations (5800-5899)
	Register(5801, ComponentWorker, "Scanner execution task processing failed")
	Register(5802, ComponentWorker, "Scanner history scavenging failed")
	Register(5803, ComponentWorker, "Scanner task queue scavenging failed")
	Register(5804, ComponentWorker, "Scanner build IDs scavenging failed")

	// Migration Operations (5900-5999)
	Register(5901, ComponentWorker, "Migration activity execution failed")
	Register(5902, ComponentWorker, "Migration workflow processing failed")
	Register(5903, ComponentWorker, "Migration validation failed")

	// Delete Namespace Operations (6000-6099)
	Register(6001, ComponentWorker, "Delete namespace workflow failed")
	Register(6002, ComponentWorker, "Delete namespace executions failed")
	Register(6003, ComponentWorker, "Delete namespace reclaim resources failed")
	Register(6004, ComponentWorker, "Delete namespace activities failed")

	// Common Components Range: 7000-7999

	// Persistence Telemetry (7000-7099)
	Register(7001, ComponentCommon, "Persistence telemetry OTEL span serialization failed")

	// Archival Components (7100-7199)
	Register(7101, ComponentArchiver, "History archival operation failed")
	Register(7102, ComponentArchiver, "Visibility archival operation failed")
	Register(7103, ComponentArchiver, "Archival upload failed")
	Register(7104, ComponentArchiver, "Archival download failed")

	// Tools and Utilities (7200-7299)
	Register(7201, ComponentTools, "Cassandra tool operation failed")
	Register(7202, ComponentTools, "SQL tool operation failed")
	Register(7203, ComponentTools, "Migration tool operation failed")
	Register(7204, ComponentTools, "CQL client operation failed")
	Register(7205, ComponentTools, "Database connection validation failed")
	Register(7206, ComponentTools, "Schema embed operation failed")

	// Component Operations (7300-7399)
	Register(7301, ComponentCommon, "Nexus operation handler failed")
	Register(7302, ComponentCommon, "Callback invocation failed")
	Register(7303, ComponentCommon, "Scheduler executor failed")
	Register(7304, ComponentCommon, "Component initialization failed")

	// Infrastructure Operations (7400-7499)
	Register(7401, ComponentInfra, "Service lifecycle operation failed")
	Register(7402, ComponentInfra, "Metrics provider operation failed")
	Register(7403, ComponentInfra, "Dynamic config operation failed")
	Register(7404, ComponentInfra, "Task scheduler operation failed")

	// Persistence Operations (7500-7599)
	Register(7501, ComponentPersist, "Elasticsearch processor operation failed")
	Register(7502, ComponentPersist, "Visibility store operation failed")
	Register(7503, ComponentPersist, "NDC history importer operation failed")
	Register(7504, ComponentPersist, "History manager operation failed")
	Register(7505, ComponentPersist, "DLQ metrics emitter operation failed")

	// History Service Operations (7800-7899)
	Register(7801, ComponentHistory, "Workflow transaction operation failed")
	Register(7802, ComponentHistory, "Replication task processor operation failed")
	Register(7803, ComponentHistory, "Replication task executor operation failed")
	Register(7804, ComponentHistory, "API get history utility operation failed")
	Register(7805, ComponentHistory, "Respond workflow task completed operation failed")

	// Membership Operations (7900-7999)
	Register(7901, ComponentCommon, "Ringpop test cluster operation failed")
	Register(7902, ComponentCommon, "Ringpop monitor operation failed")

	// Namespace Operations (8000-8099)  
	Register(8001, ComponentCommon, "Namespace replication task executor operation failed")

	// Test and Development Operations (7600-7699)
	Register(7601, ComponentTest, "Test data converter operation failed")
	Register(7602, ComponentTest, "Test execution failure")
	Register(7603, ComponentTest, "Development utility failed")

	// Log Migration Operations (7700-7799)
	Register(7701, ComponentCommon, "Log migration operation failed")

	// Additional Common Components (8100-8199)
	Register(8101, ComponentCommon, "Metrics operation failed")
	Register(8102, ComponentCommon, "Dynamic config operation failed")
	Register(8103, ComponentCommon, "Task scheduler operation failed")
	Register(8104, ComponentCommon, "Finalizer operation failed")
	Register(8105, ComponentCommon, "Soft assert operation failed")
	Register(8106, ComponentCommon, "Nexus endpoint registry operation failed")
	Register(8107, ComponentCommon, "XDC cache operation failed")
	Register(8108, ComponentCommon, "Visibility manager metrics operation failed")
	Register(8109, ComponentCommon, "Persistence metric client operation failed")
	Register(8110, ComponentCommon, "Nexus endpoint manager operation failed")
	Register(8111, ComponentCommon, "Namespace registry operation failed")
	Register(8112, ComponentCommon, "DLQ message handler operation failed")
	Register(8113, ComponentCommon, "Utility operation failed")
	Register(8114, ComponentCommon, "Finalizer timeout")
	Register(8115, ComponentCommon, "Deadlock detected")
	Register(8116, ComponentCommon, "Deadlock profile not found")
	Register(8117, ComponentCommon, "Deadlock profile failed")
	Register(8118, ComponentPersist, "DLQ list failed")
	Register(8119, ComponentPersist, "DLQ process queue name failed")
	Register(8120, ComponentPersist, "DLQ category not found")
	Register(8121, ComponentPersist, "DLQ history service lookup failed")
	Register(8122, ComponentCommon, "XDC cache events truncated")

	// Schema and Embedding (8200-8299)
	Register(8201, ComponentSchema, "Schema embed operation failed")

	// Additional Archival (8300-8399)
	Register(8301, ComponentArchiver, "S3 visibility archiver operation failed")
	Register(8302, ComponentArchiver, "Additional archival operation failed")

	// Additional History Service Operations (8400-8499)
	Register(8401, ComponentHistory, "Task priority unknown key")
	Register(8402, ComponentHistory, "Task priority unknown type")
	Register(8403, ComponentHistory, "Archive target failed")
	Register(8404, ComponentWorker, "Scheduler result size exceeded")
	Register(8405, ComponentWorker, "Scheduler failure size exceeded")
	Register(8406, ComponentHistory, "History events cache retrieve failed")
	Register(8407, ComponentHistory, "History events data corruption")

	// Namespace Replication Operations (8500-8599)
	Register(8501, ComponentCommon, "Namespace replication operation failed")
	Register(8502, ComponentCommon, "Namespace replication UUID collision")
	Register(8503, ComponentCommon, "Namespace replication creation UUID collision")
	Register(8504, ComponentCommon, "Namespace replication creation error")
	Register(8505, ComponentCommon, "Namespace replication creation name collision")

	// TLS/Encryption Operations (8600-8699)
	Register(8601, ComponentCommon, "TLS per-host provider lookup error")
	Register(8602, ComponentCommon, "TLS certificate expiration check error")
	Register(8603, ComponentCommon, "TLS certificate expired")

	// Nexus Operations (8700-8799)
	Register(8701, ComponentCommon, "Common Nexus operation failed")

	// Persistence Operations (8800-8899)
	Register(8801, ComponentPersist, "Cassandra drop keyspace error")
	Register(8802, ComponentPersist, "Cassandra create keyspace error")
	Register(8803, ComponentPersist, "SQL close database error")
	Register(8804, ComponentPersist, "SQL transaction rollback error")

	// Worker Operations (8900-8999)
	Register(8901, ComponentWorker, "Add search attributes ES mapping retryable")
	Register(8902, ComponentWorker, "Add search attributes ES mapping non-retryable")
	Register(8903, ComponentWorker, "Add search attributes ES status failed")
	Register(8904, ComponentWorker, "Child workflow error")
	Register(8905, ComponentWorker, "Task submit to executor failed")

	// Common Archiver Operations (9000-9099)
	Register(9001, ComponentArchiver, "Common archiver operation failed")
	
	// Additional Missing Operations (9100-9199)
	Register(9101, ComponentWorker, "Replication task generation failed")
	Register(9102, ComponentCommon, "Namespace get for replication failed")
	Register(9103, ComponentWorker, "Worker scheduler state inconsistent")
	Register(9104, ComponentWorker, "Task queue user data update failed")
	Register(9105, ComponentWorker, "Task queue user data sync failed")
	Register(9106, ComponentWorker, "Task queue user data wait failed")
	Register(9107, ComponentWorker, "Deployment client error")
	Register(9108, ComponentWorker, "Worker deployment operation failed")
	Register(9109, ComponentWorker, "Task queue add rate fetch error")
	Register(9110, ComponentWorker, "Workflow count for drainage error")
	Register(9111, ComponentCommon, "Membership listener removal failed")
	Register(9112, ComponentHistory, "History event get failed")
	Register(9113, ComponentHistory, "Import action commit failed")
	Register(9114, ComponentHistory, "History get error")
	Register(9115, ComponentHistory, "Event replication failed")
	Register(9116, ComponentHistory, "History queue unknown alert type")
	Register(9117, ComponentHistory, "History queue task range complete failed")
	Register(9118, ComponentHistory, "History queue update state failed")
	Register(9119, ComponentHistory, "History queue task load failed")
	Register(9120, ComponentHistory, "History queue rate limiter configuration")
	Register(9121, ComponentHistory, "History queue task retrieve failed")
	Register(9122, ComponentHistory, "History scheduler creation failed")
	Register(9123, ComponentHistory, "Replication task reader error")
	Register(9124, ComponentHistory, "Stream close error")
	Register(9125, ComponentHistory, "Replication ack level update failed")
	Register(9126, ComponentHistory, "Replication task save failed")
	Register(9127, ComponentHistory, "History replication task conversion failed")
	Register(9128, ComponentHistory, "Replication service error")
	Register(9129, ComponentHistory, "Rate limiter wait failed")
	Register(9130, ComponentHistory, "Version history get failed")
	Register(9131, ComponentHistory, "History workflow size constraint violation")
	Register(9132, ComponentMatching, "Matching persistent store failure")
	Register(9133, ComponentHistory, "Replication task fetch failed")
	Register(9134, ComponentHistory, "History replication cleanup failed")
	Register(9135, ComponentHistory, "History replication DLQ operation failed")
	Register(9136, ComponentWorker, "Task queue delete error")
	Register(9137, ComponentWorker, "Scavenger delete handler error")
	Register(9138, ComponentWorker, "Task queue list error")
	Register(9139, ComponentHistory, "History sync workflow state retrieve failed")
	Register(9140, ComponentHistory, "History sync workflow state update failed")
	Register(9141, ComponentWorker, "Workflow start error")
}

// Constants for easy access
const (
	// Infrastructure Range: 1000-1999
	InfraDBConnectionFailed                        = 1101
	InfraTransactionFailed                         = 1102
	InfraPersistenceTimeout                        = 1103
	InfraNexusEndpointCreateUpdateFailed           = 1104
	InfraNexusEndpointGetFailed                    = 1105
	InfraNextPageTokenSerializationFailed          = 1106
	InfraMutableStateRetrievalFailed               = 1107
	InfraHistoryBranchTrimFailed                   = 1108
	InfraNexusEndpointsTableVersionIncrementFailed = 1109
	InfraNexusEndpointDeleteOperationFailed        = 1110
	InfraNexusEndpointDeleteRowsAffectedError      = 1111
	InfraServiceStartupFailed                      = 1112
	InfraServiceShutdownFailed                     = 1113
	InfraClusterMetadataConfigurationError         = 1114
	InfraServiceLifecycleHookFailed                = 1115
	InfraClusterMembershipChange                   = 1201
	InfraRingMembershipFailed                      = 1202
	InfraRingRefreshFailed                         = 1203
	InfraClusterMetadataRefreshFailed              = 1204
	InfraRingRefreshScheduledEventFailed           = 1205
	InfraListenerNotificationChannelFull           = 1206
	InfraRingRefreshByRequestFailed                = 1207
	InfraPeriodicRingRefreshFailed                 = 1208

	// Frontend Service Range: 2000-2999
	FrontendInvalidWorkflowID         = 2001
	FrontendInvalidNamespace          = 2002
	FrontendMissingParameter          = 2003
	FrontendInvalidRequestFormat      = 2004
	FrontendAuthRequired              = 2101
	FrontendInvalidToken              = 2102
	FrontendInsufficientPerms         = 2103
	FrontendTokenExpired              = 2104
	FrontendRateLimitExceeded         = 2201
	FrontendNamespaceRateLimit        = 2202
	FrontendUserRateLimit             = 2203
	FrontendNamespaceNotFound         = 2301
	FrontendNamespaceExists           = 2302
	FrontendNamespaceOpFailed         = 2303
	FrontendWorkerHeartbeatFailed     = 2401
	FrontendMatchingServiceCallFailed = 2402
	FrontendWorkerReachabilityFailed  = 2403
	FrontendScheduleDecodingFailed    = 2404
	FrontendNexusTaskQueuePollFailed  = 2405

	// Admin Operations (2500-2599)
	FrontendSearchAttributesGetFailed            = 2501
	FrontendReplicationMessagesServerCloseFailed = 2502
	FrontendHistoryHostDescribeFailed            = 2503

	// Schedule Operations (2600-2699)
	FrontendScheduleMemoEncodingFailed = 2601

	// Nexus Operations (2700-2799)
	FrontendNexusOperationPanicCaptured            = 2701
	FrontendNexusPayloadSizeExceedsLimit           = 2702
	FrontendNexusLinkURLParsingFailed              = 2703
	FrontendNexusForwardedStartOperationFailed     = 2704
	FrontendNexusForwardedCancelOperationFailed    = 2705
	FrontendNexusHTTPClientCreationFailed          = 2706
	FrontendNexusServiceBaseURLConstructionFailed  = 2707
	FrontendNexusFailureMarshalingFailed           = 2708
	FrontendNexusResponseBodyWriteFailed           = 2709
	FrontendNexusInvalidURLProvided                = 2710
	FrontendNexusInvalidNamespaceName              = 2711
	FrontendNexusClaimsRetrievalFailed             = 2712
	FrontendNexusInvalidEndpointID                 = 2713
	FrontendNexusNamespaceLookupFailed             = 2714
	FrontendOpenAPISpecReaderInitFailed            = 2715
	FrontendOpenAPISpecSendFailed                  = 2716
	FrontendOpenAPISpecChecksumVerificationFailed  = 2717
	FrontendNexusEndpointsPersistenceListingFailed = 2718
	FrontendNexusEndpointClientGenericError        = 2719

	// Warnings (2900-2999)
	FrontendTaskQueueKindUnspecified = 2901

	// History Service Range: 3000-3999
	HistoryWorkflowNotFound     = 3001
	HistoryInvalidState         = 3002
	HistoryWorkflowExists       = 3003
	HistoryTaskProcessingFailed = 3004
	HistoryActivityNotFound     = 3101
	HistoryActivityFailed       = 3102
	HistoryActivityTimeout      = 3103
	HistoryShardOwnershipLost   = 3301
	HistoryCurrentBranchChanged = 3302
	HistoryConditionFailed      = 3303

	// Transfer Queue Operations (3400-3499)
	HistoryTransferRequestCancelFailed = 3401
	HistoryTransferSignalFailed        = 3402
	HistoryTransferChildWorkflowFailed = 3403
	HistoryTransferAutoResetFailed     = 3404

	// Workflow State Management (3500-3599)
	HistoryMutableStateDirtyTransaction   = 3501
	HistoryDataInconsistency              = 3502
	HistorySyncVersionedTransitionMissing = 3503

	// Queue Processing (3600-3699)
	HistoryTaskProcessorNotRegistered = 3601
	HistoryEngineRetrievalFailed      = 3602
	HistoryDLQReplicationTasksFailed  = 3603

	// Archival Operations (3700-3799)
	HistoryVisibilityURIParsingFailed = 3701
	HistoryArchivalURIParsingFailed   = 3702

	// Queue Executable Operations (3800-3899)
	HistoryTaskPayloadSerializationFailed = 3801
	HistoryQueueExecutablePanicCaptured   = 3802

	// Matching Service Range: 4000-4999
	MatchingTaskQueueNotFound                  = 4001
	MatchingPartitionFailed                    = 4002
	MatchingForwardingFailed                   = 4003
	MatchingTaskDroppedNonRetryable            = 4004
	MatchingDeploymentRegistrationError        = 4005
	MatchingPartitionForceLoadFailed           = 4006
	MatchingNoWorkers                          = 4101
	MatchingWorkerPollingFailed                = 4102
	MatchingWorkerAssignFailed                 = 4103
	MatchingTaskStoreOperationFailed           = 4201
	MatchingUserDataFetchFailed                = 4301
	MatchingUserDataVersionMismatch            = 4302
	MatchingReplicationTaskPublishFailed       = 4303
	MatchingUserDataUpdateFailed               = 4304
	MatchingUserDataPushFailed                 = 4305
	MatchingUserDataVersionExceeded            = 4306
	MatchingNamespaceLookupFailed              = 4401
	MatchingNexusEndpointCreationFailed        = 4402
	MatchingNexusEndpointUpdateFailed          = 4403
	MatchingNexusEndpointDeletionFailed        = 4404
	MatchingNexusEndpointsOwnershipCheckFailed = 4405
	MatchingNexusEndpointsOwnershipNotOwner    = 4406
	MatchingDeploymentVersionRegistrationError = 4501
	MatchingDeploymentVersionWaitTimeout       = 4502

	// Worker Service Range: 5000-5999
	WorkerStartupFailed                                    = 5001
	WorkerTaskProcessingFailed                             = 5002
	WorkerShutdownTimeout                                  = 5003
	WorkerMembershipListenerUnregisterFailed               = 5004
	WorkerExistingTimerFoundError                          = 5005
	WorkerSDKStartFailedOutOfRetries                       = 5006
	WorkerHostLookupFailed                                 = 5007
	WorkerSDKNonRetryableError                             = 5008
	WorkerHistoryArchivalFailed                            = 5101
	WorkerVisibilityArchivalFailed                         = 5102
	WorkerScannerHeartbeatRecoveryFailed                   = 5201
	WorkerReplicationTasksFetchFailed                      = 5301
	WorkerReplicationTasksApplyFailed                      = 5302
	WorkerReplicationTasksDLQPutFailed                     = 5303
	WorkerNamespaceReplicationTaskProcessingFailed         = 5304
	WorkerTaskQueueUserDataReplicationTaskProcessingFailed = 5305
	WorkerDeleteNamespaceChildWorkflowError                = 5401

	// Batcher Operations (5500-5599)
	WorkerBatchOperationCompletionFailed               = 5501
	WorkerBatchOperationNamespaceMismatch              = 5502
	WorkerBatchResetOptionsDeserializationFailed       = 5503
	WorkerBatchPostResetOperationDeserializationFailed = 5504
	WorkerBatchHeartbeatRecoveryFailed                 = 5505
	WorkerBatchWorkflowCountEstimationFailed           = 5506
	WorkerBatchOperationTaskProcessingFailed           = 5507
	WorkerBatchWorkflowHistoryReverseReadFailed        = 5508
	WorkerBatchWorkflowHistoryReadFailed               = 5509

	// Worker Deployment Operations (5600-5699)
	WorkerDeploymentQueryHandlerSetupFailed        = 5601
	WorkerDeploymentWorkflowLockAcquisitionFailed  = 5602
	WorkerDeploymentUpdateCanceledBeforeStart      = 5603
	WorkerDeploymentVersionCheckFailed             = 5604
	WorkerDeploymentVersionWorkflowExecutionFailed = 5605
	WorkerDeploymentVersionRegistrationFailed      = 5606
	WorkerDeploymentVersionPollingFailed           = 5607
	WorkerDeploymentVersionValidationFailed        = 5608
	WorkerDeploymentVersionTimeout                 = 5609

	// Scheduler Operations (5700-5799)
	WorkerSchedulerWorkflowError           = 5701
	WorkerSchedulerActionExecutionFailed   = 5702
	WorkerSchedulerPolicyValidationFailed  = 5703
	WorkerSchedulerTriggerEvaluationFailed = 5704
	WorkerSchedulerStateTransitionFailed   = 5705

	// Scanner Operations (5800-5899)
	WorkerScannerExecutionTaskProcessingFailed = 5801
	WorkerScannerHistoryScavengingFailed       = 5802
	WorkerScannerTaskQueueScavengingFailed     = 5803
	WorkerScannerBuildIDsScavengingFailed      = 5804

	// Migration Operations (5900-5999)
	WorkerMigrationActivityExecutionFailed  = 5901
	WorkerMigrationWorkflowProcessingFailed = 5902
	WorkerMigrationValidationFailed         = 5903

	// Delete Namespace Operations (6000-6099)
	WorkerDeleteNamespaceWorkflowFailed         = 6001
	WorkerDeleteNamespaceExecutionsFailed       = 6002
	WorkerDeleteNamespaceReclaimResourcesFailed = 6003
	WorkerDeleteNamespaceActivitiesFailed       = 6004

	// Common Components Range: 7000-7999

	// Persistence Telemetry (7000-7099)
	CommonPersistenceTelemetryOTELSerializationFailed = 7001

	// Archival Components (7100-7199)
	CommonHistoryArchivalOperationFailed = 7101
	CommonVisibilityArchivalOperationFailed = 7102
	CommonArchivalUploadFailed = 7103
	CommonArchivalDownloadFailed = 7104

	// Tools and Utilities (7200-7299)
	ToolsCassandraOperationFailed = 7201
	ToolsSQLOperationFailed = 7202
	ToolsMigrationOperationFailed = 7203
	ToolsCQLClientOperationFailed = 7204
	ToolsDatabaseConnectionValidationFailed = 7205
	ToolsSchemaEmbedOperationFailed = 7206

	// Component Operations (7300-7399)
	ComponentNexusOperationHandlerFailed = 7301
	ComponentCallbackInvocationFailed = 7302
	ComponentSchedulerExecutorFailed = 7303
	ComponentInitializationFailed = 7304

	// Infrastructure Operations (7400-7499)
	InfraServiceLifecycleOperationFailed = 7401
	InfraMetricsProviderOperationFailed = 7402
	InfraDynamicConfigOperationFailed = 7403
	InfraTaskSchedulerOperationFailed = 7404

	// Persistence Operations (7500-7599)
	PersistElasticsearchProcessorOperationFailed = 7501
	PersistVisibilityStoreOperationFailed = 7502
	PersistNDCHistoryImporterOperationFailed = 7503
	PersistHistoryManagerOperationFailed = 7504
	PersistDLQMetricsEmitterOperationFailed = 7505

	// History Service Operations (7800-7899)
	HistWorkflowTransactionOperationFailed = 7801
	HistReplicationTaskProcessorOperationFailed = 7802
	HistReplicationTaskExecutorOperationFailed = 7803
	HistAPIGetHistoryUtilityOperationFailed = 7804
	HistRespondWorkflowTaskCompletedOperationFailed = 7805

	// Membership Operations (7900-7999)
	MemberRingpopTestClusterOperationFailed = 7901
	MemberRingpopMonitorOperationFailed = 7902

	// Namespace Operations (8000-8099)
	NSReplicationTaskExecutorOperationFailed = 8001

	// Test and Development Operations (7600-7699)
	TestDataConverterOperationFailed = 7601
	TestExecutionFailure = 7602
	TestDevelopmentUtilityFailed = 7603

	// Log Migration Operations (7700-7799)
	CommonLogMigrationOperationFailed = 7701

	// Additional Common Components (8100-8199)
	CommonMetricsOperationFailed = 8101
	CommonDynamicConfigOperationFailed = 8102
	CommonTaskSchedulerOperationFailed = 8103
	CommonFinalizerOperationFailed = 8104
	CommonSoftAssertOperationFailed = 8105
	CommonNexusEndpointRegistryOperationFailed = 8106
	CommonXDCCacheOperationFailed = 8107
	XDCCacheEventsTruncated = 8122
	CommonVisibilityManagerMetricsOperationFailed = 8108
	CommonPersistenceMetricClientOperationFailed = 8109
	CommonNexusEndpointManagerOperationFailed = 8110
	CommonNamespaceRegistryOperationFailed = 8111
	CommonDLQMessageHandlerOperationFailed = 8112
	CommonUtilityOperationFailed = 8113
	CommonFinalizerTimeout = 8114
	DeadlockDetected = 8115
	DeadlockProfileNotFound = 8116
	DeadlockProfileFailed = 8117
	PersistenceDLQListFailed = 8118
	PersistenceDLQProcessQueueNameFailed = 8119
	PersistenceDLQCategoryNotFound = 8120
	PersistenceDLQHistoryServiceLookupFailed = 8121

	// Schema and Embedding (8200-8299)
	SchemaEmbedOperationFailed = 8201

	// Additional Archival (8300-8399)
	ArchiverS3VisibilityOperationFailed = 8301
	ArchiverAdditionalOperationFailed = 8302

	// Additional History Service Operations (8400-8499)
	HistoryTaskPriorityUnknownKey = 8401
	HistoryTaskPriorityUnknownType = 8402
	HistoryArchiveTargetFailed = 8403
	WorkerSchedulerResultSizeExceeded = 8404
	WorkerSchedulerFailureSizeExceeded = 8405
	HistoryEventsCacheRetrieveFailed = 8406
	HistoryEventsDataCorruption = 8407

	// Namespace Replication Operations (8500-8599)
	CommonNamespaceReplicationOperationFailed = 8501
	NamespaceReplicationUUIDCollision = 8502
	NamespaceReplicationCreationUUIDCollision = 8503
	NamespaceReplicationCreationError = 8504
	NamespaceReplicationCreationNameCollision = 8505

	// TLS/Encryption Operations (8600-8699)
	TLSPerHostProviderLookupError = 8601
	TLSCertExpirationCheckError = 8602
	TLSCertExpired = 8603

	// Nexus Operations (8700-8799)
	CommonNexusOperationFailed = 8701

	// Persistence Operations (8800-8899)
	PersistenceCassandraDropKeyspaceError = 8801
	PersistenceCassandraCreateKeyspaceError = 8802
	PersistenceSQLCloseDatabaseError = 8803
	PersistenceSQLTransactionRollbackError = 8804

	// Worker Operations (8900-8999)
	WorkerAddSearchAttributesESMappingRetryable = 8901
	WorkerAddSearchAttributesESMappingNonRetryable = 8902
	WorkerAddSearchAttributesESStatusFailed = 8903
	ChildWorkflowError = 8904
	TaskSubmitToExecutorFailed = 8905

	// Common Archiver Operations (9000-9099)
	CommonArchiverOperationFailed = 9001
	
	// Additional Missing Operations (9100-9199)
	ReplicationTaskGenerationFailed = 9101
	NamespaceGetForReplicationFailed = 9102
	WorkerSchedulerStateInconsistent = 9103
	TaskQueueUserDataUpdateFailed = 9104
	TaskQueueUserDataSyncFailed = 9105
	TaskQueueUserDataWaitFailed = 9106
	DeploymentClientError = 9107
	WorkerDeploymentOperationFailed = 9108
	TaskQueueAddRateFetchError = 9109
	WorkflowCountForDrainageError = 9110
	MembershipListenerRemovalFailed = 9111
	HistoryEventGetFailed = 9112
	ImportActionCommitFailed = 9113
	HistoryGetError = 9114
	EventReplicationFailed = 9115
	HistoryQueueUnknownAlertType = 9116
	HistoryQueueTaskRangeCompleteFailed = 9117
	HistoryQueueUpdateStateFailed = 9118
	HistoryQueueTaskLoadFailed = 9119
	HistoryQueueRateLimiterConfiguration = 9120
	HistoryQueueTaskRetrieveFailed = 9121
	HistorySchedulerCreationFailed = 9122
	ReplicationTaskReaderError = 9123
	StreamCloseError = 9124
	ReplicationAckLevelUpdateFailed = 9125
	ReplicationTaskSaveFailed = 9126
	HistReplicationTaskConversionFailed = 9127
	ReplicationServiceError = 9128
	RateLimiterWaitFailed = 9129
	VersionHistoryGetFailed = 9130
	HistoryWorkflowSizeConstraintViolation = 9131
	MatchingPersistentStoreFailure = 9132
	ReplicationTaskFetchFailed = 9133
	HistReplicationCleanupFailed = 9134
	HistReplicationDLQOperationFailed = 9135
	TaskQueueDeleteError = 9136
	ScavengerDeleteHandlerError = 9137
	TaskQueueListError = 9138
	HistorySyncWorkflowStateRetrieveFailed = 9139
	HistorySyncWorkflowStateUpdateFailed = 9140
	WorkflowStartError = 9141
)
