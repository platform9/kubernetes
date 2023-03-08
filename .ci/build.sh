source ~/.bashrc
GO_VERSION=${GO_VERSION:-1.19}

configure_go() {
  if [ -n "${USE_SYSTEM_GO:-}" ] ; then
    echo "\$USE_SYSTEM_GO set, using system go instead of gimme"
    return 0
  else
    which gimme > /dev/null || (echo "error: missing required command 'gimme'" && exit 1)
    eval "$(GIMME_GO_VERSION=${GO_VERSION} gimme)"
  fi
  which go
  go version
}

configure_go

KUBE_BUILD_PLATFORMS=linux/amd64 make all WHAT=cmd/kubelet GOFLAGS=-v