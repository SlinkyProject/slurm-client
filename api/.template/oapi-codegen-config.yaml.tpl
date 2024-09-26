---
package: {{SLURM_GO_MODULE}}
generate:
  client: true
  embedded-spec: true
  models: true
output: slurm.gen.go
output-options:
  skip-fmt: false
  skip-prune: true
