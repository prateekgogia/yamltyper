package frontend

func Run() error {
	return nil
}

/*

Two modes we need to support init and editing an existing YAML file in Kubernetes
 cmds supported
 > yamltyper init --kubernetes --file="filepath" --kubeconfig=""
 > yamltyper edit --kubernetes

 `yamltyper init` lets you create a new file
 -- To create new file we need to get all the fields for that mode (kubernetes)
 -- provide all the options to the user and in the end save/write the file

*/

type Extractor interface {
	NextField() (Provider, error)
}

type Provider interface {
	FieldName() (string, error)
	Options() (string, error)
	Type() (string, error)
	Validator(interface{}) error
}

// renderer to renders everything on the terminal
// have a prompt, which reads user output/actions and writes them to a writer
//
