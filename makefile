# Команда для сборки Docker-образа
build:
	docker-compose build

# Команда для запуска контейнеров
up:
	docker-compose up

# Команда для остановки контейнеров
down:
	docker-compose down

# Команда для просмотра логов
logs:
	docker-compose logs -f
