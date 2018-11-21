SOLC     = solc
ABIGEN   = abigen
CONTRACT = BirrContract
PACKAGE  = contract

.PHONY: all
all: abigen solc build

.PHONY: solc
solc: $(CONTRACT).sol
	$(SOLC) $(CONTRACT).sol --abi --bin -o $(CONTRACT)

.PHONY: abigen
abigen: $(CONTRACT).sol
	abigen --sol $(CONTRACT).sol --pkg $(PACKAGE) --out $(CONTRACT).go

.PHONY: build
build: abigen
	go build

.PHONY: clean
clean:
	rm -rf $(CONTRACT)
	go clean
