
#include <iostream>
#include "simple.h"
#include "polynomial.h"

using namespace std;

int main(int argc, char **argv) {
    pnm_t poly, d_poly;
    int fib_n = 15;

    cout << "tri_area_geon():" << "\n";
    cout << tri_area_geron(3.5, 4.1, 1.4) << "\n\n";

    cout << "tri_area_geron_det():" << "\n";
    cout << tri_area_geron_det(3.5, 4.1, 1.4) << "\n\n";

    cout << "print_pnm():" << "\n";
    poly.push_back(make_pair(0, 32.6));
    poly.push_back(make_pair(1, 7.2));
    poly.push_back(make_pair(2, 0.5));
    poly.push_back(make_pair(3, 0.353));
    poly.push_back(make_pair(4, -0.44));
    print_pnm(poly);
    cout << "\n";

    cout << "pnm_get_derivative():" << "\n";
    d_poly = pnm_get_derivative(poly);
    print_pnm(d_poly);
    cout << "\n";
    
    cout << "fib_by_mtrx():" << "\n";
    cout << "fib_n = " << fib_n << "\n";
    cout << fib_by_mtrx(fib_n) << "\n\n";
}
