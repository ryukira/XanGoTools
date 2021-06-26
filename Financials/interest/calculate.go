package interest

import "math"

func SayHello(name string) string {
	return "hello " + name
}

// FuturePresent || Given Present Return Future
func FuturePresent(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * math.Pow(1+interest, time)
	return result
}

// PresentFuture || Given Future Return Present
func PresentFuture(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * math.Pow(1+interest, -time)
	return result
}

// FutureAnnual || Given Annual Return Future
func FutureAnnual(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * ((math.Pow(1+interest, time) - 1) / interest)
	return result
}

// AnnualFuture || Given Future Return Annual
func AnnualFuture(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * (interest / (math.Pow(1+interest, time) - 1))
	return result
}

// AnnualPresent || Given Present Return Annual
func AnnualPresent(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * ((interest * math.Pow(1+interest, time)) / (math.Pow(1+interest, time) - 1))
	return result
}

// PresentAnnual || Given Annual Return Present
func PresentAnnual(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * ((math.Pow(1+interest, time) - 1) / (interest * math.Pow(1+interest, time)))
	return result
}

// AnnualGradient || Given Gradient Return Annual
func AnnualGradient(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * ((1 / interest) - (time / (math.Pow(1+interest, time) - 1)))
	return result
}

// PresentGradient || Given Gradient Return Present
func PresentGradient(value float64, time float64, interest float64) float64 {
	var result float64
	result = value * ((math.Pow(1+interest, time) - interest*time - 1) / (math.Pow(interest, 2) * math.Pow(1+interest, time)))
	return result
}
