build:
	@go build -o bin/iptables-go

run: build
	@sudo ./bin/iptables-go


