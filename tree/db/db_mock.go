package db

import (
	"context"
	"tree/models"

	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) StoreNode(ctx context.Context, node models.Node) error {
	rets := m.Called()
	return rets.Error(0)
}

func (m *MockStore) FindNodeChilds(ctx context.Context, id int) ([]models.Node, error) {
	rets := m.Called()
	return rets.Get(0).([]models.Node), rets.Error(1)
}

func (m *MockStore) UpdateParent(ctx context.Context, nodeId int, parent int) error {
	rets := m.Called()
	return rets.Error(0)
}

func (m *MockStore) Close() {
}

func InitMockStore() *MockStore {
	store := new(MockStore)
	SetRepository(store)
	return store
}
