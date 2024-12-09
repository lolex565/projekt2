package menu

import (
	"projekt2/graph"
	"projekt2/solver/bf"
	"projekt2/solver/bnb"
	"projekt2/solver/dp"
	"projekt2/solver/gr"
	"projekt2/solver/sa"
	"projekt2/solver/ts"
)

type Menu struct {
	bfATSPSolver  bf.BFATSPSolver
	bnbATSPSolver bnb.BNBATSPSolver
	dpATSPSolver  dp.DPATSPSolver
	grATSPSolver  gr.GRATSPSolver
	saATSPSolver  sa.SaATSPSolver
	tsATSPSolver  ts.TsATSPSolver
	graph         graph.Graph
}

func (m *Menu) BfATSPSolver() bf.BFATSPSolver {
	return m.bfATSPSolver
}

func (m *Menu) GetBfATSPSolverPtr() *bf.BFATSPSolver {
	return &m.bfATSPSolver
}

func (m *Menu) SetBfATSPSolver(bfATSPSolver bf.BFATSPSolver) {
	m.bfATSPSolver = bfATSPSolver
}

func (m *Menu) BnbATSPSolver() bnb.BNBATSPSolver {
	return m.bnbATSPSolver
}

func (m *Menu) GetBnbATSPSolverPtr() *bnb.BNBATSPSolver {
	return &m.bnbATSPSolver
}

func (m *Menu) SetBnbATSPSolver(bnbATSPSolver bnb.BNBATSPSolver) {
	m.bnbATSPSolver = bnbATSPSolver
}

func (m *Menu) DpATSPSolver() dp.DPATSPSolver {
	return m.dpATSPSolver
}

func (m *Menu) GetDpATSPSolverPtr() *dp.DPATSPSolver {
	return &m.dpATSPSolver
}

func (m *Menu) SetDpATSPSolver(dpATSPSolver dp.DPATSPSolver) {
	m.dpATSPSolver = dpATSPSolver
}

func (m *Menu) GrATSPSolver() gr.GRATSPSolver {
	return m.grATSPSolver
}

func (m *Menu) GetGrATSPSolverPtr() *gr.GRATSPSolver {
	return &m.grATSPSolver
}

func (m *Menu) SetGrATSPSolver(grATSPSolver gr.GRATSPSolver) {
	m.grATSPSolver = grATSPSolver
}

func (m *Menu) SaATSPSolver() sa.SaATSPSolver {
	return m.saATSPSolver
}

func (m *Menu) GetSaATSPSolverPtr() *sa.SaATSPSolver {
	return &m.saATSPSolver
}

func (m *Menu) SetSaATSPSolver(saATSPSolver sa.SaATSPSolver) {
	m.saATSPSolver = saATSPSolver
}

func (m *Menu) TsATSPSolver() ts.TsATSPSolver {
	return m.tsATSPSolver
}

func (m *Menu) GetTsATSPSolverPtr() *ts.TsATSPSolver {
	return &m.tsATSPSolver
}

func (m *Menu) SetTsATSPSolver(tsATSPSolver ts.TsATSPSolver) {
	m.tsATSPSolver = tsATSPSolver
}

func (m *Menu) Graph() graph.Graph {
	return m.graph
}

func (m *Menu) SetGraph(graph graph.Graph) {
	m.graph = graph
}

func NewMenu() *Menu {
	return &Menu{}
}

func NewDefaultMenu(g graph.Graph) *Menu {
	return &Menu{
		graph: g,
	}
}
