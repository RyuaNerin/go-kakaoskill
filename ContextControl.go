package skill

// context control 필드는 블록에서 생성한 outputContext의 lifeSpan, params 등을 제어할 수 있습니다.
type ContextControl struct {
	Values []ContextValue `json:"values"`
}

type ContextValue struct {
	Name     string            `json:"name"`             // 수정하려는 output 컨텍스트의 이름
	LifeSpan string            `json:"lifeSpan"`         // 수정하려는 ouptut 컨텍스트의 lifeSpan
	Params   map[string]string `json:"params,omitempty"` // output 컨텍스트에 저장하는 추가 데이터
}
