package ellipticcurves

import (
	"fmt"
	"math.io/crath/modular"
	"math/big"
)

// SimpleCurve - elliptic curve with definition:
// 		Em (A, B) == y^2 = x^3 + Ax + B (mod M)
type SimpleCurve struct {
	A, B, M *big.Int
	points  []Point
}

type Point struct {
	X, Y *big.Int
}

// CalculateVal - calculate a value of a function with provided x value
func (c *SimpleCurve) CalculateVal(x *big.Int) *big.Int {
	result := big.NewInt(0)
	return result.
		Mul(x, c.A).
		Add(result, c.B).
		Add(result, big.NewInt(0).Exp(x, big.NewInt(3), c.M)).
		Mod(result, c.M)
}

func (c *SimpleCurve) CalculatePoints() []Point {
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
			if yPoints, isIn := squares[c.CalculateVal(i).String()]; isIn {
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

func (c *SimpleCurve) InitPoints() {
	c.points = c.CalculatePoints()
}

func (c *SimpleCurve) AddPoints(a, b Point) (Point, error) {
	if len(c.points) < 1 {
		c.points = c.CalculatePoints()
	}

	result := Point{
		X: big.NewInt(0),
		Y: big.NewInt(0),
	}

	if a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0 {
		k := big.NewInt(0)

		two := big.NewInt(2)
		two.Mul(two, a.Y)

		k.Exp(a.X, big.NewInt(2), c.M).Mul(k, big.NewInt(3)).Add(k, c.A).Mul(k, modular.ModInverse(two, c.M)).Mod(k, c.M)

		x3 := big.NewInt(0)
		x3.Exp(k, big.NewInt(2), c.M)

		sub := big.NewInt(2)
		sub.Mul(sub, a.X)

		x3.Sub(x3, sub)
		x3.Mod(x3, c.M)

		result.X.Set(x3)

		y2 := big.NewInt(0).Set(a.X)
		y2.Sub(y2, x3).Mod(y2, c.M)
		y2.Mul(y2, k)
		y2.Sub(y2, a.Y).Mod(y2, c.M)

		result.Y.Set(y2)
	} else if a.X.Cmp(b.X) == 0 { // x1 must not be equal to x2
		return result, fmt.Errorf("infinitely distant point")
	} else {
		k := big.NewInt(0).Set(b.Y)
		k.Sub(k, a.Y)

		div := big.NewInt(0).Set(b.X)
		div.Sub(div, a.X)

		k.Mul(k, modular.ModInverse(div, c.M)).Mod(k, c.M)

		x3 := big.NewInt(0)
		x3.Exp(k, big.NewInt(2), c.M)
		x3.Sub(x3, a.X).Sub(x3, b.X).Mod(x3, c.M)

		result.X.Set(x3)

		y2 := big.NewInt(0).Set(a.X)
		y2.Sub(y2, x3).Mod(y2, c.M)
		y2.Mul(y2, k)
		y2.Sub(y2, a.Y).Mod(y2, c.M)

		result.Y.Set(y2)
	}

	return result, nil
}

// TODO: Incorrect order calculations - do some tests
// PointOrd - point order == how many times point must be added to itself to get infinitely distant point
func (c *SimpleCurve) PointOrd(p Point) *big.Int {
	ord := big.NewInt(1)
	one := big.NewInt(1)
	two := big.NewInt(2)
	pt := Point{
		X: p.X,
		Y: p.Y,
	}

	y2 := big.NewInt(0)
	var err error = nil
	for y2.Set(pt.Y).Exp(y2, two, c.M); c.CalculateVal(pt.X).Cmp(y2) == 0; {
		pt, err = c.AddPoints(pt, p)

		ord.Add(ord, one)

		if err != nil {
			break
		}
	}

	return ord
}
