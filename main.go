package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"shapesBasicLevel/shapes/dto"
	"shapesBasicLevel/shapes/model"
)

func main() {
	shapes, err := load("data/shapeLibrary.json")
	if err != nil {
		panic("Could not load shapes from file - " + err.Error())
	}
	for _, s := range shapes {
		s.Print()
		fmt.Printf("Area: %f\n\n", s.CalculateArea())
	}
}

func load(path string) (map[string]model.Shape, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("Unable to read file - " + err.Error())
	}

	var shapesFromFile []dto.ShapeFile
	err = json.Unmarshal(file, &shapesFromFile)
	if err != nil {
		return nil, errors.New("Unable to deserialize file - " + err.Error())
	}

	shapes := make(map[string]model.Shape)
	for _, s := range shapesFromFile {
		id, err := uuid.Parse(s.Id)
		if err != nil {
			return nil, errors.New("Unable to parse shape id - " + err.Error())
		}
		switch s.Type {
		case "ELLIPSE":
			shapes[s.Id] = model.Ellipse{
				Id:     id,
				RadioA: s.A,
				RadioB: s.B,
			}
		case "TRIANGLE":
			shapes[s.Id] = model.Triangle{
				Id:     id,
				Base:   s.A,
				Height: s.B,
			}

		case "RECTANGLE":
			shapes[s.Id] = model.Rectangle{
				Id:     id,
				Length: s.A,
				Width:  s.B,
			}

		default:
			return nil, errors.New("unrecognized shape type " + s.Type)
		}
	}
	return shapes, nil
}
