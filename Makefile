dep:
	dep ensure -v -vendor-only

pretty:
	gofmt -s -w .

bin: pretty dep
	go build -o bukarehat app/bukarehatbot/main.go

run: bin
	./bukarehat

deploy:
	sudo cp bukarehat.service /lib/systemd/system/bukarehat.service

# Only for development
dev:
	go build -o bukarehat app/bukarehatbot/main.go
	./bukarehat