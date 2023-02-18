package k8s

import (
	"fmt"

	"github.com/darchlabs/infra/pkg/project"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create namespace for k8s cluster
func Namespace(ctx *pulumi.Context, projectName string, environment project.Environment) error {
	name := fmt.Sprintf("dl-%s-ns", projectName)
	args := &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(name),
		},
	}

	_, err := corev1.NewNamespace(ctx, name, args)
	if err != nil {
		return err
	}

	return nil
}
