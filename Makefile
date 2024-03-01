.PHONY: setup generate generate-proto fix-proto

SERVICE_NAME = protoc
NETWORK_NAME = chat-grpc

setup:
	docker network create ${NETWORK_NAME} || true
	docker-compose up -d

generate:
	@$(MAKE) generate-proto
	@echo "[OK] search proto generated!"
	@$(MAKE) fix-proto
	@echo "[OK] search.pb.go fixed!"

generate-proto:
	@docker compose exec ${SERVICE_NAME} sh -c 'protoc -I . --proto_path=proto --go_out=. --go-grpc_out=. $$(find proto -name "*.proto")'


fix-proto:
	@for file in ./chatgrpc/*.go; do \
		perl -i -pe 'next if /prev/; s/,omitempty//g' $$file; \
	done

