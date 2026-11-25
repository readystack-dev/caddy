# Caddy: ReadyStack build kit

The offline build kit for the Docker image `readystack/caddy:v2.11.9-CE-ubuntu24.04-r3`. It holds what the image is built from (upstream source, base images, toolchain, pinned dependencies) and the documents, so the image can be rebuilt with no network access (`docker build --network=none`).

- Release: [v2.11.9-CE-ubuntu24.04-r3](https://github.com/readystack-dev/caddy/releases/tag/v2.11.9-CE-ubuntu24.04-r3)
- Image: [`docker.io/readystack/caddy:v2.11.9-CE-ubuntu24.04-r3`](https://hub.docker.com/r/readystack/caddy/tags), manifest digest `sha256:b36d6ee63c6c76857f17623c09923ec0b93a5b0a27a4b465eb51016b4552314b`
- Product page: https://readystack.dev/products/caddy
- Defects: this repository's Issues (defects recorded from 25 November 2026 on) and https://readystack.dev/defects.html?product=caddy

Caddy is distributed under its upstream licence (Apache-2.0); the kit contains upstream's source under that licence, and each vendored component under its own.

## Download

From the [v2.11.9-CE-ubuntu24.04-r3 release](https://github.com/readystack-dev/caddy/releases/tag/v2.11.9-CE-ubuntu24.04-r3): `ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.tar.gz` (440.2 MB), SHA-256 `8472b5ea3e03502cf529b8087a2795c92e77cebb08b24f7939ecdadc036b14b5`.

```sh
sha256sum --check --ignore-missing SHA256SUMS
```

## Verify and rebuild

The rebuild statement is signed with the ReadyStack release key. Check the signature before you rely on it:

```sh
cosign verify-blob --key https://readystack.dev/keys/cosign.pub \
  --bundle ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.rebuild.json.sigstore ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.rebuild.json
```

Unpack with `-p` so file modes are kept, then follow the kit's own README.md and docs/REBUILD.md:

```sh
tar -xpzf ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.tar.gz
cd v2.11.9-CE-ubuntu24.04-r3
less README.md docs/REBUILD.md
```

The build trail (`ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.trail.json`, signed the same way, and its Markdown copy) records how this release was built and checked.

## This repository

`recipe/` holds the kit's recipe files (Dockerfiles, build scripts, pins) as published at https://readystack.dev/products/caddy/recipe. The complete kit is the release download. Each release is a tag here.
