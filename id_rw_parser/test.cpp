
mtrx3_t	mtrx_invert(const mtrx3_t &m) {
	mtrx3_t inverse, rt;
	float det, invDet;

	inverse[id_rw(0, 0, 3)] = m[id_rw(1, 1, 3)] * m[id_rw(2, 2, 3)] - m[id_rw(1, 2, 3)] * m[id_rw(2, 1, 3)];
	inverse[id_rw(1, 0, 3)] = m[id_rw(1, 2, 3)] * m[id_rw(2, 0, 3)] - m[id_rw(1, 0, 3)] * m[id_rw(2, 2, 3)];
	inverse[id_rw(2, 0, 3)] = m[id_rw(1, 0, 3)] * m[id_rw(2, 1, 3)] - m[id_rw(1, 1, 3)] * m[id_rw(2, 0, 3)];

	det = m[id_rw(0, 0, 3)] * inverse[id_rw(0, 0, 3)] + m[id_rw(0, 1, 3)] * inverse[id_rw(1, 0, 3)] + 
		  m[id_rw(0, 2, 3)] * inverse[id_rw(2, 0, 3)];

	if (fabs(det) < f_eps) {
		return mtrx3_t();
	}

	invDet = 1.0f / det;

	inverse[id_rw(0, 1, 3)] = m[id_rw(0, 2, 3)] * m[id_rw(2, 1, 3)] - m[id_rw(0, 1, 3)] * m[id_rw(2, 2, 3)];
	inverse[id_rw(0, 2, 3)] = m[id_rw(0, 1, 3)] * m[id_rw(1, 2, 3)] - m[id_rw(0, 2, 3)] * m[id_rw(1, 1, 3)];
	inverse[id_rw(1, 1, 3)] = m[id_rw(0, 0, 3)] * m[id_rw(2, 2, 3)] - m[id_rw(0, 2, 3)] * m[id_rw(2, 0, 3)];
	inverse[id_rw(1, 2, 3)] = m[id_rw(0, 2, 3)] * m[id_rw(1, 0, 3)] - m[id_rw(0, 0, 3)] * m[id_rw(1, 2, 3)];
	inverse[id_rw(2, 1, 3)] = m[id_rw(0, 1, 3)] * m[id_rw(2, 0, 3)] - m[id_rw(0, 0, 3)] * m[id_rw(2, 1, 3)];
	inverse[id_rw(2, 2, 3)] = m[id_rw(0, 0, 3)] * m[id_rw(1, 1, 3)] - m[id_rw(0, 1, 3)] * m[id_rw(1, 0, 3)];

	rt[id_rw(0, 0, 3)] = inverse[id_rw(0, 0, 3)] * invDet;
	rt[id_rw(0, 1, 3)] = inverse[id_rw(0, 1, 3)] * invDet;
	rt[id_rw(0, 2, 3)] = inverse[id_rw(0, 2, 3)] * invDet;

	rt[id_rw(1, 0, 3)] = inverse[id_rw(1, 0, 3)] * invDet;
	rt[id_rw(1, 1, 3)] = inverse[id_rw(1, 1, 3)] * invDet;
	rt[id_rw(1, 2, 3)] = inverse[id_rw(1, 2, 3)] * invDet;

	rt[id_rw(2, 0, 3)] = inverse[id_rw(2, 0, 3)] * invDet;
	rt[id_rw(2, 1, 3)] = inverse[id_rw(2, 1, 3)] * invDet;
	rt[id_rw(2, 2, 3)] = inverse[id_rw(2, 2, 3)] * invDet;

	return rt;
}
