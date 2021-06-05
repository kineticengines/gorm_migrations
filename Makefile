SHELL := /bin/bash
dir=$(shell echo $(CURDIR) )
home=$(shell echo $(HOME) )
dt=$(shell echo $(date))
version=v0.0.1

.PHONY: install-ci-deps install-packages checker run-tests build install ci

install-ci-deps:
	export PATH=$PATH:$GOPATH/bin
	go install honnef.co/go/tools/cmd/staticcheck@latest && \
	go get -u github.com/securego/gosec/v2/cmd/gosec && \
	go get -u -v github.com/ory/go-acc && \
	go get -u -v github.com/axw/gocov/gocov && \
	go get github.com/fzipp/gocyclo/cmd/gocyclo && \
	go get -u github.com/client9/misspell/cmd/misspell

install-packages:
	go env && go get -u -v ./... && go mod download -x && go mod tidy -v && go mod verify 	

checker:	
	staticcheck ./... && go vet ./...  && gosec -exclude=G101,G404 ./... && gocyclo . && misspell .

run-tests:
	go-acc -o coverage.txt ./... -- -timeout 10m && \
	go tool cover -html=coverage.txt -o coverage.html && gocov convert coverage.txt > coverage.json && \
	gocov report coverage.json > coverage_report.txt && tail coverage_report.txt && rm -rf cover* 

# ci: install-ci-deps install-packages run-tests 
ci: install-ci-deps install-packages # TODO: restore ci
	
build:
	go build -a -o gormgx -ldflags="-X 'github.com/kineticengines/gorm-migrations/pkg/commands.version=$(version)' -X 'github.com/kineticengines/gorm-migrations/pkg/commands.buildtime=$(dt)'" -trimpath -race github.com/kineticengines/gorm-migrations/pkg/

install:build	
	mv $(dir)/gormgx $(home)/go/bin/     


