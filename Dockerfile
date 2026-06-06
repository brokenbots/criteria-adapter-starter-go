# Optional: opt into publishing a runnable container image (D12).
#
# By default this starter publishes an artifact-only OCI package (the single Go
# binary) and does NOT build an image. Uncomment and set `with_image: "true"`
# in .github/workflows/publish.yml only if your adapter needs a runnable image
# (e.g. to run standalone in Kubernetes/ECS as a remote phone-home adapter).
#
# FROM gcr.io/distroless/static-debian12
# COPY out/adapter /usr/local/bin/adapter
# ENTRYPOINT ["/usr/local/bin/adapter"]
