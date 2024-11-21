package calc

type calculator struct{}

func NewCalculator() *calculator {
	return &calculator{}
}

func (c *calculator) Calculate(num1, num2 float64, operatorstr string) float64 {
	switch operatorstr {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0 
	}
}
