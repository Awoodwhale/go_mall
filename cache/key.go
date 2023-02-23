package cache

import (
	"fmt"
	"strconv"
)

// ProductViewKey
// @Description: 获取redis中view的key
// @param uid uint
// @return string
func ProductViewKey(uid uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(uid)))
}
