package loopfunction

import (
	"fmt"
	"math"
)

func SqrtDemo(x float64) float64 {
	z := 1.0 // initial guess
	fmt.Printf("计算 √%.1f 的过程:\n", x)
	fmt.Printf("初始猜测: z = %.6f\n", z)

	for i := 1; i <= 5; i++ {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("第%d次迭代: z = %.6f (误差: %.6f)\n",
			i, z, math.Abs(z*z-x))

		// 如果连续两次结果很接近，可以提前结束
		if math.Abs(z-oldZ) < 1e-10 {
			fmt.Printf("收敛了！\n")
			break
		}
	}

	fmt.Printf("最终结果: %.6f\n", z)
	fmt.Printf("标准库结果: %.6f\n", math.Sqrt(x))
	fmt.Printf("误差: %.10f\n\n", math.Abs(z-math.Sqrt(x)))

	return z
}

func TestDemo() {
	SqrtDemo(2.0)
	SqrtDemo(9.0)
	SqrtDemo(16.0)
}
