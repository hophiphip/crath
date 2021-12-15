package ellipticcurves

import (
	"math/big"
	"testing"
)

type InputOutput struct {
	input, output *big.Int
}

type CalculateTestCase struct {
	curve  SimpleCurve
	values []InputOutput
}

type GetPointsTestCase struct {
	curve  SimpleCurve
	points []Point
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

func TestCalculateVal(t *testing.T) {
	buffer := big.NewInt(0)
	for _, testSample := range curveTestSamples {
		for _, testValue := range testSample.values {
			buffer.Set(testSample.curve.calculateVal(testValue.input))
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

func TestGetPoints(t *testing.T) {
	for _, testSample := range pointsTestSamples {
		points := testSample.curve.getPoints()

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
