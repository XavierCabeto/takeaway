// Code generated by MockGen. DO NOT EDIT.
// Source: product.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	application "github.com/XavierCabeto/takeaway/application"
	gomock "github.com/golang/mock/gomock"
)

// MockProductInterface is a mock of ProductInterface interface.
type MockProductInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductInterfaceMockRecorder
}

// MockProductInterfaceMockRecorder is the mock recorder for MockProductInterface.
type MockProductInterfaceMockRecorder struct {
	mock *MockProductInterface
}

// NewMockProductInterface creates a new mock instance.
func NewMockProductInterface(ctrl *gomock.Controller) *MockProductInterface {
	mock := &MockProductInterface{ctrl: ctrl}
	mock.recorder = &MockProductInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductInterface) EXPECT() *MockProductInterfaceMockRecorder {
	return m.recorder
}

// GetID mocks base method.
func (m *MockProductInterface) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockProductInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockProductInterface)(nil).GetID))
}

// GetName mocks base method.
func (m *MockProductInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockProductInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockProductInterface)(nil).GetName))
}

// GetPrice mocks base method.
func (m *MockProductInterface) GetPrice() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockProductInterfaceMockRecorder) GetPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockProductInterface)(nil).GetPrice))
}

// IsValid mocks base method.
func (m *MockProductInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockProductInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockProductInterface)(nil).IsValid))
}

// MockProductServiceInterface is a mock of ProductServiceInterface interface.
type MockProductServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceInterfaceMockRecorder
}

// MockProductServiceInterfaceMockRecorder is the mock recorder for MockProductServiceInterface.
type MockProductServiceInterfaceMockRecorder struct {
	mock *MockProductServiceInterface
}

// NewMockProductServiceInterface creates a new mock instance.
func NewMockProductServiceInterface(ctrl *gomock.Controller) *MockProductServiceInterface {
	mock := &MockProductServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProductServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceInterface) EXPECT() *MockProductServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductServiceInterface) Create(name string, price float64) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceInterfaceMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductServiceInterface)(nil).Create), name, price)
}

// Get mocks base method.
func (m *MockProductServiceInterface) Get(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductServiceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductServiceInterface)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockProductServiceInterface) GetAll() ([]application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductServiceInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductServiceInterface)(nil).GetAll))
}

// MockProductReader is a mock of ProductReader interface.
type MockProductReader struct {
	ctrl     *gomock.Controller
	recorder *MockProductReaderMockRecorder
}

// MockProductReaderMockRecorder is the mock recorder for MockProductReader.
type MockProductReaderMockRecorder struct {
	mock *MockProductReader
}

// NewMockProductReader creates a new mock instance.
func NewMockProductReader(ctrl *gomock.Controller) *MockProductReader {
	mock := &MockProductReader{ctrl: ctrl}
	mock.recorder = &MockProductReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductReader) EXPECT() *MockProductReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProductReader) Get(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductReader)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockProductReader) GetAll() ([]application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductReaderMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductReader)(nil).GetAll))
}

// MockProductWriter is a mock of ProductWriter interface.
type MockProductWriter struct {
	ctrl     *gomock.Controller
	recorder *MockProductWriterMockRecorder
}

// MockProductWriterMockRecorder is the mock recorder for MockProductWriter.
type MockProductWriterMockRecorder struct {
	mock *MockProductWriter
}

// NewMockProductWriter creates a new mock instance.
func NewMockProductWriter(ctrl *gomock.Controller) *MockProductWriter {
	mock := &MockProductWriter{ctrl: ctrl}
	mock.recorder = &MockProductWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductWriter) EXPECT() *MockProductWriterMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockProductWriter) Save(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductWriterMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductWriter)(nil).Save), product)
}

// MockProductPersistenceInterface is a mock of ProductPersistenceInterface interface.
type MockProductPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductPersistenceInterfaceMockRecorder
}

// MockProductPersistenceInterfaceMockRecorder is the mock recorder for MockProductPersistenceInterface.
type MockProductPersistenceInterfaceMockRecorder struct {
	mock *MockProductPersistenceInterface
}

// NewMockProductPersistenceInterface creates a new mock instance.
func NewMockProductPersistenceInterface(ctrl *gomock.Controller) *MockProductPersistenceInterface {
	mock := &MockProductPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockProductPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductPersistenceInterface) EXPECT() *MockProductPersistenceInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProductPersistenceInterface) Get(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductPersistenceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductPersistenceInterface)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockProductPersistenceInterface) GetAll() ([]application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductPersistenceInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductPersistenceInterface)(nil).GetAll))
}

// Save mocks base method.
func (m *MockProductPersistenceInterface) Save(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductPersistenceInterfaceMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductPersistenceInterface)(nil).Save), product)
}
