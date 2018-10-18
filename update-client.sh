chmod +x ./vendor/k8s.io/code-generator/generate-groups.sh
./vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/hidevopsio/mioclient/pkg/client \
github.com/hidevopsio/mioclient/pkg/apis \
mio:v1alpha1