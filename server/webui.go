package server

import (
	"bytes"
	"html/template"
)

const htmlStyleTemplate = `
<style>
body {
	background-color: #f6f6f6;
}
input[type="text"],
textarea {
	padding: 0.8em;
	outline: none;
	border: solid 1px lightgray;
	border-radius: 5px;
	font-size: 16px;
	width: 95%;
}

.config_panel {
	padding: 4px 12px 8px 12px;
	width: 600px;
	margin-bottom: 8px;
	border: solid 1px lightgray;
	border-radius: 5px;
}

.buttons_area {
	position: absolute;
	left: 536px;
}
.buttons_area button {
	margin-left: 4px;
	border: solid 1px lightgray;
	border-radius: 6px;
	width: 27px;
	height: 27px;
}
.buttons_area button.danger {
	color: #ffffff;
	background-color: #da4f49;
}

.forms_area {
}
.config_item div {
	padding-left: 8px;
}
.config_item h4 {
	margin: 0;
}
</style>
`

const htmlPanelTemplate = `
<div class="config_panel">
	<div class="buttons_area">
		<button>&#8593;</button><button>&#8595;</button><button class="danger">X</button>
	</div>
	<div class="forms_area">
		<div class="config_item">
			<h4>request</h4>
			<div>
				route<br/>
				<input type="text" value="{{ .Request.Route }}"/>
			</div>
			<div>
				method<br/>
				<input type="text" value="{{ .Request.Method }}"/>
			</div>
		</div>
		<div class="config_item">
			<h4>command</h4>
			<div>
				env<br/>
				<textarea>{{ .Command.Env }}</textarea>
			</div>
			<div>
				path<br/>
				<textarea>{{ .Command.Env }}</textarea>
			</div>
			<div>
				args<br/>
				<textarea>{{ .Command.Args }}</textarea>
			</div>
		</div>
		<div class="config_item">
			<h4>response</h4>
			<div>
				status<br/>
				<input type="text" value="{{ .Response.Status }}"/>
			</div>
			<div>
				headers<br/>
				<textarea>{{ .Response.Headers }}</textarea>
			</div>
			<div>
				body<br/>
				<input type="text" value="{{ .Response.Body }}"/>
			</div>
			<div>
				template<br/>
				<input type="text" value="{{ .Response.Template }}"/>
			</div>
			<div>
				file<br/>
				<input type="text" value="{{ .Response.File }}"/>
			</div>
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
