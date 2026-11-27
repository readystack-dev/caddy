[![ReadyStack](https://readystack.dev/logo.svg)](https://readystack.dev)

**Stable, Open-source Releases, Automated** · [readystack.dev](https://readystack.dev)

[![Docker Hub](https://img.shields.io/badge/Docker%20Hub-readystack%2Fcaddy-2496ED?logo=docker&logoColor=white)](https://hub.docker.com/r/readystack/caddy) [![Docker pulls](https://img.shields.io/docker/pulls/readystack/caddy)](https://hub.docker.com/r/readystack/caddy) [![Grade at release](https://img.shields.io/badge/grade%20at%20release-A%20%2894%25%29-2ea44f)](https://readystack.dev/products/caddy) [![Signed: cosign](https://img.shields.io/badge/signed-cosign-2ea44f)](https://readystack.dev/keys/cosign.pub) ![SBOM: SPDX](https://img.shields.io/badge/SBOM-SPDX-2ea44f) ![Provenance: SLSA v1](https://img.shields.io/badge/provenance-SLSA%20v1-2ea44f) ![Offline rebuild: verified](https://img.shields.io/badge/offline%20rebuild-verified-2ea44f)

# Caddy: ReadyStack build kit

The offline build kit for the Docker image `readystack/caddy:v2.11.9-CE-ubuntu24.04-r3`. It holds what the image is built from (upstream source, base images, toolchain, pinned dependencies) and the documents, so the image can be rebuilt with no network access (`docker build --network=none`).

- Release: [v2.11.9-CE-ubuntu24.04-r3](https://github.com/readystack-dev/caddy/releases/tag/v2.11.9-CE-ubuntu24.04-r3)
- Image: [`docker.io/readystack/caddy:v2.11.9-CE-ubuntu24.04-r3`](https://hub.docker.com/r/readystack/caddy/tags), manifest digest `sha256:b36d6ee63c6c76857f17623c09923ec0b93a5b0a27a4b465eb51016b4552314b`
- Product page: https://readystack.dev/products/caddy
- Source and build recipe: https://readystack.dev/products/caddy/recipe
- Build trail: https://readystack.dev/products/caddy/trail/v2.11.9-CE-ubuntu24.04-r3
- Defects: this repository's Issues (defects recorded from 25 November 2026 on) and https://readystack.dev/defects.html?product=caddy

## The image

Single container, single static binary, foreground process. No sidecars, no DB.

## Who uses Caddy and why

Developers and system administrators who serve websites, APIs and apps and want HTTPS without handling certificates by hand.

Caddy is a web server and reverse proxy with HTTPS on by default: it gets and renews certificates for public names and runs its own local certificate authority for internal names.

A typical use: A small team runs several web apps on one server. A short Caddyfile puts each app behind its own domain name as a reverse proxy; Caddy obtains and renews the certificates and redirects plain HTTP to HTTPS, so nobody copies certificate files around or tracks expiry dates.

From the upstream project's own description: https://github.com/caddyserver/caddy, https://caddyserver.com/docs/, https://caddyserver.com/docs/automatic-https.

## This release

- Docker Scout policy grade at release: **A** (94%), graded 2026-11-17 with `docker scout policy` (CLI 1.26.0, docker-default configuration).
  - Docker Scout counts 9 packages with copyleft licences (GPL, LGPL, MPL), all of them operating-system packages (it counts only packages whose licence it can read; the SBOM lists every component).
- Signed with ReadyStack's release key (fingerprint `04daa92344a52dbc`): https://readystack.dev/keys/cosign.pub.
- Runs as a non-root user (`caddy`).
- Attestations on the image: an SPDX SBOM and SLSA provenance v1.
- Rebuilt from this kit with network egress blocked on 2026-11-17; the rebuilt image digest is the published one.
- Base: Ubuntu 24.04.

The current grade, measured again every day, is on https://readystack.dev/products/caddy.

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

`recipe/` holds the kit's recipe files (Dockerfiles, build scripts, pins) as published at https://readystack.dev/products/caddy/recipe. The complete kit is the release download. Each release is a tag here. How to report a vulnerability: [SECURITY.md](SECURITY.md).

## About ReadyStack

**Stable, Open-source Releases, Automated.** ReadyStack releases popular open-source apps every week with every available critical and high security fix applied, plus full documentation, reference data, tests and a build recipe that recreates each release bit-for-bit. Each release is rebuilt from its pinned inputs with the network off, checked against its published digest, signed, and scanned and graded again every day. The images and build kits are free, and ReadyStack is soon to be a nonprofit.

[readystack.dev](https://readystack.dev) · [All products](https://readystack.dev/products.html) · [Build kits on GitHub](https://github.com/readystack-dev) · [Defect log](https://readystack.dev/defects.html) · contact@readystack.dev

— ReadyStack · https://readystack.dev
