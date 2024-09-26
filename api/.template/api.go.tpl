// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package {{SLURM_GO_MODULE}}

// Ref: https://slurm.schedmd.com/rest_api.html

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@{{OAPI_CODEGEN_VERSION}} -config oapi-codegen-config.yaml slurm-openapi.gen.json
