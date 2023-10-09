
type bar1_vec3_t = Array<number>[3];

type bar2_vec3_t = {
    x: number;
    y: number;
    z: number;
};

function foo1_vec3_sum(a: bar1_vec3_t, b: bar1_vec3_t): bar1_vec3_t {
    let ret: bar1_vec3_t;

    ret[0] = a[0] + b[0];
    ret[1] = a[1] + b[1];
    ret[2] = a[2] + b[2];

    return ret;
}

function foo2_vec3_sum(a: bar2_vec3_t, b: bar2_vec3_t): bar2_vec3_t {
    let ret: bar2_vec3_t;

    ret.x = a.x + b.x;
    ret.y = a.y + b.y;
    ret.z = a.z + b.z;

    return ret;

}

function main() {
    
}