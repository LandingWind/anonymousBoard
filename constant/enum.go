package constant

type READ_LOCK_STATUS int

const (
	READ_LOCK_CLOSE READ_LOCK_STATUS = 0 // default: 完全开放
	READ_LOCK_OPEN  READ_LOCK_STATUS = 1 // 进入需要 enterKey
)

type EDIT_LOCK_STATUS int

const (
	EDIT_LOCK_CLOSE EDIT_LOCK_STATUS = 0 // 编辑开放
	EDIT_LOCK_OPEN  EDIT_LOCK_STATUS = 1 // default: 编辑需要 editKey
)
