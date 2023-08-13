package filed

const (
	NEW         = 0 + iota // 新投递
	ARRANGED               // 已经安排面试
	INTERVIEWED            // 已面试
	PASS                   // 通过
	OUT                    // 淘汰
)

var (
	First       = "first_status"
	Second      = "second_status"
	Final       = "final_status"
	First_Eva   = "first_eva"
	First_Mark  = "first_mark"
	Second_Eva  = "second_eva"
	Second_Mark = "second_mark"
)
