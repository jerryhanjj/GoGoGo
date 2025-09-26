package loopfunction

import (
	"fmt"
	"math"
)

func SqrtWithInitial(x, initial float64) {
	z := initial
	target := math.Sqrt(x)

	fmt.Printf("计算 √%.1f，初始值 z = %.2f (真实值: %.6f)\n", x, initial, target)

	for i := 1; i <= 8; i++ {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		error := math.Abs(z - target)

		fmt.Printf("第%d次: z = %.6f, 误差 = %.8f", i, z, error)

		// 显示收敛情况
		if i > 1 {
			improvement := math.Abs(oldZ-target) / error
			if improvement > 1 {
				fmt.Printf(" (改善 %.1f倍)", improvement)
			}
		}
		fmt.Println()

		// 如果达到机器精度就停止
		if error < 1e-15 {
			fmt.Printf("  → 在第%d次迭代达到机器精度！\n", i)
			break
		}
	}
	fmt.Printf("最终误差: %.2e\n\n", math.Abs(z-target))
}

func CompareInitialValues() {
	x := 2.0
	fmt.Printf("=== 比较不同初始值计算 √%.1f 的效果 ===\n\n", x)

	// 测试不同的初始值
	initials := []float64{0.1, 0.5, 1.0, 1.4, 2.0, 5.0, 10.0, 100.0}

	for _, init := range initials {
		SqrtWithInitial(x, init)
	}
}

func AnalyzeConvergence() {
	x := 16.0
	target := math.Sqrt(x)

	fmt.Printf("=== 分析不同初始值的收敛性 (计算 √%.1f = %.1f) ===\n\n", x, target)

	testCases := []struct {
		initial     float64
		description string
	}{
		{1.0, "远小于真实值"},
		{4.0, "等于真实值"},
		{8.0, "是真实值的2倍"},
		{16.0, "是真实值的4倍"},
		{64.0, "是真实值的16倍"},
	}

	for _, tc := range testCases {
		fmt.Printf("初始值 %.1f (%s):\n", tc.initial, tc.description)
		z := tc.initial

		for i := 1; i <= 5; i++ {
			z -= (z*z - x) / (2 * z)
			error := math.Abs(z - target)
			fmt.Printf("  第%d次: %.6f (误差: %.6f)\n", i, z, error)
		}
		fmt.Println()
	}
}
func BadInitialValues() {
	x := 4.0
	target := math.Sqrt(x)

	fmt.Printf("=== 测试极端初始值 (计算 √%.1f = %.1f) ===\n\n", x, target)

	extremeCases := []struct {
		initial     float64
		description string
	}{
		{0.001, "非常接近0"},
		{1000.0, "非常大的值"},
		{0.01, "很小的值"},
	}

	for _, tc := range extremeCases {
		fmt.Printf("初始值 %.3f (%s):\n", tc.initial, tc.description)
		z := tc.initial

		for i := 1; i <= 6; i++ {
			z -= (z*z - x) / (2 * z)
			error := math.Abs(z - target)
			fmt.Printf("  第%d次: %.6f (误差: %.6f)\n", i, z, error)
		}
		fmt.Println()
	}
}

func InitialValueDemo() {
	CompareInitialValues()
	fmt.Println(string(make([]byte, 80)))
	AnalyzeConvergence()
	fmt.Println(string(make([]byte, 80)))
	BadInitialValues()
}
