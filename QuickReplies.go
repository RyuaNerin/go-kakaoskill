package skill

// 바로가기 응답은 발화와 동일합니다.
// 대신, 사용자가 직접 발화를 입력하지 않아도 선택을 통해서 발화를 전달하거나 다른 블록을 호출할 수 있습니다.
type QuickReply struct {
	Label       string      `json:"label"`       // 사용자에게 노출될 바로가기 응답의 표시
	Action      string      `json:"action"`      // 바로가기 응답의 기능 (message or block)
	MessageText string      `json:"messageText"` // 사용자 측으로 노출될 발화
	BlockId     string      `json:"blockId"`     // 바로가기 응답이 'block' 일 때 연결될 블록의 id
	Extra       interface{} `json:"extra"`       // 블록을 호출 시 추가적으로 제공하는 정보
}
