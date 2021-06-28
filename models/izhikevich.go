package models

type IzhikModel struct {
	A      scalar
	B      scalar
	C      scalar
	D      scalar
	Vmax   scalar
	Vreset scalar
}

/*
pars=[0.02      0.2     -65      6       14 ;...    % tonic spiking
      0.02      0.25    -65      6       0.5 ;...   % phasic spiking
      0.02      0.2     -50      2       15 ;...    % tonic bursting
      0.02      0.25    -55     0.05     0.6 ;...   % phasic bursting
      0.02      0.2     -55     4        10 ;...    % mixed mode
      0.01      0.2     -65     8        30 ;...    % spike frequency adaptation
      0.02      -0.1    -55     6        0  ;...    % Class 1
      0.2       0.26    -65     0        0  ;...    % Class 2
      0.02      0.2     -65     6        7  ;...    % spike latency
      0.05      0.26    -60     0        0  ;...    % subthreshold oscillations
      0.1       0.26    -60     -1       0  ;...    % resonator
      0.02      -0.1    -55     6        0  ;...    % integrator
      0.03      0.25    -60     4        0;...      % rebound spike
      0.03      0.25    -52     0        0;...      % rebound burst
      0.03      0.25    -60     4        0  ;...    % threshold variability
      1         1.5     -60     0      -65  ;...    % bistability
        1       0.2     -60     -21      0  ;...    % DAP
      0.02      1       -55     4        0  ;...    % accomodation
     -0.02      -1      -60     8        80 ;...    % inhibition-induced spiking
     -0.026     -1      -45     0        80];       % inhibition-induced bursting

a=pars(1,1);
b=pars(1,2);
c=pars(1,3);
d=pars(1,4);
I=pars(1,5);
*/

func NewIzhikModel() *IzhikModel {
	return &IzhikModel{
		A:      0.02,
		B:      0.2,
		C:      -50.0,
		D:      2.0,
		Vmax:   30.0,
		Vreset: -30.0,
	}
}

func (m *IzhikModel) Step(n *Neuron, input float64) bool {
	dV := 0.04*n.V*n.V + 5*n.V + 140 - n.W + input
	dW := m.A * (m.B*n.V - n.W)
	n.V += dV
	n.W += dW
	if n.V > m.Vmax {
		n.V = m.C
		n.W += m.D
		return true
	}
	return false
}
