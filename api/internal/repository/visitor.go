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

func (r *VisitorRepository) GetByID(ctx context.Context, id uint64) (domain.Visitor, bool, error) {
	var vid uint64
	var nickname string
	var bd time.Time
	var ps int
	err := r.DB.QueryRowContext(ctx, `
		SELECT visitor_id, nickname, birth_date, party_size
		FROM visitors
		WHERE visitor_id = ?
		LIMIT 1
	`, id).Scan(&vid, &nickname, &bd, &ps)
	if err == sql.ErrNoRows {
		return domain.Visitor{}, false, nil
	}
	if err != nil {
		return domain.Visitor{}, false, err
	}
	return domain.Visitor{VisitorID: vid, Nickname: nickname, BirthDate: bd, PartySize: ps}, true, nil
}

func (r *VisitorRepository) GetByNicknameAndBirthDate(ctx context.Context, nickname string, birthDate time.Time) (domain.Visitor, bool, error) {
	var v domain.Visitor
	// birth_date は DATE 型想定（YYYY-MM-DD）
	row := r.DB.QueryRowContext(ctx, `
    SELECT visitor_id, nickname, birth_date, party_size
    FROM visitors
    WHERE nickname = ? AND birth_date = ?
    ORDER BY visitor_id ASC
    LIMIT 1
  `, nickname, birthDate.Format("2006-01-02"))
	switch err := row.Scan(&v.VisitorID, &v.Nickname, &v.BirthDate, &v.PartySize); err {
	case nil:
		return v, true, nil
	case sql.ErrNoRows:
		return domain.Visitor{}, false, nil
	default:
		return domain.Visitor{}, false, err
	}
}
