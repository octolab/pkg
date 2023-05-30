---
code: null
id: PVT_kwDOACPy4c4ABxWp
fullDatabaseId: "464297"
number: 6
url: https://github.com/orgs/octolab/projects/6
owner:
  id: MDEyOk9yZ2FuaXphdGlvbjIzNTU5Mzc=
  login: "octolab"
  type: Organization
title: "🏗 Platform"
shortDescription: "OctoLab Platform development."
public: true
closed: false
createdAt: 2022-05-11T08:35:10Z
updatedAt: 2023-10-18T13:30:11Z
closedAt: null
---

# 🏗 Platform

## Backlog

- [ ] octolab/cli#1 — feature: process manager
- [ ] octolab/cli#4 — debugger: improve code quality and prepare PoC to integrate it with templates
- [ ] octolab/cli#7 — config: compare os.UserHomeDir and github.com/mitchellh/go-homedir, make a decision about integration
- [ ] octolab/cli#18 — errors: define error with error code supporting
- [ ] octolab/protocol#12 — http: interceptor: request dumper with tracer support
- [ ] octolab/propaganda#1 — middlewares: define inventory for http.RoundTripper
- [ ] octolab/propaganda#2 — middlewares: define inventory for http.Handler
- [ ] octolab/propaganda#3 — headers: port from protocol
- [ ] octolab/propaganda#4 — spread: middlewares: support X-Features header
- [ ] octolab/propaganda#5 — spread: context: propagation mechanism over HTTP
- [ ] octolab/propaganda#6 — spread: interceptor: request debugger
- [ ] octolab/propaganda#7 — headers: add X-Request-UUID
- [ ] [[issue-25]] — review the docs
- [ ] [[issue-26]] — apply action for classified error
- [ ] [[issue-27]] — define converters
- [ ] [[issue-29]] — improve sync.Times
- [ ] [[issue-31]] — add interruptible reader/writer to io package
- [ ] [[issue-32]] — extend safe package by finalizers
- [ ] [[issue-33]] — calculate stats for classifier
- [ ] [[issue-34]] — add safe.StopTimer
- [ ] [[issue-35]] — extend errors.Is
- [ ] [[issue-36]] — testing: env: write test to reproduce errors with concurrent access to environment variables
- [ ] [[issue-37]] — embed github.com/mpvl/unique
- [ ] [[issue-38]] — add async.DelayedJob
- [ ] [[issue-45]] — improve io.Discard for defer
- [ ] [[issue-46]] — write benchmark to RepeatableReadCloser
- [ ] [[issue-47]] — add NoContent to testing
- [ ] [[issue-48]] — add ReadCloser to testing
- [ ] [[issue-49]] — port stream from testit
- [ ] [[issue-50]] — publish errors internal interfaces
- [ ] [[issue-51]] — find build tag go1.13
- [ ] [[issue-52]] — rethink SetEnvs from testing
- [ ] [[issue-54]] — extend time package based on git interface
- [ ] [[issue-55]] — port iotest and others
- [ ] [[issue-56]] — testing: env: refactor SetEnvs
- [ ] [[issue-57]] — throw instead panic
- [ ] [[issue-58]] — extend errors.Is
- [ ] [[issue-60]] — revert 19e671e when go1.18 will be available
- [ ] [[issue-61]] — adapter for embed.FS for afero.FS
- [ ] octolab/config#1 — input: improve config experience with Viper and Cobra
- [ ] [[issue-62]] — env: add helper method like a Set or Tuple for testing.Setenv
- [ ] [[issue-63]] — safe: Writer functionality
- [ ] octolab/config#3 — feature: decode from string
- [ ] octolab/config#4 — input: Regexp type for kelseyhightower/envconfig integration
- [ ] octolab/config#5 — input: Regexp type for spf13/pflag integration
- [ ] octolab/config#8 — input: Secret type support
- [ ] octolab/config#9 — feature: find way to build flexible nested feature set
- [ ] octolab/cli#19 — docs: readme: describe module and its purpose
- [ ] octolab/config#10 — docs: readme: describe module and its purpose
- [ ] [[issue-65]] — env: add analogue of `os.Getenv` and `os.Setenv`
- [ ] [[issue-66]] — env: dump to dotenv
- [ ] octolab/config#11 — input: File type for spf13/pflag integration
- [ ] octolab/cli#24 — deps: up github.com/spf13/cobra to v1.5.0
- [ ] octolab/cli#25 — command: set up NoFileCompletions for commands without args
- [ ] [[issue-67]] — safe: NoCopy
- [ ] octolab/cli#26 — cobra: version: extend by new possibilities of go 1.18
- [ ] [[issue-68]] — runtime: todo or die
- [ ] octolab/cli#27 — cobra: support new features
- [ ] [[issue-76]] — time: port code from maintainer and research jinzhu/now
- [ ] octolab/cli#28 — cobra: version: extend information

## To do

- [ ] [[issue-64]] — org: review and classify issues
- [ ] octolab/config#2 — deps: up to gopkg.in/yaml.v3
- [ ] octolab/cli#23 — deps: up github.com/spf13/viper to v1.12.0

## In progress

## Done

- [x] octolab/cli#20 — org: review and classify issues
- [x] octolab/config#7 — org: review and classify issues
- [x] octolab/propaganda#8 — org: review and classify issues
- [x] octolab/homebrew-tap#2 — broken formulas
- [x] octolab/protocol#16 — org: review and classify issues
- [x] octolab/homebrew-tap#3 — org: rename main branch and take a snapshot
- [x] octolab/homebrew-tap#4 — docs: readme: update list of binaries
- [x] octolab/homebrew-tap#5 — ci/cd: port healthcheck from the related repositories
- [x] octolab/homebrew-tap#6 — ci/cd: check installation process

## Won't do

- [x] octolab/cli#16 — make a decision about coral
- [x] octolab/cli#17 — research possibility to downgrade supported version to 1.11
