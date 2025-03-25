# ChangeLog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

### Fixed

### Changed

### Removed

## v0.2.0

### Added

- Added `Create()` and `Delete()` support for v40 JobInfo.
- Added `Create()` and `Delete()` support for v41 JobInfo, including
  "placeholder" job.
- Added aggregated errors created from Slurm API errors.
- Added job `GetStateAsSet()` helper functions.
- Added partition `GetStateAsSet()` helper functions.
- Added `Update()` support for v4{0,1} JobInfo
- Added WaitRefreshCache option for Get and List, which will wait for the next
  cache refresh before reading from cache.

### Fixed

- Fixed race condition between client cache Start and Stop.
- Fixed object.Object cast panics when using AppendItem().

### Changed

- Changed object name `Ping` => `ControllerPing`.
- Changed `fake` and `interceptor` module path.
- Changed Slurm object types to their versioned representation.
- Changed `Update()` function signature, added `req` parameter.
- Changed the `Create()` function signature, added `req` parameter.

### Removed

- Removed the skipCache client option in operations that do not read cache.
- Removed unimplemented `DeleteAllOf()` functionality.
