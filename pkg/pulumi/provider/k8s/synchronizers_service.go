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

func Synchonizers(ctx *pulumi.Context, env env.Env, namespace string) error {
	// create synchronizers configmap
	err := synchronizersConfigMap(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR synchronizersConfigMap")
		return err
	}

	// create synchronizers deployment
	err = synchronizersDeployment(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR synchronizersDeployment")
		return err
	}

	// create synchronizers service
	err = synchronizersService(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR synchronizersService")
		return err
	}

	return nil
}

// configmap
func synchronizersConfigMap(ctx *pulumi.Context, env env.Env, namespace string) error {
	// define synchronizers configmap
	name := "synchronizers-v2-config"
	args := &corev1.ConfigMapArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String(name),
			Namespace: pulumi.String(namespace),
		},
		Data: pulumi.StringMap{
			"NODE_URL":          pulumi.String(env.SynchronizersNodeURL),
			"INTERVAL_SECONDS":  pulumi.String(env.SynchronizersIntervalSeconds),
			"DATABASE_FILEPATH": pulumi.String(env.SynchronizersDatabaseFilepath),
			"PORT":              pulumi.String(env.SynchronizersPort),
		},
	}

	// create synchronizers configmap
	_, err := corev1.NewConfigMap(ctx, name, args)
	if err != nil {
		return err
	}

	return nil
}

// deployment
func synchronizersDeployment(ctx *pulumi.Context, env env.Env, namespace string) error {
	// parse port
	port, err := strconv.Atoi(env.SynchronizersPort)
	if err != nil {
		return err
	}

	// define deployment args
	args := &appsv1.DeploymentArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("synchronizers-v2-api"),
			Namespace: pulumi.String(namespace),
		},
		Spec: &appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"app": pulumi.String("synchronizers-v2-deployment"),
				},
			},
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: pulumi.StringMap{
						"app": pulumi.String("synchronizers-v2-deployment"),
					},
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						&corev1.ContainerArgs{
							Name:  pulumi.String("synchronizers-v2"),
							Image: pulumi.String("darchlabs/synchronizer-v2:1.2.0"),
							EnvFrom: corev1.EnvFromSourceArray{
								&corev1.EnvFromSourceArgs{
									ConfigMapRef: &corev1.ConfigMapEnvSourceArgs{
										Name: pulumi.String("synchronizers-v2-config"),
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

	// create synchronizers deployment
	_, err = appsv1.NewDeployment(ctx, "synchronizers-v2-deployment", args)
	if err != nil {
		return err
	}

	return nil
}

// service
func synchronizersService(ctx *pulumi.Context, env env.Env, namespace string) error {
	// parse port
	port, err := strconv.Atoi(env.SynchronizersPort)
	if err != nil {
		return err
	}

	// define synchronizers service
	args := &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("synchronizers-v2"),
			Namespace: pulumi.String(namespace),
		},
		Spec: &corev1.ServiceSpecArgs{
			Ports: &corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Port: pulumi.Int(port),
					Name: pulumi.String("synch-http"),
				},
			},
			Selector: pulumi.StringMap{
				"app": pulumi.String("synchronizers-v2-deployment"),
			},
		},
	}

	// create synchronizers service
	_, err = corev1.NewService(ctx, "synchronizers-v2-service", args)
	if err != nil {
		return err
	}

	return nil
}
