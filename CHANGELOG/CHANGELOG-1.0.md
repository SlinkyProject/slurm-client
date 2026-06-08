## v1.0.4

### Fixed

- Updated moby/buildkit to v0.28.1 to address CVE-2026-33747.
- GO-2026-4864 GO-2026-4865 GO-2026-4866 GO-2026-4869 GO-2026-4870 GO-2026-4946
  GO-2026-4947.
- GO-2026-5005 GO-2026-5006 GO-2026-5013 GO-2026-5014 GO-2026-5015 GO-2026-5016
  GO-2026-5017 GO-2026-5018 GO-2026-5019 GO-2026-5020 GO-2026-5021 GO-2026-5023
  GO-2026-5024 GO-2026-5025 GO-2026-5026 GO-2026-5027 GO-2026-5028 GO-2026-5029
  GO-2026-5030 GO-2026-5033.
- GO-2026-5037 GO-2026-5038 GO-2026-5039.

### Changed

- Upgraded testcontainers-go to v0.42.0. This removes some direct dependencies
  of docker/docker and also addresses CVE-2026-39882 and CVE-2026-39883 by
  indirectly updating opentelemetry.

## v1.0.3

### Fixed

- Updated google.golang.org/grpc to v1.79.3 to address CVE-2026-33186.

## v1.0.2

### Fixed

- Use SkipCache when returning objects that have been created or updated to
  ensure data coherency.

## v1.0.1

### Fixed

- Upgrade containerd to address CVE-2024-25621 and CVE-2025-64329.
- Upgrade crypto to address CVE-2025-58181 and CVE-2025-47914.
- Upgrade go-viper/mapstructure to 2.4.0.
- Upgrade docker to address CVE-2025-54388.
- Update go toolchain to 1.25.5.

## v1.0.0

## v1.0.0-rc1

### Added

- Added Slurm v44 API and types.

### Changed

- Regenerated Slurm API from 25.05.4.

### Removed

- Removed v40 API and types.
