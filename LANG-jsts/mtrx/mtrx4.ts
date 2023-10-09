//
// всё-таки видимо имеет смысл именно так сформировать класс
// data стала приватной и запись туда невозможна, только по сеттеру
//

type mtrx4 = Float32Array(16);

class mtrx4_t {
    // единичная матрица
    constructor () {
        let i, j: number;

	    for (i = 0; i < this.order; i++) {
	    	for (j = 0; j < this.order; j++) {
	    		if (i == j) {
	    			this.data[id_rw(i, j, this.order)] = 1.0;
	    		} else {
	    			this.data[id_rw(i, j, this.order)] = 0.0;
	    		}
	    	}
	    }
    }
    
    private readonly order: number = 4; 
    private data: number[];

    at(i: number): number {
        return this.data[i];
    }
    
    //"safe" at(), проверяет границы массива
    sat(i: number): number {
        if ((i < 0) || (i > this.order)) {
            return -1;
        }

        return this.data[i];
    }

    set_at(i: number, n: number) {
        this.data[i] = n;
    }
}

function id_rw(i: number, j: number, n: number): number {
	return (i*n + j);
};

function id_cw(i: number, j: number, n: number): number {
	return (j*n + i);
};

function mtrx4_mult(a: mtrx4_t, b: mtrx4_t): mtrx4_t {
    const mrange = 4;
	let i, j, k, tmp: number;
    let rt: mtrx4_t;

	for (i = 0; i < mrange; i++) {
		for (j = 0; j < mrange; j++) {
			tmp = 0.0;
            
            for (k = 0; k < mrange; k++) {
				tmp = tmp + a.at(id_rw(k, j, mrange)) * b.at(id_rw(i, k, mrange));
            }
            
			rt.set_at(id_rw(i, j, mrange), tmp);
		}
	}

	return rt;
}