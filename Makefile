IMAGE_NAME := "foo"

test:	
	@docker build -t goci_web .
	@docker-compose run wait
	@docker-compose run web go test 
	@docker-compose rm -f 
	@docker-compose down		