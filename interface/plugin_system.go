package interfacego

import "fmt"

// 实际应用场景3：插件系统
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	GetFee(amount float64) float64
}

// 不同的支付方式实现
type AlipayProcessor struct{}

func (a AlipayProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("通过支付宝支付 ¥%.2f\n", amount)
	return nil
}

func (a AlipayProcessor) GetFee(amount float64) float64 {
	return amount * 0.01 // 1% 手续费
}

type WeChatPayProcessor struct{}

func (w WeChatPayProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("通过微信支付 ¥%.2f\n", amount)
	return nil
}

func (w WeChatPayProcessor) GetFee(amount float64) float64 {
	return amount * 0.008 // 0.8% 手续费
}

type CreditCardProcessor struct{}

func (c CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("通过信用卡支付 ¥%.2f\n", amount)
	return nil
}

func (c CreditCardProcessor) GetFee(amount float64) float64 {
	return amount * 0.02 // 2% 手续费
}

// PayPal处理器
type PayPalProcessor struct{}

func (p PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("通过PayPal支付 $%.2f\n", amount/7) // 假设汇率1:7
	return nil
}

func (p PayPalProcessor) GetFee(amount float64) float64 {
	return amount * 0.025 // 2.5% 手续费
}

// 支付服务可以使用任何支付处理器
type PaymentService struct {
	processors map[string]PaymentProcessor
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		processors: make(map[string]PaymentProcessor),
	}
}

func (p *PaymentService) RegisterProcessor(name string, processor PaymentProcessor) {
	p.processors[name] = processor
	fmt.Printf("注册支付处理器: %s\n", name)
}

func (p *PaymentService) Pay(method string, amount float64) error {
	processor, exists := p.processors[method]
	if !exists {
		return fmt.Errorf("不支持的支付方式: %s", method)
	}

	fee := processor.GetFee(amount)
	total := amount + fee

	fmt.Printf("支付金额: ¥%.2f, 手续费: ¥%.2f, 总计: ¥%.2f\n",
		amount, fee, total)

	return processor.ProcessPayment(total)
}

func (p *PaymentService) ListAvailableMethods() {
	fmt.Println("可用的支付方式:")
	for method := range p.processors {
		fmt.Printf("  - %s\n", method)
	}
}

func PluginSystemDemo() {
	fmt.Println("=== 插件系统演示 ===")

	// 创建支付服务
	paymentService := NewPaymentService()

	// 注册不同的支付处理器（就像插件）
	paymentService.RegisterProcessor("支付宝", AlipayProcessor{})
	paymentService.RegisterProcessor("微信支付", WeChatPayProcessor{})
	paymentService.RegisterProcessor("信用卡", CreditCardProcessor{})

	fmt.Println()
	paymentService.ListAvailableMethods()

	fmt.Println("\n执行支付:")
	paymentService.Pay("支付宝", 100.0)
	paymentService.Pay("微信支付", 200.0)
	paymentService.Pay("信用卡", 300.0)

	// 新增支付方式也很容易
	fmt.Println("\n--- 新增PayPal支付 ---")

	// 使用新的支付处理器
	paymentService.RegisterProcessor("PayPal", PayPalProcessor{})
	paymentService.Pay("PayPal", 140.0)

	fmt.Println("\n优势：添加新支付方式完全不影响现有代码！")
}
