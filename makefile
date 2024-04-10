build: 
	go build -o example/main.exe -ldflags "-X 'github.com/pragmaticengineering/go-version.BuildVersion=1.0.0'" example/main.go