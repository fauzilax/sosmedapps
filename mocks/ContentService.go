// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	contents "sosmedapps/features/contents"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// ContentService is an autogenerated mock type for the ContentService type
type ContentService struct {
	mock.Mock
}

// AddContent provides a mock function with given fields: token, fileHeader, newContent
func (_m *ContentService) AddContent(token interface{}, fileHeader multipart.FileHeader, newContent contents.CoreContent) (contents.CoreContent, error) {
	ret := _m.Called(token, fileHeader, newContent)

	var r0 contents.CoreContent
	if rf, ok := ret.Get(0).(func(interface{}, multipart.FileHeader, contents.CoreContent) contents.CoreContent); ok {
		r0 = rf(token, fileHeader, newContent)
	} else {
		r0 = ret.Get(0).(contents.CoreContent)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, multipart.FileHeader, contents.CoreContent) error); ok {
		r1 = rf(token, fileHeader, newContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllContent provides a mock function with given fields:
func (_m *ContentService) AllContent() ([]contents.CoreContent, error) {
	ret := _m.Called()

	var r0 []contents.CoreContent
	if rf, ok := ret.Get(0).(func() []contents.CoreContent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contents.CoreContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContent provides a mock function with given fields: token, contentID
func (_m *ContentService) DeleteContent(token interface{}, contentID uint) error {
	ret := _m.Called(token, contentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, contentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DetailContent provides a mock function with given fields: contentID
func (_m *ContentService) DetailContent(contentID uint) (interface{}, error) {
	ret := _m.Called(contentID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(uint) interface{}); ok {
		r0 = rf(contentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(contentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContent provides a mock function with given fields: token, contentID, content
func (_m *ContentService) UpdateContent(token interface{}, contentID uint, content string) (string, error) {
	ret := _m.Called(token, contentID, content)

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}, uint, string) string); ok {
		r0 = rf(token, contentID, content)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, string) error); ok {
		r1 = rf(token, contentID, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewContentService creates a new instance of ContentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContentService(t mockConstructorTestingTNewContentService) *ContentService {
	mock := &ContentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
