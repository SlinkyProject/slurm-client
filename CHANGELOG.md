# ChangeLog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Added v43, v42, v41 Reconfigure object.
- Added ability to update the server and token used by the client.

### Fixed

- Fixed informer cache not handling v43, v42 objects, and v41 stats.
- Fixed v42 Stats not being hooked in.

### Changed

- Changed client cache start by `New()` instead of by `Start()`.

### Removed
