# go-dl-mercurial-commitments

A pure Go implementation of discrete log based mercurial commitments (using Ristretto prime-order group over Edwards25519). Based on the paper *Mercurial Commitments with Applications to Zero-Knowledge Sets* by Chase et al. [[CHLMR05](https://cs.brown.edu/~mchase/papers/merc.pdf)].


## Installing and Usage

To install the package:

```shell
go get -u github.com/smarky7CD/go-dl-mercurial-commitments
```

The package can be imported and used as such:

```go
import(
    "fmt"
    "github.com/smarky7cd/go-dl-mercurial-commitments"
)

func MCexample() {

    var x,y string
	var pi0, pi1 ristretto.Scalar

    // hard commitments

	x = "sam"
    y = "jack"
	h := GeneratePublicParameters()
    // generate hard commitment for "sam"
	c0, c1 := HardCommit(&h, []byte(x), &pi0, &pi1)

    // valid tease and open

	// check hard tease
	tau := HardTease(&pi0)
	checktease := VerTease(&c0, &c1, []byte(x), &tau)
    fmt.Println("Hard tease check is: ", checktease) //true

	// check open
	checkopen := VerOpen(&h, &c0, &c1, []byte("sam"), &pi0, &pi1)
    fmt.Println("Open check is: ", checkopen) //true

    // invalid tease and open

    // check hard tease
	tau = HardTease(&pi0)
	checkTease = VerTease(&c0, &c1, []byte("jack"), &tau)
	fmt.Println("Hard tease check is: ", checktease) //false

	// check open
	checkOpen = VerOpen(&h, &c0, &c1, []byte("jack"), &pi0, &pi1)
	fmt.Println("Open check is: ", checkopen) //true //false

    // soft commitments
    // generate soft commitment for "·"
	c0, c1 = SoftCommit(&pi0, &pi1)

	// check soft tease to "jack"
	tau = SoftTease([]byte(y), &pi0, &pi1)
	checktease = VerTease(&c0, &c1, []byte("jack"), &tau)
    fmt.Println("Soft tease check is: ", checktease) // true
    // check soft tease for "sam"
	tau = SoftTease([]byte(x), &pi0, &pi1)
	checktease = VerTease(&c0, &c1, []byte("sam"), &tau)
    fmt.Println("Soft tease check is: ", checktease) // true
}


```

## Tests

We provide a number correctness tests for a valid hard commitment, an invalid hard commitment, and a soft commitment.

They can be ran by using:

```shell
go test -v
```

## References

- [CHLMR05]: Chase, Melissa, et al. "Mercurial commitments with applications to zero-knowledge sets." Advances in Cryptology–EUROCRYPT 2005: 24th Annual International Conference on the Theory and Applications of Cryptographic Techniques, Aarhus, Denmark, May 22-26, 2005. Proceedings 24. Springer Berlin Heidelberg, 2005.