FROM openshift/origin-release:golang-1.14 AS builder

ENV BASE=github.com/openshift-knative/serverless-operator
WORKDIR ${GOPATH}/src/${BASE}

COPY . .

ENV GOFLAGS="-mod=vendor"
RUN go build -o /tmp/operator ${BASE}/serving/ingress/cmd/controller

FROM openshift/origin-base
COPY --from=builder /tmp/operator /ko-app/operator

ENTRYPOINT ["/ko-app/operator"]
