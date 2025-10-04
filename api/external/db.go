package external

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func mysqlDSNFromEnv() (driver, dsn string) {
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASSWORD", "root")
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	db := getenv("DB_NAME", "app_db")

	q := url.Values{}
	q.Set("charset", "utf8mb4")
	q.Set("parseTime", "true")
	q.Set("loc", "UTC")

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, db, q.Encode())
	return "mysql", dsn
}

func resolveDSN() (driver, dsn string) {
	if raw := os.Getenv("DATABASE_URL"); raw != "" {
		u, err := url.Parse(raw)
		if err == nil && u.Scheme == "mysql" {
			return "mysql", raw
		}
	}
	drv := getenv("DB_DRIVER", "mysql")
	switch drv {
	case "mysql":
		return mysqlDSNFromEnv()
	default:
		return mysqlDSNFromEnv()
	}
}

func OpenDB() (*sql.DB, error) {
	driver, dsn := resolveDSN()
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}
