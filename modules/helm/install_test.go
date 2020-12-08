// +build kubeall helm

// NOTE: we have build tags to differentiate kubernetes tests from non-kubernetes tests, and further differentiate helm
// tests. This is done because minikube is heavy and can interfere with docker related tests in terratest. Similarly,
// helm can overload the minikube system and thus interfere with the other kubernetes tests. To avoid overloading the
// system, we run the kubernetes tests and helm tests separately from the others.

package helm

import (
	"crypto/tls"
	"fmt"
	"strings"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
)

// Test that we can install a remote chart (e.g stable/chartmuseum)
func TestRemoteChartInstall(t *testing.T) {
	t.Parallel()

	namespaceName := fmt.Sprintf(
		"%s-%s",
		strings.ToLower(t.Name()),
		strings.ToLower(random.UniqueId()),
	)

	// Use default kubectl options to create a new namespace for this test, and then update the namespace for kubectl
	kubectlOptions := k8s.NewKubectlOptions("", "", namespaceName)

	defer k8s.DeleteNamespace(t, kubectlOptions, namespaceName)
	k8s.CreateNamespace(t, kubectlOptions, namespaceName)

	// Override service type to node port
	options := &Options{
		KubectlOptions: kubectlOptions,
		SetValues: map[string]string{
			"service.type": "NodePort",
		},
	}

	// Add the stable repo under a random name so as not to touch existing repo configs
	uniqueName := strings.ToLower(fmt.Sprintf("terratest-%s", random.UniqueId()))
	defer RemoveRepo(t, options, uniqueName)
	AddRepo(t, options, uniqueName, "https://charts.helm.sh/stable")
	helmChart := fmt.Sprintf("%s/chartmuseum", uniqueName)

	// Generate a unique release name so we can defer the delete before installing
	releaseName := fmt.Sprintf(
		"chartmuseum-%s",
		strings.ToLower(random.UniqueId()),
	)
	defer Delete(t, options, releaseName, true)

	// Test if helm.install will return an error if the chart version is incorrect
	options.Version = "notValidVersion.0.0.0"
	require.Error(t, InstallE(t, options, helmChart, releaseName))

	// Fix chart version and retry install
	options.Version = "2.3.0"
	// Test that passing extra arguments doesn't error, by changing default timeout
	options.ExtraArgs = []string{"--timeout", "5m1s"}
	require.NoError(t, InstallE(t, options, helmChart, releaseName))
	waitForRemoteChartPods(t, kubectlOptions, releaseName, 1)

	// Verify service is accessible. Wait for it to become available and then hit the endpoint.
	// Service name is RELEASE_NAME-CHART_NAME
	serviceName := fmt.Sprintf("%s-chartmuseum", releaseName)
	k8s.WaitUntilServiceAvailable(t, kubectlOptions, serviceName, 10, 1*time.Second)
	service := k8s.GetService(t, kubectlOptions, serviceName)
	endpoint := k8s.GetServiceEndpoint(t, kubectlOptions, service, 8080)

	// Setup a TLS configuration to submit with the helper, a blank struct is acceptable
	tlsConfig := tls.Config{}

	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		fmt.Sprintf("http://%s", endpoint),
		&tlsConfig,
		30,
		10*time.Second,
		func(statusCode int, body string) bool {
			return statusCode == 200
		},
	)
}

func waitForRemoteChartPods(t *testing.T, kubectlOptions *k8s.KubectlOptions, releaseName string, podCount int) {
	// Get pod and wait for it to be avaialable
	// To get the pod, we need to filter it using the labels that the helm chart creates
	filters := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=chartmuseum,release=%s", releaseName),
	}
	k8s.WaitUntilNumPodsCreated(t, kubectlOptions, filters, podCount, 30, 10*time.Second)
	pods := k8s.ListPods(t, kubectlOptions, filters)
	for _, pod := range pods {
		k8s.WaitUntilPodAvailable(t, kubectlOptions, pod.Name, 30, 10*time.Second)
	}
}
