package mapgo

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func RunMapGo() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	fmt.Println(m1)
	fmt.Println(m2)

	mutMap := make(map[string]int)

	mutMap["答案"] = 42
	fmt.Println("值：", mutMap["答案"])

	mutMap["答案"] = 48
	fmt.Println("值：", mutMap["答案"])

	delete(mutMap, "答案")
	fmt.Println("值：", mutMap["答案"])

	v, ok := mutMap["答案"]
	fmt.Println("值：", v, "是否存在？", ok)
}
