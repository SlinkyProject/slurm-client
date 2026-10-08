---
# yaml-language-server: $schema=https://raw.githubusercontent.com/oapi-codegen/oapi-codegen/{{OAPI_CODEGEN_VERSION}}/configuration-schema.json
package: {{SLURM_GO_MODULE}}
generate:
  client: true
  embedded-spec: true
  models: true
output: slurm.gen.go
output-options:
  skip-fmt: false
  skip-prune: true
