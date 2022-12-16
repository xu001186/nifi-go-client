
.PHONY: generate-code
generate-code:
	docker run --rm -v ${PWD}/:/local swaggerapi/swagger-codegen-cli-v3 generate -i /local/swagger/nifi-1.17.0.json -l go -o /local/nifi/ --additional-properties packageName=nifi
