package replication

import (
	"context"
	"fmt"
	"math"
	"time"

	"go.temporal.io/api/serviceerror"
	enumsspb "go.temporal.io/server/api/enums/v1"
	persistencespb "go.temporal.io/server/api/persistence/v1"
	replicationspb "go.temporal.io/server/api/replication/v1"
	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/persistence"
	historyi "go.temporal.io/server/service/history/interfaces"
	"go.temporal.io/server/service/history/replication"
	"go.temporal.io/server/service/history/shard"
	"go.temporal.io/server/service/history/tasks"
)

func GetTasks(
	ctx context.Context,
	shardContext historyi.ShardContext,
	replicationAckMgr replication.AckManager,
	pollingCluster string,
	ackMessageID int64,
	ackTimestamp time.Time,
	queryMessageID int64,
) (*replicationspb.ReplicationMessages, error) {
	allClusterInfo := shardContext.GetClusterMetadata().GetAllClusterInfo()
	clusterInfo, ok := allClusterInfo[pollingCluster]
	if !ok {
		return nil, serviceerror.NewInternal(
			fmt.Sprintf("missing cluster info for cluster: %v", pollingCluster),
		)
	}
	readerID := shard.ReplicationReaderIDFromClusterShardID(
		clusterInfo.InitialFailoverVersion,
		shardContext.GetShardID(),
	)

	if ackMessageID != persistence.EmptyQueueMessageID {
		if err := shardContext.UpdateReplicationQueueReaderState(
			readerID,
			&persistencespb.QueueReaderState{
				Scopes: []*persistencespb.QueueSliceScope{{
					Range: &persistencespb.QueueSliceRange{
						InclusiveMin: shard.ConvertToPersistenceTaskKey(
							tasks.NewImmediateKey(ackMessageID + 1),
						),
						ExclusiveMax: shard.ConvertToPersistenceTaskKey(
							tasks.NewImmediateKey(math.MaxInt64),
						),
					},
					Predicate: &persistencespb.Predicate{
						PredicateType: enumsspb.PREDICATE_TYPE_UNIVERSAL,
						Attributes:    &persistencespb.Predicate_UniversalPredicateAttributes{},
					},
				}},
			},
		); err != nil {
			log.ErrorWithCode(shardContext.GetLogger(), errorcode.HistoryHistoryReplicationError, "error updating replication level for shard", err, tag.OperationFailed)
		}
		shardContext.UpdateRemoteClusterInfo(pollingCluster, ackMessageID, ackTimestamp)
	}

	replicationMessages, err := replicationAckMgr.GetTasks(
		ctx,
		pollingCluster,
		queryMessageID,
	)
	if err != nil {
		log.ErrorWithCode(shardContext.GetLogger(), errorcode.HistoryHistoryReplicationFailed2, "Failed to retrieve replication messages.", err)
		return nil, err
	}

	shardContext.GetLogger().Debug("Successfully fetched replication messages.", tag.Counter(len(replicationMessages.ReplicationTasks)))
	return replicationMessages, nil
}
