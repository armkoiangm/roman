package roman

import "testing"

func TestConvertI(t *testing.T) {
	c := NewConverter()

	n := c.Convert("I")

	if n != 1 {
		t.Errorf("I converts to 1, however %v", n)
	}
}

func TestConvertII(t *testing.T) {
	c := NewConverter()

	n := c.Convert("II")

	if n != 2 {
		t.Errorf("II converts to 2, however %v", n)
	}
}

func TestConvertIV(t *testing.T) {
	c := NewConverter()

	n := c.Convert("IV")

	if n != 4 {
		t.Errorf("IV converts to 4, however %v", n)
	}
}

func TestConvertV(t *testing.T) {
	c := NewConverter()

	n := c.Convert("V")

	if n != 5 {
		t.Errorf("V converts to 5, however %v", n)
	}
}

func TestConvertVI(t *testing.T) {
	c := NewConverter()

	n := c.Convert("VI")

	if n != 6 {
		t.Errorf("VI converts to 6, however %v", n)
	}
}

func TestConvertIX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("IX")

	if n != 9 {
		t.Errorf("IX converts to 9, however %v", n)
	}
}

func TestConvertX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("X")

	if n != 10 {
		t.Errorf("X converts to 10, however %v", n)
	}
}

func TestConvertXX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("XX")

	if n != 20 {
		t.Errorf("XX converts to 20, however %v", n)
	}
}

func TestConvertXXX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("XXX")

	if n != 30 {
		t.Errorf("XXX converts to 30, however %v", n)
	}
}

func TestConvertXL(t *testing.T) {
	c := NewConverter()

	n := c.Convert("XL")

	if n != 40 {
		t.Errorf("XL converts to 40, however %v", n)
	}
}

func TestConvertL(t *testing.T) {
	c := NewConverter()

	n := c.Convert("L")

	if n != 50 {
		t.Errorf("L converts to 50, however %v", n)
	}
}

func TestConvertLX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("LX")

	if n != 60 {
		t.Errorf("LX converts to 60, however %v", n)
	}
}

func TestConvertLXX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("LXX")

	if n != 70 {
		t.Errorf("LXX converts to 70, however %v", n)
	}
}

func TestConvertXC(t *testing.T) {
	c := NewConverter()

	n := c.Convert("XC")

	if n != 90 {
		t.Errorf("XC converts to 90, however %v", n)
	}
}

func TestConvertC(t *testing.T) {
	c := NewConverter()

	n := c.Convert("C")

	if n != 100 {
		t.Errorf("C converts to 100, however %v", n)
	}
}

func TestConvertCC(t *testing.T) {
	c := NewConverter()

	n := c.Convert("CC")

	if n != 200 {
		t.Errorf("CC converts to 200, however %v", n)
	}
}

func TestConvertCD(t *testing.T) {
	c := NewConverter()

	n := c.Convert("CD")

	if n != 400 {
		t.Errorf("CD converts to 400, however %v", n)
	}
}

func TestConvertD(t *testing.T) {
	c := NewConverter()

	n := c.Convert("D")

	if n != 500 {
		t.Errorf("D converts to 500, however %v", n)
	}
}

func TestConvertDC(t *testing.T) {
	c := NewConverter()

	n := c.Convert("DC")

	if n != 600 {
		t.Errorf("DC converts to 600, however %v", n)
	}
}

func TestConvertCM(t *testing.T) {
	c := NewConverter()

	n := c.Convert("CM")

	if n != 900 {
		t.Errorf("CM converts to 900, however %v", n)
	}
}

func TestConvertM(t *testing.T) {
	c := NewConverter()

	n := c.Convert("M")

	if n != 1000 {
		t.Errorf("M converts to 1000, however %v", n)
	}
}

func TestConvertMM(t *testing.T) {
	c := NewConverter()

	n := c.Convert("MM")

	if n != 2000 {
		t.Errorf("MM converts to 2000, however %v", n)
	}
}

func TestConvertXXXIX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("XXXIX")

	if n != 39 {
		t.Errorf("XXXIX converts to 39, however %v", n)
	}
}

func TestConvertCCXLVI(t *testing.T) {
	c := NewConverter()

	n := c.Convert("CCXLVI")

	if n != 246 {
		t.Errorf("CCXLVI converts to 246, however %v", n)
	}
}

func TestConvertDCCLXXXIX(t *testing.T) {
	c := NewConverter()

	n := c.Convert("DCCLXXXIX")

	if n != 789 {
		t.Errorf("DCCLXXXIX converts to 789, however %v", n)
	}
}

func TestConvertMMCDXXI(t *testing.T) {
	c := NewConverter()

	n := c.Convert("MMCDXXI")

	if n != 2421 {
		t.Errorf("MMCDXXI converts to 2421, however %v", n)
	}
}

func TestConvertMCMLIV(t *testing.T) {
	c := NewConverter()

	n := c.Convert("MCMLIV")

	if n != 1954 {
		t.Errorf("MCMLIV converts to 1954, however %v", n)
	}
}

func TestConvertMLXVI(t *testing.T) {
	c := NewConverter()

	n := c.Convert("MLXVI")

	if n != 1066 {
		t.Errorf("MLXVI converts to 1066, however %v", n)
	}
}

func TestConvertMLXIV(t *testing.T) {
	c := NewConverter()

	n := c.Convert("MLXIV")

	if n != 1064 {
		t.Errorf("MLXIV converts to 1064, however %v", n)
	}
}

func TestConvertMLxiV(t *testing.T) {
	c := NewConverter()

	n := c.Convert("MLxiV")

	if n != 1064 {
		t.Errorf("MLxiV converts to 1064, however %v", n)
	}
}
