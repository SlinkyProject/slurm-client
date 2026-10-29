## v1.3.0-rc1

### Added

- Added test cases for reservation create, update, and delete methods.
- Added v0045 API.
- File based token rotation.
- Added `errors` module to check known returned client errors.
- Expand Slurm host lists with correct padding and ordering.
- Compress host lists while preserving node names and ordering.
- Added RuntimeObject interface.
- Added method to expose the raw underlying versioned client interface.
- Support indexed lookups in Slurm informer caches.
- Return accepted object identities without a follow-up read.
- Support setting object identity through SetKey.
- Allow skipping readback after updates and deletions.

### Fixed

- Updated google.golang.org/grpc to v1.79.3 to address CVE-2026-33186.
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
- Fix security vulnerabilities.
- Clear syncErrorGet map entry on cache sync.
- GO-2026-4970 GO-2026-5856 GO-2026-5942 GO-2026-5970.
- GO-2026-5026 GO-2026-5942 GO-2026-5972 GO-2026-6088 GO-2026-6089 GO-2026-6090
  GO-2026-6091 GO-2026-6218.
- Fix setting duplicate ContentType on POST requests.
- Return recognizable not-found errors from live object lookups.
- Drastically improved performance of new type pattern object creation from
  Slurm API.
- GO-2026-6629.
- Convert Int to String with `strconv.Itoa()` instead of
  `fmt.Sprintf("%d", num)`.
- GO-2026-6599 GO-2026-6600 GO-2026-6603 GO-2026-6604 GO-2026-6605 GO-2026-6607
  GO-2026-6608 GO-2026-6609 GO-2026-6610 GO-2026-6611 GO-2026-6612 GO-2026-6613
  GO-2026-6617.
- Evict missing cached objects, emit deletions, and preserve update event
  values.
- Classify invalid job IDs returned by acknowledged submissions.
- Report type mismatches separately from unsupported client operations.
- Avoid a data race while dispatching Slurm informer add events.
- Avoid races and nil callbacks when changing informer handlers.
- GO-2026-6809.

### Changed

- Upgraded testcontainers-go to v0.42.0. This removes come direct dependencies
  of docker/docker and also addresses CVE-2026-39882 and CVE-2026-39883 by
  indirectly updating opentelemetry.
- Changed comparison function from `Semantic.DeepEqual` to `reflect.DeepEqual`
  to detect when fetched objects are different than the cached objects.
- Collect errors with standard `errors.Join()` instead of Kubernetes
  `utilerrors.NewAggregate()`.
- Rename the not-found error sentinel to ErrNotFound.
- RefreshCache reads detect completion in up to 50ms instead of up to
  syncPeriod/30 (1s at the default 30s syncPeriod).
- Changed object interfaces, RuntimeObject now has DeepCopyObject and ObjectList
  drops DeepCopyObjectList for DeepCopyObject.
- Changed GetStateAsSet from object to generic utility function.
- New type pattern (e.g., `pkg/types/*`) was removed, Slurm API object can be
  used directly. List were renamed to ObjectList to avoid name collision with
  existing types.
- Reduce allocations in cached lists and informer index lookups.
- Return a plain error for an acknowledged submission without a valid key.
- Keep HTTP 501 separate from unsupported client errors.
- Single-object cache refreshes no longer wait up to 1s.
- List cache refreshes no longer wait up to 1s.
- Delete refreshes the cached object in the background.
- Release explicit refresh callers when their fetch completes.
- Wake cache sync waiters on fetch completion instead of polling.

### Removed

- Removed v0041 API.
- Config.AuthToken removed in favor of Config.TokenProvider.
- Removed `V004(2|3|4|5)ControllerPingPinged(UP|DOWN)` constants.
- Removed `*Reconfigure` object.
