package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
)

func main() {

	// t := reflect.TypeOf(sampleDS{})
	// v := reflect.Indirect(reflect.ValueOf(&sampleDS{}))
	// options := []string{}
	// for i := 0; i < t.NumField(); i++ {
	// 	reflect.TypeOf(t.Field(i)).Kind().String()
	// 	fmt.Println(t.Field(i).Name)
	// 	field := t.Field(i).Tag.Get(`yaml`)
	// 	options = append(options, field)
	// }
	// fmt.Println(options, len(options))
	options := getAllFields("", sampleDS{})
	fmt.Println()
	// return
	v := reflect.Indirect(reflect.ValueOf(&sampleDS{}))

	for i, _ := range options {
		prompt := &survey.Select{
			Options: options[i:],
		}
		fieldSelected := ""
		survey.AskOne(prompt, &fieldSelected)

		answer := askUser(fieldSelected)
		v.Field(i).Set(reflect.ValueOf(answer))
	}

	fmt.Printf("Value is %+v\n", v)
}

func getAllFields(parentStruct string, obj interface{}) []string {
	fields := make([]string, 0)
	ifv := reflect.ValueOf(obj)
	ift := reflect.TypeOf(obj)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			parent := ift.Name()
			if parentStruct != "" {
				parent = parentStruct + "." + parent
			}
			fields = append(fields, getAllFields(parent, v.Interface())...)
		default:
			field := ift.Field(i).Tag.Get(`yaml`)
			if parentStruct != "" {
				field = parentStruct + "." + field
			}
			fields = append(fields, field)
		}
	}
	return fields
}

func askUser(field string) string {
	result := ""
	// perform the questions
	err := survey.AskOne(
		&survey.Input{
			Message: fmt.Sprintf("%v : ", field),
		},
		&result)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return result
}

var ds = &sampleDS{}

func init() {
	data, err := ioutil.ReadFile("sample.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(data, ds); err != nil {
		log.Fatalln(err)
	}
}

type sampleDS struct {
	// Test
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   string `yaml:"metadata"`
}

type Test struct {
	Test2
	Value int `yaml:"value"`
}
type Test2 struct {
	Value int `yaml:"value"`
}
