

run-db:
	sh deploy/run/db.sh

stop-db:
	docker rm -f db-simple-blog

run-app:
	POSTGRES_HOST=changeme POSTGRES_PORT=changeme POSTGRES_USER=changeme POSTGRES_PASS=changeme POSTGRES_DBNAME=changeme POSTGRES_MAX_IDLE_CONNS=changeme POSTGRES_MAX_OPEN_CONNS=changeme POSTGRES_MAX_LIFE_TIME_CONNS=changeme POSTGRES_MAX_IDLE_TIME_CONNS=changeme APP_PORT=changeme go run main.go