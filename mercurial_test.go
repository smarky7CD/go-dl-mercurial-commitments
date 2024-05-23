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

	// check hard tease
	tau := HardTease(&pi0)
	checkTease := VerTease(&c0, &c1, []byte(x), &tau)
	assert.True(t, checkTease, "checkTease should be true.")

	// check open
	checkopen := VerOpen(&h, &c0, &c1, []byte("sam"), &pi0, &pi1)
	assert.True(t, checkopen, "checkOpen should be true.")
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

	// check hard tease
	tau := HardTease(&pi0)
	checkTease := VerTease(&c0, &c1, []byte("jack"), &tau)
	assert.False(t, checkTease, "checkTease should be false.")

	// check open
	checkOpen := VerOpen(&h, &c0, &c1, []byte("jack"), &pi0, &pi1)
	assert.False(t, checkOpen, "checkOpen should be false.")
}

func TestSoftCommit(t *testing.T) {
	var x, y string
	var pi0, pi1 ristretto.Scalar

	x = "sam"
	y = "jack"

	pi0.Rand()
	pi1.Rand()

	// generate soft commitment for "·"
	c0, c1 := SoftCommit(&pi0, &pi1)

	// check soft tease to "jack"
	tau := SoftTease([]byte(y), &pi0, &pi1)
	checkTease := VerTease(&c0, &c1, []byte("jack"), &tau)
	assert.True(t, checkTease, "checkTease should be true.")
	// check soft tease to "sam"
	tau = SoftTease([]byte(x), &pi0, &pi1)
	checkTease = VerTease(&c0, &c1, []byte("sam"), &tau)
	assert.True(t, checkTease, "checkTease should be true.")
}
