package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
)

func IsJobFinished(j *batchv1.Job) (bool, batchv1.JobConditionType) {
	for _, c := range j.Status.Conditions {
		if c.Type == batchv1.JobComplete && c.Status == corev1.ConditionTrue {
			return true, c.Type
		}

		if c.Type == batchv1.JobFailed {
			return true, c.Type
		}
	}

	return false, ""
}

func GetJobLogs(iface kubernetes.Interface, job *batchv1.Job) (string, error) {
	pods, err := getPodsForJob(iface, job)
	if err != nil {
		return "", err
	}
	if len(pods) == 0 {
		return "", fmt.Errorf("job %s:not find pods", job.Name)
	}

	return getPodLogs(iface, pods[0])
}

func getPodLogs(iface kubernetes.Interface, pod corev1.Pod) (string, error) {
	req := iface.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &corev1.PodLogOptions{})
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", fmt.Errorf(":%s: error in opening stream", pod.Name)
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", fmt.Errorf(":%s: error in copy information from podLogs to buf", pod.Name)
	}
	str := buf.String()
	return str, nil
}

func getPodsForJob(iface kubernetes.Interface, job *batchv1.Job) ([]corev1.Pod, error) {

	r, err := labels.NewRequirement("controller-uid", selection.Equals, []string{string(job.GetUID())})
	if err != nil {
		return nil, err
	}

	selector := labels.NewSelector().Add(*r)
	lists, err := iface.CoreV1().Pods(job.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		return nil, err
	}
	return lists.Items, err
}

func ContainerAddMounter(container *corev1.Container, mounter corev1.VolumeMount) {
	if container.VolumeMounts == nil {
		container.VolumeMounts = []corev1.VolumeMount{mounter}
	} else {
		container.VolumeMounts = append(container.VolumeMounts, mounter)
	}
}

func PodSpecAddVolume(spec *corev1.PodSpec, volume corev1.Volume) {
	if spec.Volumes == nil {
		spec.Volumes = []corev1.Volume{volume}
	} else {
		spec.Volumes = append(spec.Volumes, volume)
	}
}
