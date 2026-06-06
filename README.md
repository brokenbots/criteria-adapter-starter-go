# criteria-adapter-starter-go

A starter template for building a [Criteria](https://github.com/brokenbots/criteria)
adapter in Go using
[`criteria-go-adapter-sdk`](https://github.com/brokenbots/criteria-go-adapter-sdk).

> **Use this template** (green button on GitHub) or
> `gh repo create my-adapter --template brokenbots/criteria-adapter-starter-go`.

## Quickstart

```bash
# 1. Clone your new repo and update the module path in go.mod.

# 2. Edit main.go — change Info(), the schemas, and Execute() logic.

# 3. Inspect the generated manifest.
go run . --emit-manifest

# 4. Build the adapter binary.
go build -o out/adapter .

# 5. Publish: push a tag. The publish workflow builds, signs, and pushes an
#    OCI artifact to GHCR under your org.
git tag v0.1.0 && git push origin v0.1.0
```

The published adapter can then be pulled by a workflow:
`criteria adapter pull ghcr.io/<your-org>/<repo>:0.1.0`.

## What's in here

| Path | Purpose |
|------|---------|
| [`main.go`](main.go) | Entrypoint + `Service` implementation |
| [`.github/workflows/publish.yml`](.github/workflows/publish.yml) | Tag → build → sign → push (artifact-only) |
| [`Dockerfile`](Dockerfile) | Commented; uncomment to also publish a runnable image |
| [`examples/remote/`](examples/remote/) | k8s / docker-compose / systemd manifests for remote (phone-home) mode |

## Running remotely

To run the adapter remotely instead of letting the host launch it, call
`adapterhost.ServeRemote(&greeterAdapter{}, opts)` and deploy with one of the
manifests in [`examples/remote/`](examples/remote/). See the
[SDK README](https://github.com/brokenbots/criteria-go-adapter-sdk#running-as-a-remote-adapter)
for the full `ServeRemote` API.

## License

MIT
