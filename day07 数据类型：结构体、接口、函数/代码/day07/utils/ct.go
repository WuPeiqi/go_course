/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package utils

// 私有
type file struct {
	fd   int
	name string
}

// 公有
func NewFile(fd int, name string) *file {
	// ...
	return &file{fd, name}
}
