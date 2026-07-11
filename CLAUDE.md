# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About

A Go CLI that acts as a drop-in `sendmail` replacement, translating sendmail-style invocation (args + a piped RFC822 email on stdin) into an `msmtp` invocation. It exists so that programs which shell out to `sendmail` can transparently deliver mail via `msmtp` instead.

## Commands

- Compile: `./step-compile.sh` (or `go build -o build/bin/sendmail ./sendmail`) → binary at `build/bin/sendmail`
- Test all: `./step-test.sh` (or `go test ./sendmail`)
- Test single: `go test ./sendmail -run TestName`
- Clean: `./step-clean.sh` (removes `build/`)
- Full local build+test: `./create-local-release.sh` (clean → compile → test)
- Local build without tests: `./create-local-release-no-tests.sh`
- Public release (tags, builds `.deb`, requires clean state and a version arg): `./create-public-release.sh X.X.X`

The release scripts are composed of the `step-*.sh` scripts, each runnable standalone; `step-debian-create.sh` and `step-git-tag.sh` expect a `$VERSION` env var (set by the calling `create-*` script).

## Architecture

Single Go module (`github.com/foilen/sendmail-to-msmtp`), single package `main` under `sendmail/`:

- `main.go` — entry point. Detects `/etc/sendmail-to-msmtp.json` if present, calls `process()` to build the msmtp command, optionally copies the spooled email into a debug dump directory, then execs `msmtp` with the built args and the spooled file piped as stdin. The config path, `msmtp` binary path, and `msmtp` config path (`-C`) can each be overridden with an env var — see "Environment variables" below.
- `process.go` — the core logic (`process(ctx *ProcessContext) []string`). Parses sendmail-style CLI args (`-r`/`-f` for sender, `-t` to read recipients from headers, trailing non-flag args as recipients), streams stdin to a temp file line-by-line, parses headers to extract `From:` (handling folded/multi-line headers and both `Name <addr>` and bare `addr` forms), and returns the full `msmtp` argv. All I/O and side effects go through `ProcessContext` rather than globals, which is what makes this function unit-testable.
- `process_context.go` — `ProcessContext` struct: carries CLI args, the stdin reader, and mutable state (config path, msmtp binary/config path overrides, spooled file path, dump file prefix) between `main.go` and `process.go`.
- `config-file.go` — reads/unmarshals the optional JSON config file (`DefaultFrom`, `EmailDumpDirectory`).
- `process_test.go` + `sendmail/testdata/*` — table-driven-style tests feeding raw `.txt`/`.json` fixtures through `process()` and asserting the resulting msmtp argv and spooled file contents.

### "From" precedence

The sender address is resolved in increasing priority (later overrides earlier), per README: config file `defaultFrom` → `-r` arg → `-f` arg → `From:` header. This ordering is implemented directly in the arg-parsing loop and header-parsing loop in `process.go` — preserve it when touching that logic.

### Environment variables

`main.go` reads these to override defaults, mirroring one convention (`SENDMAIL_TO_MSMTP_*`):

- `SENDMAIL_TO_MSMTP_CONFIG_PATH` — overrides `/etc/sendmail-to-msmtp.json`
- `SENDMAIL_TO_MSMTP_MSMTP_PATH` — overrides `/usr/bin/msmtp`
- `SENDMAIL_TO_MSMTP_MSMTPRC_PATH` — if set, passed to `msmtp` as `-C <path>` (`ctx.msmtpConfigPath` in `process.go`)

### Debug email dump

If `emailDumpDirectory` is set in the config, every run writes `<timestamp>-raw.eml` (untouched stdin) and, from `main.go`, a `<timestamp>-sendmail.eml` (the exact spooled file handed to msmtp). These are never cleaned up automatically.
