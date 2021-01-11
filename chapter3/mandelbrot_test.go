package chapter3

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -run TestComplexAtPixel -v
func TestComplexAtPixel(t *testing.T) {
	t.Parallel()
	t.Run("0+0i=>1 pixels rect", func(t *testing.T) {
		z := complex(0, 0)
		zb := complexAtPixel(z)
		z1 := zb(-1, 0)
		z2 := zb(1, 0)
		z3 := zb(0, -1)
		z4 := zb(0, 1)
		t.Logf("0+0i: z1: %v, z2: %v, z3: %v, z4: %v", z1, z2, z3, z4)
	})
	t.Run("-2-2i=>1024 pixels rect", func(t *testing.T) {
		z := complex(-2, -2)
		zb := complexAtPixel(z)
		z1 := zb(0, 0)
		z2 := zb(0, MBHeight)
		z3 := zb(MBWidth, MBHeight)
		z4 := zb(MBWidth, 0)
		t.Logf("-2-2i: z1: %v, z2: %v, z3: %v, z4: %v", z1, z2, z3, z4)
	})

}

// go test -run TestScale -v
func TestScale(t *testing.T) {
	x := scaleX(0)
	y := scaleY(0)
	z1 := complex(x, y)
	x = scaleX(MBHeight)
	y = scaleY(0)
	z2 := complex(x, y)
	x = scaleX(MBHeight)
	y = scaleY(MBWidth)
	z3 := complex(x, y)
	x = scaleX(0)
	y = scaleY(MBWidth)
	z4 := complex(x, y)
	t.Logf("z1: %v, z2: %v, z3: %v, z4: %v", z1, z2, z3, z4)
}

// go test -run TestMBGraphHandler -v
func TestMBGraphHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/mb", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MBGraphHandler)
	handler.ServeHTTP(rr, r)
	if rr.Code != http.StatusOK {
		t.Errorf("MBGraphHandler status got: %d, want: %d", rr.Code, http.StatusOK)
	}
}

// go test -run TestMBGraphHandler -v
func TestMBGraphBWHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/mb", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MBGraphBWHandler)
	handler.ServeHTTP(rr, r)
	if rr.Code != http.StatusOK {
		t.Errorf("MBGraphBWHandler status got: %d, want: %d", rr.Code, http.StatusOK)
	}
}
