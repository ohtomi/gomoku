package server

import (
	"bytes"
	"html/template"
)

const htmlPanelTemplate = `
<div>
	<div>
		<button>&#8593;</button><button>&#8595;</button><button>X</button>
	</div>
	<div>
		<h4>request</h4>
		route: <input type="text" value="{{ .Request.Route }}"/>
		method: <input type="text" value="{{ .Request.Method }}"/>
		<h4>command</h4>
		env: <input type="text" value="{{ .Command.Env }}"/>
		path: <input type="text" value="{{ .Command.Env }}"/>
		args: <input type="text" value="{{ .Command.Env }}"/>
		<h4>response</h4>
		status: <input type="text" value="{{ .Response.Status }}"/>
		headers: <input type="text" value="{{ .Response.Headers }}"/>
		body: <input type="text" value="{{ .Response.Body }}"/>
		template: <input type="text" value="{{ .Response.Template }}"/>
		file: <input type="text" value="{{ .Response.File }}"/>
	</div>
</div>
`

func (c *Config) ToHtmlPanel() (string, error) {
	panel := ""

	for _, element := range *c {
		buf := &bytes.Buffer{}
		t, err := template.New("panel").Parse(htmlPanelTemplate)
		if err != nil {
			return "", err
		}
		if err := t.Execute(buf, element); err != nil {
			return "", err
		}
		panel += buf.String()
	}

	return panel, nil
}
