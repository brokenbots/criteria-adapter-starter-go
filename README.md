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
| [`Makefile`](Makefile) | `make build` / `make publish` — registry-agnostic local path |
| [`.github/workflows/publish.yml`](.github/workflows/publish.yml) | Tag → build → keyless-sign → push (GitHub Actions) |
| [`.gitlab-ci.yml.example`](.gitlab-ci.yml.example) | Same flow for GitLab CI (copy to `.gitlab-ci.yml`) |
| [`Dockerfile`](Dockerfile) | Commented; uncomment to also publish a runnable image |
| [`examples/remote/`](examples/remote/) | k8s / docker-compose / systemd manifests for remote (phone-home) mode |

## Publishing

Three equivalent paths produce the same signed OCI artifact:

- **GitHub Actions** — push a `v*` tag; [`publish.yml`](.github/workflows/publish.yml)
  builds, **keyless-signs** (Sigstore, via the job's OIDC identity), and pushes.
- **GitLab CI** — copy [`.gitlab-ci.yml.example`](.gitlab-ci.yml.example) to
  `.gitlab-ci.yml`; it does the same, signing keyless via GitLab's `id_tokens`.
- **Local / other CI** — `make publish REGISTRY=ghcr.io/you/your-adapter:0.1.0`.
  Requires the [`criteria`](https://github.com/brokenbots/criteria) CLI on PATH.
  Publishes unsigned by default; set `SIGN_KEY=/path/to/cosign.key` for
  explicit-key signing (interactive keyless is a CI-only path).

To also ship a runnable container image (for `environment.runtime = "docker"`),
build and push it from your own CI (its `Dockerfile`), then record it: add
`image: ghcr.io/you/your-adapter:0.1.0-image` to the action, or
`criteria adapter publish … --image <ref>` locally.

## Running remotely

To run the adapter remotely instead of letting the host launch it, call
`adapterhost.ServeRemote(&greeterAdapter{}, opts)` and deploy with one of the
manifests in [`examples/remote/`](examples/remote/). See the
[SDK README](https://github.com/brokenbots/criteria-go-adapter-sdk#running-as-a-remote-adapter)
for the full `ServeRemote` API.

## License

MIT
