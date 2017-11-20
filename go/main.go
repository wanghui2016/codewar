package main
import (
	"fmt"
	"codewar/go/practic"
	// "sync"
)

func main(){
	// myPool := &sync.Pool{
	// 	New: func() interface{} {
	// 		fmt.Println("Creating new instance.")
	// 		return struct{}{}
	// 	},
	// }
	
	// myPool.Get()
	// instance := myPool.Get()
	// myPool.Put(instance)
	// myPool.Get()

	str := "aabbccabc"
	res := practic.Compress(str)
	fmt.Println(res)
}

