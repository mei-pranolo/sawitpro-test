package repository

import (
	"context"

	m "github.com/SawitProRecruitment/UserService/types"
)

func (r *Repository) GetEstateByID(ctx context.Context, id string) (estate m.Estate, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT * FROM estate WHERE id = $1", id).Scan(&estate.ID, &estate.Length, &estate.Width)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateEstate(ctx context.Context, length int, width int) (id string, err error) {
	err = r.Db.QueryRow(`INSERT INTO estate (length, width) VALUES($1, $2) RETURNING id`, length, width).Scan(&id)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateTree(ctx context.Context, estateID string, tree m.Tree) (id string, err error) {
	sqlStatement := `INSERT INTO tree (estate_id, x, y, height) VALUES($1, $2, $3, $4)`
	_, err = r.Db.Exec(sqlStatement, estateID, tree.X, tree.Y, tree.Height)
	if err != nil {
		return
	}
	return estateID, err
}

func (r *Repository) GetTree(ctx context.Context, estateID string) (trees []m.Tree, err error) {
	rows, err := r.Db.Query("SELECT x,y,height FROM tree WHERE estate_id = $1", estateID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tree m.Tree
		err = rows.Scan(&tree.X, &tree.Y, &tree.Height)
		if err != nil {
			continue
		}
		trees = append(trees, tree)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return
	}

	return
}
