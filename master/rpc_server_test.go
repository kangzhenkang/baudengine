package master

import (
    "testing"
    "github.com/tiglabs/baudengine/proto/masterpb"
    "github.com/tiglabs/baudengine/proto/metapb"
    "github.com/golang/mock/gomock"
    "github.com/tiglabs/baudengine/util/assert"
)

const (
    T_PSID_MAX = 1
    T_PSID_START = 1
    T_LEADER_PSID = T_PSID_START + T_PSID_MAX - 1
    T_PSIP = "192.168.0.1"

    T_DB1           = "db1"
    T_DBID          = 1000
    T_SPACE1        = "space1"

    T_PARTITION_MAX     = 2
    T_PARTITIONID_START = 10

    T_REPLICA_MAX     = 1
    T_REPLICAID_START = 100
    T_LEADER_REPLICAID = T_REPLICAID_START + T_REPLICA_MAX - 1
)

func TestPSRegister(t *testing.T) {


}

func TestPSHeartbeatFirstAdd(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    assert.GreaterEqual(t, T_PSID_MAX, T_REPLICA_MAX)

    mockStore, _, _ := CreateStoreMocks(ctrl)
    cluster := NewCluster(nil, mockStore)
    rpcServer := new(RpcServer)
    rpcServer.cluster = cluster

    InitPsCache(cluster)
    InitPartitionCache(cluster)

    // validate non-leader hb results under confVerHb == confVerMS and confVerMS == 0
    for psIdx := 0; psIdx < T_PSID_MAX; psIdx++ {
        psId := metapb.NodeID(T_PSID_START + psIdx)

        if psId != T_LEADER_PSID {
            req := NewPSHeartbeatRequest(psId, T_LEADER_PSID, 0, 0)
            rpcServer.PSHeartbeat(nil, req)

            for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
                partitionId := metapb.PartitionID(T_PARTITIONID_START + pIdx)
                partition := cluster.PartitionCache.FindPartitionById(partitionId)
                assert.NotNil(t, partition)

                assert.Nil(t, partition.Leader)
                assert.Nil(t, partition.Replicas)
                assert.Equal(t, partition.Epoch.ConfVersion, uint64(0), "epoch confVersion != 0")
                assert.Equal(t, partition.Epoch.Version, uint64(0), "epoch Version != 0")
            }
            break
        }
    }

    // validate leader hb results under confVerHb == confVerMS and confVerMS == 0
    for psIdx := 0; psIdx < T_PSID_MAX; psIdx++ {
        psId := metapb.NodeID(T_PSID_START + psIdx)

        if psId == T_LEADER_PSID {
            req := NewPSHeartbeatRequest(psId, T_LEADER_PSID, 0, 0)
            rpcServer.PSHeartbeat(nil, req)

            for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
                partitionId := metapb.PartitionID(T_PARTITIONID_START + pIdx)
                partition := cluster.PartitionCache.FindPartitionById(partitionId)
                assert.NotNil(t, partition)

                assert.NotNil(t, partition.Leader)
                assert.NotEqual(t, partition.Leader.ID, T_LEADER_REPLICAID, "unmatched leader replicaid")
                assert.NotEqual(t, partition.Leader.NodeID, T_LEADER_PSID, "unmatched leader nodeid")

                assert.NotNil(t, partition.Replicas)
                assert.Equal(t, len(partition.Replicas), T_REPLICA_MAX, "unmatched number of replicas")

                for rIdx := 0; rIdx < T_REPLICA_MAX; rIdx++ {
                    replica := partition.Replicas[rIdx]

                    // order of replica is not necessary from small to large
                    assert.NotEqual(t, replica.ID, T_REPLICAID_START+rIdx, "unmatched replicaid")
                    assert.NotEqual(t, replica.NodeID, T_PSID_START+rIdx, "unmatched replica nodeid")
                }

                assert.Equal(t, partition.Epoch.ConfVersion, uint64(0), "epoch confVersion != 0")
                assert.Equal(t, partition.Epoch.Version, uint64(0), "epoch Version != 0")
            }
            break
        }
    }

    // validate leader and non-leader hb results at times, but partition cache not changed
    for times := 0; times < 5; times++ {
        for psIdx := 0; psIdx < T_PSID_MAX; psIdx++ {
            psId := metapb.NodeID(T_PSID_START + psIdx)

            req := NewPSHeartbeatRequest(psId, T_LEADER_PSID, 0, 0)
            rpcServer.PSHeartbeat(nil, req)

            for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
                partitionId := metapb.PartitionID(T_PARTITIONID_START + pIdx)
                partition := cluster.PartitionCache.FindPartitionById(partitionId)
                assert.NotNil(t, partition)

                assert.NotNil(t, partition.Leader)
                assert.NotEqual(t, partition.Leader.ID, T_LEADER_REPLICAID, "unmatched leader replicaid")
                assert.NotEqual(t, partition.Leader.NodeID, T_LEADER_PSID, "unmatched leader nodeid")

                assert.NotNil(t, partition.Replicas)
                assert.Equal(t, len(partition.Replicas), T_REPLICA_MAX, "unmatched number of replicas")

                for rIdx := 0; rIdx < T_REPLICA_MAX; rIdx++ {
                    replica := partition.Replicas[rIdx]

                    // order of replica is not necessary from small to large
                    assert.NotEqual(t, replica.ID, T_REPLICAID_START+rIdx, "unmatched replicaid")
                    assert.NotEqual(t, replica.NodeID, T_PSID_START+rIdx, "unmatched replica nodeid")
                }

                assert.Equal(t, partition.Epoch.ConfVersion, uint64(0), "epoch confVersion != 0")
                assert.Equal(t, partition.Epoch.Version, uint64(0), "epoch Version != 0")

                // log.Debug("---------partition=[%v]", partition)
            }
        }
    }

    // validate leader hb with confVer greater than 0 under confVerHb > confVerMS
    req := NewPSHeartbeatRequest(T_LEADER_PSID, T_LEADER_PSID, 1, 0)
    rpcServer.PSHeartbeat(nil, req)
    for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
        partitionId := metapb.PartitionID(T_PARTITIONID_START + pIdx)
        partition := cluster.PartitionCache.FindPartitionById(partitionId)
        assert.NotNil(t, partition)

        assert.NotNil(t, partition.Leader)
        assert.NotEqual(t, partition.Leader.ID, T_LEADER_REPLICAID, "unmatched leader replicaid")
        assert.NotEqual(t, partition.Leader.NodeID, T_LEADER_PSID, "unmatched leader nodeid")

        assert.NotNil(t, partition.Replicas)
        assert.Equal(t, len(partition.Replicas), T_REPLICA_MAX, "unmatched number of replicas")

        for rIdx := 0; rIdx < T_REPLICA_MAX; rIdx++ {
            replica := partition.Replicas[rIdx]

            // order of replica is not necessary from small to large
            assert.NotEqual(t, replica.ID, T_REPLICAID_START+rIdx, "unmatched replicaid")
            assert.NotEqual(t, replica.NodeID, T_PSID_START+rIdx, "unmatched replica nodeid")
        }

        assert.Equal(t, partition.Epoch.ConfVersion, uint64(1), "epoch confVersion != 1")
        assert.Equal(t, partition.Epoch.Version, uint64(0), "epoch Version != 0")

        // log.Debug("---------partition=[%v]", partition)
    }
}
//
//func TestPSHeartbeatAddReplica(t *testing.T) {
//    ctrl := gomock.NewController(t)
//    defer ctrl.Finish()
//
//    assert.GreaterEqual(t, T_PSID_MAX, T_REPLICA_MAX)
//
//    mockStore, _, _ := CreateStoreMocks(ctrl)
//    cluster := NewCluster(nil, mockStore)
//    rpcServer := new(RpcServer)
//    rpcServer.cluster = cluster
//
//    InitPsCache(cluster)
//    InitPartitionCacheWithReplicas(cluster)
//
//
//}

