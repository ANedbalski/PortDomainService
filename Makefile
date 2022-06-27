vendor:
	go mod vendor

coverage:
	./.ci/coverage.sh

generate:
	go generate ./...

# for CI to exclude slow tests
ci-test:
	go test -p 1  -short ./...

test:
	go test -p 1 ./..
