<!-- readystack:kit-readme begin (written at packaging; the kit's own README follows the end marker) -->

# Build kit for `readystack/caddy:v2.11.9-CE-ubuntu24.04-r3`

This archive rebuilds the linux/amd64 image `readystack/caddy:v2.11.9-CE-ubuntu24.04-r3` from its own contents, with the network off, and checks the result against the image digest ReadyStack published, `sha256:b36d6ee63c6c76857f17623c09923ec0b93a5b0a27a4b465eb51016b4552314b` (also in `TARGET_DIGEST.txt`). It holds everything the build needs, not the image itself.

## Rebuild it

You need linux/amd64 Docker with `buildx` 0.30 or later, `jq` and `tar` (docs/REBUILD.md lists the rest, if any, and the hardware ReadyStack used).

```
tar -xpzf ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.tar.gz    # -p keeps each file's mode
cd v2.11.9-CE-ubuntu24.04-r3
./BUILD_OFFLINE.sh            # builds with the network off, from this archive alone
./VERIFY_OFFLINE.sh           # compares the rebuilt digest with the published one, then checks the image
```

`./VERIFY_OFFLINE.sh --online` also checks that the Docker Hub tag holds the digest you rebuilt.

File modes are build input: a Dockerfile `COPY` carries them into the image. Unpacked without `-p`, your umask changes them (664 becomes 644 under umask 022). `FILE_MODES.tsv` records the mode of every file and directory as packed; `BUILD_OFFLINE.sh` puts back any that differ before it builds, and `./BUILD_OFFLINE.sh --restore-modes` does only that.

## Check ReadyStack's signature (optional)

ReadyStack signs this archive's `rebuild.json` with its release key once the finished archive has passed the release gate, and records the signature in the public Sigstore transparency log. The signature is made after the archive is sealed, so it is published beside it rather than inside it. Download it and the public key next to this directory, not into it: a file added inside the kit changes the tree hash that `rebuild.json` records, and `validate` then refuses the kit.

```
curl -fsSLo ../cosign.pub https://readystack.dev/keys/cosign.pub
curl -fsSLo ../ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.rebuild.json.sigstore https://readystack.dev/api/kit/caddy/v2.11.9-CE-ubuntu24.04-r3/files/ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.rebuild.json.sigstore
python3 docs/rebuild-bundle/rebuild_bundle.py validate . --key ../cosign.pub --signature ../ReadyStack-caddy-v2.11.9-CE-ubuntu24.04-r3.rebuild.json.sigstore
```

`validate` needs Python 3 and cosign. Without building anything it checks the signature over `rebuild.json`, every file's checksum, the tree hash over every file's path, mode and content, each base image layout's digest and the builder archive. `verify`, with the same arguments, also rebuilds with the builder's network egress blocked and exits 0 only when the rebuilt digest is the published one (docs/REBUILD.md).

## What else is here

- `docs/REBUILD.md`: the full rebuild guide: requirements, what `rebuild.json` records, the verifier's exit codes, and SLSA provenance on your own build platform
- `docs/VERIFY.md`: checking the published image, its signature and its attestations without rebuilding
- `docs/PROVENANCE.md`: where every input came from
- `docs/PINS.md`: every pinned input and why it was chosen
- `docs/RELEASE_NOTES.md`: what this release changed, and its known issues
- `docs/admin-guide/ADMIN_HANDBOOK.md`: running the product

<!-- readystack:kit-readme end -->

# ReadyStack Caddy build kit

Source-rebuild-go kit for **Caddy v2.11.9** (Apache-2.0) — the cleanest
single-binary validator of the KitForge BUILD track.

## What this builds

One static Go binary (`caddy`) on a chiselled Ubuntu 24.04 runtime (`FROM scratch` with only the files
Canonical's chisel cut from a few Ubuntu packages: no apt, dpkg, bash or curl), running **non-root** with
`cap_net_bind_service` so it can bind :80/:443 without root or `--privileged`.

| File | Role |
|---|---|
| `main.go` | xcaddy-pattern entrypoint: depends on `caddy/v2@v2.11.9` + `modules/standard`. caddy reads its version from this dependency, so the binary self-reports `v2.11.9` with no `-ldflags` hack. |
| `Caddyfile` | Default config: a host-reachable `/health` 200 endpoint + a placeholder site. Replace with your own. |
| `Dockerfile.vendor` | Online: resolves deps, raises `grpc x/crypto x/net x/sys x/text go-jose/v4 klauspost/compress` to exact minimum versions (never lowers one), checks that caddy itself resolved to `v2.11.9`, `go mod vendor`; cuts the runtime rootfs with Canonical's chisel (pinned release and chisel-releases commit, checksummed) from Ubuntu's own .debs and generates its dpkg status. Exports `vendored/{caddy-src,chisel}`. |
| `Dockerfile.offline.pinned` | `docker buildx build --network=none`: recompiles caddy and rs-probe (the HEALTHCHECK helper, `rs-probe.go`) from the vendored tree, lays the chiselled rootfs out with the caddy user (uid 999), gives `/usr/bin/caddy` cap_net_bind_service with the cut's own setcap, and builds the image `FROM scratch` from it. No network. |
| `Dockerfile`, `Dockerfile.offline` | Copies of `Dockerfile.offline.pinned`. |
| `rs-probe.go`, `rs-probe.mod` | The HEALTHCHECK helper: `rs-probe URL` exits 0 on a 2xx/3xx answer (curl -f), `-body` prints the body; standard library only. |

## Build (online, quick)

```
docker build -f Dockerfile.vendor --target export --output type=local,dest=./vendored .
docker build -t caddy-slim .
docker run -d --name caddy -p 80:80 -p 443:443 -v caddy_data:/data -v caddy_config:/config caddy-slim
curl -fsS http://127.0.0.1/health   # -> OK (from the host)
```

## Offline rebuild (the customer promise)

From the root of the extracted ReadyStack archive, with the network off:

```
./BUILD_OFFLINE.sh      # loads the pinned BuildKit, builds Dockerfile.offline.pinned against the vendored base layouts in bases/
./VERIFY_OFFLINE.sh     # compares the rebuilt digest with TARGET_DIGEST.txt and boots the rebuilt image
```

`docs/REBUILD.md` in the archive describes both scripts. A bare `docker build` still resolves base images from a
registry and does not reproduce the published digest.

## Non-root / Grade A notes

- Runs as the `caddy` user (uid != 0). No root, no `--privileged`, no
  runtime `--cap-add`. The capability is baked onto the file with
  `setcap cap_net_bind_service=+ep /usr/bin/caddy`.
- Static binary (`CGO_ENABLED=0`) built with the go1.26.10 toolchain
  (`golang:1.26.10-bookworm`, pinned by digest) → no CGO system-lib CVE
  surface; stdlib CVEs cleared by the toolchain.
- Every build stage checks that the compiled binary reports `v2.11.9` and
  stops otherwise.
- OS packages: only what the chisel slices need (base-files, base-passwd, ca-certificates, tzdata, dash, part of
  coreutils, libcap2-bin's getcap/setcap and their libraries; `vendored/chisel/packages.tsv`), cut from the current
  Ubuntu archive at vendor time; `Dockerfile.vendor` fails if a package floor of the kit is not met.

## Verify

```
docker run --rm --entrypoint caddy <img> version     # v2.11.9
docker run --rm --entrypoint getcap <img> /usr/bin/caddy  # cap_net_bind_service=ep
docker inspect -f '{{.Config.User}}' <img>             # caddy (uid 999; the image has no id binary)
```
