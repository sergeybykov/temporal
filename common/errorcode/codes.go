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
	// Register all error codes (Total: 242 codes across all components)
	// Infrastructure: 20 codes, Frontend: 34 codes, History: 21 codes
	// Matching: 19 codes, Worker: 148 codes (includes multiple components)
	registerInfraCodes()
	registerFrontendCodes()
	registerHistoryCodes()
	registerMatchingCodes()
	registerWorkerCodes()
}

func registerInfraCodes() {
	// Total: 20 error codes
	// Persistence (1100-1199)
	Register(1101, ComponentInfra, "Database connection failed")
	Register(1102, ComponentInfra, "Transaction commit failed")
	Register(1103, ComponentInfra, "Nexus endpoint create/update failed")
	Register(1104, ComponentInfra, "Nexus endpoint get failed")
	Register(1105, ComponentInfra, "Next page token serialization failed")
	Register(1106, ComponentInfra, "Mutable state retrieval failed")
	Register(1107, ComponentInfra, "History branch trim failed")
	Register(1108, ComponentInfra, "Nexus endpoints table version increment failed")
	Register(1109, ComponentInfra, "Nexus endpoint delete operation failed")
	Register(1110, ComponentInfra, "Nexus endpoint delete rows affected error")
	Register(1111, ComponentInfra, "Service startup failed")
	Register(1112, ComponentInfra, "Service shutdown failed")
	Register(1113, ComponentInfra, "Cluster metadata configuration error")
	Register(1114, ComponentInfra, "Service lifecycle hook failed")

	// Clustering (1200-1299)
	Register(1201, ComponentInfra, "Ring refresh failed")
	Register(1202, ComponentInfra, "Cluster metadata refresh failed")
	Register(1203, ComponentInfra, "Ring refresh on scheduled event failed")
	Register(1204, ComponentInfra, "Listener notification channel full")
	Register(1205, ComponentInfra, "Ring refresh by request failed")
	Register(1206, ComponentInfra, "Periodic ring refresh failed")
}

func registerFrontendCodes() {
	// Total: 34 error codes
	// Request Validation (2000-2099)
	Register(2001, ComponentFrontend, "Invalid workflow ID provided")
	Register(2002, ComponentFrontend, "Invalid request format")

	// Authentication/Authorization (2100-2199)
	Register(2101, ComponentFrontend, "Authentication required")
	Register(2102, ComponentFrontend, "Insufficient permissions")

	// Rate Limiting (2200-2299)

	// Namespace Operations (2300-2399)
	Register(2301, ComponentFrontend, "Namespace operation failed")

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
	// Total: 21 error codes
	// Workflow Execution (3000-3099)
	Register(3001, ComponentHistory, "Workflow execution not found")
	Register(3002, ComponentHistory, "Workflow in invalid state")
	Register(3003, ComponentHistory, "Workflow execution already exists")
	Register(3004, ComponentHistory, "Workflow task processing failed")

	// Activity Management (3100-3199)

	// State Management
	Register(3005, ComponentHistory, "Current branch changed")
	Register(3006, ComponentHistory, "Condition check failed")

	// Transfer Queue Operations
	Register(3007, ComponentHistory, "Transfer request cancel failed")
	Register(3008, ComponentHistory, "Transfer signal workflow failed")
	Register(3009, ComponentHistory, "Transfer child workflow start failed")
	Register(3010, ComponentHistory, "Transfer auto-reset workflow failed")

	// Workflow State Management
	Register(3011, ComponentHistory, "Mutable state dirty transaction error")
	Register(3012, ComponentHistory, "Database data inconsistency detected")
	Register(3013, ComponentHistory, "Sync versioned transition task missing")

	// Queue Processing
	Register(3014, ComponentHistory, "Task processor not registered")
	Register(3015, ComponentHistory, "Engine retrieval failed")
	Register(3016, ComponentHistory, "DLQ replication tasks fetch failed")

	// Archival Operations
	Register(3017, ComponentHistory, "Visibility URI parsing failed")
	Register(3018, ComponentHistory, "History URI parsing failed")

	// Queue Executable Operations
	Register(3019, ComponentHistory, "Task payload serialization for OTEL span failed")
	Register(3020, ComponentHistory, "Queue executable panic captured")
	Register(3021, ComponentHistory, "Outbound task executor failed")
}

