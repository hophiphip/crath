package ellipticcurves

import (
	"fmt"
	"math/big"
	"testing"
)

type InputOutput struct {
	input, output *big.Int
}

type TwoInputOutput struct {
	input1, input2, output Point
}

type CalculateTestCase struct {
	curve  SimpleCurve
	values []InputOutput
}

type GetPointsTestCase struct {
	curve  SimpleCurve
	points []Point
}

type AddPointsTestCase struct {
	curve  SimpleCurve
	values []TwoInputOutput
}

var curveTestSamples = []CalculateTestCase{
	// Curve: E7 (2, 3) ==> y^2 = x^3 + 2x + 3 (mod 7)
	{
		curve: SimpleCurve{
			A: big.NewInt(2),
			B: big.NewInt(3),
			M: big.NewInt(7),
		},
		values: []InputOutput{
			{big.NewInt(0), big.NewInt(3)},
			{big.NewInt(1), big.NewInt(6)},
			{big.NewInt(2), big.NewInt(1)},
			{big.NewInt(3), big.NewInt(1)},
			{big.NewInt(4), big.NewInt(5)},
			{big.NewInt(5), big.NewInt(5)},
			{big.NewInt(6), big.NewInt(0)},
		},
	},
	// Curve: E5 (3, 2) ==> y^2 = x^3 + 3x + 2 (mod 5)
	{
		curve: SimpleCurve{
			A: big.NewInt(3),
			B: big.NewInt(2),
			M: big.NewInt(5),
		},
		values: []InputOutput{
			{big.NewInt(0), big.NewInt(2)},
			{big.NewInt(1), big.NewInt(1)},
			{big.NewInt(2), big.NewInt(1)},
			{big.NewInt(3), big.NewInt(3)},
			{big.NewInt(4), big.NewInt(3)},
		},
	},
}

var pointsTestSamples = []GetPointsTestCase{
	{
		curve: SimpleCurve{
			A: big.NewInt(2),
			B: big.NewInt(3),
			M: big.NewInt(7),
		},
		points: []Point{
			{
				X: big.NewInt(2),
				Y: big.NewInt(1),
			},
			{
				X: big.NewInt(2),
				Y: big.NewInt(6),
			},
			{
				X: big.NewInt(3),
				Y: big.NewInt(1),
			},
			{
				X: big.NewInt(3),
				Y: big.NewInt(6),
			},
			{
				X: big.NewInt(6),
				Y: big.NewInt(0),
			},
		},
	},
	{
		curve: SimpleCurve{
			A: big.NewInt(3),
			B: big.NewInt(2),
			M: big.NewInt(5),
		},
		points: []Point{
			{
				X: big.NewInt(1),
				Y: big.NewInt(1),
			},
			{
				X: big.NewInt(1),
				Y: big.NewInt(4),
			},
			{
				X: big.NewInt(2),
				Y: big.NewInt(1),
			},
			{
				X: big.NewInt(2),
				Y: big.NewInt(4),
			},
		},
	},
}

var addPointsTestSamples = []AddPointsTestCase{
	{
		curve: SimpleCurve{
			A:      big.NewInt(3),
			B:      big.NewInt(2),
			M:      big.NewInt(5),
			points: nil,
		},
		values: []TwoInputOutput{
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(0),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(1),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(0),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				output: Point{
					X: big.NewInt(0),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(4),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(0),
					Y: big.NewInt(0),
				},
			},
		},
	},
	{
		curve: SimpleCurve{
			A:      big.NewInt(2),
			B:      big.NewInt(3),
			M:      big.NewInt(7),
			points: nil,
		},
		values: []TwoInputOutput{
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			{
				input1: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
				input2: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				output: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
				output: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
				input2: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
				output: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
			},
			{
				input1: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				output: Point{
					X: big.NewInt(0),
					Y: big.NewInt(0),
				},
			},
			{
				input1: Point{
					X: big.NewInt(3),
					Y: big.NewInt(1),
				},
				input2: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(1),
				},
			},
			{
				input1: Point{
					X: big.NewInt(3),
					Y: big.NewInt(6),
				},
				input2: Point{
					X: big.NewInt(6),
					Y: big.NewInt(0),
				},
				output: Point{
					X: big.NewInt(2),
					Y: big.NewInt(6),
				},
			},
		},
	},
}

