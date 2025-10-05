package repository

import (
	"context"
	"database/sql"
	"time"

	"ticket-app/internal/domain"
)

type TicketRepository struct{ DB *sql.DB }

func NewTicketRepository(db *sql.DB) *TicketRepository { return &TicketRepository{DB: db} }

func (r *TicketRepository) Create(ctx context.Context, t domain.Ticket) (uint64, error) {
	res, err := r.DB.ExecContext(ctx, `
        INSERT INTO tickets (visitor_id, project_id, status, entry_start_time, entry_end_time)
        VALUES (?, ?, ?, ?, ?)
    `, t.VisitorID, t.ProjectID, t.Status, t.EntryStartTime, t.EntryEndTime)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (r *TicketRepository) ListAll(ctx context.Context) ([]domain.Ticket, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT ticket_id, visitor_id, project_id, status, entry_start_time, entry_end_time
        FROM tickets
        ORDER BY ticket_id DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Ticket
	for rows.Next() {
		var id, vid, pid uint64
		var status string
		var est, eet sql.NullTime
		if err := rows.Scan(&id, &vid, &pid, &status, &est, &eet); err != nil {
			return nil, err
		}
		var estp, eetp *time.Time
		if est.Valid {
			tmp := est.Time
			estp = &tmp
		}
		if eet.Valid {
			tmp := eet.Time
			eetp = &tmp
		}
		out = append(out, domain.Ticket{TicketID: id, VisitorID: vid, ProjectID: pid, Status: status, EntryStartTime: estp, EntryEndTime: eetp})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *TicketRepository) ListByVisitor(ctx context.Context, visitorID uint64) ([]domain.Ticket, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT ticket_id, visitor_id, project_id, status, entry_start_time, entry_end_time
        FROM tickets
        WHERE visitor_id = ?
        ORDER BY ticket_id DESC
    `, visitorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Ticket
	for rows.Next() {
		var id, vid, pid uint64
		var status string
		var est, eet sql.NullTime
		if err := rows.Scan(&id, &vid, &pid, &status, &est, &eet); err != nil {
			return nil, err
		}
		var estp, eetp *time.Time
		if est.Valid {
			tmp := est.Time
			estp = &tmp
		}
		if eet.Valid {
			tmp := eet.Time
			eetp = &tmp
		}
		out = append(out, domain.Ticket{TicketID: id, VisitorID: vid, ProjectID: pid, Status: status, EntryStartTime: estp, EntryEndTime: eetp})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *TicketRepository) GetByID(ctx context.Context, id uint64) (domain.Ticket, bool, error) {
	var tid, vid, pid uint64
	var status string
	var est, eet sql.NullTime
	err := r.DB.QueryRowContext(ctx, `
        SELECT ticket_id, visitor_id, project_id, status, entry_start_time, entry_end_time
        FROM tickets
        WHERE ticket_id = ?
        LIMIT 1
    `, id).Scan(&tid, &vid, &pid, &status, &est, &eet)
	if err == sql.ErrNoRows {
		return domain.Ticket{}, false, nil
	}
	if err != nil {
		return domain.Ticket{}, false, err
	}
	var estp, eetp *time.Time
	if est.Valid {
		tmp := est.Time
		estp = &tmp
	}
	if eet.Valid {
		tmp := eet.Time
		eetp = &tmp
	}
	return domain.Ticket{TicketID: tid, VisitorID: vid, ProjectID: pid, Status: status, EntryStartTime: estp, EntryEndTime: eetp}, true, nil
}

func (r *TicketRepository) UpdateStatus(ctx context.Context, id uint64, status string) error {
	_, err := r.DB.ExecContext(ctx, `
        UPDATE tickets SET status = ? WHERE ticket_id = ?
    `, status, id)
	return err
}
