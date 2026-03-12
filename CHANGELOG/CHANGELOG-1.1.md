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
