package main

import (
	"fmt"

	"k8s.io/kube-openapi/pkg/util/proto"
)

// typeName finds the name of a schema
type typeName struct {
	Name string
}

var _ proto.SchemaVisitor = &typeName{}

// VisitArray adds the [] prefix and recurses.
func (t *typeName) VisitArray(a *proto.Array) {
	s := &typeName{}
	a.SubType.Accept(s)
	t.Name = fmt.Sprintf("[]%s", s.Name)
}

// VisitKind just returns "Object".
func (t *typeName) VisitKind(k *proto.Kind) {
	t.Name = "Object"
}

// VisitMap adds the map[string] prefix and recurses.
func (t *typeName) VisitMap(m *proto.Map) {
	s := &typeName{}
	m.SubType.Accept(s)
	t.Name = fmt.Sprintf("map[string]%s", s.Name)
}

// VisitPrimitive returns the name of the primitive.
func (t *typeName) VisitPrimitive(p *proto.Primitive) {
	t.Name = p.Type
}

// VisitReference is just a passthrough.
func (t *typeName) VisitReference(r proto.Reference) {
	r.SubSchema().Accept(t)
}

// GetTypeName returns the type of a schema.
func GetTypeName(schema proto.Schema) string {
	t := &typeName{}
	schema.Accept(t)
	return t.Name
}
