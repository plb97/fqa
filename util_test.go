// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fqa

import (
	"testing"
)

type test_divent_t struct {
	n, d, q, r int
}
func Test_divent(t *testing.T) {
	test := "divent"
	var tt = []test_divent_t{
		{n:3,d:2,q:1,r:1},
		{n:-3,d:2,q:-2,r:1},
	}
	var obtenu test_divent_t 
	for _, v := range tt {
		obtenu.q, obtenu.r = Divent(v.n, v.d)
		if obtenu.q != v.q {
			t.Errorf(test+" %v: attendu q %v, obtenu %v", v, v.q, obtenu.d)
		}
		if obtenu.r != v.r {
			t.Errorf(test+" %v: attendu r %v, obtenu %v",v , v.r, obtenu.r)
		}
	}
}

type test_ent_t struct {
	f, r float64
	n int
}
func Test_ent(t *testing.T) {
	test := "ent"
	const prec = 1e-15
	var tt = []test_ent_t{
		{f:1.2,n:1,r:0.2},
		{f:-1.2,n:-2,r:0.8},
		{f:1.0,n:1,r:0.0},
		{f:-1.0,n:-1,r:0.0},
	}
	var obtenu test_ent_t
	for _, v := range tt {
		obtenu.n, obtenu.r = Ent(v.f)
		if obtenu.n != v.n {
			t.Errorf(test+" %v: attendu n %v, obtenu %v", v, v.n, obtenu.n)
		}
		if !Egal_f(obtenu.r,v.r,prec) {
			t.Errorf(test+" %v: attendu r %.15f, obtenu %.15f", v, v.r, obtenu.r)
		}
	}
}

type test_corrig_t struct {
	a, m, b, n int
}
func Test_corrig(t *testing.T) {
	test := "corrig"
	var tt = []test_corrig_t{
		{a:10,m:1,b:10,n:1},
		{a:10,m:13,b:11,n:1},
		{a:10,m:25,b:12,n:1},
		{a:10,m:0,b:9,n:12},
		{a:10,m:-1,b:9,n:11},
		{a:10,m:-12,b:8,n:12},
		{a:10,m:-13,b:8,n:11},
	}
	var obtenu test_corrig_t
	for _, v := range tt {
		obtenu.b, obtenu.n = Corrig_am(v.a, v.m)
		if obtenu.b != v.b {
			t.Errorf(test+" %v: attendu b %v, obtenu %v", v, v.b, obtenu.b)
		}
		if obtenu.n != v.n {
			t.Errorf(test+" %v: attendu n %v, obtenu %v", v, v.n, obtenu.n)
		}
	}
}

type test_norm_t struct {
	a, m, b, n int
}
func Test_norm(t *testing.T) {
	test := "norm"
	var tt = []test_norm_t{
		{a:10,m:3,b:10,n:3},
		{a:10,m:1,b:9,n:13},
		{a:10,m:2,b:9,n:14},
	}
	var obtenu test_norm_t
	for _, v := range tt {
		obtenu.b, obtenu.n = Norm_am(v.a, v.m)
		if obtenu.b != v.b {
			t.Errorf(test+" %v: attendu b %v, obtenu %v", v, v.b, obtenu.b)
		}
		if obtenu.n != v.n {
			t.Errorf(test+" %v: attendu n %v, obtenu %v", v, v.n, obtenu.n)
		}
		obtenu.a, obtenu.m = DNorm_am(v.b, v.n)
		if obtenu.a != v.a {
			t.Errorf(test+" %v: attendu a %v, obtenu %v", v, v.a, obtenu.a)
		}
		if obtenu.n != v.n {
			t.Errorf(test+" %v: attendu m %v, obtenu %v", v, v.m, obtenu.m)
		}
	}
}
