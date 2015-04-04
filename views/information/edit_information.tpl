{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Edit Information - SYSU Train System</title>
{{end}}

{{define "body"}}
	{{template "information_script" .}}
	{{template "information_body" .}}
{{end}}

{{template "information/script_edit_information.tpl" .}}

{{template "information/body_edit_information.tpl" .}}
