package master

import (
	"github.com/pkg/errors"
)

//master global error definitions
var (
	ErrSuc           = errors.New("success")
	ErrInternalError = errors.New("internal error")
	ErrSysBusy       = errors.New("system busy")
	ErrParamError    = errors.New("param error")
	ErrInvalidCfg    = errors.New("config error")

	ErrDupDb              = errors.New("duplicated database")
	ErrDbNotExists        = errors.New("db not exists")
	ErrDupSpace           = errors.New("duplicated space")
	ErrSpaceNotExists     = errors.New("space not exists")
	ErrPSNotExists        = errors.New("partition server is not exists")
	ErrGenIdFailed        = errors.New("generate id is failed")
	ErrBoltDbOpsFailed    = errors.New("boltdb operation error")
	ErrUnknownRaftCmdType = errors.New("unknown raft command type")
	ErrRouteNotFound      = errors.New("route not found")
	//ErrEntryNotFound		    = errors.New("storage entry not found")

	ErrGrpcInvalidResp      = errors.New("invalid grpc response")
	ErrGrpcInvokeFailed     = errors.New("invoke grpc is failed")
	ErrGrpcParamError       = errors.New("grpc param error")
	ErrGrpcEmptyFollowers   = errors.New("reported empty followers")
	ErrGrpcInvalidFollowers = errors.New("reported invalid followers")

	ErrRaftNotRegHandler          = errors.New("have no register raft handler")
	ErrRaftInvalidNode            = errors.New("invalid raft node")
	ErrRaftUnknownType            = errors.New("unknown raft socket type")
	ErrRaftNoApplyHandler         = errors.New("raft group not register apply handler")
	ErrRaftNoPeerChangeHandler    = errors.New("raft group not register peer change handler")
	ErrRaftNoLeaderChangeHandler  = errors.New("raft group not register leader change handler")
	ErrRaftNoFatalEventHandler    = errors.New("raft group not register fatal event handler")
	ErrRaftNoSnapshotHandler      = errors.New("raft group not register snapshot handler")
	ErrRaftNoApplySnapshotHandler = errors.New("raft group not register apply snapshot handler")
	ErrRaftUnknownResponseType    = errors.New("unknown response type")
)

// http response error code and error message definitions
const (
	ERRCODE_SUCCESS = iota
	ERRCODE_INTERNAL_ERROR
	ERRCODE_SYSBUSY
	ERRCODE_PARAM_ERROR

	ERRCODE_DUP_DB
	ERRCODE_DB_NOTEXISTS
	ERRCODE_DUP_SPACE
	ERRCODE_SPACE_NOTEXISTS
	ERRCODE_PS_NOTEXISTS

	ERRCODE_GENID_FAILED
	ERRCODE_BOLTDB_OPTFAILED

//	ERRCODE_UNKNOWN_RAFTCMDTYPE
)

var Err2CodeMap = map[error]int32{
	ErrSuc:           ERRCODE_SUCCESS,
	ErrInternalError: ERRCODE_INTERNAL_ERROR,
	ErrSysBusy:       ERRCODE_SYSBUSY,
	ErrParamError:    ERRCODE_PARAM_ERROR,

	ErrDupDb:          ERRCODE_DUP_DB,
	ErrDbNotExists:    ERRCODE_DB_NOTEXISTS,
	ErrDupSpace:       ERRCODE_DUP_SPACE,
	ErrSpaceNotExists: ERRCODE_SPACE_NOTEXISTS,
	ErrPSNotExists:    ERRCODE_PS_NOTEXISTS,

	ErrGenIdFailed:     ERRCODE_GENID_FAILED,
	ErrBoltDbOpsFailed: ERRCODE_BOLTDB_OPTFAILED,
}
