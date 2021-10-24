
fmt:
	@echo "go fmt lrucache"	
	go fmt ./...

vet:
	@echo "go vet lrucache"	
	go vet ./...

lint:
	@echo "go lint lrucache"	
	golint ./...

test:
	@echo "Testing lrucache"	
	go test -v --cover ./...
