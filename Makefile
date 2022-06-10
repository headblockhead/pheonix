build:
	docker build -f Dockerfile.gen -t headblockhead/pheonix/gen:0.1 .
	docker build -f Dockerfile.print -t headblockhead/pheonix/print:0.1 .
run: build
	docker run -v phoenix_images:/phoenix_images -it headblockhead/pheonix/gen:0.1
	docker run -v phoenix_images:/phoenix_images --net=host -it headblockhead/pheonix/print:0.1