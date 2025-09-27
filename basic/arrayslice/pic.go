package arrayslice

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	magic := make([][]uint8, dy)
	for x := range magic {
		blue := make([]uint8, dx)
		for y := range blue {
			blue[y] = uint8(x*y - 1)
		}
		magic[x] = blue
	}
	return magic
}

func ShowPic() {
	pic.Show(Pic)
}

// 如果不想使用外部依赖，可以使用这个简化版本
func ShowPicSimple() {
	result := Pic(8, 8)
	for i, row := range result {
		fmt.Printf("Row %d: %v\n", i, row)
	}
}
