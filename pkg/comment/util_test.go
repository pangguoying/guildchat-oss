package comment

import (
	"testing"
)

/**
*如果想每次执行结果都不一样的话，需要命令行假如 “-count=1” 禁用 cache
 */
func TestRandString(t *testing.T) {
	var length int = 6
	str := RandString(length)
	if len(str) != length {
		t.Error("测试失败")
	}
}

func TestOpenFile(t *testing.T) {
	_, err := OpenFile("./hello/a.txt")
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
}
