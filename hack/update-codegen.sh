#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -x

#SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
#CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
CODEGEN_PKG="${GOPATH}/src/github.com/kubernetes/code-generator"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
#bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \

# for unit
#  ../upmio/upm-pkg/pkg/client/unit \
#bash "${CODEGEN_PKG}"/generate-groups.sh \
#  all \
#  upm-pkg/pkg/client/unit \
#  upm-pkg/pkg/apis \
#  unit:v1alpha1 \
#  --output-base "$(dirname "${BASH_SOURCE[0]}")/../.." \
#  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# for unitset
#bash "${CODEGEN_PKG}"/generate-groups.sh \
#  all \
#  upm-pkg/pkg/client/unitset \
#  upm-pkg/pkg/apis \
#  unitset:v1alpha1 \
#  --output-base "$(dirname "${BASH_SOURCE[0]}")/../.." \
#  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# for unitset
${CODEGEN_PKG}/generate-groups.sh all github.com/upmio/upm-pkg/pkg/client/unitset github.com/upmio/upm-pkg/pkg/apis unitset:v1alpha1 --go-header-file ./boilerplate.go.txt

# for unit
#${CODEGEN_PKG}/generate-groups.sh all github.com/upmio/upm-pkg/pkg/client/unit github.com/upmio/upm-pkg/pkg/apis unit:v1alpha1 --go-header-file ./boilerplate.go.txt

# for test
#${CODEGEN_PKG}/generate-groups.sh all  github.com/upmio/upm-pkg/pkg/client/unitset-test  github.com/upmio/upm-pkg/pkg/apis unitset:v1alpha1 --go-header-file ./boilerplate.go.txt