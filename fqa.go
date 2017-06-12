// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fqa

import (
	"fmt"
)
// fqa(a,b,r)(x) = [(ax + r) /b]
type Fqa_t struct {
	a int
	b int
	r int
}
// 'Creer' renvoie une nouvelle forme quasi affine de parametres a, b, c
func Creer(a, b, r int) *Fqa_t {
	return &Fqa_t{a,b,r}
}
// 'Elmt' retourne les parametres a, b, r
func (f *Fqa_t) Elmt() (int, int, int) {
	return f.a, f.b, f.r
}
// 'Valeur' renvoie la valeur en 'n'
func (f *Fqa_t)Valeur(n int) (int) { // v = [(a*n+r)/b]
	q,_ := Divent(f.a*n+f.r,f.b)
	return q
}
// 'Inverse' renvoie la valeur inverse en 'n'
func (f *Fqa_t)Inverse(n int) (int) { // v = [(b*n+b-r-1)/a]
	q,_ := Divent(f.b*n+f.b-f.r-1,f.a)
	return q
}
// 'Div_fqa' retourne l'equivalent du quotient q
// et du reste r de la 'division' de 'n' par 'f'
func (f *Fqa_t) Div_fqa(n int) (q int, r int) {
	q = f.Inverse(n)		// equivalent [n/f]
	r = n - f.Valeur(q)		// equivalent n - f*[n/f]
	return q,r
}
// 'Egal' retourne 'true' les parametres des formes quasi affines 'f' et 't sont respectivement egaux et 'false sinon
func (f *Fqa_t)Egal(t *Fqa_t) bool {
	if nil == t {return nil == f}
	return f.a == t.a && f.b == t.b && f.r == t.r
}
func (f Fqa_t)String() string {
	return fmt.Sprintf("[a:%d b:%d r:%d]",f.a,f.b,f.r)
}

// 'minmax' retourne le statut ok :
// 'false' si le tableau t est vide 
// 'true' sinon, accompagne
// du min et du max du tableau
func minmax(t []int) (bool, int, int) {
	if 1 > len(t) {
		return false,0,0
	}	
	min := t[0]
	max := t[0]
	for _, v := range t {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}	
	return true, min, max
}

// [Droites discretes et calendriers (Albert Troesch)](https://mathinfo.unistra.fr/fileadmin/upload/IREM/Publications/L_Ouvert/n071/o_71_27-42.pdf)
// mise en oeuvre de l'algorithme decrit dans le document cite plus haut
func algo(min int,t []int) ([]int, int, bool) {
	var c, l = make([]int,len(t)), make([]int,len(t))
	j := 0
	l[j] = 0
	echange := false
	for i := range t {
		c[i] = t[i] - min
		if 0 == i && 0 != c[i] {
			echange = true
		}
		if echange {
			c[i] = 1 - c[i]
		}
		l[j]++
		if 1 == c[i] {
			j++
			l[j] = 0
		}
	}
	// le dernier segment n'est pas vide
	if 0 < l[j] {
		l[j]++
		j++
	}
	// on supprime le premier segment
	g := l[0]
	l = l[1:j]
	return l, g, echange
}

// premiere etape de l'algorithme decrit dans le document cite plus haut
func etape1(t []int) (bool, int, int, int) {
	var (
		ok bool
		min, max int
		a, b, r int
	)
	ok, min, max = minmax(t)
	if !ok {
		return false, 0, 0, 0
	}
	if max - min > 1 { // pas une droite discrete
		// il y a au moins deux codes puisque min != max
		// si le dernier code est le min
		// alors on peut (tenter de) le supprimer (segment exterieur)
		// pour recommencer (ex. le 28 ou 29 fevrier)
		if min == t[len(t)-1] {
			return etape1(t[0:len(t)-1])
		}
		return false, 0, 0, 0
	}
	// critere de fin avec succes
	if min == max {
		return true, min, 1, 0 // point 1
	}

	l, g, echange := algo(min,t)

	// on recommence l'etape 1 
	// recursivement jusqu'à ce que  
	// tous les codes soient egaux
	// (min == max critere de fin)
	// ou qu'une erreur se produise
	ok, a, b, r = etape1(l)
	if !ok { // une erreur s'est produite on s'arrete
		return false, 0, 0, 0
	}
	// aucune erreur ne s'est produite
	// on effectue l'etape 2 a rebours
	// en depilant les appels recursifs
	// on vient soit du point 1 plus haut
	// (min == max, critere de fin avec succes)
	// soit du point 2 ci-dessous
	a, b, r = etape2(a,b,r,min,g,echange)
	return true, a, b, r // point 2
}

// deuxieme etape de l'algorithme decrit dans le document cite plus haut
func etape2(a, b, r, p, g int, echange bool) (int, int, int) {

	// Operation 3 (symetrie orthogonale : x' = y ; y' = x)
	a, b, r = b, a, b - r - 1
	if 0 < g {
		r = r - a*g + b
	}

	// Operation 2 (symetrie oblique : x' = x ; y' = x - y)
	if echange {
		a, b, r = b - a, b, b - r - 1
	}

	// Operation 1 (transvection : x' = x ; y' = y - px)
	a, b, r = a + p*b, b, r

	return a, b, r
}

// DEFINITION : si 'f'(x) = [(ax + r) / b] est une forme quasi affine
// alors le code en 'x' = f(x + 1) - f(x)
// 'Codes' renvoie un indicateur et un pointeur *Fqa_t
// si l'indicateur = 'false' le pointeur est nul
// si l'indicateur = 'true' le pointeur represente la forme quasi
// affine obtenue a partir des codes
func Codes(c []int, x0, y0 int) (bool, *Fqa_t) {
	ok, a, b, r := etape1(c)
	if !ok {
		return false, nil
	}
	r += b*y0 - a*x0
	return ok, Creer(a, b, r)
}
