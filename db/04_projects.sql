USE app_db;
CREATE TABLE IF NOT EXISTS projects (
  project_id        INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  project_name      VARCHAR(255) NOT NULL,
  building_id       INT UNSIGNED NOT NULL,
  requires_ticket   BOOLEAN NOT NULL DEFAULT FALSE,
  remaining_tickets INT UNSIGNED NOT NULL DEFAULT 300,
  start_time        DATETIME NOT NULL,
  end_time          DATETIME NOT NULL,
  CONSTRAINT fk_projects_building
    FOREIGN KEY (building_id) REFERENCES buildings(building_id)
    ON DELETE RESTRICT ON UPDATE CASCADE,
  CHECK (end_time >= start_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;