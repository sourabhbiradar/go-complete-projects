P_DIR := proto
S_DIR := server
C_DIR := client

PROTO_FILES := $(P_DIR)/blog.proto

protoc:
	protoc --go-grpc_out=./$(P_DIR) --go_out=./$(P_DIR) $(PROTO_FILES)

serverrun:
	go run $(S_DIR)/*.go

clientrun:
	go run $(C_DIR)/*.go

.PHONY: protoc serverrun clientrun