package goweb

import (

    "github.com/julienschmidt/httprouter"
)

func RoutesConfig() *httprouter.Router {

    r := httprouter.New()

    return r
}
