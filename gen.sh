export GO_POST_PROCESS_FILE="`command -v gofmt` -w"

INPUT=cockroach-cloud-v1-2022-03-31.json
GIT_USER=jess-sheneberger
GIT_REPO=cockroach-cloud-golang
PACKAGE=ccloud

openapi-generator generate \
    --enable-post-process-file \
    -i $INPUT -g go --remove-operation-id-prefix \
    --git-repo-id $GIT_REPO --git-user-id $GIT_USER \
    --package-name $PACKAGE -o .

# for some reason the generated code is missing some curly braces, was too lazy to figure out why
git apply patch-build-error-1.patch