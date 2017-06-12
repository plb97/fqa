// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fqa

import (
	"testing"
)

type test_Fqa_t struct {
	i, n int
	f *Fqa_t
}
var tt = []test_Fqa_t {
		{i:20,n:2451605,f:Creer(146097,4,6884480)},	// siecle calendier gregorien
		{i:16,n:5844,f:Creer(1461,4,0)},		// annee calendier gregorien
		{i:7,n:122,f:Creer(153,5,-457)},		// mois calendrier gregorien
		{i:8,n:7,f:Creer(1,1,-1)},			// jour calendrier gregorien
}

func Test_fqa_valeur(t *testing.T) {
	test := "valeur"
	for _, v := range tt {
		obtenu := v.f.Valeur(v.i)
		attendu := v.n
		if obtenu != attendu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}	
}

func Test_fqa_inverse(t *testing.T) {
	test := "inverse"
	for _, v := range tt {
		obtenu := v.f.Inverse(v.n)
		attendu := v.i
		if obtenu != attendu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}	
}

type donnees_t struct {
	c []int // tableau des codes
	x0, y0 int // variable, valeur
}

type test_t struct {
	ok bool
	f *Fqa_t
}

func Test_codes(t *testing.T) {
	test := "codes"
	var tests_codes = []struct {
		donnees donnees_t
		attendu test_t
	}{
		{ // exemple
			donnees_t{[]int{2,2,1,2,2,2,1,2,2,1,2,2,2,1,2,2,1,2,2,2,1},0,0},
			test_t{ok:true,f:Creer(12,7,5)},
		},
		{ // calendrier musulman (cycle des annees)
			donnees_t{[]int{3,2,3,3,3,2,3,3,2,3,3,3,2,3,3,3,2,3,3,2,3,3},0,2},
			test_t{ok:true,f:Creer(30,11,26)},
		},
		{ // calendrier juif (annees embolismiques)
			donnees_t{[]int{3,3,2,3,3,3,2,3,3,2,3,3,3,2},0,0},
			test_t{ok:true,f:Creer(19,7,5)},
		},
		{ // jours (ecarts ou codes successifs entre j+1 et j)
			donnees_t{[]int{1,1,1,1,1,1,1,1},1,0},
			test_t{ok:true,f:Creer(1,1,-1)},
		},
		{ // mois (duree des mois de mars a fevrier sur une annee glissante) 
		  // remarque : le mois de fevrier n'a pas d'effet sur le resultat
			donnees_t{[]int{31,30,31,30,31,31,30,31,30,31,31,28},3,0},
			test_t{ok:true,f:Creer(153,5,-457)},
		},
		{ // calendrier julien (annees)
		  // (durees des annees sur 2 periodes de 4 ans) 
		  // (origine -4712 a minuit cad jd+0.5)
		  // 4712 * 365.25 = 1721058
		  // le premier mars seules les annees comptent
		  // l'annee 0 (la variable x0) est bissextile : 
		  // du 1 janvier 0 au 1 mars 0 = 31+29=60
		  // 1721058 + 60 = 1721118 (la valeur y0)
			donnees_t{[]int{365,365,365,366,365,365,365,366},0,1721118},
			test_t{ok:true,f:Creer(1461,4,6884472)},
		},
		{ // calendrier gregorien (siecles)
		  // (durees des siecles sur deux periodes de 400 ans)
		  // on fait commencer le jd a minuit et non pas a midi (jd+0.5)
		  // on impose la continuite du jd : 
		  //   jd gregorien(15 octobre 1582) = jd julien(4 octobre 1582) + 1
		  // remarque : l'annee gregorienne 1582 n'a que 355 jours
		  //            de meme l'annee julienne 46 avant JC eut 445 jours
		  // remarque : John Herschel (1792-1871)) proposa de retirer 
		  //            au calendrier gregorien une annee bissextile tous
		  //            les 4000 ans (les annees divisibles par 4000 ne seraient
		  //            plus bissextiles) mais a ce jour cela n'a pas ete retenu
		  //            et heureusement car cela preparerait le "bug" de l'an 4000...
		  // jd julien(4 octobre 1582) = 2299160
		  // du 4 octobre 1582 au 15 octobre 1582 = 1
		  // du 15 octobre au 31 decembre = 78
		  // annee 1583 = 365
		  // du 1 janvier 1584 au 31 decembre 1599 = 5844
		  // du 1 janvier 1600 au 29 fevrier 1600 = 60
		  // 2299160 + 1 + 78 + 365 + 5844 + 60 = 2305508
		  // le 1 mars 1600, seul le siecle s intervient dans le resultat (jd)
		  // s = 1600 / 100 = 16 (variable x0)
		  // jd gregorien (1 mars 1600) = 2305508 (valeur y0)
		  // [(a*s + r) / b] = jj = [(a*16 + r) / b] = 2305508
			donnees_t{[]int{36524,36524,36524,36525,36524,36524,36524,36525},16,2305508},
			test_t{ok:true,f:Creer(146097,4,6884480)},
		},
	}
	
	var obtenu = new(test_t)
	for _, tt := range tests_codes {
		obtenu.ok, obtenu.f = Codes(tt.donnees.c, tt.donnees.x0, tt.donnees.y0)
		if !(tt.attendu.ok == obtenu.ok) {
			t.Errorf(test+"(%v): attendu ok %t, obtenu %t",tt.donnees, tt.attendu.ok, obtenu.ok)
		}
		if tt.attendu.ok && !obtenu.f.Egal(tt.attendu.f) {
			t.Errorf(test+"(%v): attendu fqa %v, obtenu %v",tt.donnees, tt.attendu.f, obtenu.f)
		}
	}	
}

