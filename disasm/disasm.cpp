
#include <iostream>

int32_t global;

class tst_arr_c {
	public:
		int32_t arr[4];

	tst_arr_c():
		arr{0, 0, 0, 0} {};
};

tst_arr_c foo(int32_t foo_in) {
	//int32_t local = 0;
	//int32_t alocal[4];
	tst_arr_c arr;

	//local = local + foo_in;

	return arr;//local;
}

int main(int argc, char **argv) {
	tst_arr_c bar;

	bar = foo(10);

	std::cout << "global = " << std::endl;

	return 0;
}