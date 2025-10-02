USE app_db;
CREATE TABLE buildings (
  building_id    INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  building_name  VARCHAR(255) NOT NULL,
  latitude       DOUBLE,
  longitude      DOUBLE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;