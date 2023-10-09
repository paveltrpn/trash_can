var dec;
(function (dec) {
    dec[dec["x"] = 0] = "x";
    dec[dec["y"] = 1] = "y";
})(dec || (dec = {}));
;
var x_size = 4;
var y_size = 4;
var field = x_size * y_size;
/*	multidimensional array mapping, array[i][j]
    row-wise (C, C++):
    (0	1)
    (2	3)

    column-wise (Fortran, Matlab):
    (0	2)
    (1	3)
*/
function id_rw(i, j, n) {
    return (i * n + j);
}
;
function id_cw(i, j, n) {
    return (j * n + i);
}
;
function brasenham_line(start, end) {
    var dX = Math.abs(end[dec.x] - start[dec.x]);
    var dY = Math.abs(end[dec.y] - start[dec.y]);
    var signX;
    var signY;
    var err = dX - dY;
    var rt = [];
    var now_point = start;
    if (start[dec.x] < end[dec.x]) {
        signX = 1;
    }
    else {
        signX = -1;
    }
    if (start[dec.y] < end[dec.y]) {
        signY = 1;
    }
    else {
        signY = -1;
    }
    rt.push(end);
    while ((now_point[dec.x] != end[dec.x]) || (now_point[dec.y] != end[dec.y])) {
        rt.push([now_point[0], now_point[1]]);
        var err2 = err * 2;
        if (err2 > -dY) {
            err -= dY;
            now_point[dec.x] += signX;
        }
        if (err2 < dX) {
            err += dX;
            now_point[dec.y] += signY;
        }
    }
    return rt;
}
function main() {
    var canvas = new Array(field);
    console.log("brasenham line");
    var line = brasenham_line([0, 0], [3, 3]);
    for (var i = 0; i < x_size; i++) {
        for (var j = 0; j < y_size; j++) {
            canvas[id_rw(i, j, x_size)] = 0;
            for (var _i = 0, line_1 = line; _i < line_1.length; _i++) {
                var pt = line_1[_i];
                if ((pt[dec.x] == i) && (pt[dec.y] == j)) {
                    canvas[id_rw(i, j, x_size)] = 1;
                }
            }
        }
    }
    var screen = canvas.toString();
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
