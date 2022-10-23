// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/repo.go

// Package data is a generated GoMock package.
package data

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// RepoMock is a mock of Repo interface.
type RepoMock struct {
	ctrl     *gomock.Controller
	recorder *RepoMockMockRecorder
}

// RepoMockMockRecorder is the mock recorder for RepoMock.
type RepoMockMockRecorder struct {
	mock *RepoMock
}

// NewRepoMock creates a new mock instance.
func NewRepoMock(ctrl *gomock.Controller) *RepoMock {
	mock := &RepoMock{ctrl: ctrl}
	mock.recorder = &RepoMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *RepoMock) EXPECT() *RepoMockMockRecorder {
	return m.recorder
}

// AddCase mocks base method.
func (m *RepoMock) AddCase(ctx context.Context, date time.Time, amount int, sluggedPrefecture string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCase", ctx, date, amount, sluggedPrefecture)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCase indicates an expected call of AddCase.
func (mr *RepoMockMockRecorder) AddCase(ctx, date, amount, sluggedPrefecture interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCase", reflect.TypeOf((*RepoMock)(nil).AddCase), ctx, date, amount, sluggedPrefecture)
}

// AddCounty mocks base method.
func (m *RepoMock) AddCounty(ctx context.Context, county County) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCounty", ctx, county)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCounty indicates an expected call of AddCounty.
func (mr *RepoMockMockRecorder) AddCounty(ctx, county interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCounty", reflect.TypeOf((*RepoMock)(nil).AddCounty), ctx, county)
}

// AddFullInfo mocks base method.
func (m *RepoMock) AddFullInfo(ctx context.Context, fi *FullInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFullInfo", ctx, fi)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFullInfo indicates an expected call of AddFullInfo.
func (mr *RepoMockMockRecorder) AddFullInfo(ctx, fi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFullInfo", reflect.TypeOf((*RepoMock)(nil).AddFullInfo), ctx, fi)
}

// AddMunicipality mocks base method.
func (m *RepoMock) AddMunicipality(ctx context.Context, name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMunicipality", ctx, name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMunicipality indicates an expected call of AddMunicipality.
func (mr *RepoMockMockRecorder) AddMunicipality(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMunicipality", reflect.TypeOf((*RepoMock)(nil).AddMunicipality), ctx, name)
}

// AddYearlyDeath mocks base method.
func (m *RepoMock) AddYearlyDeath(ctx context.Context, munId, deaths, year int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddYearlyDeath", ctx, munId, deaths, year)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddYearlyDeath indicates an expected call of AddYearlyDeath.
func (mr *RepoMockMockRecorder) AddYearlyDeath(ctx, munId, deaths, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddYearlyDeath", reflect.TypeOf((*RepoMock)(nil).AddYearlyDeath), ctx, munId, deaths, year)
}

// GetCases mocks base method.
func (m *RepoMock) GetCases(ctx context.Context, filter CasesFilter) ([]Case, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCases", ctx, filter)
	ret0, _ := ret[0].([]Case)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCases indicates an expected call of GetCases.
func (mr *RepoMockMockRecorder) GetCases(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCases", reflect.TypeOf((*RepoMock)(nil).GetCases), ctx, filter)
}

// GetCounties mocks base method.
func (m *RepoMock) GetCounties(ctx context.Context) ([]County, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounties", ctx)
	ret0, _ := ret[0].([]County)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounties indicates an expected call of GetCounties.
func (mr *RepoMockMockRecorder) GetCounties(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounties", reflect.TypeOf((*RepoMock)(nil).GetCounties), ctx)
}

// GetFromTimeline mocks base method.
func (m *RepoMock) GetFromTimeline(ctx context.Context, filter DatesFilter) ([]FullInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromTimeline", ctx, filter)
	ret0, _ := ret[0].([]FullInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromTimeline indicates an expected call of GetFromTimeline.
func (mr *RepoMockMockRecorder) GetFromTimeline(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromTimeline", reflect.TypeOf((*RepoMock)(nil).GetFromTimeline), ctx, filter)
}
