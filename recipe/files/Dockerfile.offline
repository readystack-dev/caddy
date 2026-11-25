# =====================================================================
# HERMETIC OFFLINE BUILD, chiselled runtime (canonical, attestation-producing): built by `docker buildx build
#   --network=none --attest type=provenance,mode=max --attest type=sbom`. Every input is vendored or pinned by
#   digest: the Go builder and ubuntu:24.04 (the rootfs assembly stage only), vendored/caddy-src (the fully
#   vendored module tree, GOPROXY=off) and vendored/chisel (the runtime rootfs Canonical's chisel cut from
#   Ubuntu's own .debs: rootfs.tar, debs/, packages.tsv, slices.txt). The image is FROM scratch with that rootfs
#   alone: no apt, dpkg, bash or curl. Dockerfile and Dockerfile.offline are copies of this file.
# =====================================================================
FROM golang:1.26.11-bookworm@sha256:dc40a13b61304647ed7e5d581d246b86413dfc29d63c6d018310feb93cd98fc0 AS build
WORKDIR /src
COPY vendored/caddy-src/ .
ENV CGO_ENABLED=0 GOTOOLCHAIN=local GOFLAGS="-mod=vendor -buildvcs=false" GOPROXY=off
RUN go build -trimpath -ldflags='-s -w' -o /caddy .
RUN /caddy version | tee /tmp/caddy.version && grep -Eq '^v2\.11\.9( |$)' /tmp/caddy.version || \
    { echo "built caddy does not report v2.11.9" >&2; exit 1; }

# ---- rs-probe: the HEALTHCHECK helper (standard library only, from the kit's rs-probe.go, the engine's
# generators/gen_dockerfile.py RS_PROBE_GO): the image has no curl, whose library closure alone costs ~21 medium findings
FROM golang:1.26.11-bookworm@sha256:dc40a13b61304647ed7e5d581d246b86413dfc29d63c6d018310feb93cd98fc0 AS build_rs_probe
WORKDIR /probe
COPY rs-probe.go main.go
COPY rs-probe.mod go.mod
ENV CGO_ENABLED=0 GOTOOLCHAIN=local GOFLAGS="-buildvcs=false" GOPROXY=off SOURCE_DATE_EPOCH=1794879874
RUN go build -trimpath -ldflags='-s -w' -o /rs-probe .

# ---- rootfs: the chiselled Ubuntu 24.04 runtime (vendored/chisel, cut by Dockerfile.vendor's chiselcut stage from
# Ubuntu's own .debs), the service user, the binary and its configuration, laid out on the pinned base. Only /rootfs
# reaches the image. setcap runs from the cut itself (libcap2-bin_bins) under chroot: caddy binds :80/:443 as the
# non-root caddy user also where the runtime keeps net.ipv4.ip_unprivileged_port_start at 1024; the file
# capability travels as the security.capability xattr.
FROM ubuntu:24.04@sha256:2198e4c65076e9cb0b3d695219df667c0e3b0072da64ca3c2648aca17dcecaf5 AS rootfs
ARG SOURCE_DATE_EPOCH=1794879874
ADD vendored/chisel/rootfs.tar /rootfs/
RUN printf 'caddy:x:999:999::/var/lib/caddy:/bin/sh\n' >> /rootfs/etc/passwd && \
    printf 'caddy:x:999:\n' >> /rootfs/etc/group && \
    mkdir -p /rootfs/data /rootfs/config /rootfs/etc/caddy /rootfs/var/lib/caddy && chown -R 999:999 /rootfs/data /rootfs/config /rootfs/etc/caddy /rootfs/var/lib/caddy && chmod 750 /rootfs/var/lib/caddy
COPY --from=build /caddy /rootfs/usr/bin/caddy
COPY --from=build_rs_probe /rs-probe /rootfs/usr/local/bin/rs-probe
COPY Caddyfile /rootfs/etc/caddy/Caddyfile
RUN chmod 755 /rootfs/usr/bin/caddy /rootfs/usr/local/bin/rs-probe && \
    chown 999:999 /rootfs/etc/caddy/Caddyfile && \
    chroot /rootfs /usr/sbin/setcap cap_net_bind_service=+ep /usr/bin/caddy && \
    { chroot /rootfs /usr/sbin/getcap /usr/bin/caddy | grep -q 'cap_net_bind_service=ep' || \
      { echo "setcap did not give /usr/bin/caddy cap_net_bind_service" >&2; exit 1; }; }
# The scanners' inputs must be in the image: /etc/os-release and a dpkg status that names every package of the cut
# (chisel writes none; Dockerfile.vendor generates it). The wall-clock and host-derived paths go.
RUN --mount=type=bind,source=vendored/chisel,target=/chisel set -eu; \
    [ -e /rootfs/etc/os-release ] || { echo "no /etc/os-release: base-files_release-info is missing" >&2; exit 1; }; \
    [ "$(grep -c '^Package: ' /rootfs/var/lib/dpkg/status)" = "$(wc -l < /chisel/packages.tsv)" ] || \
      { echo "/var/lib/dpkg/status does not list every package of vendored/chisel/packages.tsv" >&2; exit 1; }; \
    rm -rf /rootfs/var/log/dpkg.log /rootfs/var/log/alternatives.log /rootfs/var/log/apt /rootfs/var/cache/ldconfig/aux-cache /rootfs/var/cache/fontconfig

FROM scratch
ARG SOURCE_DATE_EPOCH=1794879874
COPY --from=rootfs /rootfs/ /
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin XDG_DATA_HOME=/data XDG_CONFIG_HOME=/config
HEALTHCHECK --interval=30s --timeout=10s --start-period=10s \
  CMD ["/usr/local/bin/rs-probe", "http://127.0.0.1:80/health"]
USER caddy
EXPOSE 80 443 2019
ENTRYPOINT ["caddy","run","--config","/etc/caddy/Caddyfile","--adapter","caddyfile"]
# >>> readystack OCI labels: strategies/caddy.json image_labels plus the standard OCI labels,
# the set a generated kit carries (gen_dockerfile.label_lines). Written by scripts/kit_migrations.py;
# scripts/retarget.py re-renders it with every retarget. Never edit it by hand.
LABEL org.opencontainers.image.licenses="Apache-2.0"
LABEL org.opencontainers.image.title="Caddy"
LABEL org.opencontainers.image.description="Caddy serves as a modern web server that automatically manages HTTPS certificates for your applications. It simplifies secure traffic routing without requiring complex configuration or external dependencies."
LABEL org.opencontainers.image.source="https://github.com/caddyserver/caddy"
LABEL org.opencontainers.image.version="v2.11.9-CE-ubuntu24.04-r3"
LABEL org.opencontainers.image.revision="0d98caea58f794f88e709b1bef2880996a471bca"
# <<< readystack OCI labels
