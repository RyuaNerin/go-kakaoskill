package skill

// 기본 카드형 출력 요소입니다.
// 기본 카드는 소셜, 썸네일, 프로필 등을 통해서 사진이나 글, 인물 정보 등을 공유할 수 있습니다.
// 기본 카드는 제목과 설명 외에 썸네일 그룹, 프로필, 버튼 그룹, 소셜 정보를 추가로 포함합니다.
type BasicCard struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Buttons     []Button   `json:"buttons,omitempty"`
	// profile
	// social
}
