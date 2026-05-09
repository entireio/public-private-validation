# dummy

A placeholder Go CLI used to validate Entire's split-storage workflow: source code lives in this **public** repository, while checkpoints (sessions, transcripts, intermediate snapshots) are pushed to a separate **private** repository.

The CLI itself does nothing useful — it compiles and prints help. The interesting bits are the surrounding scaffold (mise, golangci-lint, GitHub Actions) so the repo behaves like a real Go project when checkpoints are recorded against it.

See [docs.entire.io](https://docs.entire.io) for the public/private checkpoint setup.

## Usage

```sh
mise run build
./dummy help
```

## Common tasks

| Task                | Command            |
| ------------------- | ------------------ |
| Build               | `mise run build`   |
| Format              | `mise run fmt`     |
| Lint                | `mise run lint`    |
| Test                | `mise run test`    |
| Format + lint + test | `mise run check`  |
