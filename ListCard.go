package skill

// 리스트 카드형 출력 요소입니다.
// 리스트 카드는 표현하고자 하는 대상이 다수일 때, 이를 효과적으로 전달할 수 있습니다.
// 헤더와 아이템을 포함하며, 헤더는 리스트 카드의 상단에 위치합니다.
// 리스트 상의 아이템은 각각의 구체적인 형태를 보여주며, 제목과 썸네일, 상세 설명을 포함합니다.
type ListCard struct {
	Header  ListItemHeader  `json:"header"`            // 카드의 상단 항목
	Items   []ListItemItems `json:"items"`             // 카드의 각각 아이템 (최대 5개)
	Buttons []Button        `json:"buttons,omitempty"` // 최대 2개
}

type ListItemHeader struct {
	Title    string `json:"title"`    // listCard의 제목
	ImageUrl string `json:"imageUrl"` // listCard 제목의 배경
	Link     Link   `json:"link"`     // 클릭시 작동하는 링크입니다.
}

type ListItemItems struct {
	Title       string `json:"title"`                 // 제목
	Description string `json:"description,omitempty"` // 설명
	ImageUrl    string `json:"imageUrl"`              // 우측 안내 사진
	Link        Link   `json:"link"`                  // 클릭시 작동하는 링크입니다.
}
