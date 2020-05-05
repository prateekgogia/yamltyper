package kubernetes

import (
	"encoding/json"
	"io"
	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// If config provided fetch from server
// if save flag set save to JSON file,
// how do I know whats the version I am suppose to save with?
// If config is not provided read from pre-existing JSON file
type resources struct {
	kubeconfig string
	cacheJSON  bool
	rw         io.ReadWriter
}

func (r *resources) getServerPreferredResources() ([]*metav1.APIResourceList, error) {

	resourcesList := make([]*metav1.APIResourceList, 0)
	var err error
	if r.kubeconfig != "" {
		resourcesList, err = r.fromServer()
		if err != nil {
			return nil, err
		}
		if r.cacheJSON {
			if err := r.writeToFile(resourcesList); err != nil {
				return nil, err
			}
		}
		return resourcesList, nil
	}
	return r.fromCachedFile()
}

func (r *resources) writeToFile(list []*metav1.APIResourceList) error {

	content, err := json.Marshal(list)
	if err != nil {
		return err
	}
	if _, err := r.rw.Write(content); err != nil {
		return err
	}
	return nil
}

func (r *resources) fromCachedFile() ([]*metav1.APIResourceList, error) {

	contentBytes, err := ioutil.ReadAll(r.rw)
	if err != nil {
		return nil, err
	}
	list := make([]*metav1.APIResourceList, 0)
	if err := json.Unmarshal(contentBytes, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *resources) fromServer() ([]*metav1.APIResourceList, error) {

	clientSet, err := getKubeClientSet(r.kubeconfig)
	if err != nil {
		return nil, err
	}
	lists, err := clientSet.ServerPreferredResources()
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func getKubeClientSet(kubeconfig string) (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}
