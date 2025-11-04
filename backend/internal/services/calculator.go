package services

import (
	customeerrors "interface_lesson/internal/customerrors"

	"go.uber.org/zap"
)

type CalculatorService interface {
	Addition(a, b int) int
	Subtraction(a, b int) int
	GetOperation() int
	Division(a, b int) (int, *customeerrors.Wrapper)
}

type calculatorServiceImpl struct {
	log   *zap.Logger
	count int
}

// Addition implements CalculatorService.
func (c *calculatorServiceImpl) Addition(a int, b int) int {
	c.count++
	return a + b
}

// GetOperation implements CalculatorService.
func (c *calculatorServiceImpl) GetOperation() int {
	return c.count
}

// Subtraction implements CalculatorService.
func (c *calculatorServiceImpl) Subtraction(a int, b int) int {
	c.count++
	return a - b
}

func (c *calculatorServiceImpl) Division(a int, b int) (int, *customeerrors.Wrapper) {
	c.count++

	if b == 0 {
		c.log.Warn(
			"Attempted to divide by zero",
			zap.Int("a", a),
			zap.Int("b", b),
		)
		return 0, customeerrors.NewErrorWrapper(
			customeerrors.ErrDivisionByZero,
			"Whyyy..",
		)
	}
	return a / b, nil
}

func NewCalculatorService(log *zap.Logger) CalculatorService {
	return &calculatorServiceImpl{log: log}
}
