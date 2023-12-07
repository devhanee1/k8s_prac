package job

import (
	"context"
	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type JellybeanzJobHandler struct {
	ctx    context.Context
	client client.Client
	Logger logr.Logger
}

func NewJellybeanzJobHandler(ctx context.Context, client client.Client) *JellybeanzJobHandler {
	jellyBeanzJobHandler := &JellybeanzJobHandler{
		ctx:    ctx,
		client: client,
		Logger: logr.Logger{},
	}
	return jellyBeanzJobHandler
}

func (r *JellybeanzJobHandler) DeployJob() error {

	jobResource := batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "jellybeanz-job",
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Name:  "jellybeanz-job",
						Image: "hello-world:latest",
					}},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}
	return r.client.Create(r.ctx, &jobResource)
}
