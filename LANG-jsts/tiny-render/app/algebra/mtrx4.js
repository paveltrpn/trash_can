import * as vec3 from "./vec3.js";
import { idRw } from "./common.js";
export class mtrx4 {
    data;
    order = 4;
    constructor(src) {
        if ((src) || (src instanceof mtrx4)) {
            this.data = new Float32Array(16);
            this.fromMtrx4(src);
        }
        else if ((src) || (src instanceof Float32Array)) {
            this.data = new Float32Array(16);
            this.fromArray(src);
        }
        else {
            let i, j;
            this.data = new Float32Array(16);
            for (i = 0; i < this.order; i++) {
                for (j = 0; j < this.order; j++) {
                    if (i == j) {
                        this.data[idRw(i, j, this.order)] = 1.0;
                    }
                    else {
                        this.data[idRw(i, j, this.order)] = 0.0;
                    }
                }
            }
        }
    }
    setIdtt() {
        let i, j;
        for (i = 0; i < this.order; i++) {
            for (j = 0; j < this.order; j++) {
                if (i == j) {
                    this.data[idRw(i, j, this.order)] = 1.0;
                }
                else {
                    this.data[idRw(i, j, this.order)] = 0.0;
                }
            }
        }
    }
    fromMtrx4(src) {
        for (let i = 0; i < this.order * this.order; i++) {
            this.data[i] = src.data[i];
        }
    }
    fromArray(src) {
        for (let i = 0; i < this.order * this.order; i++) {
            this.data[i] = src[i];
        }
    }
    fromQtnn(src) {
        let wx, wy, wz, xx, yy, yz, xy, xz, zz, x2, y2, z2;
        x2 = src.data[0] + src.data[0];
        y2 = src.data[1] + src.data[1];
        z2 = src.data[2] + src.data[2];
        xx = src.data[0] * x2;
        xy = src.data[0] * y2;
        xz = src.data[0] * z2;
        yy = src.data[1] * y2;
        yz = src.data[1] * z2;
        zz = src.data[2] * z2;
        wx = src.data[3] * x2;
        wy = src.data[3] * y2;
        wz = src.data[3] * z2;
        this.data[0] = 1.0 - (yy + zz);
        this.data[1] = xy - wz;
        this.data[2] = xz + wy;
        this.data[3] = 0.0;
        this.data[4] = xy + wz;
        this.data[5] = 1.0 - (xx + zz);
        this.data[6] = yz - wx;
        this.data[7] = 0.0;
        this.data[8] = xz - wy;
        this.data[9] = yz + wx;
        this.data[10] = 1.0 - (xx + yy);
        this.data[11] = 0.0;
        this.data[12] = 0.0;
        this.data[13] = 0.0;
        this.data[14] = 0.0;
        this.data[15] = 1.0;
    }
    mult(a) {
        let i, j, k, tmp;
        let rt = new mtrx4();
        for (i = 0; i < this.order; i++) {
            for (j = 0; j < this.order; j++) {
                tmp = 0.0;
                for (k = 0; k < this.order; k++) {
                    tmp = tmp + this.data[(idRw(k, j, this.order))] * a.data[(idRw(i, k, this.order))];
                }
                rt.data[(idRw(i, j, this.order))] = tmp;
            }
        }
        this.fromMtrx4(rt);
    }
    setPerspective(fovy, aspect, near, far) {
        let f = 1.0 / Math.tan(fovy / 2);
        let nf;
        this.data[0] = f / aspect;
        this.data[1] = 0;
        this.data[2] = 0;
        this.data[3] = 0;
        this.data[4] = 0;
        this.data[5] = f;
        this.data[6] = 0;
        this.data[7] = 0;
        this.data[8] = 0;
        this.data[9] = 0;
        this.data[11] = -1;
        this.data[12] = 0;
        this.data[13] = 0;
        this.data[15] = 0;
        if (far != null && far !== Infinity) {
            nf = 1 / (near - far);
            this.data[10] = (far + near) * nf;
            this.data[14] = 2 * far * near * nf;
        }
        else {
            this.data[10] = -1;
            this.data[14] = -2 * near;
        }
    }
    setTranslate(vec) {
        this.setIdtt();
        this.data[12] = vec.data[0];
        this.data[13] = vec.data[1];
        this.data[14] = vec.data[2];
    }
    invert() {
        let inv = new mtrx4();
        let i, det;
        inv.data[0] = this.data[5] * this.data[10] * this.data[15] -
            this.data[5] * this.data[11] * this.data[14] -
            this.data[9] * this.data[6] * this.data[15] +
            this.data[9] * this.data[7] * this.data[14] +
            this.data[13] * this.data[6] * this.data[11] -
            this.data[13] * this.data[7] * this.data[10];
        inv.data[4] = -this.data[4] * this.data[10] * this.data[15] +
            this.data[4] * this.data[11] * this.data[14] +
            this.data[8] * this.data[6] * this.data[15] -
            this.data[8] * this.data[7] * this.data[14] -
            this.data[12] * this.data[6] * this.data[11] +
            this.data[12] * this.data[7] * this.data[10];
        inv.data[8] = this.data[4] * this.data[9] * this.data[15] -
            this.data[4] * this.data[11] * this.data[13] -
            this.data[8] * this.data[5] * this.data[15] +
            this.data[8] * this.data[7] * this.data[13] +
            this.data[12] * this.data[5] * this.data[11] -
            this.data[12] * this.data[7] * this.data[9];
        inv.data[12] = -this.data[4] * this.data[9] * this.data[14] +
            this.data[4] * this.data[10] * this.data[13] +
            this.data[8] * this.data[5] * this.data[14] -
            this.data[8] * this.data[6] * this.data[13] -
            this.data[12] * this.data[5] * this.data[10] +
            this.data[12] * this.data[6] * this.data[9];
        inv.data[1] = -this.data[1] * this.data[10] * this.data[15] +
            this.data[1] * this.data[11] * this.data[14] +
            this.data[9] * this.data[2] * this.data[15] -
            this.data[9] * this.data[3] * this.data[14] -
            this.data[13] * this.data[2] * this.data[11] +
            this.data[13] * this.data[3] * this.data[10];
        inv.data[5] = this.data[0] * this.data[10] * this.data[15] -
            this.data[0] * this.data[11] * this.data[14] -
            this.data[8] * this.data[2] * this.data[15] +
            this.data[8] * this.data[3] * this.data[14] +
            this.data[12] * this.data[2] * this.data[11] -
            this.data[12] * this.data[3] * this.data[10];
        inv.data[9] = -this.data[0] * this.data[9] * this.data[15] +
            this.data[0] * this.data[11] * this.data[13] +
            this.data[8] * this.data[1] * this.data[15] -
            this.data[8] * this.data[3] * this.data[13] -
            this.data[12] * this.data[1] * this.data[11] +
            this.data[12] * this.data[3] * this.data[9];
        inv.data[13] = this.data[0] * this.data[9] * this.data[14] -
            this.data[0] * this.data[10] * this.data[13] -
            this.data[8] * this.data[1] * this.data[14] +
            this.data[8] * this.data[2] * this.data[13] +
            this.data[12] * this.data[1] * this.data[10] -
            this.data[12] * this.data[2] * this.data[9];
        inv.data[2] = this.data[1] * this.data[6] * this.data[15] -
            this.data[1] * this.data[7] * this.data[14] -
            this.data[5] * this.data[2] * this.data[15] +
            this.data[5] * this.data[3] * this.data[14] +
            this.data[13] * this.data[2] * this.data[7] -
            this.data[13] * this.data[3] * this.data[6];
        inv.data[6] = -this.data[0] * this.data[6] * this.data[15] +
            this.data[0] * this.data[7] * this.data[14] +
            this.data[4] * this.data[2] * this.data[15] -
            this.data[4] * this.data[3] * this.data[14] -
            this.data[12] * this.data[2] * this.data[7] +
            this.data[12] * this.data[3] * this.data[6];
        inv.data[10] = this.data[0] * this.data[5] * this.data[15] -
            this.data[0] * this.data[7] * this.data[13] -
            this.data[4] * this.data[1] * this.data[15] +
            this.data[4] * this.data[3] * this.data[13] +
            this.data[12] * this.data[1] * this.data[7] -
            this.data[12] * this.data[3] * this.data[5];
        inv.data[14] = -this.data[0] * this.data[5] * this.data[14] +
            this.data[0] * this.data[6] * this.data[13] +
            this.data[4] * this.data[1] * this.data[14] -
            this.data[4] * this.data[2] * this.data[13] -
            this.data[12] * this.data[1] * this.data[6] +
            this.data[12] * this.data[2] * this.data[5];
        inv.data[3] = -this.data[1] * this.data[6] * this.data[11] +
            this.data[1] * this.data[7] * this.data[10] +
            this.data[5] * this.data[2] * this.data[11] -
            this.data[5] * this.data[3] * this.data[10] -
            this.data[9] * this.data[2] * this.data[7] +
            this.data[9] * this.data[3] * this.data[6];
        inv.data[7] = this.data[0] * this.data[6] * this.data[11] -
            this.data[0] * this.data[7] * this.data[10] -
            this.data[4] * this.data[2] * this.data[11] +
            this.data[4] * this.data[3] * this.data[10] +
            this.data[8] * this.data[2] * this.data[7] -
            this.data[8] * this.data[3] * this.data[6];
        inv.data[11] = -this.data[0] * this.data[5] * this.data[11] +
            this.data[0] * this.data[7] * this.data[9] +
            this.data[4] * this.data[1] * this.data[11] -
            this.data[4] * this.data[3] * this.data[9] -
            this.data[8] * this.data[1] * this.data[7] +
            this.data[8] * this.data[3] * this.data[5];
        inv.data[15] = this.data[0] * this.data[5] * this.data[10] -
            this.data[0] * this.data[6] * this.data[9] -
            this.data[4] * this.data[1] * this.data[10] +
            this.data[4] * this.data[2] * this.data[9] +
            this.data[8] * this.data[1] * this.data[6] -
            this.data[8] * this.data[2] * this.data[5];
        det = this.data[0] * inv.data[0] + this.data[1] * inv.data[4] + this.data[2] * inv.data[8] + this.data[3] * inv.data[12];
        if (det == 0) {
            this.setIdtt();
            return;
        }
        det = 1.0 / det;
        for (i = 0; i < 16; i++)
            this.data[i] = inv.data[i] * det;
    }
    transpose() {
        let tmp = new mtrx4;
        tmp.data[0] = this.data[0];
        tmp.data[1] = this.data[4];
        tmp.data[2] = this.data[8];
        tmp.data[3] = this.data[12];
        tmp.data[4] = this.data[1];
        tmp.data[5] = this.data[5];
        tmp.data[6] = this.data[9];
        tmp.data[7] = this.data[13];
        tmp.data[8] = this.data[2];
        tmp.data[9] = this.data[6];
        tmp.data[10] = this.data[10];
        tmp.data[11] = this.data[14];
        tmp.data[12] = this.data[3];
        tmp.data[13] = this.data[7];
        tmp.data[14] = this.data[11];
        tmp.data[15] = this.data[15];
        for (let i = 0; i < 16; i++) {
            this.data[i] = tmp.data[i];
        }
    }
    setAxisAngl(axis, phi) {
        let cosphi, sinphi, vxvy, vxvz, vyvz, vx, vy, vz;
        let ax = new vec3.vec3();
        ax = vec3.vec3Normalize(axis);
        cosphi = Math.cos(phi);
        sinphi = Math.sin(phi);
        vxvy = ax.data[0] * ax.data[1];
        vxvz = ax.data[0] * ax.data[2];
        vyvz = ax.data[1] * ax.data[2];
        vx = ax.data[0];
        vy = ax.data[1];
        vz = ax.data[2];
        this.data[0] = cosphi + (1.0 - cosphi) * vx * vx;
        this.data[1] = (1.0 - cosphi) * vxvy - sinphi * vz;
        this.data[2] = (1.0 - cosphi) * vxvz + sinphi * vy;
        this.data[3] = 0.0;
        this.data[4] = (1.0 - cosphi) * vxvy + sinphi * vz;
        this.data[5] = cosphi + (1.0 - cosphi) * vy * vy;
        this.data[6] = (1.0 - cosphi) * vyvz - sinphi * vx;
        this.data[7] = 0.0;
        this.data[8] = (1.0 - cosphi) * vxvz - sinphi * vy;
        this.data[9] = (1.0 - cosphi) * vyvz + sinphi * vx;
        this.data[10] = cosphi + (1.0 - cosphi) * vz * vz;
        this.data[11] = 0.0;
        this.data[12] = 0.0;
        this.data[13] = 0.0;
        this.data[14] = 0.0;
        this.data[15] = 1.0;
    }
    setEuler(yaw, pitch, roll) {
        let cosy, siny, cosp, sinp, cosr, sinr;
        cosy = Math.cos(yaw);
        siny = Math.sin(yaw);
        cosp = Math.cos(pitch);
        sinp = Math.sin(pitch);
        cosr = Math.cos(roll);
        sinr = Math.sin(roll);
        this.data[0] = cosy * cosr - siny * cosp * sinr;
        this.data[1] = -cosy * sinr - siny * cosp * cosr;
        this.data[2] = siny * sinp;
        this.data[3] = 0.0;
        this.data[4] = siny * cosr + cosy * cosp * sinr;
        this.data[5] = -siny * sinr + cosy * cosp * cosr;
        this.data[6] = -cosy * sinp;
        this.data[7] = 0.0;
        this.data[8] = sinp * sinr;
        this.data[9] = sinp * cosr;
        this.data[10] = cosp;
        this.data[11] = 0.0;
        this.data[12] = 0.0;
        this.data[13] = 0.0;
        this.data[14] = 0.0;
        this.data[15] = 1.0;
    }
    multTranslate(a, v) {
        let x = v[0], y = v[1], z = v[2];
        let a00, a01, a02, a03;
        let a10, a11, a12, a13;
        let a20, a21, a22, a23;
        if (a.data === this.data) {
            this.data[12] = a.data[0] * x + a.data[4] * y + a.data[8] * z + a.data[12];
            this.data[13] = a.data[1] * x + a.data[5] * y + a.data[9] * z + a.data[13];
            this.data[14] = a.data[2] * x + a.data[6] * y + a.data[10] * z + a.data[14];
            this.data[15] = a.data[3] * x + a.data[7] * y + a.data[11] * z + a.data[15];
        }
        else {
            a00 = a.data[0];
            a01 = a.data[1];
            a02 = a.data[2];
            a03 = a.data[3];
            a10 = a.data[4];
            a11 = a.data[5];
            a12 = a.data[6];
            a13 = a.data[7];
            a20 = a.data[8];
            a21 = a.data[9];
            a22 = a.data[10];
            a23 = a.data[11];
            this.data[0] = a00;
            this.data[1] = a01;
            this.data[2] = a02;
            this.data[3] = a03;
            this.data[4] = a10;
            this.data[5] = a11;
            this.data[6] = a12;
            this.data[7] = a13;
            this.data[8] = a20;
            this.data[9] = a21;
            this.data[10] = a22;
            this.data[11] = a23;
            this.data[12] = a00 * x + a10 * y + a20 * z + a.data[12];
            this.data[13] = a01 * x + a11 * y + a21 * z + a.data[13];
            this.data[14] = a02 * x + a12 * y + a22 * z + a.data[14];
            this.data[15] = a03 * x + a13 * y + a23 * z + a.data[15];
        }
    }
}
export function mtrx4Transpose(m) {
    const mrange = 4;
    let i, j, tmp;
    let rt = new mtrx4(m);
    for (i = 0; i < mrange; i++) {
        for (j = 0; j < i; j++) {
            tmp = rt.data[idRw(i, i, mrange)];
            rt.data[idRw(i, j, mrange)] = rt.data[idRw(j, i, mrange)];
            rt.data[idRw(j, i, mrange)] = tmp;
        }
    }
    return rt;
}
export function mtrx4Mult(a, b) {
    const mrange = 4;
    let i, j, k, tmp;
    let rt = new mtrx4();
    for (i = 0; i < mrange; i++) {
        for (j = 0; j < mrange; j++) {
            tmp = 0.0;
            for (k = 0; k < mrange; k++) {
                tmp = tmp + a.data[(idRw(k, j, mrange))] * b.data[(idRw(i, k, mrange))];
            }
            rt.data[(idRw(i, j, mrange))] = tmp;
        }
    }
    return rt;
}
