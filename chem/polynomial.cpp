#include "polynomial.h"

#include <iostream>
#include <vector>
#include <utility>

using namespace std;

void print_pnm(pnm_t pnm) {
    cout << "P(x) = ";

    for (auto &p:pnm) {
        cout << get<1>(p) << "*" << "x^" << get<0>(p);

        if ( p != pnm.back()) {
            cout << " + ";
        }
    }

    cout << "\n";
}

pnm_t pnm_simplify(pnm_t pnm) {

}

pnm_t pnm_get_derivative(pnm_t pnm) {
    float dk;
    int de;
    pnm_t rt;

    for (auto p:pnm) {
        de = get<0>(p) - 1;

        if (de < 0) {
            continue;
        }

        dk = get<1>(p)*get<0>(p);

        rt.push_back(make_pair(de,dk));
    }

    return rt;
}

vector<float> pnm_solve_dichotomy(pnm_t p, float left, float right) {

}

vector<float> pnm_solve_newton(pnm_t p, float left, float right) {

}
