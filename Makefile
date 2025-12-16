.PHONY: proto clean

proto:
	@echo "Generating Go code from proto files..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/game_agent.proto
	@echo "Proto files generated successfully"

clean:
	@echo "Cleaning generated proto files..."
	rm -f proto/*.pb.go
	@echo "Clean complete"
