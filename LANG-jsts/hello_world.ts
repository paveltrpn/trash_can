
function hello_world(print: string) {
    console.log(print);
}

function fib(n: number): number[] {
    let fib: number[] = new Array(n+1);

    fib[0] = 0;
    fib[1] = 1;
    fib[2] = 1;

    for (let i = 3; i < n; i++) {
        fib[i] = fib[i-1] + fib[i-2];
    }

    return fib;
}

function mtrx2_mult(a: number[], b: number[]): number[] {
    let rt: number[] = new Array(4);

    rt[0] = a[0]*b[0] + a[1]*b[2];
    rt[1] = a[0]*b[1] + a[1]*b[3];
    rt[2] = a[2]*b[0] + a[3]*b[2];
    rt[3] = a[2]*b[1] + a[3]*b[3];

    return rt;
}

function fib_mtrx(n: number): number {
    const Q: number[] = [1, 1, 1, 0];
    let M:number[];

    if (n < 2) {
        return n;
    }
   
    if (n == 2) {
        return 1;
    }

    M = Q;
    for (let i = 0; i < n-2; i++) {
        M = mtrx2_mult(M, Q);
    }

    return M[0];
}

function erato_sieve(n: number): boolean[] {
    let prime: boolean[] = new Array(n+1);
    let j,i: number;

    for(i = 0; i < prime.length; i++) {
        prime[i] = true;
    }

    prime[0] = prime[1] = false;
    
    for (i = 2; i <= n; ++i) {
        if (prime[i]) {
            if ((i * i) <= n) {
                for (j = i*i; j <= n; j += i) {
                    prime[j] = false;
                }
            }
        }
    }
    
    return prime
}

function main() {
    /*
    let sieve: boolean[];
    console.log("erato_sieve");
    sieve = erato_sieve(100);
    for (let i = 0; i < sieve.length - 1; i++) {
        if (sieve[i] == true) {
            console.log(i);
        }
    }
    */
       
    let f: number[];
    console.log("fibonacci");
    f = fib(15);
    for (let i = 0; i < f.length - 1; i++) {
        console.log(f[i], " ", fib_mtrx(i));
    }
    
}

main()

