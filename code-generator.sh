#!/usr/bin/env bash

cd vendor/k8s.io/code-generator/

chmod 777 generate-groups.sh

./generate-groups.sh all \
github.com/hidevopsio/mioclient/pkg/client \
github.com/hidevopsio/mioclient/pkg/apis \
mio:v1alpha1