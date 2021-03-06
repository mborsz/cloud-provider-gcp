package main

import (
	"flag"
	"fmt"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/pkg/apis/clientauthentication"
	clientauthv1alpha1 "k8s.io/client-go/pkg/apis/clientauthentication/v1alpha1"
)

var (
	audience = flag.String("audience", "", "Audience field of for the VM ID token. Must be a URI.")

	scheme       = runtime.NewScheme()
	codecs       = serializer.NewCodecFactory(scheme)
	groupVersion = schema.GroupVersion{
		Group:   "client.authentication.k8s.io",
		Version: "v1alpha1",
	}
)

func init() {
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	clientauthv1alpha1.AddToScheme(scheme)
	clientauthentication.AddToScheme(scheme)
}

func main() {
	// Override the default in glog. There's verbosity flag for suppressing output.
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *audience == "" {
		glog.Exit("--audience must be set")
	}

	token, err := metadata.Get(fmt.Sprintf("instance/service-accounts/default/identity?audience=%s&format=full", *audience))
	if err != nil {
		glog.Exit(err)
	}
	token = "vmid-" + token

	if err := writeResponse(token); err != nil {
		glog.Exit(err)
	}
}

func writeResponse(token string) error {
	resp := &clientauthentication.ExecCredential{
		Status: &clientauthentication.ExecCredentialStatus{
			// Make Kubelet poke us every hour, we'll cache the cert for longer.
			ExpirationTimestamp: &metav1.Time{time.Now().Add(time.Hour)},
			Token:               token,
		},
	}
	data, err := runtime.Encode(codecs.LegacyCodec(groupVersion), resp)
	if err != nil {
		return err
	}
	fmt.Print(string(data))
	return nil
}
