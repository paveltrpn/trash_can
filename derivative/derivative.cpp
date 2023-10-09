
#include <iostream>
#include <vector>
#include <cmath>
#include "plplot/plplot.h"
#include "plplot/plstream.h"

template <typename T>
void swap_arg(T &a, T &b) {
	T c;
	
	c = b;
	b = a;
	a = c;
}

class derivative_c {
    private:
        std::vector<double> fn;
        std::vector<double> first_drvt;
        std::vector<double> second_drvt;
        std::vector<double> grid;
        
        void set_grid(double begin,double end,uint32_t res);
    public:
        derivative_c(const double begin, const double end, const uint32_t res);
        ~derivative_c();

        void find_drvt();
        void show_plot(int argc, char** argv);
};

derivative_c::derivative_c(const double begin, const double end, const uint32_t res) {
    set_grid(begin,end,res);

    fn.resize(res);
    first_drvt.resize(res);
    second_drvt.resize(res);
}

derivative_c::~derivative_c() {

}

void derivative_c::set_grid(double begin, double end, const uint32_t res) {
    double ival;
    uint32_t step = 0;

    if (begin > end) swap_arg(begin,end);

    ival = (end - begin)/(res-1);

    grid.resize(res);

    for (auto &i:grid) {
        i = begin + step*ival;
        step++;
    };
}

void derivative_c::find_drvt() {
    uint32_t step = 0;

    for (auto &x:grid) {
        fn[step]            = x*x*x;
        first_drvt[step]    = 3*x*x;
        second_drvt[step]   = 6*x;
        step++; 
    } 
}

void derivative_c::show_plot(int argc, char** argv) {
    plstream *pls = new plstream();

    // Parse and process command line arguments
    pls->parseopts( &argc, argv, PL_PARSE_FULL );

    // Initialize plplot
    pls->init();

    // Create a labelled box to hold the plot.
    pls->env( -3.0, 3.0, -3.0, 3.0, 0, 0 );
    pls->lab( "x", "y", "2D line plot" );

    // Plot the data that was prepared above.
    pls->line( fn.size(), &grid[0], &fn[0] );
    pls->line( fn.size(), &grid[0], &first_drvt[0] );
    pls->line( fn.size(), &grid[0], &second_drvt[0] );

    // In C++ we don't call plend() to close PLplot library
    // this is handled by the destructor
    delete pls;
}

int main(int argc, char** argv) {
    derivative_c drv(-1,1,20);

    drv.find_drvt();
    drv.show_plot(argc,argv);

    return 0;
}