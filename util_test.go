package fqa

import (
	"testing"
)

type test_divent_t struct {
	n, d, q, r int
}
func Test_divent(t *testing.T) {
	var tt = []test_divent_t{
		{n:3,d:2,q:1,r:1},
		{n:-3,d:2,q:-2,r:1},
	}
	var actual test_divent_t 
	for _, v := range tt {
		actual.q, actual.r = Divent(v.n, v.d)
		if actual.q != v.q {
			t.Errorf("divent %v: expected q %v, actual %v", v, v.q, actual.d)
		}
		if actual.r != v.r {
			t.Errorf("divent %v: expected r %v, actual %v",v , v.r, actual.r)
		}
	}
}

type test_ent_t struct {
	f, r float64
	n int
}
func Test_ent(t *testing.T) {
	const prec = 1e-15
	var tt = []test_ent_t{
		{f:1.2,n:1,r:0.2},
		{f:-1.2,n:-2,r:0.8},
		{f:1.0,n:1,r:0.0},
		{f:-1.0,n:-1,r:0.0},
	}
	var actual test_ent_t
	for _, v := range tt {
		actual.n, actual.r = Ent(v.f)
		if actual.n != v.n {
			t.Errorf("divent %v: expected n %v, actual %v", v, v.n, actual.n)
		}
		if !Equal_f(actual.r,v.r,prec) {
			t.Errorf("divent %v: expected r %.15f, actual %.15f", v, v.r, actual.r)
		}
	}
}

type test_corrig_t struct {
	a, m, b, n int
}
func Test_corrig(t *testing.T) {
	var tt = []test_corrig_t{
		{a:10,m:1,b:10,n:1},
		{a:10,m:13,b:11,n:1},
		{a:10,m:25,b:12,n:1},
		{a:10,m:0,b:9,n:12},
		{a:10,m:-1,b:9,n:11},
		{a:10,m:-12,b:8,n:12},
		{a:10,m:-13,b:8,n:11},
	}
	var actual test_corrig_t
	for _, v := range tt {
		actual.b, actual.n = Corrig_am(v.a, v.m)
		if actual.b != v.b {
			t.Errorf("corrig %v: expected b %v, actual %v", v, v.b, actual.b)
		}
		if actual.n != v.n {
			t.Errorf("corrig %v: expected n %v, actual %v", v, v.n, actual.n)
		}
	}
}

type test_norm_t struct {
	a, m, b, n int
}
func Test_norm(t *testing.T) {
	var tt = []test_norm_t{
		{a:10,m:3,b:10,n:3},
		{a:10,m:1,b:9,n:13},
		{a:10,m:2,b:9,n:14},
	}
	var actual test_norm_t
	for _, v := range tt {
		actual.b, actual.n = Norm_am(v.a, v.m)
		if actual.b != v.b {
			t.Errorf("norm %v: expected b %v, actual %v", v, v.b, actual.b)
		}
		if actual.n != v.n {
			t.Errorf("norm %v: expected n %v, actual %v", v, v.n, actual.n)
		}
		actual.a, actual.m = DNorm_am(v.b, v.n)
		if actual.a != v.a {
			t.Errorf("norm %v: expected a %v, actual %v", v, v.a, actual.a)
		}
		if actual.n != v.n {
			t.Errorf("norm %v: expected m %v, actual %v", v, v.m, actual.m)
		}
	}
}
