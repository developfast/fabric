// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	common "github.com/hyperledger/fabric-protos-go-apiv2/common"
	ledger "github.com/hyperledger/fabric/core/ledger"

	mock "github.com/stretchr/testify/mock"
)

// Committer is an autogenerated mock type for the Committer type
type Committer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Committer) Close() {
	_m.Called()
}

// CommitLegacy provides a mock function with given fields: blockAndPvtData, commitOpts
func (_m *Committer) CommitLegacy(blockAndPvtData *ledger.BlockAndPvtData, commitOpts *ledger.CommitOptions) error {
	ret := _m.Called(blockAndPvtData, commitOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ledger.BlockAndPvtData, *ledger.CommitOptions) error); ok {
		r0 = rf(blockAndPvtData, commitOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommitPvtDataOfOldBlocks provides a mock function with given fields: reconciledPvtdata, unreconciled
func (_m *Committer) CommitPvtDataOfOldBlocks(reconciledPvtdata []*ledger.ReconciledPvtdata, unreconciled ledger.MissingPvtDataInfo) ([]*ledger.PvtdataHashMismatch, error) {
	ret := _m.Called(reconciledPvtdata, unreconciled)

	var r0 []*ledger.PvtdataHashMismatch
	var r1 error
	if rf, ok := ret.Get(0).(func([]*ledger.ReconciledPvtdata, ledger.MissingPvtDataInfo) ([]*ledger.PvtdataHashMismatch, error)); ok {
		return rf(reconciledPvtdata, unreconciled)
	}
	if rf, ok := ret.Get(0).(func([]*ledger.ReconciledPvtdata, ledger.MissingPvtDataInfo) []*ledger.PvtdataHashMismatch); ok {
		r0 = rf(reconciledPvtdata, unreconciled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ledger.PvtdataHashMismatch)
		}
	}

	if rf, ok := ret.Get(1).(func([]*ledger.ReconciledPvtdata, ledger.MissingPvtDataInfo) error); ok {
		r1 = rf(reconciledPvtdata, unreconciled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoesPvtDataInfoExistInLedger provides a mock function with given fields: blockNum
func (_m *Committer) DoesPvtDataInfoExistInLedger(blockNum uint64) (bool, error) {
	ret := _m.Called(blockNum)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (bool, error)); ok {
		return rf(blockNum)
	}
	if rf, ok := ret.Get(0).(func(uint64) bool); ok {
		r0 = rf(blockNum)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocks provides a mock function with given fields: blockSeqs
func (_m *Committer) GetBlocks(blockSeqs []uint64) []*common.Block {
	ret := _m.Called(blockSeqs)

	var r0 []*common.Block
	if rf, ok := ret.Get(0).(func([]uint64) []*common.Block); ok {
		r0 = rf(blockSeqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*common.Block)
		}
	}

	return r0
}

// GetConfigHistoryRetriever provides a mock function with given fields:
func (_m *Committer) GetConfigHistoryRetriever() (ledger.ConfigHistoryRetriever, error) {
	ret := _m.Called()

	var r0 ledger.ConfigHistoryRetriever
	var r1 error
	if rf, ok := ret.Get(0).(func() (ledger.ConfigHistoryRetriever, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ledger.ConfigHistoryRetriever); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ledger.ConfigHistoryRetriever)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentBlockHash provides a mock function with given fields:
func (_m *Committer) GetCurrentBlockHash() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMissingPvtDataTracker provides a mock function with given fields:
func (_m *Committer) GetMissingPvtDataTracker() (ledger.MissingPvtDataTracker, error) {
	ret := _m.Called()

	var r0 ledger.MissingPvtDataTracker
	var r1 error
	if rf, ok := ret.Get(0).(func() (ledger.MissingPvtDataTracker, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ledger.MissingPvtDataTracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ledger.MissingPvtDataTracker)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPvtDataAndBlockByNum provides a mock function with given fields: seqNum
func (_m *Committer) GetPvtDataAndBlockByNum(seqNum uint64) (*ledger.BlockAndPvtData, error) {
	ret := _m.Called(seqNum)

	var r0 *ledger.BlockAndPvtData
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*ledger.BlockAndPvtData, error)); ok {
		return rf(seqNum)
	}
	if rf, ok := ret.Get(0).(func(uint64) *ledger.BlockAndPvtData); ok {
		r0 = rf(seqNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ledger.BlockAndPvtData)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(seqNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPvtDataByNum provides a mock function with given fields: blockNum, filter
func (_m *Committer) GetPvtDataByNum(blockNum uint64, filter ledger.PvtNsCollFilter) ([]*ledger.TxPvtData, error) {
	ret := _m.Called(blockNum, filter)

	var r0 []*ledger.TxPvtData
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, ledger.PvtNsCollFilter) ([]*ledger.TxPvtData, error)); ok {
		return rf(blockNum, filter)
	}
	if rf, ok := ret.Get(0).(func(uint64, ledger.PvtNsCollFilter) []*ledger.TxPvtData); ok {
		r0 = rf(blockNum, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ledger.TxPvtData)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, ledger.PvtNsCollFilter) error); ok {
		r1 = rf(blockNum, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LedgerHeight provides a mock function with given fields:
func (_m *Committer) LedgerHeight() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCommitter creates a new instance of Committer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommitter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Committer {
	mock := &Committer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
