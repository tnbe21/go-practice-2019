package terror

import (
	"fmt"
	"time"
)

func Print(err error) {
	fmt.Printf("[error] %s %s \n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
}
