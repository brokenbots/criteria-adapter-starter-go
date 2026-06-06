# Registry-agnostic build + publish for the Go adapter starter.
#
# This mirrors what .github/workflows/publish.yml does in CI, for users not on
# GitHub Actions (GitLab CI, local, other CI). The publish step shells out to
# the `criteria` CLI, which performs the same manifest-emit -> OCI artifact ->
# (optional) cosign signature -> push as the reusable publish action.
.PHONY: build emit-manifest clean publish

build:
	mkdir -p out
	go build -o out/adapter .

emit-manifest:
	go run . --emit-manifest

clean:
	rm -rf out

# --- Publishing ---------------------------------------------------------------
# Publish the built adapter as an OCI artifact. Requires the `criteria` CLI on
# PATH (https://github.com/brokenbots/criteria).
#
#   make publish REGISTRY=ghcr.io/you/your-adapter:0.1.0
#
# Local runs publish UNSIGNED by default. Set SIGN_KEY=/path/to/cosign.key for
# explicit-key cosign signing. CI signs keyless via Sigstore — see
# .github/workflows/publish.yml and .gitlab-ci.yml.example.
REGISTRY ?= ghcr.io/$(USER)/criteria-adapter-starter-go:dev
CRITERIA ?= criteria
SIGN_KEY  ?=

publish: build
	$(CRITERIA) adapter publish out/adapter --registry "$(REGISTRY)" \
		$(if $(SIGN_KEY),--sign-key "$(SIGN_KEY)",)
