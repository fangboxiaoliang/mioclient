#!/usr/bin/env bash

cd vendor/k8s.io/code-generator/

chmod o+x generate-groups.sh

./generate-groups.sh all \
hidevops.io/mioclient/pkg/client \
hidevops.io/mioclient/pkg/apis \
mio:v1alpha1