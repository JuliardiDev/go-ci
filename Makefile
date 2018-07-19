IMAGE_NAME := "foo"

test:
	@(docker-compose rm -f)
	@(docker-compose down)
	@(docker-compose run wait)	
	@(docker-compose build)
	@(docker-compose run web go test)