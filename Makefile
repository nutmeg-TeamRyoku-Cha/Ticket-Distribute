visitor:
	docker exec -it mysql mysql -uroot -proot app_db -e "SELECT * FROM visitors;"
session:
	docker exec -it mysql mysql -uroot -proot app_db -e "SELECT * FROM login_sessions;"
building:
	docker exec -it mysql mysql -uroot -proot app_db -e "SELECT * FROM buildings;"
project:
	docker exec -it mysql mysql -uroot -proot app_db -e "SELECT * FROM projects;"
ticket:
	docker exec -it mysql mysql -uroot -proot app_db -e "SELECT * FROM tickets;"