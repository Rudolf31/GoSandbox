package services

type CalculatorService interface {
	Addition(a, b int) int
	Subtraction(a, b int) int
	GetOpperation() int
}

type calculatorServiceImpl struct {
	count int
}

// Addition implements CalculatorService.
func (c *calculatorServiceImpl) Addition(a int, b int) int {
	c.count++
	return a + b
}

// GetOpperation implements CalculatorService.
func (c *calculatorServiceImpl) GetOpperation() int {
	return c.count
}

// Subtraction implements CalculatorService.
func (c *calculatorServiceImpl) Subtraction(a int, b int) int {
	c.count++
	return a - b
}

func NewCalculatorService() CalculatorService {
	return &calculatorServiceImpl{}
}
