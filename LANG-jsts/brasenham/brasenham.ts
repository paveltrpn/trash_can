
type point_t = [number, number];
enum dec {x, y};

const x_size = 4;
const y_size = 4;
const field = x_size * y_size;

/*	multidimensional array mapping, array[i][j]
	row-wise (C, C++):
	(0	1)
	(2	3)

	column-wise (Fortran, Matlab):
	(0	2)
	(1	3)
*/

function id_rw(i: number, j: number, n: number): number {
	return (i*n + j);
};

function id_cw(i: number, j: number, n: number): number {
	return (j*n + i);
};

function brasenham_line(start: point_t, end: point_t): point_t[] {
    let dX = Math.abs(end[dec.x] - start[dec.x]);
    let dY = Math.abs(end[dec.y] - start[dec.y]);
    let signX: number;
    let signY: number;
    let err = dX - dY;
    let rt: point_t[] = [];
    let now_point = start;

    if (start[dec.x] < end[dec.x]) {
        signX = 1;
    } else {
        signX = -1;
    }

    if (start[dec.y] < end[dec.y]) {
        signY = 1;
    } else {
        signY = -1;
    }

    rt.push(end);

    while ((now_point[dec.x] != end[dec.x]) || (now_point[dec.y] != end[dec.y])) {
        rt.push([now_point[0], now_point[1]]);

        const err2: number = err * 2;

        if(err2 > -dY) 
        {
            err -= dY;
            now_point[dec.x] += signX;
        }

        if(err2 < dX) 
        {
            err += dX;
            now_point[dec.y] += signY;
        }
    }
    
    return rt; 
}

function main() {
    let canvas: number[] = new Array(field);

    console.log("brasenham line");

    let line = brasenham_line([0, 0], [3, 3]);

    for (let i = 0; i < x_size; i++) {
        for (let j = 0; j < y_size; j++) {
            canvas[id_rw(i,j,x_size)] = 0;
            for (let pt of line) {
                if ((pt[dec.x] == i) && (pt[dec.y] == j)) {
                    canvas[id_rw(i,j,x_size)] = 1;
                }
            }  
        }
    }
    
    let screen: string = canvas.toString();

    console.log(screen);
}

main();

/*
void drawLine(int x1, int y1, int x2, int y2) {
    const int deltaX = abs(x2 - x1);
    const int deltaY = abs(y2 - y1);
    const int signX = x1 < x2 ? 1 : -1;
    const int signY = y1 < y2 ? 1 : -1;
    //
    int error = deltaX - deltaY;
    //
    setPixel(x2, y2);
    while(x1 != x2 || y1 != y2) 
   {
        setPixel(x1, y1);
        const int error2 = error * 2;
        //
        if(error2 > -deltaY) 
        {
            error -= deltaY;
            x1 += signX;
        }
        if(error2 < deltaX) 
        {
            error += deltaX;
            y1 += signY;
        }
    }

}
*/