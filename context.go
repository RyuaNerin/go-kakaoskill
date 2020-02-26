package skill

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Context struct {
	w       http.ResponseWriter
	Reqeust *http.Request
	Payload *SkillPayload

	sent bool
}

func (ctx *Context) WriteSimpleText(str string) {
	res := SkillResponse{
		Version: "2.0",
		Template: SkillTemplate{
			Outputs: []Component{
				Component{
					SimpleText: &SimpleText{
						Text: str,
					},
				},
			},
		},
	}

	ctx.WriteResponse(&res)
}

func (ctx *Context) WriteSimpleImage(imageUrl string, altText string) {
	res := SkillResponse{
		Version: "2.0",
		Template: SkillTemplate{
			Outputs: []Component{
				Component{
					SimpleImage: &SimpleImage{
						ImageUrl: imageUrl,
						AltText:  altText,
					},
				},
			},
		},
	}

	ctx.WriteResponse(&res)
}

func (ctx *Context) WriteResponse(res *SkillResponse) {
	ctx.sent = true
	ctx.w.WriteHeader(http.StatusOK)
	_ = jsoniter.NewEncoder(ctx.w).Encode(&res)
}
