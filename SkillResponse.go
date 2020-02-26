package skill

type SkillResponse struct {
	Version  string                 `json:"version"`           // must "2.0"
	Template SkillTemplate          `json:"template"`          // 스킬 응답의 출력 모양을 담고 있는 항목입니다. 출력으로 사용할 요소 그룹, 바로가기 응답 그룹 등을 포함합니다.
	Context  *ContextControl        `json:"context,omitempty"` // 블록의 context 설정을 제어할 수 있습니다.
	Data     map[string]interface{} `json:"data,omitempty"`    //필요에 따라 커스텀한 데이터를 넣을 수 있는 항목입니다.
}
