// Code generated by mockery v2.32.4. DO NOT EDIT.

package xMock

import (
	context "context"
	entity "noneland/backend/interview/internal/entity"

	mock "github.com/stretchr/testify/mock"

	pkg "noneland/backend/interview/pkg"
)

// MockExchangeQryService is an autogenerated mock type for the ExchangeQryService type
type MockExchangeQryService struct {
	mock.Mock
}

type MockExchangeQryService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExchangeQryService) EXPECT() *MockExchangeQryService_Expecter {
	return &MockExchangeQryService_Expecter{mock: &_m.Mock}
}

// GetBalanceByUserId provides a mock function with given fields: ctx, usrId
func (_m *MockExchangeQryService) GetBalanceByUserId(ctx context.Context, usrId string) (entity.BalanceResponse, error) {
	ret := _m.Called(ctx, usrId)

	var r0 entity.BalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.BalanceResponse, error)); ok {
		return rf(ctx, usrId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.BalanceResponse); ok {
		r0 = rf(ctx, usrId)
	} else {
		r0 = ret.Get(0).(entity.BalanceResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, usrId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExchangeQryService_GetBalanceByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalanceByUserId'
type MockExchangeQryService_GetBalanceByUserId_Call struct {
	*mock.Call
}

// GetBalanceByUserId is a helper method to define mock.On call
//   - ctx context.Context
//   - usrId string
func (_e *MockExchangeQryService_Expecter) GetBalanceByUserId(ctx interface{}, usrId interface{}) *MockExchangeQryService_GetBalanceByUserId_Call {
	return &MockExchangeQryService_GetBalanceByUserId_Call{Call: _e.mock.On("GetBalanceByUserId", ctx, usrId)}
}

func (_c *MockExchangeQryService_GetBalanceByUserId_Call) Run(run func(ctx context.Context, usrId string)) *MockExchangeQryService_GetBalanceByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockExchangeQryService_GetBalanceByUserId_Call) Return(_a0 entity.BalanceResponse, _a1 error) *MockExchangeQryService_GetBalanceByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExchangeQryService_GetBalanceByUserId_Call) RunAndReturn(run func(context.Context, string) (entity.BalanceResponse, error)) *MockExchangeQryService_GetBalanceByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// GetSpotTransactionListByUserId provides a mock function with given fields: ctx, userId, dtoPage, tRange
func (_m *MockExchangeQryService) GetSpotTransactionListByUserId(ctx context.Context, userId string, dtoPage pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThan) (pkg.ListResponse[entity.ExchangeTransactionResponse], error) {
	ret := _m.Called(ctx, userId, dtoPage, tRange)

	var r0 pkg.ListResponse[entity.ExchangeTransactionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, pkg.PageParam, pkg.TimestampRangeEndTimeLessThan) (pkg.ListResponse[entity.ExchangeTransactionResponse], error)); ok {
		return rf(ctx, userId, dtoPage, tRange)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, pkg.PageParam, pkg.TimestampRangeEndTimeLessThan) pkg.ListResponse[entity.ExchangeTransactionResponse]); ok {
		r0 = rf(ctx, userId, dtoPage, tRange)
	} else {
		r0 = ret.Get(0).(pkg.ListResponse[entity.ExchangeTransactionResponse])
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, pkg.PageParam, pkg.TimestampRangeEndTimeLessThan) error); ok {
		r1 = rf(ctx, userId, dtoPage, tRange)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExchangeQryService_GetSpotTransactionListByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSpotTransactionListByUserId'
type MockExchangeQryService_GetSpotTransactionListByUserId_Call struct {
	*mock.Call
}

// GetSpotTransactionListByUserId is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - dtoPage pkg.PageParam
//   - tRange pkg.TimestampRangeEndTimeLessThan
func (_e *MockExchangeQryService_Expecter) GetSpotTransactionListByUserId(ctx interface{}, userId interface{}, dtoPage interface{}, tRange interface{}) *MockExchangeQryService_GetSpotTransactionListByUserId_Call {
	return &MockExchangeQryService_GetSpotTransactionListByUserId_Call{Call: _e.mock.On("GetSpotTransactionListByUserId", ctx, userId, dtoPage, tRange)}
}

func (_c *MockExchangeQryService_GetSpotTransactionListByUserId_Call) Run(run func(ctx context.Context, userId string, dtoPage pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThan)) *MockExchangeQryService_GetSpotTransactionListByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(pkg.PageParam), args[3].(pkg.TimestampRangeEndTimeLessThan))
	})
	return _c
}

func (_c *MockExchangeQryService_GetSpotTransactionListByUserId_Call) Return(_a0 pkg.ListResponse[entity.ExchangeTransactionResponse], _a1 error) *MockExchangeQryService_GetSpotTransactionListByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExchangeQryService_GetSpotTransactionListByUserId_Call) RunAndReturn(run func(context.Context, string, pkg.PageParam, pkg.TimestampRangeEndTimeLessThan) (pkg.ListResponse[entity.ExchangeTransactionResponse], error)) *MockExchangeQryService_GetSpotTransactionListByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExchangeQryService creates a new instance of MockExchangeQryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExchangeQryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExchangeQryService {
	mock := &MockExchangeQryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
