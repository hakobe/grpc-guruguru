.PHONY: gen-pb
gen-pb:
	protoc ./guruguru.proto --go_out=plugins=grpc:./boss/guruguru
	protoc ./guruguru.proto --go_out=plugins=grpc:./go_worker/guruguru