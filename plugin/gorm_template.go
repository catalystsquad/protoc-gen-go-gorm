package plugin

import "text/template"

var messageTemplate = template.Must(template.New("message").Funcs(templateFuncs).Parse(`
type {{ .Model.Name }} struct {
	{{- range .Model.Fields }}
    {{ if .ShouldGenerateBelongsToIdField }}
    {{ .Options.GetBelongsTo.Foreignkey }} *string {{ emptyTag }}
    {{ end }}
    {{ .Comments -}}
    {{ .GoName }} {{ .ModelType }} {{ .Tag -}}
	{{ end }}
}

func (m *{{ .Model.Name }}) TableName() string {
	return "{{ .Model.TableName }}"
}

func (m *{{ .Model.Name }}) ToProto() (theProto *{{.GoIdent.GoName}}, err error) {
	if m == nil {
		return
	}
	theProto = &{{.GoIdent.GoName}}{}
	{{ range .Model.Fields }}
    {{ if .IsTimestamp }}
    {{ if eq .Desc.Kind 9 }}
	if m.{{ .GoName }} != nil {
		theProto.{{ .GoName }} = m.{{ .GoName }}.Format(time.RFC3339Nano)
	}
    {{ else }}
    if m.{{ .GoName }} != nil {
		theProto.{{ .GoName }} = timestamppb.New(*m.{{ .GoName }})
	}
    {{ end }}
    {{ else if or .IsStructPb .IsJsonb }}
	if m.{{ .GoName }} != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(m.{{ .GoName }}); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theProto.{{ .GoName }}); err != nil {
			return
		}
	}
    {{ else if and .Options .Options.TimeFormatOverride }}
	if m.{{ .GoName }} != nil {
	{{- if .IsOptional }}
		theProto.{{ .GoName }} = lo.ToPtr(m.{{ .GoName }}.UTC().Format("{{ .TimeFormat }}"))
	{{- else }}
		theProto.{{ .GoName }} = m.{{ .GoName }}.UTC().Format("{{ .TimeFormat }}")
	{{- end }}
	}
    {{ else if and .IsMessage (eq .IsRepeated false) }}
	if theProto.{{ .GoName }}, err = m.{{ .GoName }}.ToProto(); err != nil {
		return
	}
    {{ else if and .IsMessage .IsRepeated }}
	if len(m.{{ .GoName }}) > 0 {
		theProto.{{ .GoName }} = []*{{ .Message.Desc.Name }}{}
        for _, item := range m.{{ .GoName }} {
			var {{ .GoName }}Proto *{{ .Message.GoIdent.GoName }}
			if {{ .GoName }}Proto, err = item.ToProto(); err != nil {
				return 
			} else {
				theProto.{{ .GoName }} = append(theProto.{{ .GoName }}, {{ .GoName }}Proto)
			}	
		}
	}
    {{ else if and .Enum ( eq .IsRepeated false) }}
	{{ if .Options.EnumAsString }}
	theProto.{{ .GoName }} = {{ .Enum.GoIdent.GoName }}({{ .Enum.GoIdent.GoName }}_value[m.{{ .GoName }}])
    {{ else }}
	theProto.{{ .GoName }} = {{ .Enum.GoIdent.GoName }}(m.{{ .GoName }})
    {{ end }}
	{{ else if and .Enum .IsRepeated }}
	{{ if .Options.EnumAsString }}
	if len(m.{{ .GoName }}) > 0 {
		theProto.{{ .GoName }} = []{{ .Enum.GoIdent.GoName }}{}
		for _, val := range m.{{ .GoName }} {
			theProto.{{ .GoName }} = append(theProto.{{ .GoName }}, {{ .Enum.GoIdent.GoName }}({{ .Enum.GoIdent.GoName }}_value[val]))
		}
	}
    {{ else }}
	if len(m.{{ .GoName }}) > 0 {
		theProto.{{ .GoName }} = []{{ .Enum.GoIdent.GoName }}{}
		for _, val := range m.{{ .GoName }} {
			theProto.{{ .GoName }} = append(theProto.{{ .GoName }}, {{ .Enum.GoIdent.GoName }}(val))
		}
	}
    {{ end }}
    {{ else }}
    theProto.{{ .GoName }} = m.{{ .GoName }}
    {{ end }}
	{{ end }}
	return
}

func (p *{{.GoIdent.GoName}}) GetProtoId() *string {
	return p.Id
}

func (p *{{.GoIdent.GoName}}) SetProtoId(id string) {
	p.Id = lo.ToPtr(id)
}

func (p *{{.GoIdent.GoName}}) ToModel() (theModel *{{ .Model.Name }}, err error) {
	if p == nil {
		return
	}
	theModel = &{{ .Model.Name }}{}
	{{ range .Model.Fields }}
    {{ if .IsTimestamp }}
	{{ if eq .Desc.Kind 9 }}
	if p.{{ .GoName }} != "" {
		var timestamp time.Time
		if timestamp, err = time.Parse(time.RFC3339Nano, p.{{ .GoName }}); err != nil {
			return
		}
		theModel.{{ .GoName }} = &timestamp
	}
    {{ else }}
    if p.{{ .GoName }} != nil {
		theModel.{{ .GoName }} = lo.ToPtr(p.{{ .GoName }}.AsTime())
	}
    {{ end }}
    {{ else if and .Options .Options.TimeFormatOverride }}
	{{- if .IsOptional }}
	if p.{{ .GoName }} != nil {
		var date time.Time
		if date, err = time.Parse("{{ .TimeFormat }}", *p.{{ .GoName }}); err != nil {
			return
		}
	{{- else }}
	if p.{{ .GoName }} != "" {
		var date time.Time
		if date, err = time.Parse("{{ .TimeFormat }}", p.{{ .GoName }}); err != nil {
			return
		}
	{{- end }}
		dateUTC := date.UTC()
		theModel.{{ .GoName }} = &dateUTC
	}
    {{ else if or .IsStructPb .IsJsonb }}
	if p.{{ .GoName }} != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(p.{{ .GoName }}); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theModel.{{ .GoName }}); err != nil {
			return
		}
	}
	{{ else if and .IsMessage (eq .IsRepeated false) }}
	if theModel.{{ .GoName }}, err = p.{{ .GoName }}.ToModel(); err != nil {
		return
	}
    {{ if .Options.BelongsTo }}
    // if the object is present, the object's id overrides the existing id field value
    if p.{{ .GoName }} != nil {
	  {{ if .Options.BelongsTo.Foreignkey -}}
      theModel.{{ .Options.BelongsTo.Foreignkey }} = p.{{ .GoName }}.Id
      {{ else }}
      theModel.{{ .GoName }}Id = p.{{ .GoName }}.Id
      {{ end -}}
	}
    {{ end }}
	{{ else if and .IsMessage .IsRepeated }}
	if len(p.{{ .GoName }}) > 0 {
		theModel.{{ .GoName }} = {{ .ModelType }}{}
        for _, item := range p.{{ .GoName }} {
			var {{ .GoName }}Model {{ .ModelSingularType }}
			if {{ .GoName }}Model, err = item.ToModel(); err != nil {
				return 
			} else {
				theModel.{{ .GoName }} = append(theModel.{{ .GoName }}, {{ .GoName }}Model)
			}	
		}
	}
    {{ else if and .Enum (eq .IsRepeated false) }}
	{{ if .Options.EnumAsString }}
	theModel.{{ .GoName }} = p.{{ .GoName }}.String()
	{{ else }}
	theModel.{{ .GoName }} = int(p.{{ .GoName }})
	{{ end }}
	{{ else if and .Enum .IsRepeated }}
	{{ if .Options.EnumAsString }}
	if len(p.{{ .GoName }}) > 0 {
		theModel.{{ .GoName }} = pq.StringArray{}
		for _, val := range p.{{ .GoName }} {
			theModel.{{ .GoName }} = append(theModel.{{ .GoName }}, val.String())
		}
	}
    {{ else }}
	if len(p.{{ .GoName }}) > 0 {
		theModel.{{ .GoName }} = pq.Int32Array{}
		for _, val := range p.{{ .GoName }} {
			theModel.{{ .GoName }} = append(theModel.{{ .GoName }}, int32(val))
		}
	}
	{{ end }}
    {{ else }}
    theModel.{{ .GoName }} = p.{{ .GoName }}
    {{ end }}
	{{ end }}
	return
}
`))
