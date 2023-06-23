.DEFAULT_GOAL := build

clean:
	go clean
	rm -rf build
build:
	go clean
	go build -o ./tokensale main.go

contract: clean abi abigen
abi:
	solc --abi third_party/tokenManager.sol -o build
abigen:
	abigen --abi build/TokenManager.abi --pkg=tokensale --out=third_party/tokenManager.go