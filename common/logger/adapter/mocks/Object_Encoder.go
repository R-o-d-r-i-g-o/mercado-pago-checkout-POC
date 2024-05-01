// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	zapcore "go.uber.org/zap/zapcore"
)

// ObjectEncoder is an autogenerated mock type for the ObjectEncoder type
type ObjectEncoder struct {
	mock.Mock
}

// AddArray provides a mock function with given fields: key, marshaler
func (_m *ObjectEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	ret := _m.Called(key, marshaler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, zapcore.ArrayMarshaler) error); ok {
		r0 = rf(key, marshaler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddBinary provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddBinary(key string, value []byte) {
	_m.Called(key, value)
}

// AddBool provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddBool(key string, value bool) {
	_m.Called(key, value)
}

// AddByteString provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddByteString(key string, value []byte) {
	_m.Called(key, value)
}

// AddComplex128 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddComplex128(key string, value complex128) {
	_m.Called(key, value)
}

// AddComplex64 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddComplex64(key string, value complex64) {
	_m.Called(key, value)
}

// AddDuration provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddDuration(key string, value time.Duration) {
	_m.Called(key, value)
}

// AddFloat32 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddFloat32(key string, value float32) {
	_m.Called(key, value)
}

// AddFloat64 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddFloat64(key string, value float64) {
	_m.Called(key, value)
}

// AddInt provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddInt(key string, value int) {
	_m.Called(key, value)
}

// AddInt16 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddInt16(key string, value int16) {
	_m.Called(key, value)
}

// AddInt32 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddInt32(key string, value int32) {
	_m.Called(key, value)
}

// AddInt64 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddInt64(key string, value int64) {
	_m.Called(key, value)
}

// AddInt8 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddInt8(key string, value int8) {
	_m.Called(key, value)
}

// AddObject provides a mock function with given fields: key, marshaler
func (_m *ObjectEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	ret := _m.Called(key, marshaler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, zapcore.ObjectMarshaler) error); ok {
		r0 = rf(key, marshaler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddReflected provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddReflected(key string, value interface{}) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddString provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddString(key string, value string) {
	_m.Called(key, value)
}

// AddTime provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddTime(key string, value time.Time) {
	_m.Called(key, value)
}

// AddUint provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddUint(key string, value uint) {
	_m.Called(key, value)
}

// AddUint16 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddUint16(key string, value uint16) {
	_m.Called(key, value)
}

// AddUint32 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddUint32(key string, value uint32) {
	_m.Called(key, value)
}

// AddUint64 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddUint64(key string, value uint64) {
	_m.Called(key, value)
}

// AddUint8 provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddUint8(key string, value uint8) {
	_m.Called(key, value)
}

// AddUintptr provides a mock function with given fields: key, value
func (_m *ObjectEncoder) AddUintptr(key string, value uintptr) {
	_m.Called(key, value)
}

// OpenNamespace provides a mock function with given fields: key
func (_m *ObjectEncoder) OpenNamespace(key string) {
	_m.Called(key)
}

type mockConstructorTestingTNewObjectMarshaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewObjectEncoder creates a new instance of ObjectMarshaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewObjectEncoder(t mockConstructorTestingTNewObjectMarshaller) *ObjectEncoder {
	mock := &ObjectEncoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
