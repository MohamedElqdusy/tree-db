package db

import (
	"context"
	"database/sql"
	"tree/logger"
	"tree/models"

	_ "github.com/lib/pq"
)

type PostgreRepository struct {
	db *sql.DB
}

func NewPostgre(url string) (*PostgreRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgreRepository{
		db,
	}, nil
}

func (p *PostgreRepository) Close() {
	p.db.Close()
}

func (p *PostgreRepository) StoreNode(ctx context.Context, node models.Node) error {
	_, err := p.db.Exec("INSERT INTO nodes(id, content, tree_root, parent, path) VALUES($1,$2,$3,$4,$5)",
		node.ID, node.Content, node.Root, node.Parent, node.Path)
	return err
}

func (p *PostgreRepository) UpdateParent(ctx context.Context, nodeId int, parent int) error {
	_, err := p.db.Exec("SELECT change_Parent($1,$2)", nodeId, parent)
	return err
}

func (p *PostgreRepository) FindNodeChilds(ctx context.Context, id int) ([]models.Node, error) {
	rows, err := p.db.Query("SELECT id, content, tree_root, parent, height(id), path from nodes WHERE path LIKE '%' || $1 || '%' ;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractNodeChilds(rows)
}

func extractNodeChilds(rows *sql.Rows) ([]models.Node, error) {
	var childs []models.Node
	for rows.Next() {
		node := models.Node{}
		err := rows.Scan(&node.ID, &node.Content, &node.Root, &node.Parent, &node.Height, &node.Path)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		childs = append(childs, node)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return childs, nil
}
