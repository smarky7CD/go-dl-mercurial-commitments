package mercurial

import (
	"math/big"

	"github.com/bwesterb/go-ristretto"
)

// prime order of the Edwards25519 base point
var n25519, _ = new(big.Int).SetString("7237005577332262213973186563042994240857116359379907606001950938285454250989", 10)

// g is Edwards25519 basepoint
// h (returned here) is a randomly selected point
func generatePublicParameters() ristretto.Point {
	var random ristretto.Scalar
	var h ristretto.Point
	random.Rand()
	h.ScalarMultBase(&random)
	return h
}

func hardCommit(h *ristretto.Point, x []byte, r0 *ristretto.Scalar, r1 *ristretto.Scalar) (ristretto.Point, ristretto.Point) {
	var c0, t01, t02, t03, c1 ristretto.Point
	var xScalar ristretto.Scalar
	xScalar.Derive(x)
	t01.ScalarMultBase(&xScalar)
	t02.ScalarMult(h, r1)
	t03.ScalarMult(&t02, r0)
	c0.Add(&t01, &t03)
	c1.Set(&t02)
	return c0, c1
}

func softCommit(r0 *ristretto.Scalar, r1 *ristretto.Scalar) (ristretto.Point, ristretto.Point) {
	var c0, c1 ristretto.Point
	c0.ScalarMultBase(r0)
	c1.ScalarMultBase(r1)
	return c0, c1
}

func hardTease(r0 *ristretto.Scalar) ristretto.Scalar {
	return *r0
}

func softTease(x []byte, r0 *ristretto.Scalar, r1 *ristretto.Scalar) ristretto.Scalar {
	var t, xScalar, subR0X ristretto.Scalar
	var tINT, modINTR1 big.Int
	xScalar.Derive(x)
	subR0X.Sub(r0, &xScalar)
	modINTR1.ModInverse(r1.BigInt(), n25519)
	tINT.Mul(subR0X.BigInt(), &modINTR1)
	t.SetBigInt(&tINT)
	return t
}

func verTease(c0 *ristretto.Point, c1 *ristretto.Point, x []byte, tau *ristretto.Scalar) bool {
	var cC, t0, t1 ristretto.Point
	var xScalar ristretto.Scalar
	xScalar.Derive(x)
	t0.ScalarMultBase(&xScalar)
	t1.ScalarMult(c1, tau)
	cC.Add(&t0, &t1)
	return cC.Equals(c0)
}

func open(r0 *ristretto.Scalar, r1 *ristretto.Scalar) (ristretto.Scalar, ristretto.Scalar) {
	return *r0, *r1
}

func verOpen(h *ristretto.Point, c0 *ristretto.Point, c1 *ristretto.Point, x []byte, pi0 *ristretto.Scalar, pi1 *ristretto.Scalar) bool {
	var cC0, t00, t01, cC1 ristretto.Point
	var xScalar ristretto.Scalar
	xScalar.Derive(x)
	t00.ScalarMultBase(&xScalar)
	t01.ScalarMult(c1, pi0)
	cC0.Add(&t00, &t01)
	cC1.ScalarMult(h, pi1)
	return cC0.Equals(c0) && cC1.Equals(c1)
}
