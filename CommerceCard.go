package skill

// 케로셀은 여러 장의 카드를 하나의 메세지에 일렬로 포함하는 타입입니다.
type CommerceCard struct {
	Type   string         `json:"type"`             // 케로셀의 타입입니다.	("basicCard" 혹은 "commerceCard")
	Items  []interface{}  `json:"items"`            // 케로셀 아이템입니다 (BasicCard or CommerceCard, 최대 10개)
	Header CarouselHeader `json:"header,omitempty"` // 케로셀의 커버를 제공합니다 (CommerceCard 만 지원하고 있습니다. 추후 BasicCard 도 지원할 예정입니다.)
}
