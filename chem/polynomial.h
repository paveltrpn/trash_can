#pragma once

#include <vector>
#include <utility>

using namespace std;

typedef vector<pair<int, float>> pnm_t;

void print_pnm(pnm_t p);
pnm_t pnm_simplify(pnm_t pnm);
pnm_t pnm_get_derivative(pnm_t pnm);
vector<float> pnm_solve_dichotomy(pnm_t p, float left, float right);
vector<float> pnm_solve_newton(pnm_t p, float left, float right);
