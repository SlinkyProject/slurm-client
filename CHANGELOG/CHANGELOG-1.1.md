## v1.1.1

### Fixed

- GO-2026-4864 GO-2026-4865 GO-2026-4866 GO-2026-4869 GO-2026-4870 GO-2026-4946
  GO-2026-4947.
- GO-2026-5005 GO-2026-5006 GO-2026-5013 GO-2026-5014 GO-2026-5015 GO-2026-5016
  GO-2026-5017 GO-2026-5018 GO-2026-5019 GO-2026-5020 GO-2026-5021 GO-2026-5023
  GO-2026-5024 GO-2026-5025 GO-2026-5026 GO-2026-5027 GO-2026-5028 GO-2026-5029
  GO-2026-5030 GO-2026-5033.
- GO-2026-5037 GO-2026-5038 GO-2026-5039.
- Fixed cases where a full buffered informer cache sync request channel would
  cause deadlock.
- Fix informerCache Get() and List() methods to avoid a possible hang when
  context is cancelled while waiting to write to sync channels.

### Changed

- Upgraded testcontainers-go to v0.42.0. This removes some direct dependencies
  of docker/docker and also addresses CVE-2026-39882 and CVE-2026-39883 by
  indirectly updating opentelemetry.

## v1.1.0

### Added

- Added test cases for reservation create, update, and delete methods.

### Fixed

- Updated google.golang.org/grpc to v1.79.3 to address CVE-2026-33186.
- Updated moby/buildkit to v0.28.1 to address CVE-2026-33747.

## v1.1.0-rc1

### Added

- Added CRUD methods for reservations.
- Add new endpoint for adding nodes to slurm.
- Added cache refresh batching to avoid redundant rapid refresh cache requests.
- Add basic fake client Start() and Stop() implementation.

### Fixed

- Upgrade containerd to address CVE-2024-25621 and CVE-2025-64329.
- Upgrade crypto to address CVE-2025-58181 and CVE-2025-47914.
- Upgrade go-viper/mapstructure to 2.4.0.
- Upgrade docker to address CVE-2025-54388.
- Update go toolchain to 1.25.5.
- Use SkipCache when returning objects that have been created or updated to
  ensure data coherency.
- Fixed cases where a cache refresh request on multiple objects simultaneously
  would incorrectly indicate all were refreshed when only the first was.
- Refresh cache after mutating object.
- Fixed stale objects in cache when they are deleted.
- Fixed cache mutation detection by comparing the values of pointers, not their
  addresses.
- Fixed case where a ReservationInfo and its Update request disagreed on the
  reservation name to be updated.
- CVE-2025-62725 Docker Compose path traversal (arbitrary file overwrite).

### Changed

- Internal processes will wait with timeout for cache to synchronize.
- Process RefreshCache requests in parallel.
- Buffer internal informer channels to help with internal queue processing.
- Start() now blocks, aligned with community conventions.
- Terminate WaitForCache loops early if cache informers are stopped mid
  operation.
- Informer `WaitForSync{List,Get}()` are no longer interface functions.

### Removed

- Remove Allocate from `Create()` options. This only was compatible with JobInfo
  creation.