func TestCalculateVal(t *testing.T) {
	buffer := big.NewInt(0)
	for _, testSample := range curveTestSamples {
		for _, testValue := range testSample.values {
			buffer.Set(testSample.curve.CalculateVal(testValue.input))
			if buffer.Cmp(testValue.output) != 0 {
				t.Errorf("For the curve: E(%5.5s, %5.5s) expected: %10.5s, but got: %10.5s\n",
					testSample.curve.A.String(),
					testSample.curve.B.String(),
					testValue.output.String(),
					buffer.String(),
				)
			}
		}
	}
}

func TestCalculatePoints(t *testing.T) {
	for _, testSample := range pointsTestSamples {
		points := testSample.curve.CalculatePoints()

		if len(points) != len(testSample.points) {
			t.Errorf("For the curve: E(%5.5s, %5.5s) expected amount of points: %d, but got: %d\n",
				testSample.curve.A.String(),
				testSample.curve.B.String(),
				len(testSample.points),
				len(points),
			)
		}

		for i, point := range testSample.points {
			if point.X.Cmp(points[i].X) != 0 || point.Y.Cmp(points[i].Y) != 0 {
				t.Errorf("For the curve: E(%5.5s, %5.5s) expected  point: (%3.3s,%3.3s), but got: (%3.3s,%3.3s)\n",
					testSample.curve.A.String(),
					testSample.curve.B.String(),
					point.X.String(),
					point.Y.String(),
					points[i].X,
					points[i].Y,
				)
			}
		}
	}
}

func TestAddPoints(t *testing.T) {
	for _, testCase := range addPointsTestSamples {
		for _, testValue := range testCase.values {
			result, err := testCase.curve.AddPoints(testValue.input1, testValue.input2)

			if err == nil {
				y2 := big.NewInt(0).Set(result.Y)
				y2.Exp(y2, big.NewInt(2), testCase.curve.M)
				if testCase.curve.CalculateVal(result.X).Cmp(y2) != 0 {
					fmt.Printf("NOTE: For the curve: E(%5.5s, %5.5s) for points (%5.5s, %5.5s) and (%5.5s, %5.5s) addition result infinitely distant point\n",
						testCase.curve.A.String(),
						testCase.curve.B.String(),
						testValue.input1.X.String(),
						testValue.input1.Y.String(),
						testValue.input2.X.String(),
						testValue.input2.Y.String(),
					)
				} else {
					if result.X.Cmp(testValue.output.X) != 0 || result.Y.Cmp(testValue.output.Y) != 0 {
						t.Errorf("For the curve: E(%5.5s, %5.5s) for points (%5.5s, %5.5s) and (%5.5s, %5.5s) addition result expected to be (%5.5s, %5.5s) but got (%5.5s, %5.5s)\n",
							testCase.curve.A.String(),
							testCase.curve.B.String(),
							testValue.input1.X.String(),
							testValue.input1.Y.String(),
							testValue.input2.X.String(),
							testValue.input2.Y.String(),
							testValue.output.X.String(),
							testValue.output.Y.String(),
							result.X.String(),
							result.Y.String(),
						)
					}
				}
			} else {
				fmt.Printf("NOTE: For the curve: E(%5.5s, %5.5s) for points (%5.5s, %5.5s) and (%5.5s, %5.5s) addition result infinitely distant point\n",
					testCase.curve.A.String(),
					testCase.curve.B.String(),
					testValue.input1.X.String(),
					testValue.input1.Y.String(),
					testValue.input2.X.String(),
					testValue.input2.Y.String(),
				)
			}
		}
	}
}

type PointAddition struct {
	p1, p2 Point
	result *Point
}

