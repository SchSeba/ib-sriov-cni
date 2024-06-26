// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import ns "github.com/containernetworking/plugins/pkg/ns"
import types "github.com/k8snetworkplumbingwg/ib-sriov-cni/pkg/types"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// ApplyVFConfig provides a mock function with given fields: conf
func (_m *Manager) ApplyVFConfig(conf *types.NetConf) error {
	ret := _m.Called(conf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.NetConf) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReleaseVF provides a mock function with given fields: conf, podifName, cid, netns
func (_m *Manager) ReleaseVF(conf *types.NetConf, podifName string, cid string, netns ns.NetNS) error {
	ret := _m.Called(conf, podifName, cid, netns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.NetConf, string, string, ns.NetNS) error); ok {
		r0 = rf(conf, podifName, cid, netns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetVFConfig provides a mock function with given fields: conf
func (_m *Manager) ResetVFConfig(conf *types.NetConf) error {
	ret := _m.Called(conf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.NetConf) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupVF provides a mock function with given fields: conf, podifName, cid, netns
func (_m *Manager) SetupVF(conf *types.NetConf, podifName string, cid string, netns ns.NetNS) error {
	ret := _m.Called(conf, podifName, cid, netns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.NetConf, string, string, ns.NetNS) error); ok {
		r0 = rf(conf, podifName, cid, netns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
