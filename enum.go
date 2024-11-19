package pica

// 排序方式
const (
	SortDefault    = "ua"
	SortTimeNewest = "dd" // 新到旧
	SortTimeOldest = "da" // 旧到新
	SortLikeMost   = "ld" // 最多爱心
	SortViewMost   = "vd" // 最多绅士指名
)

// 图片质量
const (
	ImageQualityOriginal = "original"
	ImageQualityLow      = "low"
	ImageQualityMedium   = "medium"
	ImageQualityHigh     = "high"
)

// 请求结果
const (
	ActionLike        = "like"
	ActionUnlike      = "unlike"
	ActionFavourite   = "favourite"
	ActionUnFavourite = "un_favourite"
)

// 排行榜类型
const (
	LeaderboardH24 = "H24"
	LeaderboardD7  = "D7"
	LeaderboardD30 = "D30"
)
