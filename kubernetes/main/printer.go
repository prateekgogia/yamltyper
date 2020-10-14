package main

import (
	"fmt"

	"k8s.io/kube-openapi/pkg/util/proto"
)

type fieldsPrinterBuilder struct {
	Recursive bool
}

type fieldsPrinter interface {
	PrintFields(proto.Schema) error
}

// BuildFieldsPrinter builds the appropriate fieldsPrinter.
func (f fieldsPrinterBuilder) BuildFieldsPrinter(writer *Formatter) fieldsPrinter {
	if f.Recursive {
		return &recursiveFieldsPrinter{
			Writer: writer,
		}
	}

	return &regularFieldsPrinter{
		Writer: writer,
	}
}

// indentDesc is the level of indentation for descriptions.
const indentDesc = 2

// regularFieldsPrinter prints fields with their type and description.
type regularFieldsPrinter struct {
	Writer *Formatter
	Error  error
}

// var _ proto.SchemaVisitor = &regularFieldsPrinter{}
// var _ fieldsPrinter = &regularFieldsPrinter{}

// VisitArray prints a Array type. It is just a passthrough.
func (f *regularFieldsPrinter) VisitArray(a *proto.Array) {
	a.SubType.Accept(f)
}

// VisitKind prints a Kind type. It prints each key in the kind, with
// the type, the required flag, and the description.
func (f *regularFieldsPrinter) VisitKind(k *proto.Kind) {
	for _, key := range k.Keys() {
		v := k.Fields[key]
		required := ""
		if k.IsRequired(key) {
			required = " -required-"
		}
		fmt.Printf("PG VisitKind is %+v\n", key)
		if err := f.Writer.Write("%s\t<%s>%s", key, GetTypeName(v), required); err != nil {
			f.Error = err
			return
		}
		if err := f.Writer.Indent(indentDesc).WriteWrapped("%s", v.GetDescription()); err != nil {
			f.Error = err
			return
		}
		if err := f.Writer.Write(""); err != nil {
			f.Error = err
			return
		}
	}
}

// VisitMap prints a Map type. It is just a passthrough.
func (f *regularFieldsPrinter) VisitMap(m *proto.Map) {
	m.SubType.Accept(f)
}

// VisitPrimitive prints a Primitive type. It stops the recursion.
func (f *regularFieldsPrinter) VisitPrimitive(p *proto.Primitive) {
	// Nothing to do. Shouldn't really happen.
}

// VisitReference prints a Reference type. It is just a passthrough.
func (f *regularFieldsPrinter) VisitReference(r proto.Reference) {
	r.SubSchema().Accept(f)
}

// PrintFields will write the types from schema.
func (f *regularFieldsPrinter) PrintFields(schema proto.Schema) error {
	schema.Accept(f)
	return f.Error
}

const indentPerLevel = 3

// recursiveFieldsPrinter recursively prints all the fields for a given
// schema.
type recursiveFieldsPrinter struct {
	Writer *Formatter
	Error  error
}

var _ proto.SchemaVisitor = &recursiveFieldsPrinter{}
var _ fieldsPrinter = &recursiveFieldsPrinter{}
var visitedReferences = map[string]struct{}{}

// VisitArray is just a passthrough.
func (f *recursiveFieldsPrinter) VisitArray(a *proto.Array) {
	a.SubType.Accept(f)
}

// VisitKind prints all its fields with their type, and then recurses
// inside each of these (pre-order).
func (f *recursiveFieldsPrinter) VisitKind(k *proto.Kind) {
	for _, key := range k.Keys() {
		v := k.Fields[key]
		f.Writer.Write("%s\t<%s>", key, GetTypeName(v))
		subFields := &recursiveFieldsPrinter{
			Writer: f.Writer.Indent(indentPerLevel),
		}
		if err := subFields.PrintFields(v); err != nil {
			f.Error = err
			return
		}
	}
}

// VisitMap is just a passthrough.
func (f *recursiveFieldsPrinter) VisitMap(m *proto.Map) {
	m.SubType.Accept(f)
}

// VisitPrimitive does nothing, since it doesn't have sub-fields.
func (f *recursiveFieldsPrinter) VisitPrimitive(p *proto.Primitive) {
	// Nothing to do.
}

// VisitReference is just a passthrough.
func (f *recursiveFieldsPrinter) VisitReference(r proto.Reference) {
	if _, ok := visitedReferences[r.Reference()]; ok {
		return
	}
	visitedReferences[r.Reference()] = struct{}{}
	r.SubSchema().Accept(f)
	delete(visitedReferences, r.Reference())
}

// PrintFields will recursively print all the fields for the given
// schema.
func (f *recursiveFieldsPrinter) PrintFields(schema proto.Schema) error {
	schema.Accept(f)
	return f.Error
}
