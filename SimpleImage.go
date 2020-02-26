package skill

// 간단한 이미지형 출력 요소입니다.
// 이미지 링크 주소를 포함하면 이를 스크랩하여 사용자에게 전달합니다.
// 이미지 링크 주소가 유효하지 않을 수 있기 때문에, 대체 텍스트를 꼭 포함해야 합니다.
type SimpleImage struct {
	ImageUrl string `json:"imageUrl"` // 전달하고자 하는 이미지의 url입니다 (URL 형식)
	AltText  string `json:"altText"`  // url이 유효하지 않은 경우, 전달되는 텍스트입니다 (1000자 이내)
}
