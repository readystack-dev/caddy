# Caddy v2.11.9 — ReadyStack release `v2.11.9-CE-ubuntu24.04-r3`

Generated 2026-11-17 from this release's build records and strategy. Statements that come from a single record name it in brackets; the others come from this release's strategy and build log.

## Release

- **Image:** `readystack/caddy:v2.11.9-CE-ubuntu24.04-r3` (linux/amd64). Index digest `sha256:06aa670e1a64b773e994558f4690e612bf89f6a52071382bc99b8aeaf73be299` [`docs/provenance/LOCKED_DIGEST.txt`].
- **Image manifest digest:** `sha256:b36d6ee63c6c76857f17623c09923ec0b93a5b0a27a4b465eb51016b4552314b` [`TARGET_DIGEST.txt`].
- **Upstream:** Caddy v2.11.9 built from source at `v2.11.9` of https://github.com/caddyserver/caddy — upstream notes: https://github.com/caddyserver/caddy/releases/tag/v2.11.9 [`docs/provenance/BUILD_RECIPE.json`].
- **Base image:** none (`FROM scratch`): a chiselled Ubuntu 24.04 runtime, the files Canonical's chisel cut from Ubuntu's own packages (`vendored/chisel/slices.txt`, `vendored/chisel/packages.tsv`); runtime user `caddy`.
- **Build:** `source-rebuild-go`, compiled offline with `--network=none`.
- **Supersedes:** ReadyStack release `v2.11.9-CE-ubuntu24.04-r2`.

## What changed in this release

- Upstream version unchanged (v2.11.9); ReadyStack rebuild.
- Minimum dependency versions ReadyStack enforces on top of upstream (raised only where upstream's own version was lower); shipped versions from the SBOM [`docs/sbom/SBOM.json`]:
  - `golang.org/x/crypto` raised from upstream's `v0.57.0` to at least `v0.58.0` (shipped: `v0.58.0`)
  - `golang.org/x/net` raised from upstream's `v0.60.0` to at least `v0.62.0` (shipped: `v0.62.0`)
  - `golang.org/x/sys` raised from upstream's `v0.48.0` to at least `v0.49.0` (shipped: `v0.49.0`)
  - `golang.org/x/text` raised from upstream's `v0.42.0` to at least `v0.43.0` (shipped: `v0.43.0`)
  - `google.golang.org/grpc` at least `v1.83.2`: not raised, upstream already pins `v1.83.2` (shipped: `v1.83.2`)
  - `github.com/klauspost/compress` raised from upstream's `v1.20.1` to at least `v1.20.3` (shipped: `v1.20.3`)
- Caddy v2.11.9 is unchanged from the release this supersedes (v2.11.9-CE-ubuntu24.04-r2); this release rebuilds it. Changes: libc6 moves from 2.39-0ubuntu8.9 to 2.39-0ubuntu8.12 (USN-9065-1, UBUNTU-CVE-2026-84243), and the Go toolchain it is built with from 1.26.10 to 1.26.11, the newest patch of its line. Settings, ports and data volumes are those of v2.11.9-CE-ubuntu24.04-r2: replace the container with this image on the same volumes.

## Security status at release (measured 2026-11-17)

- **Docker Scout:** 0 fixable critical/high vulnerabilities; runs as `caddy` (non-root); SBOM + provenance attestations present (SPDX SBOM + SLSA v1) [`docs/provenance/SCOUT_VERDICT.json`].
- **OSV cross-check:** no fixable critical/high vulnerability beyond Docker Scout's findings.
- **Signature:** cosign signature verified by ReadyStack's release gate before publishing; check it yourself with `cosign verify --key https://readystack.dev/keys/cosign.pub readystack/caddy:v2.11.9-CE-ubuntu24.04-r3` (signatures are recorded in the public Sigstore transparency log).
- **ReadyStack release check at release:** passed (ReadyStack's own test, not a Docker Scout grade). Vulnerabilities are disclosed every day and the image never changes, so it is re-checked daily, beside Docker Scout's own dated grade, at https://readystack.dev/queue.

## Reproducibility

Rebuilds offline from this archive to the identical image digest `sha256:b36d6ee63c6c76857f17623c09923ec0b93a5b0a27a4b465eb51016b4552314b`: run `./BUILD_OFFLINE.sh`, then `./VERIFY_OFFLINE.sh` (add `--online` to confirm the Docker Hub tag holds the same digest). Before publishing, ReadyStack rebuilt the kit with `./BUILD_OFFLINE.sh` and compared the result with the published digest [`docs/provenance/REPRO_VERIFICATION.json`], then ran `./BUILD_OFFLINE.sh` and `./VERIFY_OFFLINE.sh` from this archive (without `--online`).

## Validation

- 2 of 2 use cases passed (3 authored, 2 kept in the suite; 1 not exercised, see below) against the running image [`docs/validation/test-results.json`].
- The cases are `GET /health` (200), `GET /config/` (200); the suite checks nothing else.
- 1 further authored check was not exercised and is not counted above: it (`GET /`) was answered with the same page the build's test container served for a made-up path, so it could not show the path exists (listed under Not exercised in `docs/validation/FAILED_TESTS.md`).

## Known issues

- None found (reviewed by held semantic review (Claude) on 2026-11-17).

## Verify this archive

```
sha256sum -c CHECKSUMS.sha256
./BUILD_OFFLINE.sh && ./VERIFY_OFFLINE.sh --online
cosign verify --key https://readystack.dev/keys/cosign.pub readystack/caddy:v2.11.9-CE-ubuntu24.04-r3
```

ReadyStack records every signature in the public Sigstore transparency log (Rekor), so the command needs no extra flags.

