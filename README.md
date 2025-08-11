# Slurm Client

<div align="center">

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg?style=for-the-badge)](./LICENSES/Apache-2.0.txt)
[![Tag](https://img.shields.io/github/v/tag/SlinkyProject/slurm-client?style=for-the-badge)](https://github.com/SlinkyProject/slurm-client/tags/)
[![Go-Version](https://img.shields.io/github/go-mod/go-version/SlinkyProject/slurm-client?style=for-the-badge)](./go.mod)
[![Last-Commit](https://img.shields.io/github/last-commit/SlinkyProject/slurm-client?style=for-the-badge)](https://github.com/SlinkyProject/slurm-client/commits/)

</div>

[OpenAPI] Golang client library for [Slurm REST API][rest-api]. A [Slinky]
project.

## Table of Contents

<!-- mdformat-toc start --slug=github --no-anchors --maxlevel=6 --minlevel=1 -->

- [Slurm Client](#slurm-client)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Features](#features)
  - [Limitations](#limitations)
  - [Installation](#installation)
    - [Usage](#usage)
      - [Create](#create)
      - [Delete](#delete)
      - [Get](#get)
      - [List](#list)
      - [Update](#update)
  - [Upgrades](#upgrades)
    - [0.X Releases](#0x-releases)
  - [Support and Development](#support-and-development)
  - [License](#license)

<!-- mdformat-toc end -->

## Overview

This repository provides generated [slurmrestd][rest-api] endpoints for Golang
clients, and client wrapper library inspired by the [controller-runtime] client
for Kubernetes.

## Features

- **Multiple Slurm Versions**: can interact with multiple versions of Slurm.
- **Caching**: client-side caching mechanism to reduce number of server calls.

## Limitations

Currently, the client wrapper implements a limited amount of Slurm endpoints,
specifically only Node and Job endpoints. The number of implemented endpoints
may expand in the future, by request or community efforts.

- **Slurm Version**: >=
  [24.05](https://www.schedmd.com/slurm-version-24-05-0-is-now-available/)

## Installation

```bash
go get -v github.com/SlinkyProject/slurm-client
```

### Usage

Create a Slurm client handle.

```golang
config := &client.Config{
	Server:    "http://<host>:6820",
	AuthToken: "<`auth/jwt` token>",
}
slurmClient, err := client.NewClient(config)
if err != nil {
	return err
}
```

Start Slurm client cache.

```golang
ctx := context.Background()
go slurmClient.Start(ctx)
```

#### Create

Create Slurm resources via client handle.

```golang
// Create job via V0042 endpoint
jobInfo := &types.V0042JobInfo{}
req := v0042.V0042JobSubmitReq{
	Job: &v0042.V0042JobDescMsg{
		CurrentWorkingDirectory: ptr.To("/tmp"),
		Environment: &v0042.V0042StringArray{
			"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
		},
		Script: ptr.To("#!/usr/bin/env bash\nsleep 30"),
	},
}
if err := slurmClient.Create(ctx, jobInfo, req); err != nil {
	return err
}
```

```golang
// Create job via V0041 endpoint
jobInfo := &types.V0041JobInfo{}
req := v0041.V0041JobSubmitReq{
	Job: &v0041.V0041JobDescMsg{
		CurrentWorkingDirectory: ptr.To("/tmp"),
		Environment: &v0041.V0041StringArray{
			"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
		},
		Script: ptr.To("#!/usr/bin/env bash\nsleep 30"),
	},
}
if err := slurmClient.Create(ctx, jobInfo, req); err != nil {
	return err
}
```

```golang
// Create job via V0040 endpoint
jobInfo := &types.V0040JobInfo{}
req := v0040.V0040JobSubmitReq{
	Job: &v0040.V0040JobDescMsg{
		CurrentWorkingDirectory: ptr.To("/tmp"),
		Environment: &v0040.V0040StringArray{
			"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
		},
		Script: ptr.To("#!/usr/bin/env bash\nsleep 30"),
	},
}
if err := slurmClient.Create(ctx, jobInfo, req); err != nil {
	return err
}
```

#### Delete

Delete Slurm resource via client handle.

```golang
// Delete job via V0042 endpoint
jobInfo := &types.V0042JobInfo{
	V0042JobInfo: v0042.V0042JobInfo{
		JobId: ptr.To("1"),
	},
}
if err := slurmClient.Delete(ctx, jobInfo); err != nil {
	return err
}
```

```golang
// Delete job via V0041 endpoint
jobInfo := &types.V0041JobInfo{
	V0041JobInfo: v0041.V0041JobInfo{
		JobId: ptr.To("1"),
	},
}
if err := slurmClient.Delete(ctx, jobInfo); err != nil {
	return err
}
```

```golang
// Delete job via V0040 endpoint
jobInfo := &types.V0040JobInfo{
	V0040JobInfo: v0040.V0040JobInfo{
		JobId: ptr.To("1"),
	},
}
if err := slurmClient.Delete(ctx, jobInfo); err != nil {
	return err
}
```

#### Get

Get Slurm resource via client handle.

```golang
// Fetch node via V0042 endpoint
node := &types.V0042Node{}
key := object.ObjectKey("node-0")
if err := slurmClient.Get(ctx, key, node); err != nil {
	return err
}
```

```golang
// Fetch node via V0041 endpoint
node := &types.V0041Node{}
key := object.ObjectKey("node-0")
if err := slurmClient.Get(ctx, key, node); err != nil {
	return err
}
```

```golang
// Fetch node via V0040 endpoint
node := &types.V0040Node{}
key := object.ObjectKey("node-0")
if err := slurmClient.Get(ctx, key, node); err != nil {
	return err
}
```

#### List

List Slurm resources via client handle.

```golang
// Fetch list of nodes via V0042 endpoint
nodeList := &types.V0042NodeList{}
if err := slurmClient.List(ctx, nodeList); err != nil {
	return err
}
```

```golang
// Fetch list of nodes via V0041 endpoint
nodeList := &types.V0041NodeList{}
if err := slurmClient.List(ctx, nodeList); err != nil {
	return err
}
```

```golang
// Fetch list of nodes via V0040 endpoint
nodeList := &types.V0040NodeList{}
if err := slurmClient.List(ctx, nodeList); err != nil {
	return err
}
```

#### Update

Update Slurm resource via client handle.

```golang
// Update job via V0042 endpoint
jobInfo := &types.V0042JobInfo{}
req := &v0042.V0042JobDescMsg{
	Comment: ptr.To("updated comment")
}
if err := slurmClient.Update(ctx, jobInfo, req); err != nil {
	return err
}
```

```golang
// Update job via V0041 endpoint
jobInfo := &types.V0041JobInfo{}
req := &v0041.V0041JobDescMsg{
	Comment: ptr.To("updated comment")
}
if err := slurmClient.Update(ctx, jobInfo, req); err != nil {
	return err
}
```

```golang
// Update job via V0040 endpoint
jobInfo := &types.V0040JobInfo{}
req := &v0040.V0040JobDescMsg{
	Comment: ptr.To("updated comment")
}
if err := slurmClient.Update(ctx, jobInfo, req); err != nil {
	return err
}
```

## Upgrades

```bash
go get -u github.com/SlinkyProject/slurm-client
```

### 0.X Releases

The source tree may evolve more aggressively during these release versions, so
upgrades may require updated paths in addition to the version.

## Support and Development

Feature requests, code contributions, and bug reports are welcome!

Github/Gitlab submitted issues and PRs/MRs are handled on a best effort basis.

The SchedMD official issue tracker is at <https://support.schedmd.com/>.

To schedule a demo or simply to reach out, please
[contact SchedMD][contact-schedmd].

## License

Copyright (C) SchedMD LLC.

Licensed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0) you
may not use project except in compliance with the license.

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.

<!-- Links -->

[contact-schedmd]: https://www.schedmd.com/slurm-resources/contact-schedmd/
[controller-runtime]: https://github.com/kubernetes-sigs/controller-runtime
[openapi]: https://www.openapis.org/
[rest-api]: https://slurm.schedmd.com/rest_api.html
[slinky]: https://slinky.ai/
