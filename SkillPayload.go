package skill

type SkillPayload struct {
	UserRequest UserRequest `json:"userRequest"`
	Bot         Bot         `json:"bot"`
	Action      Action      `json:"action"`
}

type UserRequest struct {
	Timezone  string `json:"timezone"`  // 사용자의 시간대를 반환합니다.
	Block     Block  `json:"block"`     // 사용자의 발화에 반응한 블록의 정보입니다.
	Utterance string `json:"utterance"` // 봇 시스템에 전달된 사용자의 발화입니다.
	Lang      string `json:"lang"`      // 사용자의 언어를 반환합니다.
	User      User   `json:"user"`      // 사용자의 정보입니다.
}

type Block struct {
	Id   string `json:"id"`   // 블록 ID
	Name string `json:"name"` // 블록 이름
}

type Bot struct {
	Id   string `json:"id"`   // 봇의 고유한 식별자입니다.
	Name string `json:"name"` // 설정된 봇의 이름입니다.
}

type Action struct {
	Id           string                 `json:"id"`           // 스킬의 고유한 식별자입니다.
	Name         string                 `json:"name"`         // 설정된 스킬의 이름입니다.
	Params       map[string]interface{} `json:"params"`       // 사용자의 발화로부터 추출한 파라미터 정보입니다.
	DetailParams map[string]DetailParam `json:"detailParams"` // 사용자의 발화로부터 추출한 엔티티의 상세 정보입니다.
	ClientExtra  map[string]interface{} `json:"clientExtra"`  // 사용자의 발화의 추가적인 정보
}

type User struct {
	Id         string       `json:"id"`         // 사용자를 식별할 수 있는 key로 최대 70자의 값을 가지고 있습니다.
	Type       string       `json:"type"`       // 현재는 botUserKey만 제공합니다.
	Properties UserProperty `json:"properties"` // 추가적으로 제공하는 사용자의 속성 정보입니다.
}

type UserProperty struct {
	PlusFriendUserKey string `json:"plusfriendUserKey"` // 카카오톡 채널에서 제공하는 사용자 식별키 입니다.
	AppUserId         string `json:"appUserId"`         // 봇 설정에서 앱키를 설정한 경우에만 제공되는 사용자 정보입니다.
}

type DetailParam struct {
	Origin    interface{} `json:"origin"`
	Value     interface{} `json:"value"`
	GroupName string      `json:"groupName"`
}
