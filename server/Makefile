build:
	go build -o goboardgames main.go

run: build
	./goboardgames

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run