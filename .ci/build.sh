source ~/.bashrc
GO_VERSION=${GO_VERSION:-1.19}
VERSION=$(git describe --tags --abbrev=0)
MAJOR_MINOR_PATCH=$(echo $VERSION | cut -d'-' -f1)
BUILD_NUMBER=${BUILD_NUMBER:-000}

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

make WHAT=cmd/kubelet KUBE_GIT_VERSION=$MAJOR_MINOR_PATCH-emp.$BUILD_NUMBER KUBE_GIT_COMMIT=$(git rev-parse HEAD) KUBE_GIT_TREE_STATE="clean" GOFLAGS=-v