func TestAdditionTable(t *testing.T) {
	curves := []SimpleCurve{
		{
			A:      big.NewInt(3),
			B:      big.NewInt(2),
			M:      big.NewInt(5),
			points: nil,
		},
		{
			A:      big.NewInt(2),
			B:      big.NewInt(3),
			M:      big.NewInt(7),
			points: nil,
		},
		{
			A:      big.NewInt(4),
			B:      big.NewInt(2),
			M:      big.NewInt(11),
			points: nil,
		},
	}

	for _, c := range curves {

		c.InitPoints()

		var table []PointAddition

		for i := 0; i < len(c.points); i++ {
			for j := i; j < len(c.points); j++ {
				res, err := c.AddPoints(c.points[i], c.points[j])
				if err == nil {
					table = append(table, PointAddition{
						p1:     Point{X: c.points[i].X, Y: c.points[i].Y},
						p2:     Point{X: c.points[j].X, Y: c.points[j].Y},
						result: &res,
					})
				} else {
					table = append(table, PointAddition{
						p1:     Point{X: c.points[i].X, Y: c.points[i].Y},
						p2:     Point{X: c.points[j].X, Y: c.points[j].Y},
						result: nil,
					})
				}
			}
		}

		fmt.Printf("For curve: E(%5.5s) (%5.5s, %5.5s):\n", c.M.String(), c.A.String(), c.B.String())

		fmt.Print("Points: ")
		for _, pt := range c.points {
			fmt.Printf("(%s, %s)", pt.X, pt.Y)
		}
		fmt.Println("")

		for _, p := range table {
			if p.result == nil {
				fmt.Printf("\tp1: (%5.5s, %5.5s) and p2: (%5.5s, %5.5s) addition is infinitely distant point (ord: 1)\n",
					p.p1.X.String(),
					p.p1.Y.String(),
					p.p2.X.String(),
					p.p2.Y.String(),
				)
			} else {
				y2 := big.NewInt(0).Set(p.result.Y)
				y2.Exp(y2, big.NewInt(2), c.M)

				if y2.Cmp(c.CalculateVal(p.result.X)) != 0 {
					fmt.Printf("\tp1: (%5.5s, %5.5s) and p2: (%5.5s, %5.5s) addition is infinitely distant point (ord: 1)\n",
						p.p1.X.String(),
						p.p1.Y.String(),
						p.p2.X.String(),
						p.p2.Y.String(),
					)
				} else {
					fmt.Printf("\tp1: (%5.5s, %5.5s) and p2: (%5.5s, %5.5s) addition is: (%5.5s, %5.5s) (ord: %5.5s)\n",
						p.p1.X.String(),
						p.p1.Y.String(),
						p.p2.X.String(),
						p.p2.Y.String(),
						p.result.X.String(),
						p.result.Y.String(),
						c.PointOrd(*p.result).String(),
					)
				}
			}
		}
	}
}

func TestPointOrd(t *testing.T) {
	curves := []SimpleCurve{
		{
			A: big.NewInt(3),
			B: big.NewInt(2),
			M: big.NewInt(5),
		},
	}

	for _, c := range curves {
		if ord := c.PointOrd(Point{
			X: big.NewInt(1),
			Y: big.NewInt(1),
		}); ord.Cmp(big.NewInt(5)) != 0 {
			t.Errorf("For curve: E(%5.5s) (%5.5s, %5.5s) and point (%5.5s, %5.5s) expected order %5.5s but got %5.5s\n",
				c.M.String(),
				c.A.String(),
				c.B.String(),
				"1",
				"1",
				"5",
				ord.String(),
			)
		}

		if ord := c.PointOrd(Point{
			X: big.NewInt(0),
			Y: big.NewInt(0),
		}); ord.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("For curve: E(%5.5s) (%5.5s, %5.5s) and point (%5.5s, %5.5s) expected order %5.5s but got %5.5s\n",
				c.M.String(),
				c.A.String(),
				c.B.String(),
				"1",
				"1",
				"1",
				ord.String(),
			)
		}
	}
}
