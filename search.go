package algorithms

import "golang.org/x/exp/constraints"

func maximalSuffix[T constraints.Ordered](x []T, p *int) int {
	m := len(x)
	ms := -1
	j := 0
	k := 1
	*p = 1
	for j+k < m {
		a := x[j+k]
		b := x[ms+k]
		if a < b {
			j += k
			k = 1
			*p = j - ms
		} else if a == b {
			if k != *p {
				k++
			} else {
				j += *p
				k = 1
			}
		} else { /* a > b */
			ms = j
			j = ms + 1
			k = 1
			*p = 1
		}
	}
	return ms
}

func computeMaximalSuffix[T constraints.Ordered](x []T, p *int) int {
	m := len(x)
	ms := -1
	j := 0
	k := 1
	*p = 1
	for j+k < m {
		a := x[j+k]
		b := x[ms+k]
		if a > b {
			j += k
			k = 1
			*p = j - ms
		} else if a == b {
			if k != *p {
				k++
			} else {
				j += *p
				k = 1
			}
		} else { /* a < b */
			ms = j
			j = ms + 1
			k = 1
			*p = 1
		}
	}
	return ms
}

func Compare[T constraints.Ordered](x, y []T) int {
	for i := range x {
		if x[i] < y[i] {
			return -1
		} else if x[i] > y[i] {
			return 1
		}
	}
	return 0
}

//func Search[T comparable](x, y []T) int {
//	p := 0
//	q := 0
//	ell := 0
//	per := 0
//
//	/* Preprocessing */
//	i := maximalSuffix(x, &p);
//	j := computeMaximalSuffix(x, &q);
//	if i > j {
//		ell = i
//		per = p
//	} else {
//		ell = j
//		per = q
//	}
//
//	/* Searching */
//	if (Compare(x[0:ell+1], x[per:per+ell+1]) == 0) {
//		j = 0
//		memory = -1
//		for j <= n - m {
//			i = MAX(ell, memory) + 1;
//			for i < m && x[i] == y[i + j] {
//				i++;
//			}
//			if (i >= m) {
//				i = ell;
//				while (i > memory && x[i] == y[i + j])
//				--i;
//				if (i <= memory)
//					OUTPUT(j);
//				j += per;
//				memory = m - per - 1;
//			}
//			else {
//				j += (i - ell);
//				memory = -1;
//			}
//		}
//	}
//	else {
//		per = MAX(ell + 1, m - ell - 1) + 1;
//		j = 0;
//		while (j <= n - m) {
//			i = ell + 1;
//			while (i < m && x[i] == y[i + j])
//			++i;
//			if (i >= m) {
//				i = ell;
//				while (i >= 0 && x[i] == y[i + j])
//				--i;
//				if (i < 0)
//					OUTPUT(j);
//				j += per;
//			}
//			else
//			j += (i - ell);
//		}
//	}
//}
