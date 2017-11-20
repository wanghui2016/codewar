package practic
import (
	"fmt"
)

// Compress 压缩重复字符
func Compress(str string) (res string){
	var ch uint8 = 0
	count := 1
	for i:=0; i < len(str);i++{
		if ch == 0 {
			ch = str[i]
			continue
		}
		if ch == str[i] {
			count++
		} else {
			res += fmt.Sprintf("%c%d",ch,count)
			ch = 0
			count = 1
		}
	}
	if ch != 0{
		res += fmt.Sprintf("%c%d",ch,count)
	}
	return
}