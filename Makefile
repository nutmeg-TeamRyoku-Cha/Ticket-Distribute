run_app:
	docker compose up --build -d
kill_app:
	docker compose down -v

conn_db:
	docker exec -it mysql mysql -uroot -proot app_db

visitor:
	docker exec -it mysql mysql -uroot -proot --default-character-set=utf8mb4 app_db -e "SELECT * FROM visitors;"
session:
	docker exec -it mysql mysql -uroot -proot --default-character-set=utf8mb4 app_db -e "SELECT * FROM login_sessions;"
building:
	docker exec -it mysql mysql -uroot -proot --default-character-set=utf8mb4 app_db -e "SELECT * FROM buildings;"
project:
	docker exec -it mysql mysql -uroot -proot --default-character-set=utf8mb4 app_db -e "SELECT * FROM projects;"
ticket:
	docker exec -it mysql mysql -uroot -proot --default-character-set=utf8mb4 app_db -e "SELECT * FROM tickets;"