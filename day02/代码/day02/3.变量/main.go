package main

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	n1 = iota
	n2
	n3
)

func main() {

	// iota
	// 示例1
	/*
		const (
			v1 = 1
			v2 = 2
			v3 = 3
			v4 = 4
			v5 = 5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例2
	/*
		const (
			v1 = iota
			v2
			v3
			v4
			v5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例3
	/*
		const (
			v1 = iota + 2
			v2
			v3
			v4
			v5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例4：
	/*
		const (
			v1 = iota + 2
			_
			v2
			v3
			v4
			v5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

}
