export GO_POST_PROCESS_FILE="`command -v gofmt` -w"

INPUT=cockroach-cloud-v1-2022-03-31.json

openapi-generator generate -i $INPUT -g go -o .
