package k8s

import (
	"fmt"

	"github.com/darchlabs/infra/internal/env"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	networkingv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/networking/v1"
)

// TODO(ca): should implements support to use domain with SSL

func Ingress(ctx *pulumi.Context, env env.Env, namespace string, url string) error {
	// // create ingress configuration
	// err := ingressRelease(ctx, namespace)
	// if err != nil {
	// 	return err
	// }

	// // create ingress routes definitions
	// err = ingressServices(ctx, namespace, url)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func ingressRelease(ctx *pulumi.Context, namespace string) error {
	// define ingress args
	args := &helmv3.ReleaseArgs{
		Chart:     pulumi.String("nginx-ingress"),
		Namespace: pulumi.String(namespace),
		RepositoryOpts: &helmv3.RepositoryOptsArgs{
			Repo: pulumi.String("https://helm.nginx.com/stable"),
		},
		SkipCrds: pulumi.Bool(true),
		Values: pulumi.Map{
			"controller": pulumi.Map{
				"service": pulumi.Map{
					"type": pulumi.String("LoadBalancer"),
				},
			},
		},
	}

	fmt.Println("BEFORE!!!")
	fmt.Println("BEFORE!!!")
	fmt.Println("BEFORE!!!")
	fmt.Println("BEFORE!!!")

	// create new helm release
	_, err := helmv3.NewRelease(ctx, "dl-ingress", args)
	if err != nil {
		fmt.Println("KAKAKAKAKA")
		fmt.Println("KAKAKAKAKA")
		fmt.Println("KAKAKAKAKA")
		fmt.Println("KAKAKAKAKA")
		fmt.Println("KAKAKAKAKA")
		fmt.Println("KAKAKAKAKA", err.Error())
		return err
	}

	fmt.Println("AFTER!!!")
	fmt.Println("AFTER!!!")
	fmt.Println("AFTER!!!")
	fmt.Println("AFTER!!!")

	return nil
}

func ingressServices(ctx *pulumi.Context, namespace string, url string) error {
	// define meta ingress args
	metaArgs := &metav1.ObjectMetaArgs{
		Name:      pulumi.String("dl-ingress"),
		Namespace: pulumi.String(namespace),
		Annotations: pulumi.StringMap{
			"nginx.ingress.kubernetes.io/ssl-redirect":   pulumi.String("false"),
			"nginx.ingress.kubernetes.io/rewrite-target": pulumi.String("/"),
		},
	}

	// define rules ingress args
	rulesArgs := networkingv1.IngressRuleArgs{
		Host: pulumi.String(url),
		Http: &networkingv1.HTTPIngressRuleValueArgs{
			Paths: networkingv1.HTTPIngressPathArray{
				networkingv1.HTTPIngressPathArgs{
					Path:     pulumi.String("/synchronizers"),
					PathType: pulumi.String("Prefix"),
					Backend: networkingv1.IngressBackendArgs{
						Service: &networkingv1.IngressServiceBackendArgs{
							Name: pulumi.String("synchronizers-v2"),
							Port: networkingv1.ServiceBackendPortArgs{
								Name: pulumi.String("synch-http"),
								// Number: pulumi.Int(5555),
							},
						},
					},
				},
			},
		},
	}

	// define ingress args
	args := &networkingv1.IngressArgs{
		Metadata: metaArgs,
		Spec: networkingv1.IngressSpecArgs{
			Rules: networkingv1.IngressRuleArray{
				rulesArgs,
			},
		},
	}

	// create ingress
	_, err := networkingv1.NewIngress(ctx, "dl-ingress", args)
	if err != nil {
		return err
	}

	return nil
}
