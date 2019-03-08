package builtin

import (
	. "github.com/gobigger/bigger"
)

func init() {

	Bigger.Router("*._doc_", Map{
		"name": "系统文档", "text": "系统文档",
		"action": func(ctx *Context){
            ctx.Data["cryptos"] = Bigger.Cryptos()
            ctx.Data["results"] = Bigger.Results()
            ctx.Data["routers"] = Bigger.Routers(ctx.Site)
			ctx.View("_doc_")
		},
	})

}


