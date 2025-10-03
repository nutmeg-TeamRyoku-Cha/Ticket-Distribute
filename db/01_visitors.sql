USE app_db;
CREATE TABLE visitors (
  visitor_id     INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  nickname       VARCHAR(255) NOT NULL,
  birth_date     DATE NOT NULL,
  party_size     INT UNSIGNED NOT NULL CHECK (party_size >= 1)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;