package main

import (
	"html/template"
	"net/http"
)

// Handlers HTTPハンドラを集めた型
type Handlers struct {
	ab *AccountBook
}

// NewHandlers Handlersを作成する
func NewHandlers(ab *AccountBook) *Handlers {
	return &Handlers{ab: ab}
}

/*
構造体の体型は以下のようになっていると思われる
hs *Handlers {
	ab *AccountBook {
		db *sql.DB
	}
}
*/

// ListHandlerで仕様するテンプレート
var listTmpl = template.Must(template.New("list").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8"/>
		<title>家計簿</title>
	</head>
	<body>
		<h1>家計簿</h1>
		<h2>最新{{len .}}件</h2>
		{{- if . -}}
		<table border="1">
			<tr><th>品目</th><th>値段</th></tr>
			{{- range .}}
			<tr><td>{{.Category}}</td><td>{{.Price}}円</td></tr>
			{{- end}}
		</table>
		{{- else}}
			データがありません
		{{- end}}
	</body>
</html>
`))

// ListHandler 最新の入力データを表示するハンドラ
func (hs *Handlers) ListHandler(w http.ResponseWriter, r *http.Request) {
	items, err := hs.ab.GetItems(10)
	if err != nil {
		// TODO: http.Errorを使って、InternalServerErrorでエラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: 取得したitemsをテンプレートに埋め込む
	err = listTmpl.Execute(w, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
