.PHONY: up gen-pb
up:
	docker-compose up
gen-pb:
	protoc ./guruguru.proto --go_out=plugins=grpc:./boss/guruguru
	protoc ./guruguru.proto --go_out=plugins=grpc:./go/guruguru