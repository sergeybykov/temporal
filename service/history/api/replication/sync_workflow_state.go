package replication

import (
	"context"

	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/service/history/replication"
)

func SyncWorkflowState(
	ctx context.Context,
	request *historyservice.SyncWorkflowStateRequest,
	replicationProgressCache replication.ProgressCache,
	syncStateRetriever replication.SyncStateRetriever,
	logger log.Logger,
) (_ *historyservice.SyncWorkflowStateResponse, retError error) {
	result, err := syncStateRetriever.GetSyncWorkflowStateArtifact(ctx, request.GetNamespaceId(), request.Execution, request.VersionedTransition, request.VersionHistories)
	if err != nil {
		log.ErrorWithCode(logger, errorcode.HistorySyncWorkflowStateRetrieveFailed, "SyncWorkflowState failed to retrieve sync state artifact", err,
			tag.WorkflowNamespaceID(request.NamespaceId),
			tag.WorkflowID(request.Execution.WorkflowId),
			tag.WorkflowRunID(request.Execution.RunId))
		return nil, err
	}

	err = replicationProgressCache.Update(request.Execution.RunId, request.TargetClusterId, result.VersionedTransitionHistory, result.SyncedVersionHistory.Items)
	if err != nil {
		log.ErrorWithCode(logger, errorcode.HistorySyncWorkflowStateUpdateFailed, "SyncWorkflowState failed to update progress cache", err,
			tag.WorkflowNamespaceID(request.NamespaceId),
			tag.WorkflowID(request.Execution.WorkflowId),
			tag.WorkflowRunID(request.Execution.RunId))
	}

	return &historyservice.SyncWorkflowStateResponse{
		VersionedTransitionArtifact: result.VersionedTransitionArtifact,
	}, nil
}
