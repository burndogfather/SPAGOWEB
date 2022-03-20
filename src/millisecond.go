package millisecond

import (
	"time"
	"strconv"
)

func millisecond()(string) {
	ymdhis := time.Now().Format("20060102150405")
	microtimeInt64 := time.Now().UnixNano() / int64(time.Millisecond)
	microtimeUnix := strconv.FormatInt(microtimeInt64, 10)
	return ymdhis + microtimeUnix[10:12]
}
