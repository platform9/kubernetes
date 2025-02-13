FROM ubuntu:20.04
ARG TEAMCITY_BUILD_BRANCH
ARG BUILD_NUMBER

# Set the environment variables for use in the container
ENV TEAMCITY_BUILD_BRANCH=$TEAMCITY_BUILD_BRANCH
ENV BUILD_NUMBER=$BUILD_NUMBER
# Set environment variable to prevent prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    git \
    curl \
    python3 \
    python3-pip \
    bash \
    make \
    wget \
    ca-certificates \
    curl \
    gnupg \
    rsync \
    awscli \
    unzip \
    gcc-aarch64-linux-gnu \
    lsb-release \
    && apt-get clean

RUN pip3 install --upgrade awscli

# Install Go (Golang)
RUN curl -fsSL https://golang.org/dl/go1.22.4.linux-amd64.tar.gz -o go1.22.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz \
    && rm go1.22.4.linux-amd64.tar.gz

# Set the Go environment variables
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Install gimme (for Go version management)
RUN curl -sSL https://github.com/travis-ci/gimme/releases/download/v1.3.1/gimme-v1.3.1-linux-amd64 -o /usr/local/bin/gimme \
    && chmod +x /usr/local/bin/gimme

# Copy repo files into the container's working directory
WORKDIR /app
COPY ../ /app

# Ensure .ci/build.sh is executable
RUN chmod +x .ci/build.sh

# Expose any necessary ports (adjust as needed for your app)
EXPOSE 8080

# Set the default command to run your build script and then upload to S3
CMD /bin/bash -c "export KUBE_BUILD_PLATFORMS='KUBE_BUILD_PLATFORMS=linux/arm64' && ./.ci/build.sh && aws s3 cp --acl public-read _output/local/bin/linux/arm64/kubelet s3://emp-kubelet/$TEAMCITY_BUILD_BRANCH/arm64/kubelet-$TEAMCITY_BUILD_BRANCH-$BUILD_NUMBER"