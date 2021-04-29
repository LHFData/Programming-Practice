package BitOperation

import (
	"fmt"
	"strings"
)

func hammingWeight(num uint32) int {
	return strings.Count(fmt.Sprintf("%b",num),"1")
}

