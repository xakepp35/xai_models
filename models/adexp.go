package models

import "math"

type AdExModel struct {
	GL     scalar
	EL     scalar
	Slope  scalar
	VT     scalar
	A      scalar
	B      scalar
	Vmax   scalar
	Vreset scalar
}

func NewAdExModel() *AdExModel {
	return &AdExModel{
		GL:     0.1,
		EL:     -0.1,
		Slope:  0.35,
		VT:     0.0,
		A:      1.01,
		B:      -0.6,
		Vmax:   2.0,
		Vreset: -0.2,
	}
}

func (m *AdExModel) Step(n *Neuron, input float64) bool {
	Vrest := n.V - m.EL
	dV := -m.GL*Vrest + m.Slope*math.Exp(n.V/m.Slope) - n.W + input
	dW := m.A*Vrest - n.W
	n.V += dV
	n.W += dW
	if n.V > m.Vmax {
		n.V = m.Vreset
		n.W += m.B
		return true
	}
	return false
}
