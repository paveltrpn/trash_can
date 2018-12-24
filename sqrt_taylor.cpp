
#include <iostream>
#include <cmath>
#include <string>

/*
int factorial(int n) {
	return n <= 1 ? 1 : n * factorial(n-1);
}
*/

int factorial(int n) {
	int tmp = 1;

	for (int i = 1; i <= n; ++i) {
		tmp *= i;
	}

	return tmp;
}

float sqrt_taylor(float x, float iter) {
	float answ = 0.0;

	/*Считаем выражение sqrt(1 + x), где abs(x) < 1, через разложение в ряд Тейлора */
	for (int n = 0; n < iter; n++) {
		answ += ( (pow(-1.0,n)*factorial(2*n)) / 
					( (1.0 - 2*n)*pow(factorial(n),2)*((1<<n)*(1<<n))) )*pow(x - 1.0,n); 
					// (1<<n)*(1<<n) = 4^n
	}

	return answ;
}

void usage() {
	std::cout << "Program sqrt_taylor" << std::endl;
}

int main(int argc, char *argv[]) {
	float x = 1.6;
	int iter = 3;
	
	for (int i = 0; i < argc; i++) {
		if (std::string(argv[i]) == "-x") {
			x = std::stof(argv[i+1]);

			if ((x <= 0.0) || (x >= 2.0)) {
				std::cout << "Error in main(): x must be under 2.0 and grater than 0.0! \n";
				return 1;
			}
			i++;
			continue;
		}

		if (std::string(argv[i]) == "-i") {
			iter = std::stoi(argv[i+1]);

			if ((iter > 15) || (iter < 1))  {
				std::cout << "Error in main(): iterations count must be in [1...15]  \n";
				return 1;
			}
			i++;
			continue;
		}
	}

	std::cout << "Program sqrt_taylor. \n\n" << "input x = " << x << "\n" <<
				"itarations = " << iter << "\n\n";	
	std::cout << "cmath sqrt of " << x << " is - " << sqrt(x) << "\n";
	std::cout << "sqrt_taylor " << x << " is - " << sqrt_taylor(x,iter) << "\n";

	return 0;
}