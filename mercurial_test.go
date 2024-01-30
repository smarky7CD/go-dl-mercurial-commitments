package mercurial

import (
	"testing"

	"github.com/bwesterb/go-ristretto"
	"github.com/stretchr/testify/assert"
)

func TestHardCommit(t *testing.T) {
	var x string
	var pi0, pi1 ristretto.Scalar

	x = "sam"
	h := generatePublicParameters()

	pi0.Rand()
	pi1.Rand()

	// generate hard commitment for "sam"
	c0, c1 := hardCommit(&h, []byte(x), &pi0, &pi1)

	// check Hard tease
	tau := hardTease(&pi0)
	checkTease := verTease(&c0, &c1, []byte(x), &tau)
	assert.True(t, checkTease, "Tease should be true.")

	// check Open
	checkOpen := verOpen(&h, &c0, &c1, []byte("sam"), &pi0, &pi1)
	assert.True(t, checkOpen, "v0 hould be true.")
}

func TestBadHardCommit(t *testing.T) {
	var x string
	var pi0, pi1 ristretto.Scalar

	x = "sam"
	h := generatePublicParameters()

	pi0.Rand()
	pi1.Rand()

	// generate hard commitment for "sam"
	c0, c1 := hardCommit(&h, []byte(x), &pi0, &pi1)

	// check Hard tease
	tau := hardTease(&pi0)
	checkTease := verTease(&c0, &c1, []byte("jack"), &tau)
	assert.False(t, checkTease, "Tease should be false.")

	// check Open
	checkOpen := verOpen(&h, &c0, &c1, []byte("jack"), &pi0, &pi1)
	assert.False(t, checkOpen, "checkOpen should be false.")
}

func TestSoftCommit(t *testing.T) {
	var x string
	var pi0, pi1 ristretto.Scalar

	x = "jack"

	pi0.Rand()
	pi1.Rand()

	// generate hard commitment for "sam"
	c0, c1 := softCommit(&pi0, &pi1)

	// check soft tease
	tau := softTease([]byte(x), &pi0, &pi1)
	checkTease := verTease(&c0, &c1, []byte("jack"), &tau)
	assert.True(t, checkTease, "Tease should be true.")
}
