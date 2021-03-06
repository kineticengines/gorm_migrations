SHELL := /bin/bash
dir=$(shell echo $(CURDIR) )
home=$(shell echo $(HOME) )
dt=$(shell echo $(date))
version=v0.0.1

.PHONY: install-ci-deps install-packages checker run-tests build install

install-ci-deps:
	go install honnef.co/go/tools/cmd/staticcheck && \
	go install github.com/securego/gosec/cmd/gosec && \
	go get -u -v github.com/ory/go-acc && \
	go get -u -v github.com/axw/gocov/gocov


install-packages:
	go get -u -v ./... && go mod download -x && go mod tidy -v && go mod verify

checker: 
	staticcheck ./... && go vet ./...  && gosec -exclude=G101,G404 ./...

run-tests:
	go-acc -o coverage.txt ./... -- -timeout 10m && \
	go tool cover -html=coverage.txt -o coverage.html && gocov convert coverage.txt > coverage.json && \
	gocov report coverage.json > coverage_report.txt && tail coverage_report.txt && rm -rf cover*
	
build:
	go build -a -o gormgx -ldflags="-X 'github.com/kineticengines/gorm-migrations/pkg/commands.version=$(version)' -X 'github.com/kineticengines/gorm-migrations/pkg/commands.buildtime=$(dt)'" -trimpath -race github.com/kineticengines/gorm-migrations/pkg/

install:build	
	mv $(dir)/gormgx $(home)/go/bin/     


