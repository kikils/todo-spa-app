build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

deploy_backend:
	docker login --username=_ --password=`heroku auth:token` registry.heroku.com
	docker build --no-cache -t registry.heroku.com/sheltered-hamlet-88717/web ./backend
	docker push registry.heroku.com/sheltered-hamlet-88717/web

deploy_frontend:
	docker-compose exec frontend /bin/sh -c "yarn dev-build && firebase deploy"