# yamltyper
Yamltyper is an extensible CLI tool which lets you write YAML files easily.

## How does it work?
yamltyper prompts with the fields supported for a particular yaml type. 

## commands

`yamltyper generate -p kubernetes`
`yamltyper generate -p kubernetes --validate`

`yamltyper generate -p openapi`
`yamltyper generate -p openapi --validate`

## How to add more projects



If someone wants to create a deployment in Kubernetes. They will be able to run to as `yamltyper generate -p kubernetes`. 

The goal for this tool is to provide user with a set of fields and let the user provide the values for these YAML fields. These fields are what a particular project supports. This initial version of the tool supports Kubernetes.

## How does it work
User can run this tool and command like `yamltyper create --project openapi`, this will help user select what they want to add to their YAML file. User may wish to skip a field if its optional. This tool will provide fields as option to the user to select the field. The result of running this tool will be a YAML file with desired fields.

This tool will provide validation for fields added and maybe add validation later.
