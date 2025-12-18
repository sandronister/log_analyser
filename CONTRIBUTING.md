# Contributing

Thanks for taking the time to contribute to Log Analyser! This document explains how to report issues, propose changes, and keep the project healthy.

## Code of Conduct

Participation in this project is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). Please read it to understand our expectations before filing an issue or opening a pull request.

## Ways to Help

- **Report bugs**: Use GitHub Issues and include reproduction steps, expected behavior, actual behavior, and log snippets if possible.
- **Request features**: Describe the use case clearly; mock ups or pseudo logs are helpful.
- **Improve documentation**: Fix typos, clarify instructions, or add missing sections.
- **Submit code**: Fix bugs, add features, or improve tooling/tests.

## Development Setup

1. Fork the repository and clone your fork locally.
2. Install Go `1.24.1` or newer.
3. Download dependencies:
   ```bash
   go mod download
   ```
4. Adjust the provided `.env` file if needed (the default `FOLDER_PATH=log_files` is enough to run the sample data).

### Running the Analyzer

```bash
go run cmd/main.go
```

### Running Tests

```bash
go test ./...
```

Add tests for every bug fix or new feature whenever feasible. If you cannot add tests, explain why in the pull request.

## Pull Request Checklist

- Create a topic branch from `main`: `git checkout -b feature/short-description`.
- Follow Go formatting conventions (`go fmt ./...` will be run in CI).
- Keep commits focused. Squash fixup commits before requesting review.
- Update documentation and configuration examples when behavior changes.
- Ensure `go test ./...` passes locally.
- Link related issues in the pull request description (`Fixes #123`).

## Issue Triage

Maintainers use labels to organize work. Feel free to triage issues by:

- Verifying reproducibility and adding details.
- Suggesting labels (bug, enhancement, help wanted, good first issue).
- Mentioning related issues or pull requests.

## Release Notes

Add entries to `CHANGELOG.md` under the **Unreleased** section for any user-facing change. Group items by `Added`, `Changed`, `Deprecated`, `Removed`, `Fixed`, or `Security` when appropriate.

## Questions?

If you are unsure about anything, open a draft issue or pull request and ask for guidance. We would rather discuss an idea early than rework it later. Happy hacking!
