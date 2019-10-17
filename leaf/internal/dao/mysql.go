package dao

const (
	_all       = "select big_tag,max_id,step,desc,update_time from segment"
	_maxSeq    = "select max_id from segment where big_tag=?"
	_updateMax = "update segment set max_id=? where big_tag=? and max_id=?"
)
