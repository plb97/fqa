// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fqa

// 'Egal' retourne 'true' si les deux nombres 'a' et 'b' different strictement de moins que 'prec'
func Egal_f(a, b, prec float64) bool {
	if 0 >= prec {
		panic("Precision negative ou nulle invalide")
	}
	c := b - a
	return -prec < c && c < prec
}

// 'Divent' retourne le quotient q
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

// 'Ent' retourne la partie entiere 'e'
// correspondant au plus grand nombre
// entier inferieur ou egal a 'f'
// et la partie fractionnaire 'r'
// toujours positive ou nulle
// f = e + r
// e <= f
// 0 <= r < 1
func Ent(f float64) (int, float64) {
	e := int(f)
	r := f -float64(e)
	if 0 > r {
		e--
		r += 1
	}
	return e, r
}
// 'Corrig_am' assure que le mois 'm' soit bien compris entre '1 et '12'
// et ajuste l'annee 'a' en consequence
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
// 'Norm_am' assure que 'm' soit bien entre '3' et '14'
// et ajuste l'annee 'a' en consequence
func Norm_am(a, m int) (int, int) {
	if 3 > m {
		m += 12
		a--
	}
	return a,m
}
// 'DNorm_am' assure que 'm' soit bien entre '1' et '12'
// et ajuste l'annee 'a' en consequence
func DNorm_am(a, m int) (int, int) {
	if 12 < m {
		m -= 12
		a++
	}
	return a, m
}

