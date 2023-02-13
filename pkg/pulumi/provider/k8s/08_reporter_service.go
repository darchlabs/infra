package k8s

import (
	"github.com/darchlabs/infra/internal/env"
	// batchv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/batch/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ReporterSynchronizers(ctx *pulumi.Context, env env.Env) error {
	// create synchronizers reporter configmap
	err := reporterSynchronizersConfigmap(ctx, env)
	if err != nil {
		return err
	}

	// create synchronizers reporter cronjob

	return nil
}

// configmap
func reporterSynchronizersConfigmap(ctx *pulumi.Context, env env.Env) error {
	// define synchronizers config map
	name := "reporter-synchronizers-config"
	args := &corev1.ConfigMapArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(name),
		},
		Data: pulumi.StringMap{
			"DATABASE_URL": pulumi.String(env.ReportSynchronizersDatabaseURL),
			"SERVICE_URL":  pulumi.String(env.ReportSynchronizersServiceURL),
			"SERVICE_TYPE": pulumi.String(env.ReportSynchronizersServiceType),
		},
	}

	// create synchronizers report
	_, err := corev1.NewConfigMap(ctx, name, args)
	if err != nil {
		return err
	}

	return nil
}

// cronjob
func reporterSynchronizersCronjob(ctx *pulumi.Context) error {
	// // define sychronizers cronhjob
	// args := &batchv1.CronJobArgs{
	// 	Metadata: &metav1.ObjectMetaArgs{
	// 		Name: pulumi.String("reporter-synchronizers"),
	// 	},
	// 	Spec: pulumi.String("* * * * *"),
	// }

	// // create synchronizers
	// _, err := batchv1.NewCronJob(ctx, "reporter-synchronizers", args)
	// if err != nil {
	// 	return err
	// }

	return nil
}
