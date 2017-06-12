package fqa

func Equal_f(a, b, prec float64) bool {
	if 0 >= prec {
		panic("Invalid negative prec")
	}
	c := b - a
	return -prec < c && c < prec
}

// Divent retourne le quotient q 
// et le reste r toujours positif ou nul
// de la division de n par d 
func Divent(n,d int) (q int, r int) {
	q, r = n/d, n%d
	if 0 > r {
		q--
		r += d
	}
	return q,r
}

// Ent retourne la partie entiere e
// correspondant au plus grand nombre
// entier inferieur ou egal a f
// et la partie fractionnaire r
// toujours positive ou nulle du nombre f
func Ent(f float64) (int, float64) {
	e := int(f)
	r := f -float64(e)
	if 0 > r {
		e--
		r += 1
	}
	return e, r
}

func Corrig_am(a, m int) (int, int) {
	var i int
	switch {
		case m > 12:
			i = (m - 1) / 12
		case m < 1:
			i = m / 12 - 1
		default:
			i = 0	
	}
	a += i
	m -= i * 12
	return a,m
}
func Norm_am(a, m int) (int, int) {
	if 3 > m {
		m += 12
		a--
	}
	return a,m
}
func DNorm_am(a, m int) (int, int) {
	if 12 < m {
		m -= 12
		a++
	}
	return a, m
}

