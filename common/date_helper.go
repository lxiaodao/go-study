package common

import (
	
    "time"
    "strconv"
)



func  CurrentUnixNano() string{
    currentTime := time.Now()
    ranname :=  strconv.FormatInt(currentTime.UnixNano(), 10)
    return ranname
} 

