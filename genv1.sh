alias openapi-generator="java -jar /Users/larryli/work/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar"

openapi-generator generate -i api/v1/datadog/api/openapi.yaml -g go -o api/v1/datadog --package-name datadog -p enumClassPrefix=true,withGoCodegenComment=true,goImportAlias=datadog -t .generator/templates
rm api/v1/datadog/go.mod api/v1/datadog/go.sum
goimports -w api/v1
