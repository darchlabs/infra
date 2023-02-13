package k8s

import (
	"strconv"

	"github.com/darchlabs/infra/internal/env"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Jobs(ctx *pulumi.Context, env env.Env) error {
	// create jobs configmap
	err := jobsConfigMap(ctx, env)
	if err != nil {
		return err
	}

	// create jobs deployment
	err = jobsDeployment(ctx, env)
	if err != nil {
		return err
	}

	// create jobs service
	err = jobsService(ctx, env)
	if err != nil {
		return err
	}

	return nil
}

// configmap
func jobsConfigMap(ctx *pulumi.Context, env env.Env) error {
	// define jobs configmap
	name := "jobs-config"
	args := &corev1.ConfigMapArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(name),
		},
		Data: pulumi.StringMap{
			"DATABASE_FILEPATH": pulumi.String(env.JobsDatabaseFilepath),
			"PORT":              pulumi.String(env.JobsPort),
			"NODE_URL":          pulumi.String(env.JobsNodeURL),
			"PRIVATE_KEY":       pulumi.String(env.JobsPrivateKey),
		},
	}

	// create jobs configmap
	_, err := corev1.NewConfigMap(ctx, name, args)
	if err != nil {
		return err
	}

	return nil
}

// deployment
func jobsDeployment(ctx *pulumi.Context, env env.Env) error {
	// parse port
	port, err := strconv.Atoi(env.JobsPort)
	if err != nil {
		return err
	}

	// define deployment args
	args := &appsv1.DeploymentArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String("jobs-api"),
		},
		Spec: &appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"app": pulumi.String("jobs-deployment"),
				},
			},
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: pulumi.StringMap{
						"app": pulumi.String("jobs-deployment"),
					},
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						&corev1.ContainerArgs{
							Name:  pulumi.String("jobs"),
							Image: pulumi.String("darchlabs/jobs:1.1.0"),
							EnvFrom: corev1.EnvFromSourceArray{
								&corev1.EnvFromSourceArgs{
									ConfigMapRef: &corev1.ConfigMapEnvSourceArgs{
										Name: pulumi.String("jobs-config"),
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

	// create jobs deployment
	_, err = appsv1.NewDeployment(ctx, "deployment", args)
	if err != nil {
		return err
	}

	return nil
}

// service
func jobsService(ctx *pulumi.Context, env env.Env) error {
	// parse port
	port, err := strconv.Atoi(env.JobsPort)
	if err != nil {
		return err
	}

	// define jobs service
	args := &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String("jobs"),
		},
		Spec: &corev1.ServiceSpecArgs{
			Ports: &corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Port: pulumi.Int(port),
					Name: pulumi.String("http"),
				},
			},
			Selector: pulumi.StringMap{
				"app": pulumi.String("jobs-deployment"),
			},
		},
	}

	// create jobs service
	_, err = corev1.NewService(ctx, "service", args)
	if err != nil {
		return err
	}

	return nil
}
