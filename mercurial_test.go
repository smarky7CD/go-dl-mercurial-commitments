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
	h := GeneratePublicParameters()

	pi0.Rand()
	pi1.Rand()

	// generate hard commitment for "sam"
	c0, c1 := HardCommit(&h, []byte(x), &pi0, &pi1)

	// check Hard tease
	tau := HardTease(&pi0)
	checktease := VerTease(&c0, &c1, []byte(x), &tau)
	assert.True(t, checktease, "Tease should be true.")

	// check Open
	checkopen := VerOpen(&h, &c0, &c1, []byte("sam"), &pi0, &pi1)
	assert.True(t, checkopen, "v0 hould be true.")
}

func TestBadHardCommit(t *testing.T) {
	var x string
	var pi0, pi1 ristretto.Scalar

	x = "sam"
	h := GeneratePublicParameters()

	pi0.Rand()
	pi1.Rand()

	// generate hard commitment for "sam"
	c0, c1 := HardCommit(&h, []byte(x), &pi0, &pi1)

	// check Hard tease
	tau := HardTease(&pi0)
	checkTease := VerTease(&c0, &c1, []byte("jack"), &tau)
	assert.False(t, checkTease, "Tease should be false.")

	// check Open
	checkOpen := VerOpen(&h, &c0, &c1, []byte("jack"), &pi0, &pi1)
	assert.False(t, checkOpen, "checkOpen should be false.")
}

func TestSoftCommit(t *testing.T) {
	var x string
	var pi0, pi1 ristretto.Scalar

	x = "jack"

	pi0.Rand()
	pi1.Rand()

	// generate soft commitment for "·"
	c0, c1 := SoftCommit(&pi0, &pi1)

	// check soft tease to "jack"
	tau := SoftTease([]byte(x), &pi0, &pi1)
	checktease := VerTease(&c0, &c1, []byte("jack"), &tau)
	assert.True(t, checktease, "Tease should be true.")
}
