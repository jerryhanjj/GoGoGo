package loopfunction

import (
	"fmt"
	"math"
)

func DetailedSqrt(x float64) {
	z := 1.0
	target := math.Sqrt(x)

	fmt.Printf("计算 √%.1f (真实值: %.6f)\n", x, target)
	fmt.Printf("初始猜测: z = %.6f\n\n", z)

	for i := 1; i <= 5; i++ {
		// 显示详细计算过程
		zSquared := z * z
		diff := zSquared - x
		derivative := 2 * z
		correction := diff / derivative

		fmt.Printf("第%d次迭代:\n", i)
		fmt.Printf("  当前 z = %.6f\n", z)
		fmt.Printf("  z² = %.6f\n", zSquared)
		fmt.Printf("  z² - x = %.6f - %.1f = %.6f\n", zSquared, x, diff)
		fmt.Printf("  2z = %.6f\n", derivative)
		fmt.Printf("  修正值 = (z² - x) / (2z) = %.6f\n", correction)

		// 应用修正
		z -= correction

		fmt.Printf("  新的 z = %.6f - (%.6f) = %.6f\n", z+correction, correction, z)

		if z > target {
			fmt.Printf("  → z 比真实值大 %.6f\n", z-target)
		} else {
			fmt.Printf("  → z 比真实值小 %.6f\n", target-z)
		}
		fmt.Printf("  误差: %.6f\n\n", math.Abs(z-target))
	}

	fmt.Printf("最终结果: %.6f\n", z)
	fmt.Printf("标准库结果: %.6f\n", target)
}

func DemoRun() {
	DetailedSqrt(2.0)
	fmt.Println("=" + string(make([]byte, 50)) + "=")
	DetailedSqrt(9.0)
}
