.PHONY: up gen-pb
up:
	docker-compose up
gen-pb:
	cp guruguru.proto ./boss/guruguru.proto && \
		docker-compose run --rm --no-deps boss /protoc/bin/protoc ./guruguru.proto --go_out=plugins=grpc:guruguru
	cp guruguru.proto ./go/guruguru.proto && \
		docker-compose run --rm --no-deps go /protoc/bin/protoc ./guruguru.proto --go_out=plugins=grpc:guruguru