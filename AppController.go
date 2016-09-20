package core

import (

    "net/http"
    "html/template"
    "path"
    "encoding/json"
)

type AppController struct{}

func (c *AppController) RenderHtml(rw http.ResponseWriter, file string, data interface{}) {

    fp := path.Join("view", file)
    tmpl, err := template.ParseFiles(fp)

    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(rw, data); err != nil {

        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }
}

func (c *AppController) RenderJson(rw http.ResponseWriter, data interface{}) {

    if data != nil {

        js, err := json.Marshal(data)
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }

        rw.Header().Set("Content-Type", "application/json")
        rw.Write(js)
        return
    }

    http.Error(rw, "data null", http.StatusInternalServerError)
}
