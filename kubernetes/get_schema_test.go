package kubernetes

import (
	"io"
	"testing"

	"k8s.io/kube-openapi/pkg/util/proto"
)

func Test_resources_getSchema(t *testing.T) {
	type fields struct {
		kubeconfig string
		cacheJSON  bool
		rw         io.ReadWriter
	}
	tests := []struct {
		name    string
		fields  fields
		want    []proto.Schema
		wantErr bool
	}{
		{"Test getting schemas from API server don't write to file",
			fields{
				kubeconfig: "/Users/pgogia/.kube/config",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &resources{
				kubeconfig: tt.fields.kubeconfig,
				cacheJSON:  tt.fields.cacheJSON,
				rw:         tt.fields.rw,
			}
			_, err := r.getSchema()
			if (err != nil) != tt.wantErr {
				t.Errorf("resources.getSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
