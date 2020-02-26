package skill

type Thumbnail struct {
	ImageUrl string `json:"imageUrl"`       //이미지의 url입니다.
	Link     *Link  `json:"link,omitempty"` // 이미지 클릭시 작동하는 link입니다.
	// true: 이미지 영역을 1:1 비율로 두고 이미지의 원본 비율을 유지합니다. 이미지가 없는 영역은 흰색으로 노출합니다.
	// false: 이미지 영역을 2:1 비율로 두고 이미지의 가운데를 크롭하여 노출합니다.
	FixedRatio bool `json:"fixedRatio"`
	Width      int  `json:"width"`  // fixedRatio를 true로 설정할 경우 필요한 값입니다. 실제 이미지 사이즈와 다른 값일 경우 원본이미지와 다르게 표현될 수 있습니다.
	Height     int  `json:"height"` // fixedRatio를 true로 설정할 경우 필요한 값입니다. 실제 이미지 사이즈와 다른 값일 경우 원본이미지와 다르게 표현될 수 있습니다.
}

// 링크는 다음과 같은 우선순위를 갖습니다.
// 모바일: web < mobile < ios or android (os에 따라 적용)
// pc: web < pc < mac or win (os에 따라 적용))
type Link struct {
	Mobile  string `json:"mobile,omitempty"`  // ios와 android를 아우르는 mobile link입니다.
	Ios     string `json:"ios,omitempty"`     // ios의 웹이나 앱을 실행하는 link입니다.
	Android string `json:"android,omitempty"` // android의 웹이나 앱을 실행하는 link입니다.
	Pc      string `json:"pc,omitempty"`      // mac과 window를 아우르는 pc link입니다.
	Mac     string `json:"mac,omitempty"`     // mac의 웹이나 앱을 실행하는 link입니다.
	Win     string `json:"win,omitempty"`     // window의 웹이나 앱을 실행하는 link입니다.
	Web     string `json:"web,omitempty"`     // 모든 기기를 아우르는 link입니다.
}

// action 종류
// webLink  : 웹 브라우저를 열고 webLinkUrl 의 주소로 이동합니다.
// osLink   : osLink 값에 따라서 웹의 주소로 이동하거나 앱을 실행합니다.
// message  : 사용자의 발화로 messageText를 실행합니다. (바로가기 응답의 메세지 연결 기능과 동일)
// phone    : phoneNumber에 있는 번호로 전화를 겁니다.
// block    : blockId를 갖는 블록을 호출합니다. (바로가기 응답의 블록 연결 기능과 동일)
//            messageText가 있다면, 해당 messageText가 사용자의 발화로 나가게 됩니다.
//            messageText가 없다면, button의 label이 사용자의 발화로 나가게 됩니다.
// share    : 말풍선을 다른 유저에게 공유합니다. share action은 특히 케로셀을 공유해야 하는 경우 유용합니다.
// operator : 상담원 연결 기능을 제공합니다. (베타)
type Button struct {
	Label      string `json:"label"`                // 버튼에 적히는 문구입니다 (최대 8자)
	Action     string `json:"action"`               // 버튼 클릭시 수행될 작업입니다
	WebLinkUrl string `json:"webLinkUrl,omitempty"` // 웹 브라우저를 열고 webLinkUrl 의 주소로 이동합니다 (Action: webLink)
	OsLink     *Link  `json:"osLink,omitempty"`     // osLink 값에 따라서 웹의 주소로 이동하거나 앱을 실행합니다 (Action: osLink)
	// Action: "message" / "block"
	// message => 사용자의 발화로 messageText를 내보냅니다. (바로가기 응답의 메세지 연결 기능과 동일)
	// block => 블록 연결시 사용자의 발화로 노출됩니다.
	MessageText string                 `json:"messageText,omitempty"`
	PhoneNumber string                 `json:"phoneNumber,omitempty"` // phoneNumber에 있는 번호로 전화를 겁니다 (전화번호)
	BlockId     string                 `json:"blockId,omitempty"`     // blockId를 갖는 블록을 호출합니다 (바로가기 응답의 블록 연결 기능과 동일)
	Extra       map[string]interface{} `json:"extra,omitempty"`       // block이나 message action으로 블록 호출시, 스킬 서버에 추가적으로 제공하는 정보
}

type CarouselHeader struct {
	Title       string     `json:"title"`                 // 케로셀 헤드 제목 (최대 2줄, 한 줄에 들어갈 수 있는 글자 수는 기기에 따라 달라집니다.)
	Description string     `json:"description,omitempty"` // 케로셀 헤드 설명 (최대 3줄, 한 줄에 들어갈 수 있는 글자 수는 기기에 따라 달라집니다.)
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`   // 케로셀 헤드 배경 이미지 (현재 imageUrl 값만 지원합니다.)
}

type Profile struct {
	Nickname string `json:"nickname"` // 프로필 이름
	ImageUrl string `json:"imageUrl"` // 프로필 이미지
}
