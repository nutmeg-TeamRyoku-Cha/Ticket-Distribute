#アプリ
up: #起動
	docker compose up -d
buildup: #ビルドと起動
	docker compose up --build -d
down: #アプリを落とす
	docker compose down
restart: #アプリ再起
	docker compose restart
kill: #ボリューム消す（DBの内容消したいときは）
	docker compose down -v
log_web: #webのlog
	docker compose logs -f web
log_api: #apiのlog
	docker compose logs -f api
log_db: #dbのlog
	docker compose logs -f db

run-prod:
	docker compose -f docker-compose.prod.yml build --no-cache
	docker compose -f docker-compose.prod.yml up -d

#DBコンテナにアクセス
conn_db:
	docker exec -it mysql mysql -uroot -proot app_db

#各テーブルをのぞく
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