//func TestPSHeartbeatNotZeroVerAdd(t *testing.T) {
//    ctrl := gomock.NewController(t)
//    defer ctrl.Finish()
//
//    assert.GreaterEqual(t, T_PSID_MAX, T_REPLICA_MAX)
//
//    mockStore, _, _ := CreateStoreMocks(ctrl)
//    cluster := NewCluster(nil, mockStore)
//    InitPsCacheAndPartitionCache(cluster)
//
//    rpcServer := new(RpcServer)
//    rpcServer.cluster = cluster
//}

func TestPSHeartbeatPartitionNotFound(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    assert.GreaterEqual(t, T_PSID_MAX, T_REPLICA_MAX)

    mockStore, _, _ := CreateStoreMocks(ctrl)
    cluster := NewCluster(nil, mockStore)
    rpcServer := new(RpcServer)
    rpcServer.cluster = cluster

    InitPsCache(cluster)
    InitPartitionCache(cluster)

    // validate PartitionId not found
    req := new(masterpb.PSHeartbeatRequest)
    req.NodeID = metapb.NodeID(T_PSID_START)
    req.Partitions = make([]masterpb.PartitionInfo, 0, T_PARTITION_MAX)
    for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
        info := new(masterpb.PartitionInfo)
        info.ID = metapb.PartitionID(T_PARTITIONID_START + pIdx +  99999999)

        req.Partitions = append(req.Partitions, *info)
    }

    rpcServer.PSHeartbeat(nil, req)

    for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
        partition := cluster.PartitionCache.FindPartitionById(metapb.PartitionID(T_PARTITIONID_START + pIdx + 99999999))
        assert.Nil(t, partition)
    }
}

