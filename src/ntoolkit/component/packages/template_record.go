package packages

import (
	"ntoolkit/component"
	"github.com/go-yaml/yaml"
	"strings"
)

func convertYamlToTemplate(rawYaml string) (template *component.ObjectTemplate, err error) {
	var record component.ObjectTemplate
	rawYaml = stripTabs(rawYaml)
	if err := yaml.Unmarshal([]byte(rawYaml), &record); err != nil {
		return nil, err
	}
	forceObjectType(&record)
	return &record, nil
}

func stripTabs(raw string) string {
	lines := strings.Split(raw, "\n")
	for i := 0; i < len(lines); i++ {
		runes := []rune(lines[i])
		out := make([]rune, 0)
		rest := false
		for j := 0; j < len(runes); j++ {
			if rest || runes[j] == ' ' {
				out = append(out, runes[j])
			} else if runes[j] == '\t' {
				out = append(out, ' ')
				out = append(out, ' ')
			} else {
				out = append(out, runes[j])
				rest = true
			}
		}
		lines[i] = string(out)
	}
	return strings.Join(lines, "\n")
}

func forceObjectType(template *component.ObjectTemplate) {
	for i := 0; i < len(template.Components); i++ {
		if template.Components[i].Data != nil {
			forceComponentDataType(&template.Components[i])
		}
	}
	for i := 0; i < len(template.Objects); i++ {
		forceObjectType(&template.Objects[i])
	}
}

func forceComponentDataType(component *component.ComponentTemplate) {
	defer (func() {
		if r := recover(); r != nil {
			return
		}
	})()
	if component.Data != nil {
		old := component.Data.(map[interface{}]interface{})
		values := make(map[string]interface{})
		for k, v := range old {
			values[k.(string)] = forceArbitraryDataType(v)
		}
		component.Data = values
	}
}

func forceArbitraryDataType(data interface{}) interface{} {
	defer (func() {
		if r := recover(); r != nil {
			return
		}
	})()
	mapped, ok := data.(map[interface{}]interface{})
	if ok {
		values := make(map[string]interface{})
		for k, v := range mapped {
			values[k.(string)] = forceArbitraryDataType(v)
		}
		return values
	}
	return data
}