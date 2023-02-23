package k8s

import (
	"fmt"
	"strconv"

	"github.com/darchlabs/infra/internal/env"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
)

func webapp(ctx *pulumi.Context, env env.Env, namespace string) error {
	// create webapp configmap
	err := webappConfigmap(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR webappConfigmap")
		return err
	}

	// create webapp deployment
	err = webappDeployment(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR webappDeployment")
		return err
	}

	// create webapp service
	err = webappService(ctx, env, namespace)
	if err != nil {
		fmt.Println("ERROR webappService")
		return err
	}

	return nil
}

// configmap
func webappConfigmap(ctx *pulumi.Context, env env.Env, namespace string) error {
	// define webapp confirmap
	name := "webapp-config"
	args := &corev1.ConfigMapArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String(name),
			Namespace: pulumi.String(namespace),
		},
		Data: pulumi.StringMap{
			"REDIS_URL":             pulumi.String(env.WebappRedisURL),
			"SYNCHORONIZER_API_URL": pulumi.String(env.WebappSynchronizersURL),
			"NODE_API_URL":          pulumi.String(env.WebappNodesURL),
			"JOB_API_URL":           pulumi.String(env.WebappJobsURL),
		},
	}

	// create webaapp configmap
	_, err := corev1.NewConfigMap(ctx, name, args)
	if err != nil {
		return err
	}

	return nil
}

// deployment
func webappDeployment(ctx *pulumi.Context, env env.Env, namespace string) error {
	// parse port
	port, err := strconv.Atoi(env.WebappPort)
	if err != nil {
		return err
	}

	// define deployment args
	args := &appsv1.DeploymentArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("webapp"),
			Namespace: pulumi.String(namespace),
		},
		Spec: &appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"app": pulumi.String("webapp-deployment"),
				},
			},
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: pulumi.StringMap{
						"app": pulumi.String("webapp-deployment"),
					},
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						&corev1.ContainerArgs{
							Name:  pulumi.String("webapp"),
							Image: pulumi.String("darchlabs/webapp:0.1.0"),
							EnvFrom: corev1.EnvFromSourceArray{
								&corev1.EnvFromSourceArgs{
									ConfigMapRef: &corev1.ConfigMapEnvSourceArgs{
										Name: pulumi.String("webapp-config"),
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

	// create webapp deployment
	_, err = appsv1.NewDeployment(ctx, "webapp-deployment", args)
	if err != nil {
		return err
	}

	return nil
}

// service
func webappService(ctx *pulumi.Context, env env.Env, namespace string) error {
	// parse port
	port, err := strconv.Atoi(env.WebappPort)
	if err != nil {
		return err
	}

	// define webapp service
	args := &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String("webapp"),
			Namespace: pulumi.String(namespace),
		},
		Spec: &corev1.ServiceSpecArgs{
			Ports: &corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Port: pulumi.Int(port),
					Name: pulumi.String("webapp-http"),
				},
			},
			Selector: pulumi.StringMap{
				"app": pulumi.String("webapp-deployment"),
			},
		},
	}

	// create webapp service
	_, err = corev1.NewService(ctx, "webapp-service", args)
	if err != nil {
		return err
	}

	return nil
}
