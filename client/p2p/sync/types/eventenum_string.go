// Code generated by "stringer -type=eventEnum"; DO NOT EDIT.

package synctypes

import "strconv"

const _eventEnum_name = "Imported"

var _eventEnum_index = [...]uint8{0, 8}

func (i eventEnum) String() string {
	if i < 0 || i >= eventEnum(len(_eventEnum_index)-1) {
		return "eventEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _eventEnum_name[_eventEnum_index[i]:_eventEnum_index[i+1]]
}
