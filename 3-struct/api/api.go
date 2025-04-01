package api

import (
	"3-struct/app/config"
	"fmt"
)

func Api() {
	key := config.NewConfig().Key
	fmt.Println(key)
}
