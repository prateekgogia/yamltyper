package main

import (
	"os"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kube-openapi/pkg/util/proto"
	"k8s.io/kubectl/pkg/util/openapi"
	// "sigs.k8s.io/structured-merge-diff/v3/schema"
)

var (
	home = os.Getenv("HOME")
	// kubeconfig = filepath.Join(home, ".kube", "config")
	kubeconfig = "/Users/pgogia/workspace/golang/src/github.com/prateekgogia/aws-controllers-k8s/build/tmp-test-e7daf411/kubeconfig"
	// kubeconfig = "/Users/pgogia/.we/.kube/kubeconfig"
)

func main() {
	// if home := homeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	openAPIResources, err := getOpenAPIResources(clientSet)
	if err != nil {
		panic(err.Error())
	}
	b := fieldsPrinterBuilder{Recursive: true}
	f := &Formatter{Writer: os.Stdout, Wrap: 80}

	lists, err := clientSet.ServerPreferredResources()
	if err != nil {
		panic(err.Error())
	}
	for _, list := range lists {
		// resource.APIResources

		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err.Error())
		}
		for _, apiResource := range list.APIResources {

			gvk := schema.GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: apiResource.Kind}
			protoSchema := openAPIResources.LookupResource(gvk)
			// fmt.Printf("Resource is %+v GVK is %+v schema is %+v\n", apiResource, gvk, protoSchema)
			m := &modelPrinter{Name: "", Writer: f, Builder: b, GVK: gvk}
			protoSchema.Accept(m)
			// fmt.Printf("Schema is %T\n", protoSchema)
			// kind := protoSchema.(*proto.Kind)
			// fmt.Printf("kind spec fields are %+v\n", kind.Fields["spec"])
		}
	}

	// schema = openAPIResources.LookupResource(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"})
	// fmt.Printf("Schemes are %+v\n", schema)

	if err := getOpenAPIData(clientSet); err != nil {
		panic(err.Error())
	}
	return

}

// m := &modelPrinter{Name: name, Writer: writer, Builder: builder, GVK: gvk}
// schema.Accept(m)

func getOpenAPIData(clientSet *kubernetes.Clientset) error {
	// openapi.
	openAPIV2Document, err := clientSet.OpenAPISchema()
	if err != nil {
		return err
	}
	models, err := proto.NewOpenAPIData(openAPIV2Document)
	if err != nil {
		return err
	}
	_ = models
	// fmt.Println("Models are ", models.ListModels())
	// fmt.Println(models.LookupModel("com.example.stable.v1.CronTab"))
	return nil
}

func getOpenAPIResources(clientSet *kubernetes.Clientset) (openapi.Resources, error) {

	// fmt.Printf("getOpenAPIResources Groups are %+v\n", groups)
	// fmt.Printf("getOpenAPIResources resources are %+v\n", resources)
	return openapi.NewOpenAPIGetter(clientSet).Get()
}

INSERT INTO `client_from_emails` (`id`, `created_at`, `updated_at`, `client_id`, `from_email_id`, `active`) VALUES ('23', '2020-09-25 18:40:21', '2020-09-25 18:41:02', 'cid-77490e9bc59b3efe', 'help@wework.com', '1');