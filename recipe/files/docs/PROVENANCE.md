# Provenance — Caddy (`v2.11.9-CE-ubuntu24.04-r3`)

Every item below is copied from a record in this archive; the record is named in brackets.

## Source

- Upstream repository: https://github.com/caddyserver/caddy
- Revision built: `v2.11.9` (commit `0d98caea58f794f88e709b1bef2880996a471bca`)
- Source tree: `5d2facdc29aac86a8324199b17b926492644473e77c2f6684159598b839ed808` over 740 files [`docs/provenance/SOURCE_TREE_HASH.json`]
- Community Edition only: enterprise code is removed or refused before the build (CE audit at acquire).

## Build

- Method: `source-rebuild-go`; base image `scratch + a chiselled Ubuntu 24.04 runtime (vendored/chisel/slices.txt, packages.tsv)`
- Compiled with `--network=none` from vendored source, dependencies and base images; BuildKit pinned to `moby/buildkit:v0.31.2` (shipped in `toolchain/`).
- SOURCE_DATE_EPOCH: `1794879874` [`SOURCE_DATE_EPOCH.txt`]
- Frozen inputs: [`docs/provenance/FREEZE_MANIFEST.json`], build recipe: [`docs/provenance/BUILD_RECIPE.json`].

## Output

- Published image: `readystack/caddy:v2.11.9-CE-ubuntu24.04-r3` — index digest `sha256:06aa670e1a64b773e994558f4690e612bf89f6a52071382bc99b8aeaf73be299` [`docs/provenance/LOCKED_DIGEST.txt`]
- linux/amd64 image digest `sha256:b36d6ee63c6c76857f17623c09923ec0b93a5b0a27a4b465eb51016b4552314b` — reproduced from this archive before publishing [`docs/provenance/REPRO_VERIFICATION.json`].

## Attestations and signature

- SBOM + SLSA provenance attached to the image: attached: an SPDX SBOM and SLSA v1 provenance [`docs/provenance/SCOUT_VERDICT.json`]
- cosign signature by the ReadyStack release key: verified before publishing. Check it: `cosign verify --key https://readystack.dev/keys/cosign.pub readystack/caddy:v2.11.9-CE-ubuntu24.04-r3`

## Pinned inputs, and why

- 36 pinned inputs, each with where it came from, why this one, how it was checked and when it was decided: `docs/PINS.md` [`docs/provenance/PINS.json`] (upstream source 2, base and build images 3, image digests 2, dependency versions readystack sets 10, pins in the kit's dockerfiles 12, readystack's build engine 5, dependencies upstream chose 2).

## Build trail

- How this release passed through the ReadyStack harness, record by record (every chain, drive, stage, review round, defect and check, each naming the record it is copied from): `ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.trail.json`, published beside this archive and signed with the release key (`ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.trail.json.sigstore`), and readable at https://readystack.dev/products/caddy/trail/v2.11.9-CE-ubuntu24.04-r3.

