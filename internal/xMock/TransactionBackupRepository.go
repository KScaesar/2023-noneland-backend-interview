// Code generated by mockery v2.32.4. DO NOT EDIT.

package xMock

import (
	context "context"
	entity "noneland/backend/interview/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockTransactionBackupRepository is an autogenerated mock type for the TransactionBackupRepository type
type MockTransactionBackupRepository struct {
	mock.Mock
}

type MockTransactionBackupRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransactionBackupRepository) EXPECT() *MockTransactionBackupRepository_Expecter {
	return &MockTransactionBackupRepository_Expecter{mock: &_m.Mock}
}

// CreatBulkTransactionBackup provides a mock function with given fields: ctx, txAll
func (_m *MockTransactionBackupRepository) CreatBulkTransactionBackup(ctx context.Context, txAll []entity.TransactionBackup) error {
	ret := _m.Called(ctx, txAll)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []entity.TransactionBackup) error); ok {
		r0 = rf(ctx, txAll)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransactionBackupRepository_CreatBulkTransactionBackup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatBulkTransactionBackup'
type MockTransactionBackupRepository_CreatBulkTransactionBackup_Call struct {
	*mock.Call
}

// CreatBulkTransactionBackup is a helper method to define mock.On call
//   - ctx context.Context
//   - txAll []entity.TransactionBackup
func (_e *MockTransactionBackupRepository_Expecter) CreatBulkTransactionBackup(ctx interface{}, txAll interface{}) *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call {
	return &MockTransactionBackupRepository_CreatBulkTransactionBackup_Call{Call: _e.mock.On("CreatBulkTransactionBackup", ctx, txAll)}
}

func (_c *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call) Run(run func(ctx context.Context, txAll []entity.TransactionBackup)) *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]entity.TransactionBackup))
	})
	return _c
}

func (_c *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call) Return(_a0 error) *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call) RunAndReturn(run func(context.Context, []entity.TransactionBackup) error) *MockTransactionBackupRepository_CreatBulkTransactionBackup_Call {
	_c.Call.Return(run)
	return _c
}

// GetSpotTransactionBackupAllByUserId provides a mock function with given fields: ctx, dto
func (_m *MockTransactionBackupRepository) GetSpotTransactionBackupAllByUserId(ctx context.Context, dto *entity.QryTransactionBackupParam) ([]entity.ExchangeTransactionResponse, error) {
	ret := _m.Called(ctx, dto)

	var r0 []entity.ExchangeTransactionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.QryTransactionBackupParam) ([]entity.ExchangeTransactionResponse, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.QryTransactionBackupParam) []entity.ExchangeTransactionResponse); ok {
		r0 = rf(ctx, dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.ExchangeTransactionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.QryTransactionBackupParam) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSpotTransactionBackupAllByUserId'
type MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call struct {
	*mock.Call
}

// GetSpotTransactionBackupAllByUserId is a helper method to define mock.On call
//   - ctx context.Context
//   - dto *entity.QryTransactionBackupParam
func (_e *MockTransactionBackupRepository_Expecter) GetSpotTransactionBackupAllByUserId(ctx interface{}, dto interface{}) *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call {
	return &MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call{Call: _e.mock.On("GetSpotTransactionBackupAllByUserId", ctx, dto)}
}

func (_c *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call) Run(run func(ctx context.Context, dto *entity.QryTransactionBackupParam)) *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.QryTransactionBackupParam))
	})
	return _c
}

func (_c *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call) Return(_a0 []entity.ExchangeTransactionResponse, _a1 error) *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call) RunAndReturn(run func(context.Context, *entity.QryTransactionBackupParam) ([]entity.ExchangeTransactionResponse, error)) *MockTransactionBackupRepository_GetSpotTransactionBackupAllByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransactionBackupRepository creates a new instance of MockTransactionBackupRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactionBackupRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactionBackupRepository {
	mock := &MockTransactionBackupRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
