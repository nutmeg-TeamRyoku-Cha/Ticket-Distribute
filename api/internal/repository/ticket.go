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
			t := est.Time
			estp = &t
		}
		if eet.Valid {
			t := eet.Time
			eetp = &t
		}
		out = append(out, domain.Ticket{
			TicketID: id, VisitorID: vid, ProjectID: pid, Status: status,
			EntryStartTime: estp, EntryEndTime: eetp,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *TicketRepository) ListByVisitorWithProject(ctx context.Context, visitorID uint64) ([]domain.TicketWithProject, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
		t.ticket_id, t.visitor_id, t.status, t.entry_start_time, t.entry_end_time,
		p.project_id, p.project_name, p.requires_ticket, p.start_time, p.end_time,
		b.building_id, b.building_name, b.latitude, b.longitude
		FROM tickets t
		JOIN projects  p ON p.project_id  = t.project_id
		JOIN buildings b ON b.building_id = p.building_id
		WHERE t.visitor_id = ?
		ORDER BY t.ticket_id DESC
	`, visitorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.TicketWithProject
	for rows.Next() {
		var (
			ticketID, vID        uint64
			status               string
			entryStart, entryEnd sql.NullTime

			projectID      uint64
			projectName    sql.NullString
			requiresTicket bool
			projectStart   time.Time
			projectEnd     sql.NullTime

			buildingID   sql.NullInt64
			buildingName sql.NullString
			lat, lng     sql.NullFloat64
		)
		if err := rows.Scan(
			&ticketID, &vID, &status, &entryStart, &entryEnd,
			&projectID, &projectName, &requiresTicket, &projectStart, &projectEnd,
			&buildingID, &buildingName, &lat, &lng,
		); err != nil {
			return nil, err
		}

		var esPtr, eePtr *time.Time
		if entryStart.Valid {
			t := entryStart.Time
			esPtr = &t
		}
		if entryEnd.Valid {
			t := entryEnd.Time
			eePtr = &t
		}

		var endPtr *time.Time
		if projectEnd.Valid {
			t := projectEnd.Time
			endPtr = &t
		}

		var (
			bID            uint64
			bName          string
			latPtr, lngPtr *float64
		)
		if buildingID.Valid {
			bID = uint64(buildingID.Int64)
		}
		if buildingName.Valid {
			bName = buildingName.String
		}
		if lat.Valid {
			v := lat.Float64
			latPtr = &v
		}
		if lng.Valid {
			v := lng.Float64
			lngPtr = &v
		}

		projName := ""
		if projectName.Valid {
			projName = projectName.String
		}

		out = append(out, domain.TicketWithProject{
			TicketID:       ticketID,
			VisitorID:      vID,
			Status:         status,
			EntryStartTime: esPtr,
			EntryEndTime:   eePtr,
			Project: domain.ProjectBrief{
				ProjectID:      projectID,
				ProjectName:    projName,
				RequiresTicket: requiresTicket,
				StartTime:      projectStart,
				EndTime:        endPtr,
				Building: domain.BuildingBrief{
					BuildingID:   bID,
					BuildingName: bName,
					Latitude:     latPtr,
					Longitude:    lngPtr,
				},
			},
		})
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
