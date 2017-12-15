package server

import (
	"bytes"
	"html/template"
)

const htmlStyleTemplate = `
<style>
body {
	background-color: #efefef;
}

.config_panel {
	padding: 4px 12px 8px 12px;
	width: 600px;
}

.buttons_area {
	position: absolute;
	left: 536px;
}
.buttons_area button {
	width: 27px;
	height: 27px;
	margin-left: 4px;
}

.forms_area {
}
.config_item div {
	padding-left: 8px;
}
.config_item h4 {
	margin: 0;
}
.config_item input[type="text"] {
	font-size: medium;
}
</style>
`

const htmlPanelTemplate = `
<div class="config_panel">
	<div class="buttons_area">
		<button>&#8593;</button><button>&#8595;</button><button>X</button>
	</div>
	<div class="forms_area">
		<div class="config_item">
			<h4>request</h4>
			<div>route: <input type="text" value="{{ .Request.Route }}"/></div>
			<div>method: <input type="text" value="{{ .Request.Method }}"/></div>
		</div>
		<div class="config_item">
			<h4>command</h4>
			<div>env: <input type="text" value="{{ .Command.Env }}"/></div>
			<div>path: <input type="text" value="{{ .Command.Env }}"/></div>
			<div>args: <input type="text" value="{{ .Command.Env }}"/></div>
		</div>
		<div class="config_item">
			<h4>response</h4>
			<div>status: <input type="text" value="{{ .Response.Status }}"/></div>
			<div>headers: <input type="text" value="{{ .Response.Headers }}"/></div>
			<div>body: <input type="text" value="{{ .Response.Body }}"/></div>
			<div>template: <input type="text" value="{{ .Response.Template }}"/></div>
			<div>file: <input type="text" value="{{ .Response.File }}"/></div>
		</div>
	</div>
</div>
`

func (c *Config) ToHtmlPanel() (string, error) {
	panel := ""

	panel += htmlStyleTemplate
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
