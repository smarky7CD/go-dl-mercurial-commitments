package mercurial

import (
	"math/big"

	"github.com/bwesterb/go-ristretto"
)

// prime order of the Edwards25519 base point
var n25519, _ = new(big.Int).SetString("7237005577332262213973186563042994240857116359379907606001950938285454250989", 10)

// g is Edwards25519 basepoint
// h (returned here) is a randomly selected point
func GeneratePublicParameters() ristretto.Point {
	var random ristretto.Scalar
	var h ristretto.Point
	random.Rand()
	h.ScalarMultBase(&random)
	return h
}

// Computes a hard commitment (c0,c1) to x given random scalars(r0,r1) and  h.
func HardCommit(h *ristretto.Point, x []byte, r0 *ristretto.Scalar, r1 *ristretto.Scalar) (ristretto.Point, ristretto.Point) {
	var c0, t01, t02, t03, c1 ristretto.Point
	var xscalar ristretto.Scalar
	xscalar.Derive(x)
	t01.ScalarMultBase(&xscalar)
	t02.ScalarMult(h, r1)
	t03.ScalarMult(&t02, r0)
	c0.Add(&t01, &t03)
	c1.Set(&t02)
	return c0, c1
}

// Computes soft commitment (c0,c1) given random scalars(r0,r1).
func SoftCommit(r0 *ristretto.Scalar, r1 *ristretto.Scalar) (ristretto.Point, ristretto.Point) {
	var c0, c1 ristretto.Point
	c0.ScalarMultBase(r0)
	c1.ScalarMultBase(r1)
	return c0, c1
}

// Returns r0 as hard tease of a hard commitment.
func HardTease(r0 *ristretto.Scalar) ristretto.Scalar {
	return *r0
}

// Returns t as soft tease to value x of a soft commitment.
func SoftTease(x []byte, r0 *ristretto.Scalar, r1 *ristretto.Scalar) ristretto.Scalar {
	var t, xscalar, subr0x ristretto.Scalar
	var tint, modintr1 big.Int
	xscalar.Derive(x)
	subr0x.Sub(r0, &xscalar)
	modintr1.ModInverse(r1.BigInt(), n25519)
	tint.Mul(subr0x.BigInt(), &modintr1)
	t.SetBigInt(&tint)
	return t
}

// Verifies a t value .
// Returns true if valid, false otherwise.
func VerTease(c0 *ristretto.Point, c1 *ristretto.Point, x []byte, tau *ristretto.Scalar) bool {
	var cc, t0, t1 ristretto.Point
	var xscalar ristretto.Scalar
	xscalar.Derive(x)
	t0.ScalarMultBase(&xscalar)
	t1.ScalarMult(c1, tau)
	cc.Add(&t0, &t1)
	return cc.Equals(c0)
}

// Opens a hard commitment (returns r0,r1).
func Open(r0 *ristretto.Scalar, r1 *ristretto.Scalar) (ristretto.Scalar, ristretto.Scalar) {
	return *r0, *r1
}

// Verifies the opening of a hard commitment for a value x.
func VerOpen(h *ristretto.Point, c0 *ristretto.Point, c1 *ristretto.Point, x []byte, pi0 *ristretto.Scalar, pi1 *ristretto.Scalar) bool {
	var cc0, t00, t01, cc1 ristretto.Point
	var xscalar ristretto.Scalar
	xscalar.Derive(x)
	t00.ScalarMultBase(&xscalar)
	t01.ScalarMult(c1, pi0)
	cc0.Add(&t00, &t01)
	cc1.ScalarMult(h, pi1)
	return cc0.Equals(c0) && cc1.Equals(c1)
}