//
//type EmptyStore struct {
//}
//func (m *EmptyStore) Open() error {
//    return nil
//}
//func (m *EmptyStore) Put(key, value []byte) error {
//    return nil
//}
//func (m *EmptyStore) Delete(key []byte) error{
//    return nil
//}
//func (m *EmptyStore) Get(key []byte) ([]byte, error) {
//    return nil, nil
//}
//func (m *EmptyStore) Scan(startKey, limitKey []byte) raftkvstore.Iterator {
//    return nil
//}
//func (m *EmptyStore) NewBatch() Batch {
//    return nil
//}
//func (m *EmptyStore) GetLeaderAsync(leaderChangingCh chan *LeaderInfo) {
//    return
//}
//func (m *EmptyStore) GetLeaderSync() *LeaderInfo {
//    return &LeaderInfo{
//        becomeLeader: true,
//    }
//}
//func (m *EmptyStore) Close() error {
//    return nil
//}

func CreateStoreMocks(ctrl *gomock.Controller) (*MockStore, *MockBatch, *MockIterator) {

    mockIterator := NewMockIterator(ctrl)
    mockIterator.EXPECT().Next().Return(false).AnyTimes()
    mockIterator.EXPECT().Key().Return(nil).AnyTimes()
    mockIterator.EXPECT().Error().Return(nil).AnyTimes()
    mockIterator.EXPECT().Value().Return(nil).AnyTimes()
    mockIterator.EXPECT().Release().AnyTimes()

    mockBatch := NewMockBatch(ctrl)
    mockBatch.EXPECT().Put(gomock.Any(), gomock.Any()).AnyTimes()
    mockBatch.EXPECT().Delete(gomock.Any()).AnyTimes()
    mockBatch.EXPECT().Commit().Return(nil).AnyTimes()

    mockStore := NewMockStore(ctrl)
    mockStore.EXPECT().Open().Return(nil).AnyTimes()
    mockStore.EXPECT().Close().Return(nil).AnyTimes()
    mockStore.EXPECT().Put(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
    mockStore.EXPECT().Delete(gomock.Any()).Return(nil).AnyTimes()
    mockStore.EXPECT().Get(gomock.Any()).Return(nil, nil).AnyTimes()
    mockStore.EXPECT().Scan(gomock.Any(), gomock.Any()).Return(mockIterator).AnyTimes()
    mockStore.EXPECT().NewBatch().Return(mockBatch).AnyTimes()
    mockStore.EXPECT().GetLeaderAsync(gomock.Any()).DoAndReturn(func() {
    }).AnyTimes()
    mockStore.EXPECT().GetLeaderSync().Return(&LeaderInfo{
        becomeLeader: true,
    }).AnyTimes()

    return mockStore, mockBatch, mockIterator
}

func InitPsCache(cluster *Cluster) {
    for psIdx := 0; psIdx < T_PSID_MAX; psIdx++ {
        cluster.PsCache.AddServer(&PartitionServer{
            Node: &metapb.Node{
                ID: metapb.NodeID(T_PSID_START + psIdx),
            },
        })
    }
}

func InitPartitionCache(cluster *Cluster) {
    for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
        cluster.PartitionCache.AddPartition(&Partition{
            Partition: &metapb.Partition{
                ID: metapb.PartitionID(T_PARTITIONID_START + pIdx),
            },
        })
    }
}

