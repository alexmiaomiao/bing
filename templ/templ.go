package templ

import (
	"text/template"
)

const t1 =
`{{.Res.Word}}{{if .Res.Pronunciation}}		[美 {{.Res.Pronunciation.AmE}}]	[英 {{.Res.Pronunciation.BrE}}]{{end}}
-------------------------------------------------------------------------------------------
{{range .Res.Defs}}{{.Pos}}		{{.Def}}
{{end}}-------------------------------------------------------------------------------------------
{{if .Verbose}}{{range .Res.Sams}}{{.Eng}}
{{.Chn}}

{{end}}{{else}}使用 bing -v {{.Res.Word}} 可以查看例句
{{end}}
`

var CommandLine = template.Must(template.New("commandline").Parse(t1))
