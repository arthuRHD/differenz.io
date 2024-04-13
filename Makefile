test: install
	@go test ./...

coverage: install
	@go test ./... -coverprofile=cover.out -json

profile-mem: docker-run
	@go tool pprof http://localhost:8888/debug/pprof/heap

profile-cpu: docker-run
	@go tool pprof http://localhost:8888/debug/pprof/profile

profile-goroutines: docker-run
	@go tool pprof http://localhost:8888/debug/pprof/block 

build: install
	@go build -o server ./...

docker-build:
	@docker build -t ghcr.io/arthurrhd/differenz.io:local .

docker-run: docker-build
	@docker run -d -p 8888:8888 ghcr.io/arthurrhd/differenz.io:local 

run: install
	@go run .

install:
	@go install .

clean:
	@go clean
	@rm -rf cover.out TestResults*.json server *.pprof

