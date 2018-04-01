.PHONY: up gen-pb
build:
	docker-compose build
up:
	docker-compose up
down:
	docker-compose down
gen-pb:
	cp guruguru.proto ./python/guruguru.proto && \
		docker-compose run --rm --no-deps python python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./guruguru.proto
	cp guruguru.proto ./boss/guruguru.proto && \
		docker-compose run --rm --no-deps boss /protoc/bin/protoc ./guruguru.proto --go_out=plugins=grpc:guruguru
	cp guruguru.proto ./go/guruguru.proto && \
		docker-compose run --rm --no-deps go /protoc/bin/protoc ./guruguru.proto --go_out=plugins=grpc:guruguru
	cp guruguru.proto ./node/guruguru.proto && \
		docker-compose run --rm --no-deps node yarn run gen-pb
	cp guruguru.proto ./ruby/guruguru.proto && \
		docker-compose run --rm --no-deps ruby grpc_tools_ruby_protoc --ruby_out=. --grpc_out=. ./guruguru.proto
	cp guruguru.proto ./elixir/guruguru.proto && \
		docker-compose run --rm --no-deps elixir protoc --elixir_out=plugins=grpc:./lib guruguru.proto