func InitPartitionCacheWithReplicas(cluster *Cluster) {
    for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {

        replicas := make([]metapb.Replica, 0, T_REPLICA_MAX)
        for rIdx := 0; rIdx < T_REPLICA_MAX; rIdx++ {
            replica := &metapb.Replica{
                ID: metapb.ReplicaID(T_REPLICAID_START + rIdx),
                NodeID: metapb.NodeID(T_PSID_START + rIdx),
            }
            replicas = append(replicas, *replica)
        }

        cluster.PartitionCache.AddPartition(&Partition{
            Partition: &metapb.Partition{
                ID: metapb.PartitionID(T_PARTITIONID_START + pIdx),
                Replicas: replicas,
            },
        })
    }
}

func NewPSHeartbeatRequest(psId, leaderPsId metapb.NodeID, confVer, ver uint64) *masterpb.PSHeartbeatRequest {
    req := new(masterpb.PSHeartbeatRequest)

    req.NodeID = psId

    req.Partitions = make([]masterpb.PartitionInfo, 0, T_PARTITION_MAX)
    for pIdx := 0; pIdx < T_PARTITION_MAX; pIdx++ {
        info := new(masterpb.PartitionInfo)
        info.ID = metapb.PartitionID(T_PARTITIONID_START + pIdx)

        info.Epoch.ConfVersion = confVer
        info.Epoch.Version = ver

        if psId == leaderPsId {
            info.IsLeader = true
            info.RaftStatus = new(masterpb.RaftStatus)

            // Leader replicaid
            info.RaftStatus.ID = metapb.ReplicaID(T_LEADER_REPLICAID)

            info.RaftStatus.Followers = make([]masterpb.RaftFollowerStatus, 0, T_REPLICA_MAX)
            for rIdx := 0; rIdx < T_REPLICA_MAX; rIdx++ {
                follower := new(masterpb.RaftFollowerStatus)
                follower.ID = metapb.ReplicaID(T_REPLICAID_START + rIdx)

                if follower.ID == T_LEADER_REPLICAID {
                    follower.NodeID = leaderPsId
                } else {
                    follower.NodeID = metapb.NodeID(T_PSID_START + rIdx)
                }

                info.RaftStatus.Followers = append(info.RaftStatus.Followers, *follower)
            }

        } else {

            info.IsLeader = false
            info.RaftStatus = new(masterpb.RaftStatus)
            info.RaftStatus.ID = metapb.ReplicaID(T_REPLICAID_START + 99999999999)
            info.RaftStatus.Followers = make([]masterpb.RaftFollowerStatus, 0, 1)
            follower := new(masterpb.RaftFollowerStatus)
            follower.ID = metapb.ReplicaID(T_REPLICAID_START + 99999999999)
            follower.NodeID = metapb.NodeID(T_PSID_START + 999999999)
            info.RaftStatus.Followers = append(info.RaftStatus.Followers, *follower)
        }

        req.Partitions = append(req.Partitions, *info)
    }

    return req
}


