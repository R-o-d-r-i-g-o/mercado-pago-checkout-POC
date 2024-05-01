// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	zapcore "go.uber.org/zap/zapcore"
)

// ArrayEncoder is an autogenerated mock type for the ArrayEncoder type
type ArrayEncoder struct {
	mock.Mock
}

// AppendArray provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendArray(_a0 zapcore.ArrayMarshaler) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(zapcore.ArrayMarshaler) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppendBool provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendBool(_a0 bool) {
	_m.Called(_a0)
}

// AppendByteString provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendByteString(_a0 []byte) {
	_m.Called(_a0)
}

// AppendComplex128 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendComplex128(_a0 complex128) {
	_m.Called(_a0)
}

// AppendComplex64 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendComplex64(_a0 complex64) {
	_m.Called(_a0)
}

// AppendDuration provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendDuration(_a0 time.Duration) {
	_m.Called(_a0)
}

// AppendFloat32 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendFloat32(_a0 float32) {
	_m.Called(_a0)
}

// AppendFloat64 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendFloat64(_a0 float64) {
	_m.Called(_a0)
}

// AppendInt provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendInt(_a0 int) {
	_m.Called(_a0)
}

// AppendInt16 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendInt16(_a0 int16) {
	_m.Called(_a0)
}

// AppendInt32 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendInt32(_a0 int32) {
	_m.Called(_a0)
}

// AppendInt64 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendInt64(_a0 int64) {
	_m.Called(_a0)
}

// AppendInt8 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendInt8(_a0 int8) {
	_m.Called(_a0)
}

// AppendObject provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendObject(_a0 zapcore.ObjectMarshaler) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(zapcore.ObjectMarshaler) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppendReflected provides a mock function with given fields: value
func (_m *ArrayEncoder) AppendReflected(value interface{}) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppendString provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendString(_a0 string) {
	_m.Called(_a0)
}

// AppendTime provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendTime(_a0 time.Time) {
	_m.Called(_a0)
}

// AppendUint provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendUint(_a0 uint) {
	_m.Called(_a0)
}

// AppendUint16 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendUint16(_a0 uint16) {
	_m.Called(_a0)
}

// AppendUint32 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendUint32(_a0 uint32) {
	_m.Called(_a0)
}

// AppendUint64 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendUint64(_a0 uint64) {
	_m.Called(_a0)
}

// AppendUint8 provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendUint8(_a0 uint8) {
	_m.Called(_a0)
}

// AppendUintptr provides a mock function with given fields: _a0
func (_m *ArrayEncoder) AppendUintptr(_a0 uintptr) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewArrayEncoder interface {
	mock.TestingT
	Cleanup(func())
}

// NewArrayEncoder creates a new instance of ArrayEncoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArrayEncoder(t mockConstructorTestingTNewArrayEncoder) *ArrayEncoder {
	mock := &ArrayEncoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
