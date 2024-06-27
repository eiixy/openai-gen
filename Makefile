install-openapi-generator-cli:
	npm install @openapitools/openapi-generator-cli -g
	openapi-generator-cli version

update:
	wget https://github.com/openai/openai-openapi/raw/master/openapi.yaml

gen:
	openapi-generator-cli generate --skip-validate-spec -i ./openapi.yaml -g go -o ./ --package-name openai --git-user-id eiixy --git-repo-id openai-gen
