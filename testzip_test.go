package zip

import (
	"testing"
)

//压缩用arc
// arc archive hotel.zip "D:\\Documents\\Visual Studio 2013\\Projects\\wuliao\\酒店数据\\bin\\Debug"
func TestX(t *testing.T) {
	err := DefaultZip.Unarchive("ss.zip", "ff")
	println(err)
}
