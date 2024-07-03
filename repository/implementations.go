package repository

import "context"

func (r *Repository) GetEstateByID(ctx context.Context, id string) (estate Estate, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", id).Scan()
	if err != nil {
		return
	}
	return
}
