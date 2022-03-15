#
# Variables
#

# Specified from caller
OS_ARCH = local
EXT =

# Fixed
RM = rm -f
GOCMD = go


#
# Rules
#

.PHONY: clean

all: clean redundantimportalias

redundantimportalias: redundantimportalias.go
	@echo ">> Compiling..."
	$(GOCMD) build $<

clean:
	@echo ">> Cleaning up..."
	$(RM) redundantimportalias

.PHONY: deps
deps:
	@echo ">> Resolving dependencies..."
	go mod download
	go mod tidy
.PHONY: deps-update
deps-update:
	@echo ">> Updating dependencies..."
	go get -d -u -t ./...
	@make deps
