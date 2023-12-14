package model_test

// Basic imports
import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"shapesBasicLevel/shapes/model"
	"testing"
)

type ShapesTestSuite struct {
	suite.Suite
}
type testInput struct {
	a, b, expected float64
}

func (suite *ShapesTestSuite) Test_ellipse_area() {
	cases := []testInput{
		{15, 7, 82.46680716750001},
		{10, 20, 157.0796327},
		{0, 0, 0},
	}
	for _, testCase := range cases {
		ellipse := &model.Ellipse{
			Id:     uuid.New(),
			RadioA: testCase.a,
			RadioB: testCase.b,
		}
		suite.InDelta(testCase.expected, ellipse.CalculateArea(), 1e-6)
	}

}

func (suite *ShapesTestSuite) Test_rectangle_area() {
	cases := []testInput{
		{15, 7, 105},
		{1.366, 20, 27.32},
		{0, 0, 0},
	}
	for _, testCase := range cases {
		rectangle := &model.Rectangle{
			Id:     uuid.New(),
			Length: testCase.a,
			Width:  testCase.b,
		}
		suite.Equal(testCase.expected, rectangle.CalculateArea())
	}
}

func (suite *ShapesTestSuite) Test_triangle_area() {
	cases := []testInput{
		{15, 7, 52.5},
		{10, 20, 100},
		{0, 0, 0},
	}
	for _, testCase := range cases {
		triangle := &model.Triangle{
			Id:     uuid.New(),
			Base:   testCase.a,
			Height: testCase.b,
		}
		suite.Equal(testCase.expected, triangle.CalculateArea())
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ShapesTestSuite))
}
