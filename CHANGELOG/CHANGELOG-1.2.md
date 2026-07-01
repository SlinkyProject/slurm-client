## v1.2.0

## v1.2.0-rc1

### Added

- Added test cases for reservation create, update, and delete methods.
- Added v0045 API.

### Fixed

- Updated moby/buildkit to v0.28.1 to address CVE-2026-33747.
- GO-2026-4864 GO-2026-4865 GO-2026-4866 GO-2026-4869 GO-2026-4870 GO-2026-4946
  GO-2026-4947.
- GO-2026-5005 GO-2026-5006 GO-2026-5013 GO-2026-5014 GO-2026-5015 GO-2026-5016
  GO-2026-5017 GO-2026-5018 GO-2026-5019 GO-2026-5020 GO-2026-5021 GO-2026-5023
  GO-2026-5024 GO-2026-5025 GO-2026-5026 GO-2026-5027 GO-2026-5028 GO-2026-5029
  GO-2026-5030 GO-2026-5033.
- CVE-2026-46680.
- CWE-168.
- GO-2026-5037 GO-2026-5038 GO-2026-5039.
- Fixed cases where a full buffered informer cache sync request channel would
  cause deadlock.
- Fix informerCache Get() and List() methods to avoid a possible hang when
  context is cancelled while waiting to write to sync channels.

### Changed

- Upgraded testcontainers-go to v0.42.0. This removes come direct dependencies
  of docker/docker and also addresses CVE-2026-39882 and CVE-2026-39883 by
  indirectly updating opentelemetry.

### Removed

- Removed v0041 API.
