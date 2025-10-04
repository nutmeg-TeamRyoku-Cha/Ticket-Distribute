package repository

import (
	"context"
	"database/sql"
	"time"

	"ticket-app/internal/domain"
)

type VisitorRepository struct{ DB *sql.DB }

func NewVisitorRepository(db *sql.DB) *VisitorRepository { return &VisitorRepository{DB: db} }

func (r *VisitorRepository) Create(ctx context.Context, v domain.Visitor) (uint64, error) {
	res, err := r.DB.ExecContext(ctx, `
        INSERT INTO visitors (nickname, birth_date, party_size)
        VALUES (?, ?, ?)
    `, v.Nickname, v.BirthDate.UTC().Format("2006-01-02"), v.PartySize)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (r *VisitorRepository) List(ctx context.Context) ([]domain.Visitor, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT visitor_id, nickname, birth_date, party_size
        FROM visitors
        ORDER BY visitor_id DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Visitor
	for rows.Next() {
		var id uint64
		var nickname string
		var bd time.Time
		var ps int
		if err := rows.Scan(&id, &nickname, &bd, &ps); err != nil {
			return nil, err
		}
		out = append(out, domain.Visitor{VisitorID: id, Nickname: nickname, BirthDate: bd, PartySize: ps})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
