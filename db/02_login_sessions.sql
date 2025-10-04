USE app_db;
CREATE TABLE login_sessions (
  session_hash BINARY(32) PRIMARY KEY,
  visitor_id  INT UNSIGNED NOT NULL,
  expires_at  DATETIME NOT NULL,
  CONSTRAINT fk_sessions_visitor
    FOREIGN KEY (visitor_id) REFERENCES visitors(visitor_id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;