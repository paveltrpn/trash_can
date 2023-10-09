
#include <iostream>
#include <cmath>

#include "simple.h"
#include "algebra.h"

using namespace std;

// Area of triangle by lenght of tri sides
// using Geron's formula. Very simple.
// From [2], ch.2, str.27
float tri_area_geron(float a, float b, float c) {
   	float S;

	if ((a < 0.0) || (b < 0.0) || (c < 0.0)) {
		cout << "triArea(): side less than zero!\n";
		return -1.0;
	}

	S = (a + b + c) * 0.5;

	return sqrtf(S * (S - a) * (S - b) * (S - c));
}

// Area of triangle by length of tri sides
// using determinant calculation of some
// special matrix build from triangle sides length. 
float tri_area_geron_det(float a, float b, float c) {
    mtrx4_t M;
	float det;
	
	M = mtrx4_t(a, b, c, 0.0,
				b, a, 0.0, c,
				c, 0.0, a, b,
				0.0, c, b, a);

	det = mtrx_det_lu<_M4_>(M);

	return sqrtf(-det) / 4.0;
}

// Return N Fibonachi number.
// Calculating by producing matrix.
float fib_by_mtrx(int n) {
	mtrx2_t prod_mtrx = {0.0, 1.0, 1.0, 1.0};
	vec2_t fib_start = {0.0, 1.0}, fib_current;

	for (int i = 0; i < n; i++) {
		fib_current = mtrx_mult_vec<_M2V2_>(prod_mtrx, fib_start);
		fib_start = fib_current;
	}

	return fib_current[0];
}