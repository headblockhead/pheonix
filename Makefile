build:
	docker build -t headblockhead/phoenix/gen:0.1 gen
	docker build -t headblockhead/phoenix/print:0.1 print
print: build
	docker compose --file print.yaml up --abort-on-container-exit
decode: build
	-rm -r DecodeOutput
	docker compose --file decode_only.yaml up --abort-on-container-exit
	sudo cp -r /var/lib/docker/volumes/pheonix_DecodeOutput/_data ./DecodeOutput
lamdabuild:
	-rm -r gen/cdk/cdk.out
	-rm cdk_out_zip.zip
	cd gen/cdk; cdk synth
	zip -rj cdk_out_zip.zip gen/cdk/cdk.out/asset.*/bootstrap
	aws lambda update-function-code --zip-file fileb://cdk_out_zip.zip --function-name PhoenixGenerator-handlerE1533BD5-GaaDanuJjl28