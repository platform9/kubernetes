echo "Dockerfile found. Building Docker image and running container..."
docker build -f .ci/Dockerfile --build-arg TEAMCITY_BUILD_BRANCH=$TEAMCITY_BUILD_BRANCH --build-arg BUILD_NUMBER=$BUILD_NUMBER -t kubelet-builder .
docker run --rm -v ~/.aws:/root/.aws kubelet-builder