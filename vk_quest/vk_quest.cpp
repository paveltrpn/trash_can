
/* 
   Written using Armadillo 9.100 linear algebra library
   installing dependences (in Linux Mint 19.0):
   LAPACK - $ sudo apt-get install liblapack*
   BLAS   - $ sudo apt-get install libblas*

   Armadillo installed from sources from git:
   $ git clone https://gitlab.com/conradsnicta/armadillo-code.git
   $ ./configure 
   $ make | make install

   LAPACK and BLAS libs need for arma::solve(), arama::det() and other methods

   Compile options:
   g++ -Wall -std=c++11  vk_quest.cpp -o main -lstdc++ -larmadillo
*/

#include <iostream>
#include <armadillo>
#include <cmath>

void swap_float(float &a, float &b) {
	float c;
	
	c = b;
	b = a;
	a = c;
}

float u_float_rand() {
    return ((float) rand()) / (float) RAND_MAX;
}

float s_float_rand() {
    return 1 - 2*u_float_rand();
}

float ival_float_rand(const float min, const float max) {
    float a = min;
    float b = max;

    if (a == b) b++;

    if (b < a) {
        swap_float(a, b);
    };

    return a + (u_float_rand() * (b - a));
}

class vk_quest_c {
    private:
        bool is_solved;
        float sys_det;

    public:
        arma::fmat k_matrix;
        arma::fvec fm_vec;
        arma::fvec ue_vec;

        vk_quest_c(const uint32_t range);
        ~vk_quest_c();

        void randomize(const uint32_t ival);
        void solve();
        void show_all();
};

vk_quest_c::vk_quest_c(const uint32_t range) {
    k_matrix = arma::fmat(range,range,arma::fill::randu);
    fm_vec   = arma::fvec(range,arma::fill::randu);
    ue_vec   = arma::fvec(range);

    is_solved = false;
}

vk_quest_c::~vk_quest_c() {}

void vk_quest_c::randomize(const uint32_t ival) {
    for (auto &it : k_matrix) {
        it = ival_float_rand(12,26);
    }

    for (auto &it : fm_vec) {
        it = rand() % ival;
    }
}

void vk_quest_c::solve() {
    sys_det = arma::det(k_matrix);

    if (sys_det == 0.0) {
        std::cout << "Determinant of k_matrix is zero!" << "\n";
        is_solved = false;
        return;
    }

    ue_vec = arma::solve(k_matrix, fm_vec);
    is_solved = true;
}

void vk_quest_c::show_all() {
    std::cout << k_matrix << "\n"; 
    std::cout << fm_vec << "\n"; 
    std::cout << ue_vec << "\n";
}

int main(int argc, char * argv[]) {
    std::cout << "FILE - vk_quest.cpp" << std::endl;
    
    vk_quest_c *vk_quest = new vk_quest_c(7);

    vk_quest->randomize(100);

    vk_quest->solve();

    vk_quest->show_all();
    
    delete vk_quest;

    return 0;
}