func registerMatchingCodes() {
	// Total: 19 error codes
	// Task Queue Management (4000-4099)
	Register(4001, ComponentMatching, "Task queue partition failed")
	Register(4002, ComponentMatching, "Task dropped due to non-retryable errors")
	Register(4003, ComponentMatching, "Deployment registration channel unlocked error")
	Register(4004, ComponentMatching, "Partition force load failed")

	// Worker Management (4100-4199)

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
	// Total: 148 error codes (includes codes for multiple components)
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
	Register(5501, ComponentWorker, "Batch operation namespace mismatch")
	Register(5502, ComponentWorker, "Batch reset options deserialization failed")
	Register(5503, ComponentWorker, "Batch post-reset operation deserialization failed")
	Register(5504, ComponentWorker, "Batch heartbeat recovery failed")
	Register(5505, ComponentWorker, "Batch workflow count estimation failed")
	Register(5506, ComponentWorker, "Batch operation task processing failed")
	Register(5507, ComponentWorker, "Batch workflow history reverse read failed")
	Register(5508, ComponentWorker, "Batch workflow history read failed")

	// Worker Deployment Operations (5600-5699)
	Register(5601, ComponentWorker, "Worker deployment query handler setup failed")
	Register(5602, ComponentWorker, "Worker deployment workflow lock acquisition failed")
	Register(5603, ComponentWorker, "Worker deployment update canceled before start")
	Register(5604, ComponentWorker, "Worker deployment version check failed")
	Register(5605, ComponentWorker, "Worker deployment version workflow execution failed")
	Register(5606, ComponentWorker, "Worker deployment version registration failed")
	Register(5607, ComponentWorker, "Worker deployment version polling failed")
	Register(5608, ComponentWorker, "Worker deployment version validation failed")

	// Scheduler Operations (5700-5799)
	Register(5701, ComponentWorker, "Scheduler action execution failed")
	Register(5702, ComponentWorker, "Scheduler policy validation failed")
	Register(5703, ComponentWorker, "Scheduler trigger evaluation failed")
	Register(5704, ComponentWorker, "Scheduler state transition failed")

	// Scanner Operations (5800-5899)
	Register(5801, ComponentWorker, "Scanner execution task processing failed")
	Register(5802, ComponentWorker, "Scanner history scavenging failed")

	// Migration Operations (5900-5999)
	Register(5901, ComponentWorker, "Migration activity execution failed")

	// Delete Namespace Operations (6000-6099)
	Register(6001, ComponentWorker, "Delete namespace executions failed")
	Register(6002, ComponentWorker, "Delete namespace reclaim resources failed")
	Register(6003, ComponentWorker, "Delete namespace activities failed")

	// Common Components Range: 7000-7999

	// Persistence Telemetry (7000-7099)
	Register(7001, ComponentCommon, "Persistence telemetry OTEL span serialization failed")

	// Archival Components (7100-7199)
	Register(7101, ComponentArchiver, "History archival operation failed")
	Register(7102, ComponentArchiver, "Visibility archival operation failed")
	Register(7103, ComponentArchiver, "Archival upload failed")

	// Tools and Utilities (7200-7299)
	Register(7201, ComponentTools, "CQL client operation failed")
	Register(7202, ComponentTools, "Schema embed operation failed")

	// Component Operations (7300-7399)

	// Infrastructure Operations (7400-7499)
	Register(7401, ComponentInfra, "Service lifecycle operation failed")
	Register(7402, ComponentInfra, "Metrics provider operation failed")
	Register(7403, ComponentInfra, "Task scheduler operation failed")

	// Persistence Operations (7500-7599)
	Register(7501, ComponentPersist, "Elasticsearch processor operation failed")
	Register(7502, ComponentPersist, "Visibility store operation failed")
	Register(7503, ComponentPersist, "NDC history importer operation failed")
	Register(7504, ComponentPersist, "History manager operation failed")

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

	// Test and Development Operations (7600-7699)

	// Log Migration Operations (7700-7799)
	Register(7701, ComponentCommon, "Log migration operation failed")

	// Additional Common Components (8100-8199)
	Register(8101, ComponentCommon, "Metrics operation failed")
	Register(8102, ComponentCommon, "Dynamic config operation failed")
	Register(8103, ComponentCommon, "Task scheduler operation failed")
	Register(8104, ComponentCommon, "Finalizer operation failed")
	Register(8105, ComponentCommon, "Soft assert operation failed")
	Register(8106, ComponentCommon, "XDC cache operation failed")
	Register(8107, ComponentCommon, "Visibility manager metrics operation failed")
	Register(8108, ComponentCommon, "Persistence metric client operation failed")
	Register(8109, ComponentCommon, "Nexus endpoint manager operation failed")
	Register(8110, ComponentCommon, "Namespace registry operation failed")
	Register(8111, ComponentCommon, "Utility operation failed")
	Register(8112, ComponentCommon, "Finalizer timeout")
	Register(8113, ComponentCommon, "Deadlock detected")
	Register(8114, ComponentCommon, "Deadlock profile not found")
	Register(8115, ComponentCommon, "Deadlock profile failed")
	Register(8116, ComponentPersist, "DLQ list failed")
	Register(8117, ComponentPersist, "DLQ process queue name failed")
	Register(8118, ComponentPersist, "DLQ category not found")
	Register(8119, ComponentPersist, "DLQ history service lookup failed")
	Register(8120, ComponentCommon, "XDC cache events truncated")

	// Schema and Embedding (8200-8299)
	Register(8201, ComponentSchema, "Schema embed operation failed")

	// Additional Archival (8300-8399)

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
	Register(9123, ComponentHistory, "Stream close error")
	Register(9124, ComponentHistory, "Replication task save failed")
	Register(9125, ComponentHistory, "History replication task conversion failed")
	Register(9126, ComponentHistory, "Rate limiter wait failed")
	Register(9127, ComponentHistory, "Version history get failed")
	Register(9128, ComponentHistory, "History workflow size constraint violation")
	Register(9129, ComponentMatching, "Matching persistent store failure")
	Register(9130, ComponentHistory, "Replication task fetch failed")
	Register(9131, ComponentHistory, "History replication cleanup failed")
	Register(9132, ComponentHistory, "History replication DLQ operation failed")
	Register(9133, ComponentWorker, "Task queue delete error")
	Register(9134, ComponentWorker, "Scavenger delete handler error")
	Register(9135, ComponentWorker, "Task queue list error")
	Register(9136, ComponentHistory, "History sync workflow state retrieve failed")
	Register(9137, ComponentHistory, "History sync workflow state update failed")
	Register(9138, ComponentWorker, "Workflow start error")
}

