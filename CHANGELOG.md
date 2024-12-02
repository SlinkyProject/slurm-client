# ChangeLog

All notable changes to this project will be documented in this file.

## \[Unreleased\]

### Added

- Added `Create()` and `Delete()` support for v40 JobInfo.

### Fixed

- Fixed race condition between client cache Start and Stop.

### Changed

- Changed object name `Ping` => `ControllerPing`.
- Changed `fake` and `interceptor` module path.
- Changed Slurm object types to their versioned representation.
- Changed `Update()` function signature, added `req` parameter.
- Changed the `Create()` function signature, added `req` parameter.

### Removed

- Removed the skipCache client option in operations that do not read cache.
- Removed unimplemented `DeleteAllOf()` functionality.
