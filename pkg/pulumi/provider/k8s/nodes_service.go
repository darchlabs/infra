package k8s

import (
	"fmt"
	"strconv"

	"github.com/darchlabs/infra/internal/env"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Nodes(ctx *pulumi.Context, env env.Env, namespace string) error {
	// create nodes configmap
	err := nodesConfigMap(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR nodesConfigMap")
		return err
	}

	// create nodes deployment
	err = nodesDeployment(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR nodesDeployment")
		return err
	}

	// create nodes service
	err = nodesService(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR nodesService")
		return err
	}

	return nil
}

// configmap
func nodesConfigMap(ctx *pulumi.Context, env env.Env, namespace string) error {
	// define nodes configmap
	name := "nodes-config"
	args := &corev1.ConfigMapArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String(name),
			Namespace: pulumi.String(namespace),
		},
		Data: pulumi.StringMap{
			"NODES_CHAIN":                pulumi.String(env.NodesChain),
			"NODES_API_SERVER_PORT":      pulumi.String(env.NodesApiServerPort),
			"NODES_RPC_PORT":             pulumi.String(env.NodesRpcPort),
			"NODES_BASE_CHAIN_DATA_PATH": pulumi.String(env.NodesBaseChainDataPath),
			"NODES_NODE_URL":             pulumi.String(env.NodesNodeURL),
			"NODES_BLOCK_NUMBER":         pulumi.String(env.NodesBlockNumber),
		},
	}

	// create nodes configmap
	_, err := corev1.NewConfigMap(ctx, name, args)
	if err != nil {
		return err
	}

	return nil
}

// deployment
func nodesDeployment(ctx *pulumi.Context, env env.Env, namespace string) error {
	// parse port
	port, err := strconv.Atoi(env.NodesApiServerPort)
	if err != nil {
		return err
	}

	// define deployment args
	args := &appsv1.DeploymentArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("nodes"),
			Namespace: pulumi.String(namespace),
		},
		Spec: &appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"app": pulumi.String("nodes-deployment"),
				},
			},
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: pulumi.StringMap{
						"app": pulumi.String("nodes-deployment"),
					},
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						&corev1.ContainerArgs{
							Name:  pulumi.String("nodes"),
							Image: pulumi.String("darchlabs/nodes-ethereum:0.0.2"),
							EnvFrom: corev1.EnvFromSourceArray{
								&corev1.EnvFromSourceArgs{
									ConfigMapRef: &corev1.ConfigMapEnvSourceArgs{
										Name: pulumi.String("nodes-config"),
									},
								},
							},
							Ports: corev1.ContainerPortArray{
								&corev1.ContainerPortArgs{
									ContainerPort: pulumi.Int(port),
								},
							},
						},
					},
				},
			},
		},
	}

	// create nodes deployment
	_, err = appsv1.NewDeployment(ctx, "nodes-deployment", args)
	if err != nil {
		return err
	}

	return nil
}

// service
func nodesService(ctx *pulumi.Context, env env.Env, namespace string) error {
	// parse port
	port, err := strconv.Atoi(env.NodesApiServerPort)
	if err != nil {
		return err
	}

	// define nodes service
	args := &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("nodes"),
			Namespace: pulumi.String(namespace),
		},
		Spec: &corev1.ServiceSpecArgs{
			Ports: &corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Port: pulumi.Int(port),
					Name: pulumi.String("http"),
				},
			},
			Selector: pulumi.StringMap{
				"app": pulumi.String("nodes-deployment"),
			},
		},
	}

	// create nodes service
	_, err = corev1.NewService(ctx, "nodes-service", args)
	if err != nil {
		return err
	}

	return nil
}
