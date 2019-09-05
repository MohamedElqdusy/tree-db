package db

import (
	"context"
	"tree/models"
)

type Repository interface {
	StoreNode(ctx context.Context, node models.Node) error
	FindNodeChilds(ctx context.Context, id int) ([]models.Node, error)
	UpdateParent(ctx context.Context, nodeId int, parent int) error
	Close()
}

var repositoryImpl Repository

func StoreNode(ctx context.Context, node models.Node) error {
	return repositoryImpl.StoreNode(ctx, node)
}

func FindNodeChilds(ctx context.Context, id int) ([]models.Node, error) {
	return repositoryImpl.FindNodeChilds(ctx, id)
}

func UpdateParent(ctx context.Context, nodeId int, parent int) error {
	return repositoryImpl.UpdateParent(ctx, nodeId, parent)
}

func Close() {
	repositoryImpl.Close()
}

func SetRepository(repository Repository) {
	repositoryImpl = repository
}
