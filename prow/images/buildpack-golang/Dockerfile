# Basic golang buildpack

FROM eu.gcr.io/kyma-project/test-infra/bootstrap:v20210730-c3e07ed7

# Commit details

ARG commit
ENV IMAGE_COMMIT=$commit
LABEL io.kyma-project.test-infra.commit=$commit

# Versions

ENV ARCH amd64
ENV GO_VERSION 1.16.5
ENV DEP_RELEASE_TAG v0.5.4
ENV KUBEBUILDER_VERSION 2.3.2
ENV KUSTOMIZE_VERSION 3.8.4
ENV GOLANGCI_LINT_VERSION v1.40.1


# apt and install additional packages
RUN apt-get update && apt-get install -y --no-install-recommends \
        rsync=3.1.2-1+deb9u2 \
        procps=2:3.3.12-3+deb9u1 \
        && apt-get clean

# Install Go and Dep
ENV GOPATH /workspace/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN wget -q https://dl.google.com/go/go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    tar xzf go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    rm go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    mv go /usr/local && \
    mkdir -p ${GOPATH}/bin && \
    mkdir -p ${GOPATH}/src && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install kubebuilder
ENV PATH /usr/local/kubebuilder/bin:$PATH
ENV PATH /home/prow/go/bin:$PATH

RUN curl -L -O "https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH}.tar.gz" && \
    tar -zxvf kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH}.tar.gz && \
    rm kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH}.tar.gz && \
    mv kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH} kubebuilder && \
    mv kubebuilder /usr/local/ && \
    curl -L -O "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v${KUSTOMIZE_VERSION}/kustomize_v${KUSTOMIZE_VERSION}_linux_${ARCH}.tar.gz" && \
    tar -zxvf kustomize_v${KUSTOMIZE_VERSION}_linux_${ARCH}.tar.gz && \
    rm kustomize_v${KUSTOMIZE_VERSION}_linux_${ARCH}.tar.gz && \
    mv kustomize /usr/local/bin/kustomize

# Install golang toolbox
RUN go get golang.org/x/tools/cmd/goimports \
    golang.org/x/lint/golint \
    github.com/ericchiang/pup \
    github.com/kisielk/errcheck && \
    GO111MODULE="on" go get github.com/vektra/mockery && \
    GO111MODULE="off" go get -u github.com/maxbrunsfeld/counterfeiter && \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ${GOPATH}/bin ${GOLANGCI_LINT_VERSION}

COPY ./license-puller.sh /license-puller.sh
ENV LICENSE_PULLER_PATH=/license-puller.sh

# Prow Tools
COPY --from=eu.gcr.io/kyma-project/test-infra/prow-tools:v20210401-294e46e5 /prow-tools /prow-tools
# for better access to prow-tools
ENV PATH=$PATH:/prow-tools
