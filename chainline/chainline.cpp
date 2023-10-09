
#include <iostream>
#include <cmath>
#include <vector>
#include <plplot/plplot.h>
#include <plplot/plstream.h>

using namespace std;

int fac(int n) {
	if (n == 0) return 1;    
    return n*fac(n-1);
}

template <typename T>
void swap_arg(T &a, T &b) {
	T c;
	
	c = b;
	b = a;
	a = c;
}

template <typename fp_type>
vector<fp_type> set_grid(fp_type begin, fp_type end, const uint32_t res) {
	vector<fp_type> rt;
	fp_type ival;
    uint32_t step = 0;

    if (begin > end) swap_arg(begin,end);

    ival = (end - begin)/(res-1);

    rt.resize(res);

    for (auto &i:rt) {
        i = begin + step*ival;
        step++;
    }

	return rt;
}
template vector<float> 	set_grid(float begin, float end, uint32_t res);
template vector<double> set_grid(double begin, double end, uint32_t res);

template <typename fp_type>
fp_type sinh_taylor(fp_type arg) {
	constexpr int steps = 5;
	fp_type rt = arg;

	for (int i = 3; i < 2*steps+1; i += 2) {
		rt = rt + pow(arg, i) / fac(i); 
	}
	return rt;
}
template float	sinh_taylor<float>(float arg);
template double sinh_taylor<double>(double arg);

template <typename fp_type>
fp_type cosh_taylor(fp_type arg) {
	constexpr int steps = 5;
	fp_type rt = 1.0;
	
	for (int i = 2; i < steps*2; i += 2) {
		rt = rt + pow(arg,i) / fac(i);
	}
	return rt;
}
template float	cosh_taylor<float>(float arg);
template double cosh_taylor<double>(double arg);

template <typename fp_type>
fp_type chain_line(fp_type arg, fp_type a) {
	return a*cosh_taylor(arg/a);
}
template float	chain_line<float>(float arg, float a);
template double chain_line<double>(double arg, double a);

int main(int argc, char **argv) {
	vector<double> grid;
	vector<double> cosh_points;
	vector<double> cosh_taylor_points;
	vector<double> sinh_points;
	vector<double> sinh_taylor_points;
	vector<double> chain;

	grid = set_grid(-27.0, 27.0, 50);

	for (auto &x:grid) {
		cosh_taylor_points.push_back(cosh_taylor(x));
		cosh_points.push_back(cosh(x));

		sinh_taylor_points.push_back(sinh_taylor(x));
		sinh_points.push_back(sinh(x));

		chain.push_back(chain_line(x, 10.0));
	}

	plstream *pls = new plstream();

    // Parse and process command line arguments
    pls->parseopts( &argc, argv, PL_PARSE_FULL );

    // Initialize plplot
    pls->init();

    // Create a labelled box to hold the plot.
    pls->env( -35.0, 35.0, -10.0, 85.0, 0, 0 );
    pls->lab( "x", "y", "Chain line plot" );

    // Plot the data that was prepared above.
    pls->line( grid.size(), &grid[0], &cosh_taylor_points[0] );
    //pls->line( grid.size(), &grid[0], &cosh_points[0] );

	pls->line( grid.size(), &grid[0], &sinh_taylor_points[0] );
    //pls->line( grid.size(), &grid[0], &sinh_points[0] );

	pls->line( grid.size(), &grid[0], &chain[0] );

    // In C++ we don't call plend() to close PLplot library
    // this is handled by the destructor
    delete pls;

	return 0;
}