// Constants for easy access
const (
	// Infrastructure Range: 1000-1999
	InfraDBConnectionFailed                        = 1101
	InfraTransactionFailed                         = 1102
	InfraNexusEndpointCreateUpdateFailed           = 1103
	InfraNexusEndpointGetFailed                    = 1104
	InfraNextPageTokenSerializationFailed          = 1105
	InfraMutableStateRetrievalFailed               = 1106
	InfraHistoryBranchTrimFailed                   = 1107
	InfraNexusEndpointsTableVersionIncrementFailed = 1108
	InfraNexusEndpointDeleteOperationFailed        = 1109
	InfraNexusEndpointDeleteRowsAffectedError      = 1110
	InfraServiceStartupFailed                      = 1111
	InfraServiceShutdownFailed                     = 1112
	InfraClusterMetadataConfigurationError         = 1113
	InfraServiceLifecycleHookFailed                = 1114
	InfraRingRefreshFailed                         = 1201
	InfraClusterMetadataRefreshFailed              = 1202
	InfraRingRefreshScheduledEventFailed           = 1203
	InfraListenerNotificationChannelFull           = 1204
	InfraRingRefreshByRequestFailed                = 1205
	InfraPeriodicRingRefreshFailed                 = 1206

	// Frontend Service Range: 2000-2999 (sorted by error code)
	FrontendInvalidWorkflowID                           = 2001
	FrontendInvalidRequestFormat                        = 2002
	FrontendFrontendFxFailed                            = 2010 // service/frontend/fx.go:235 - creating gRPC server options failed
	FrontendFrontendHttpapiserverFailed                 = 2011 // service/frontend/http_api_server.go:309 - Failed to marshal error message
	FrontendFrontendNexushandlerInvalidOperationFailed  = 2012 // service/frontend/nexus_handler.go:414 - invalid input
	FrontendFrontendNexushandlerInvalidOperationFailed2 = 2013 // service/frontend/nexus_handler.go:630 - invalid Nexus cancel operation.
	FrontendFrontendServiceFailed                       = 2014 // service/frontend/service.go:427 - Failed to serve on frontend listener
	FrontendFrontendServiceFailed2                      = 2015 // service/frontend/service.go:434 - Failed to serve HTTP API server
	FrontendFrontendWorkflowhandlerOperationFailed      = 2016 // service/frontend/workflow_handler.go:4753 - Unknown batch operation type
	FrontendFrontendWorkflowhandlerError                = 2017 // service/frontend/workflow_handler.go:5436
	FrontendAuthRequired                                = 2101
	FrontendInsufficientPerms                           = 2102
	FrontendNamespaceOpFailed                           = 2301
	FrontendWorkerHeartbeatFailed                       = 2401
	FrontendMatchingServiceCallFailed                   = 2402
	FrontendWorkerReachabilityFailed                    = 2403
	FrontendScheduleDecodingFailed                      = 2404
	FrontendNexusTaskQueuePollFailed                    = 2405
	FrontendSearchAttributesGetFailed                   = 2501
	FrontendReplicationMessagesServerCloseFailed        = 2502
	FrontendHistoryHostDescribeFailed                   = 2503
	FrontendScheduleMemoEncodingFailed                  = 2601
	FrontendNexusOperationPanicCaptured                 = 2701
	FrontendNexusPayloadSizeExceedsLimit                = 2702
	FrontendNexusLinkURLParsingFailed                   = 2703
	FrontendNexusForwardedStartOperationFailed          = 2704
	FrontendNexusForwardedCancelOperationFailed         = 2705
	FrontendNexusHTTPClientCreationFailed               = 2706
	FrontendNexusServiceBaseURLConstructionFailed       = 2707
	FrontendNexusFailureMarshalingFailed                = 2708
	FrontendNexusResponseBodyWriteFailed                = 2709
	FrontendNexusInvalidURLProvided                     = 2710
	FrontendNexusInvalidNamespaceName                   = 2711
	FrontendNexusClaimsRetrievalFailed                  = 2712
	FrontendNexusInvalidEndpointID                      = 2713
	FrontendNexusNamespaceLookupFailed                  = 2714
	FrontendOpenAPISpecReaderInitFailed                 = 2715
	FrontendOpenAPISpecSendFailed                       = 2716
	FrontendOpenAPISpecChecksumVerificationFailed       = 2717
	FrontendNexusEndpointsPersistenceListingFailed      = 2718
	FrontendNexusEndpointClientGenericError             = 2719
	FrontendTaskQueueKindUnspecified                    = 2901

	// History Service Range: 3000-3999 (sorted by error code)
	HistoryWorkflowNotFound                          = 3001
	HistoryInvalidState                              = 3002
	HistoryWorkflowExists                            = 3003
	HistoryTaskProcessingFailed                      = 3004
	HistoryCurrentBranchChanged                      = 3005
	HistoryConditionFailed                           = 3006
	HistoryTransferRequestCancelFailed               = 3007
	HistoryTransferSignalFailed                      = 3008
	HistoryTransferChildWorkflowFailed               = 3009
	HistoryTransferAutoResetFailed                   = 3010
	HistoryMutableStateDirtyTransaction              = 3011
	HistoryDataInconsistency                         = 3012
	HistorySyncVersionedTransitionMissing            = 3013
	HistoryTaskProcessorNotRegistered                = 3014
	HistoryEngineRetrievalFailed                     = 3015
	HistoryDLQReplicationTasksFailed                 = 3016
	HistoryVisibilityURIParsingFailed                = 3017
	HistoryArchivalURIParsingFailed                  = 3018
	HistoryTaskPayloadSerializationFailed            = 3019
	HistoryQueueExecutablePanicCaptured              = 3020
	HistoryOutboundTaskExecutorFailed                = 3021
	HistoryHistoryDescribeworkflowError2             = 3022 // service/history/api/describeworkflow/api.go:249
	HistoryHistoryDescribeworkflowError3             = 3023 // service/history/api/describeworkflow/api.go:261
	HistoryHistoryReapplyeventsResetOperationFailed  = 3024 // service/history/api/reapplyevents/api.go:146 - Cannot reset workflow. Ignoring reapply events.
	HistoryHistoryTrimhistoryutilError               = 3025 // service/history/api/trim_history_util.go:53
	HistoryHistoryUpdateworkflowError                = 3026 // service/history/api/updateworkflow/api.go:232
	HistoryHistoryHandlerOperationFailed             = 3027 // service/history/handler.go:1674 - History engine not found for shard
	HistoryHistoryHandlerOperationFailed2            = 3028 // service/history/handler.go:1679 - History engine not found for shard
	HistoryHistoryHandlerFailed                      = 3029 // service/history/handler.go:1691 - Failed to get replication tasks for shard
	HistoryHistoryHandlerOperationFailed3            = 3030 // service/history/handler.go:1752 - History engine not found for workflow ID.
	HistoryHistoryHandlerOperationFailed4            = 3031 // service/history/handler.go:1757 - History engine not found for shard ID.
	HistoryHistoryNdcError                           = 3032 // service/history/ndc/history_replicator.go:366
	HistoryHistoryNdcError2                          = 3033 // service/history/ndc/history_replicator.go:374
	HistoryHistoryNdcError3                          = 3034 // service/history/ndc/history_replicator.go:537
	HistoryHistoryNdcError4                          = 3035 // service/history/ndc/history_replicator.go:557
	HistoryHistoryReapplyeventsFailed                = 3036 // service/history/api/reapplyevents/api.go:166 - failed to re-apply stale events
	HistoryHistoryReplicationFailed                  = 3037 // service/history/api/replication/get_dlq_tasks.go:22 - Failed to fetch DLQ replication messages.
	HistoryHistoryReplicationError                   = 3038 // service/history/api/replication/get_tasks.go:62 - error updating replication level for shard
	HistoryHistoryReplicationFailed2                 = 3039 // service/history/api/replication/get_tasks.go:73 - Failed to retrieve replication messages.
	HistoryHistoryRespondworkflowtaskcompletedError  = 3040 // service/history/api/respondworkflowtaskcompleted/api.go:228
	HistoryHistoryRespondworkflowtaskcompletedError2 = 3041 // service/history/api/respondworkflowtaskcompleted/api.go:948
	HistoryHistoryDescribeworkflowError              = 3042 // service/history/api/describeworkflow/api.go:227
	HistoryHistoryQueuesError                        = 3052 // service/history/queues/dlq_writer.go:106
	HistoryHistoryQueuesError6                       = 3059 // service/history/queues/scheduler.go:106
	HistoryServiceStartupFailed                      = 3200 // service/history/service.go:93 - Failed to serve on history listener
	HistoryArchivalWorkflowFailed                    = 3201 // service/history/archival/archiver.go:132 - failed to archive workflow

	// Matching Service Range: 4000-4999 (sorted by error code)
	MatchingPartitionFailed                    = 4001
	MatchingTaskDroppedNonRetryable            = 4002
	MatchingDeploymentRegistrationError        = 4003
	MatchingPartitionForceLoadFailed           = 4004
	MatchingMatchingMatchingengineError        = 4008 // service/matching/matching_engine.go:708
	MatchingMatchingMatchingengineError2       = 4009 // service/matching/matching_engine.go:742
	MatchingMatchingMatchingengineError3       = 4010 // service/matching/matching_engine.go:840
	MatchingMatchingMatchingengineError5       = 4012 // service/matching/matching_engine.go:926
	MatchingMatchingMatchingengineError6       = 4013 // service/matching/matching_engine.go:960
	MatchingMatchingPritaskreaderError         = 4015 // service/matching/pri_task_reader.go:328 - nonretryable error processing spooled task
	MatchingMatchingTaskreaderError            = 4017 // service/matching/task_reader.go:115 - taskReader: unexpected error dispatching task
	MatchingHeaderParsingFailed                = 4100 // service/matching/physical_task_queue_manager.go:492 - unable to parse header
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
	WorkerScannerHeartbeatRecoveryFailed                   = 5201
	WorkerReplicationTasksFetchFailed                      = 5301
	WorkerReplicationTasksApplyFailed                      = 5302
	WorkerReplicationTasksDLQPutFailed                     = 5303
	WorkerNamespaceReplicationTaskProcessingFailed         = 5304
	WorkerTaskQueueUserDataReplicationTaskProcessingFailed = 5305
	WorkerDeleteNamespaceChildWorkflowError                = 5401

	// Batcher Operations (5500-5599)
	WorkerBatchOperationNamespaceMismatch              = 5501
	WorkerBatchResetOptionsDeserializationFailed       = 5502
	WorkerBatchPostResetOperationDeserializationFailed = 5503
	WorkerBatchHeartbeatRecoveryFailed                 = 5504
	WorkerBatchWorkflowCountEstimationFailed           = 5505
	WorkerBatchOperationTaskProcessingFailed           = 5506
	WorkerBatchWorkflowHistoryReverseReadFailed        = 5507
	WorkerBatchWorkflowHistoryReadFailed               = 5508

	// Worker Deployment Operations (5600-5699)
	WorkerDeploymentQueryHandlerSetupFailed        = 5601
	WorkerDeploymentWorkflowLockAcquisitionFailed  = 5602
	WorkerDeploymentUpdateCanceledBeforeStart      = 5603
	WorkerDeploymentVersionCheckFailed             = 5604
	WorkerDeploymentVersionWorkflowExecutionFailed = 5605
	WorkerDeploymentVersionRegistrationFailed      = 5606
	WorkerDeploymentVersionPollingFailed           = 5607
	WorkerDeploymentVersionValidationFailed        = 5608

	// Scheduler Operations (5700-5799)
	WorkerSchedulerActionExecutionFailed   = 5701
	WorkerSchedulerPolicyValidationFailed  = 5702
	WorkerSchedulerTriggerEvaluationFailed = 5703
	WorkerSchedulerStateTransitionFailed   = 5704

	// Scanner Operations (5800-5899)
	WorkerScannerExecutionTaskProcessingFailed = 5801
	WorkerScannerHistoryScavengingFailed       = 5802

	// Migration Operations (5900-5999)
	WorkerMigrationActivityExecutionFailed = 5901

	// Delete Namespace Operations (6000-6099)
	WorkerDeleteNamespaceExecutionsFailed       = 6001
	WorkerDeleteNamespaceReclaimResourcesFailed = 6002
	WorkerDeleteNamespaceActivitiesFailed       = 6003

	// Common Components Range: 7000-7999

	// Persistence Telemetry (7000-7099)
	CommonPersistenceTelemetryOTELSerializationFailed = 7001

	// Archival Components (7100-7199)
	CommonHistoryArchivalOperationFailed    = 7101
	CommonVisibilityArchivalOperationFailed = 7102
	CommonArchivalUploadFailed              = 7103

	// Tools and Utilities (7200-7299)
	ToolsCQLClientOperationFailed   = 7201
	ToolsSchemaEmbedOperationFailed = 7202

	// Component Operations (7300-7399)
	ComponentNexusOperationHandlerFailed = 7301
	ComponentCallbackInvocationFailed    = 7302
	ComponentSchedulerExecutorFailed     = 7303

	// Infrastructure Operations (7400-7499)
	InfraServiceLifecycleOperationFailed = 7401
	InfraMetricsProviderOperationFailed  = 7402
	InfraTaskSchedulerOperationFailed    = 7403

	// Persistence Operations (7500-7599)
	PersistenceElasticsearchProcessorOperationFailed = 7501
	PersistenceVisibilityStoreOperationFailed        = 7502
	PersistenceNDCHistoryImporterOperationFailed     = 7503
	PersistenceHistoryManagerOperationFailed         = 7504

	// History Service Operations (7800-7899)
	HistoryWorkflowTransactionOperationFailed          = 7801
	HistoryReplicationTaskProcessorOperationFailed     = 7802
	HistoryReplicationTaskExecutorOperationFailed      = 7803
	HistoryAPIGetHistoryUtilityOperationFailed         = 7804
	HistoryRespondWorkflowTaskCompletedOperationFailed = 7805

	// Membership Operations (7900-7999)
	MemberRingpopTestClusterOperationFailed = 7901
	MemberRingpopMonitorOperationFailed     = 7902

	// Namespace Operations (8000-8099)

	// Generated Error Codes - Each logging line gets a unique error code
	// Component Service - Generated Error Codes (10 codes)
	ComponentCallbacksError                        = 7300 // components/callbacks/fx.go:59
	ComponentCallbacksFailed                       = 7306 // components/callbacks/hsm_invocation.go:86 - Callback request failed
	ComponentNexusOperationsFailed                 = 7307 // components/nexusoperations/executors.go:261 - Nexus StartOperation request failed
	ComponentNexusOperationsError                  = 7308 // components/nexusoperations/executors.go:354
	ComponentNexusOperationsError2                 = 7309 // components/nexusoperations/executors.go:356
	ComponentNexusOperationsInvalidOperationFailed = 7310 // components/nexusoperations/executors.go:367 - invalid link data type: %q
	ComponentNexusOperationsFailed2                = 7311 // components/nexusoperations/executors.go:592 - Nexus CancelOperation request failed
	ComponentsSchedulerOperationFailed             = 7316 // components/scheduler/generator_executors.go:66 - Time went backwards
	ComponentSchedulerOperationFailed2             = 7317 // components/scheduler/spec_processor.go:105 - Schedule missed catchup window
	ComponentSchedulerInvalidOperationFailed       = 7318 // components/scheduler/spec_processor.go:152 - Invalid schedule

	// Worker Service - Generated Error Codes (114 codes)
	WorkerWorkerBatcherFailed = 5009 // service/worker/batcher/activities.go:238 - Failed to complete batch operation

	// Common Service - Generated Error Codes (192 codes)
	CommonArchiverFilestoreError  = 7000 // common/archiver/filestore/history_archiver.go:122
	CommonArchiverFilestoreError2 = 7002 // common/archiver/filestore/history_archiver.go:127
	CommonArchiverFilestoreError4 = 7004 // common/archiver/filestore/history_archiver.go:150
	CommonArchiverFilestoreError5 = 7005 // common/archiver/filestore/history_archiver.go:152
	CommonArchiverFilestoreError6 = 7006 // common/archiver/filestore/history_archiver.go:158
	CommonArchiverFilestoreError7 = 7007 // common/archiver/filestore/history_archiver.go:168
	CommonArchiverFilestoreError8 = 7008 // common/archiver/filestore/history_archiver.go:174
	CommonArchiverFilestoreError9 = 7009 // common/archiver/filestore/history_archiver.go:180

	// Client Service - Generated Error Codes (10 codes)
	ClientClientHistoryError  = 8003 // client/history/caching_redirector.go:224 - Error adding listener
	ClientClientHistoryError2 = 8004 // client/history/caching_redirector.go:228 - Error removing listener
	ClientClientHistoryFailed = 8005 // client/history/client.go:148 - Failed to get replication tasks from client
	ClientClientHistoryError3 = 8006 // client/history/metric_client.go:79 - history client encountered error

	// Tools Service - Generated Error Codes (32 codes)
	ToolsCassandraOperationFailed       = 7207 // tools/cassandra/handler.go:31 - Unable to read config.
	ToolsCassandraOperationFailed2      = 7208 // tools/cassandra/handler.go:36 - Unable to establish CQL session.
	ToolsCassandraOperationFailed3      = 7209 // tools/cassandra/handler.go:41 - Unable to setup CQL schema.
	ToolsCassandraOperationFailed4      = 7210 // tools/cassandra/handler.go:52 - Unable to read config.
	ToolsCassandraOperationFailed5      = 7211 // tools/cassandra/handler.go:57 - Unable to establish CQL session.
	ToolsCassandraUpdateOperationFailed = 7212 // tools/cassandra/handler.go:62 - Unable to update CQL schema.
	ToolsCassandraOperationFailed6      = 7213 // tools/cassandra/handler.go:71 - Unable to read config.
	ToolsCassandraOperationFailed7      = 7214 // tools/cassandra/handler.go:77 - Unable to read config.
	ToolsCassandraCreateOperationFailed = 7215 // tools/cassandra/handler.go:82 - Unable to create keyspace.
	ToolsCassandraOperationFailed8      = 7216 // tools/cassandra/handler.go:91 - Unable to read config.
	ToolsCassandraOperationFailed9      = 7217 // tools/cassandra/handler.go:97 - Unable to read config.
	ToolsCassandraOperationFailed10     = 7218 // tools/cassandra/handler.go:102 - Unable to drop keyspace.
	ToolsCassandraOperationFailed11     = 7219 // tools/cassandra/handler.go:111 - Unable to read config.
	ToolsCassandraOperationFailed12     = 7220 // tools/cassandra/handler.go:119 - Unable to establish CQL session.
	ToolsCassandraError                 = 7221 // tools/cassandra/setup_task_tests.go:18 - Error creating CQLClient
	ToolsCassandraError2                = 7222 // tools/cassandra/update_task_tests.go:16 - Error creating CQLClient
	ToolsSqlClitestError                = 7223 // tools/sql/clitest/conn_tests.go:57 - error creating sql conn
	ToolsSqlClitestError2               = 7224 // tools/sql/clitest/update_task_tests.go:51 - Error creating CQLClient
	ToolsSqlOperationFailed             = 7225 // tools/sql/handler.go:23 - Unable to read config.
	ToolsSqlConnectOperationFailed      = 7226 // tools/sql/handler.go:28 - Unable to connect to SQL database.
	ToolsSqlOperationFailed2            = 7227 // tools/sql/handler.go:33 - Unable to setup SQL schema.
	ToolsSqlOperationFailed3            = 7228 // tools/sql/handler.go:44 - Unable to read config.
	ToolsSqlConnectOperationFailed2     = 7229 // tools/sql/handler.go:49 - Unable to connect to SQL database.
	ToolsSqlUpdateOperationFailed       = 7230 // tools/sql/handler.go:54 - Unable to update SQL schema.
	ToolsSqlOperationFailed4            = 7231 // tools/sql/handler.go:64 - Unable to read config.
	ToolsSqlCreateOperationFailed       = 7232 // tools/sql/handler.go:70 - Unable to create SQL database.
	ToolsSqlOperationFailed5            = 7233 // tools/sql/handler.go:91 - Unable to read config.
	ToolsSqlOperationFailed6            = 7234 // tools/sql/handler.go:97 - Unable to drop SQL database.
	ToolsTdbgFailed                     = 7235 // tools/tdbg/factory.go:117 - Failed to create connection
	ToolsTdbgFailed2                    = 7236 // tools/tdbg/factory.go:142 - Failed to load server CA certificate
	ToolsTdbgFailed3                    = 7237 // tools/tdbg/factory.go:150 - Failed to load client certificate

	// Test and Development Operations (7600-7699)

	// Log Migration Operations (7700-7799)
	CommonLogMigrationOperationFailed = 7701

	// Additional Common Components (8100-8199)
	CommonMetricsOperationFailed                  = 8101
	CommonDynamicConfigOperationFailed            = 8102
	CommonTaskSchedulerOperationFailed            = 8103
	CommonFinalizerOperationFailed                = 8104
	CommonSoftAssertOperationFailed               = 8105
	CommonXDCCacheOperationFailed                 = 8106
	CommonVisibilityManagerMetricsOperationFailed = 8107
	CommonPersistenceMetricClientOperationFailed  = 8108
	CommonNexusEndpointManagerOperationFailed     = 8109
	CommonNamespaceRegistryOperationFailed        = 8110
	CommonUtilityOperationFailed                  = 8111
	CommonFinalizerTimeout                        = 8112
	DeadlockDetected                              = 8113
	DeadlockProfileNotFound                       = 8114
	DeadlockProfileFailed                         = 8115
	PersistenceDLQListFailed                      = 8116
	PersistenceDLQProcessQueueNameFailed          = 8117
	PersistenceDLQCategoryNotFound                = 8118
	PersistenceDLQHistoryServiceLookupFailed      = 8119
	XDCCacheEventsTruncated                       = 8120

	// Schema and Embedding (8200-8299)
	SchemaEmbedOperationFailed = 8201

	// Additional Archival (8300-8399)

	// Additional History Service Operations (8400-8499)
	HistoryTaskPriorityUnknownKey      = 8401
	HistoryTaskPriorityUnknownType     = 8402
	HistoryArchiveTargetFailed         = 8403
	WorkerSchedulerResultSizeExceeded  = 8404
	WorkerSchedulerFailureSizeExceeded = 8405
	HistoryEventsCacheRetrieveFailed   = 8406
	HistoryEventsDataCorruption        = 8407

	// Namespace Replication Operations (8500-8599)
	CommonNamespaceReplicationOperationFailed = 8501
	NamespaceReplicationUUIDCollision         = 8502
	NamespaceReplicationCreationUUIDCollision = 8503
	NamespaceReplicationCreationError         = 8504
	NamespaceReplicationCreationNameCollision = 8505

	// TLS/Encryption Operations (8600-8699)
	TLSPerHostProviderLookupError = 8601
	TLSCertExpirationCheckError   = 8602
	TLSCertExpired                = 8603

	// Nexus Operations (8700-8799)
	CommonNexusOperationFailed = 8701

	// Persistence Operations (8800-8899)
	PersistenceCassandraDropKeyspaceError   = 8801
	PersistenceCassandraCreateKeyspaceError = 8802
	PersistenceSQLCloseDatabaseError        = 8803
	PersistenceSQLTransactionRollbackError  = 8804

	// Worker Operations (8900-8999)
	WorkerAddSearchAttributesESMappingRetryable    = 8901
	WorkerAddSearchAttributesESMappingNonRetryable = 8902
	WorkerAddSearchAttributesESStatusFailed        = 8903
	ChildWorkflowError                             = 8904
	TaskSubmitToExecutorFailed                     = 8905

	// Common Archiver Operations (9000-9099)
	CommonArchiverOperationFailed = 9001

	// Additional Missing Operations (9100-9199)
	ReplicationTaskGenerationFailed        = 9101
	NamespaceGetForReplicationFailed       = 9102
	WorkerSchedulerStateInconsistent       = 9103
	TaskQueueUserDataUpdateFailed          = 9104
	TaskQueueUserDataSyncFailed            = 9105
	TaskQueueUserDataWaitFailed            = 9106
	DeploymentClientError                  = 9107
	WorkerDeploymentOperationFailed        = 9108
	TaskQueueAddRateFetchError             = 9109
	WorkflowCountForDrainageError          = 9110
	MembershipListenerRemovalFailed        = 9111
	HistoryEventGetFailed                  = 9112
	ImportActionCommitFailed               = 9113
	HistoryGetError                        = 9114
	EventReplicationFailed                 = 9115
	HistoryQueueUnknownAlertType           = 9116
	HistoryQueueTaskRangeCompleteFailed    = 9117
	HistoryQueueUpdateStateFailed          = 9118
	HistoryQueueTaskLoadFailed             = 9119
	HistoryQueueRateLimiterConfiguration   = 9120
	HistoryQueueTaskRetrieveFailed         = 9121
	HistorySchedulerCreationFailed         = 9122
	StreamCloseError                       = 9123
	ReplicationTaskSaveFailed              = 9124
	HistoryReplicationTaskConversionFailed = 9125
	RateLimiterWaitFailed                  = 9126
	VersionHistoryGetFailed                = 9127
	HistoryWorkflowSizeConstraintViolation = 9128
	MatchingPersistentStoreFailure         = 9129
	ReplicationTaskFetchFailed             = 9130
	HistoryReplicationCleanupFailed        = 9131
	HistoryReplicationDLQOperationFailed   = 9132
	TaskQueueDeleteError                   = 9133
	ScavengerDeleteHandlerError            = 9134
	TaskQueueListError                     = 9135
	HistorySyncWorkflowStateRetrieveFailed = 9136
	HistorySyncWorkflowStateUpdateFailed   = 9137
	WorkflowStartError                     = 9138
)
