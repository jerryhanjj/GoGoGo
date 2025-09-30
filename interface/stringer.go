package interfacego

import "fmt"

type stPerson struct {
	Name string
	Age  int
}

func (p stPerson) String() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func StringerDemo() {
	p := stPerson{Name: "Alice", Age: 30}
	fmt.Println(p) // 自动调用 p.String()

	z := stPerson{"Bob", 25}
	fmt.Printf("%v\n", z)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func StringerIPDemo() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%s: %s\n", name, ip)
	}
}
