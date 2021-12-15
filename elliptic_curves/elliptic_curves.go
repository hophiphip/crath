package ellipticcurves

import "math/big"

// SimpleCurve - elliptic curve with definition:
// 		Em (A, B) == y^2 = x^3 + Ax + B (mod M)
type SimpleCurve struct {
	A, B, M *big.Int
}

type Point struct {
	X, Y *big.Int
}

// calculateVal - calculate a value of a function with provided x value
func (c *SimpleCurve) calculateVal(x *big.Int) *big.Int {
	result := big.NewInt(0)
	return result.
		Mul(x, c.A).
		Add(result, c.B).
		Add(result, big.NewInt(0).Exp(x, big.NewInt(3), c.M)).
		Mod(result, c.M)
}

func (c *SimpleCurve) getPoints() []Point {
	var points []Point
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	buf := big.NewInt(0)
	squares := make(map[string][]*big.Int)

	for i := big.NewInt(0); i.Cmp(c.M) < 0; i.Add(i, one) {
		if i.Cmp(zero) == 0 || buf.GCD(nil, nil, i, c.M).Cmp(one) <= 0 {
			squares[buf.Exp(i, two, c.M).String()] = append(squares[buf.Exp(i, two, c.M).String()], big.NewInt(0).Set(i))
		}
	}

	for i := big.NewInt(0); i.Cmp(c.M) < 0; i.Add(i, one) {
		if i.Cmp(zero) == 0 || buf.GCD(nil, nil, i, c.M).Cmp(one) <= 0 {
			if yPoints, isIn := squares[c.calculateVal(i).String()]; isIn {
				for _, y := range yPoints {
					points = append(points, Point{
						X: big.NewInt(0).Set(i),
						Y: big.NewInt(0).Set(y),
					})
				}
			}
		}
	}

	return points
}

func (c *SimpleCurve) addPoints(a, b Point) (Point, error) {
	result := Point{
		X: big.NewInt(0),
		Y: big.NewInt(0),
	}

	return result, nil
}
