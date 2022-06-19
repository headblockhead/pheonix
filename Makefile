build:
	docker build -t headblockhead/phoenix/gen:0.1 gen
	docker build -t headblockhead/phoenix/print:0.1 print
print: build
	docker compose --file print.yaml up --abort-on-container-exit
decode: build
	-rm -rf DecodeOutput
	docker compose --file decode_only.yaml up --abort-on-container-exit
	sudo cp -r /var/lib/docker/volumes/pheonix_DecodeOutput/_data ./DecodeOutput
lamdabuild:
	cd gen/cdk
	cdk synth
	cd cdk.out
	zip -r ../../output.zip asset.*/bootstrap