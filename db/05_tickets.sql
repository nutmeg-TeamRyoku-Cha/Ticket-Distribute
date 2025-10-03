USE app_db;
CREATE TABLE IF NOT EXISTS tickets (
  ticket_id        INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  visitor_id       INT UNSIGNED NOT NULL,
  project_id       INT UNSIGNED NOT NULL,
  status           ENUM('issued', 'used', 'cancelled') NOT NULL DEFAULT 'issued',
  entry_start_time DATETIME NULL,
  entry_end_time   DATETIME NULL,
  CONSTRAINT fk_tickets_visitor
    FOREIGN KEY (visitor_id) REFERENCES visitors(visitor_id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_tickets_project
    FOREIGN KEY (project_id) REFERENCES projects(project_id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CHECK (
    entry_start_time IS NULL
    OR entry_end_time IS NULL
    OR entry_start_time <= entry_end_time
  )
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;