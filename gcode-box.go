package main

import (
	cnc "github.com/benishor/go-cnc"
)

func pereteTransversal(tool cnc.Tool, settings cnc.MachineSettings) *cnc.Program {
	path := new(cnc.Path)
	path.Add(
		cnc.Point{0, 0},
		cnc.Point{0, 5},
		cnc.Point{-1.5, 5},
		cnc.Point{-1.5, 20},
		cnc.Point{0, 20},
		cnc.Point{0, 25},
		cnc.Point{10, 25},
		cnc.Point{10, 30},
		cnc.Point{50, 30},
		cnc.Point{50, 25},
		cnc.Point{60, 25},
		cnc.Point{60, 20},
		cnc.Point{61.5, 20},
		cnc.Point{61.5, 5},
		cnc.Point{60, 5},
		cnc.Point{60, 0},
		cnc.Point{50, 0},
		cnc.Point{50, -1.5},
		cnc.Point{10, -1.5},
		cnc.Point{10, 0},
	)
	path.Translate(1.5, 1.5)

	settings.CuttingType = cnc.CUTTING_TYPE_OUTSIDE
	program := new(cnc.Program)
	program.Add(cnc.NewProfileOperation(*path, tool, settings))

	return program
}

func pereteLateral(tool cnc.Tool, settings cnc.MachineSettings) *cnc.Program {
	outerPath := cnc.Path{}
	outerPath.Add(
		cnc.Point{0, 0},
		cnc.Point{0, 5},
		cnc.Point{1.5, 5},
		cnc.Point{1.5, 20},
		cnc.Point{0, 20},
		cnc.Point{0, 25},

		cnc.Point{10, 25},
		cnc.Point{10, 30},
		cnc.Point{90, 30},
		cnc.Point{90, 25},

		cnc.Point{100, 25},
		cnc.Point{100, 20},
		cnc.Point{98.5, 20},
		cnc.Point{98.5, 5},
		cnc.Point{100, 5},
		cnc.Point{100, 0},

		cnc.Point{90, 0},
		cnc.Point{90, -1.5},
		cnc.Point{10, -1.5},
		cnc.Point{10, 0},
	)

	fanta1 := cnc.Path{}
	fanta1.Add(
		cnc.Point{33, 4.7},
		cnc.Point{33, 20.3},
		cnc.Point{35, 20.3},
		cnc.Point{35, 4.8},
	)

	fanta2 := cnc.Path{}
	fanta2.Add(
		cnc.Point{33, 4.7},
		cnc.Point{33, 20.3},
		cnc.Point{35, 20.3},
		cnc.Point{35, 4.7},
	)
	fanta2.Translate(33.3, 0)

	program := new(cnc.Program)


	settings.CuttingType = cnc.CUTTING_TYPE_INSIDE
	program.Add(cnc.NewProfileOperation(fanta1, tool, settings))
	program.Add(cnc.NewProfileOperation(fanta2, tool, settings))

	settings.CuttingType = cnc.CUTTING_TYPE_OUTSIDE
	program.Add(cnc.NewProfileOperation(outerPath, tool, settings))

	return program
}

func capac(tool cnc.Tool, settings cnc.MachineSettings) *cnc.Program {
	outerPath := cnc.Path{}
	outerPath.Add(
		cnc.Point{0, 0},
		cnc.Point{0, 11.5},
		cnc.Point{1.5, 11.5},
		cnc.Point{1.5, 51.5},
		cnc.Point{0, 51.5},
		cnc.Point{0, 63},

		cnc.Point{10, 63},
		cnc.Point{10, 61.5},
		cnc.Point{90, 61.5},
		cnc.Point{90, 63},

		cnc.Point{100, 63},
		cnc.Point{100, 51.5},
		cnc.Point{98.5, 51.5},
		cnc.Point{98.5, 11.5},
		cnc.Point{100, 11.5},
		cnc.Point{100, 0},

		cnc.Point{90, 0},
		cnc.Point{90, 1.5},
		cnc.Point{10, 1.5},
		cnc.Point{10, 0},
	)

	fanta1 := cnc.Path{}
	fanta1.Add(
		cnc.Point{33, 11},
		cnc.Point{33, 52},
		cnc.Point{35, 52},
		cnc.Point{35, 11},
	)

	fanta2 := cnc.Path{}
	fanta2.Add(
		cnc.Point{33, 11},
		cnc.Point{33, 52},
		cnc.Point{35, 52},
		cnc.Point{35, 11},
	)
	fanta2.Translate(33.3, 0)


	//outerPath.Transpose()
	//fanta1.Transpose()
	//fanta2.Transpose()

	program := new(cnc.Program)


	settings.CuttingType = cnc.CUTTING_TYPE_INSIDE
	program.Add(cnc.NewProfileOperation(fanta1, tool, settings))
	program.Add(cnc.NewProfileOperation(fanta2, tool, settings))

	settings.CuttingType = cnc.CUTTING_TYPE_OUTSIDE
	program.Add(cnc.NewProfileOperation(outerPath, tool, settings))

	return program
}


func main() {
	tool := cnc.Tool{
		Diameter: 1.2,
		Name: "Freza de 1.2mm"}

	settings := cnc.MachineSettings{
		PlungeFeedrate: 400,
		MovementFeedrate: 400,
		SafeZ: 4,
		CuttingDepth: -2.2}

	p := pereteLateral(tool, settings)
	p.WriteTo("./perete-lateral.gcode")

	p = pereteTransversal(tool, settings)
	p.WriteTo("./perete-transversal.gcode")

	p = capac(tool, settings)
	p.WriteTo("./capac.gcode")
}