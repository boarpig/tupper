tupper: tupper.go
	6g tupper.go
	6l -o tupper tupper.6

clean:
	rm tupper.6
	rm tupper

run: tupper
	./tupper
