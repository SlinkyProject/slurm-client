# Slurm Client

This Go library is responsible for enabling communication with [Slurm]. It
relies on communication with [slurmrestd] and auth/jwt based authentication.

Communication to the versioned [rest-api] endpoints are multiplexed for seamless
request/response handling of API endpoints. This abstracts the Slurm OpenAPI
Spec endpoints and objects for easy use. It supports a number of Slurm versions.

Current supported versions are: Slurm 24.05; data parsers v40, v41

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

[rest-api]: https://slurm.schedmd.com/rest_api.html
[slurm]: https://www.schedmd.com/slurm
[slurmrestd]: https://slurm.schedmd.com/slurmrestd.html
