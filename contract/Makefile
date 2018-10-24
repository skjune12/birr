SOLC     = solc
ABIGEN   = abigen
CONTRACT = SimpleStorage
MAIN		 = main.go

.PHONY: solc
solc: $(CONTRACT).sol
	$(SOLC) $(CONTRACT).sol --abi --bin -o $(CONTRACT)

.PHONY: abigen
abigen: $(CONTRACT).sol
	abigen --sol=SimpleStorage.sol --pkg=main --out=contract.go

.PHONY: build
build: abigen
	go build

.PHONY: run
run: abigen
	go run main.go contract.go
