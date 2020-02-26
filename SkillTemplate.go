package skill

type SkillTemplate struct {
	Outputs      []Component  `json:"outputs"`                // 출력 그룹(outputs)은 여러 종류의 출력 요소(component)를 포함합니다 (1개 이상 3개 이하)
	QuickReplies []QuickReply `json:"quickReplies,omitempty"` // 바로가기 그룹(quickReplies)은 여러 개의 바로가기 요소(quickReply)를 포함합니다 (10개 이하)
}

type Component struct {
	SimpleText   *SimpleText   `json:"simpleText,omitempty"`   // 간단 텍스트
	SimpleImage  *SimpleImage  `json:"simpleImage,omitempty"`  // 간단 이미지
	BasicCard    *BasicCard    `json:"basicCard,omitempty"`    // 기본 카드 (캐로샐 가능)
	CommerceCard *CommerceCard `json:"commerceCard,omitempty"` // 커머스 카드 (캐로샐 가능)
	ListCard     *ListCard     `json:"listCard,omitempty"`     // 리스트 카드
}
