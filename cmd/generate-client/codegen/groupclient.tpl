// Code generated by generate-client. DO NOT EDIT.

package {{ .Package.Name }}

import (
    "github.com/onosproject/helm-go/pkg/kubernetes/resource"
)

type {{ .Types.Interface }} interface {
    {{- range $name, $resource := .Resources }}
    {{ $resource.Client.Types.Interface }}
    {{- end }}
}

func New{{ .Types.Interface }}(resources resource.Client, filter resource.Filter) {{ .Types.Interface }} {
	return &{{ .Types.Struct }}{
		Client: resources,
		{{- range $name, $resource := .Resources }}
        {{ $resource.Client.Types.Interface }}: New{{ $resource.Client.Types.Interface }}(resources, filter),
        {{- end }}
	}
}

type {{ .Types.Struct }} struct {
	resource.Client
    {{- range $name, $resource := .Resources }}
    {{ $resource.Client.Types.Interface }}
    {{- end }}
}
