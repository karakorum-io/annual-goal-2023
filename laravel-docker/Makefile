setup:
	@make build
	@make up
	@make create
	@make composer-update
	@make data
	@make run
create:
	docker exec laravel-docker bash -c "composer create-project laravel/laravel"
build:
	docker-compose build --no-cache --force-rm
down:
	docker-compose stop
up:
	docker-compose up -d
composer-update:
	docker exec laravel-docker bash -c cd laravel -c "composer update"
data:
	docker-compose exec laravel-docker bash -c cd laravel -c "php artisan migrate"
	docker-compose exec laravel-docker bash -c cd laravel -c "php artisan db:seed"
run:
	docker-compose exec laravel-docker bash -c cd laravel -c "php artisan serve"
into-bash:
	docker-compose exec laravel-docker bash