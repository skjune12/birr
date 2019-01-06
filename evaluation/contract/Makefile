SOLC = solc
ABIGEN = abigen
CONTRACT = BirrContract
PKGNAME = contract

.PHONY: all
all: clean solc abigen

.PHONY: solc
solc: $(CONTRACT).sol
	$(SOLC) $(CONTRACT).sol --abi --bin -o .

.PHONY: abigen
abigen: $(CONTRACT).sol
	abigen --sol $(CONTRACT).sol --pkg $(PKGNAME) --out=$(CONTRACT).go

.PHONY: clean
clean:
	rm -rf $(CONTRACT).abi $(CONTRACT).bin
