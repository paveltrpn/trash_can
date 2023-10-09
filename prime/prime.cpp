
#include <iostream>
#include <vector>
#include <cmath>

void show_sieve(const std::vector<bool> sieve) {
    uint64_t n = 0;

    for (auto i : sieve) {
        std::cout << "number - " << n << " is prime - " << i << std::endl;
        n++;
    }
}

uint64_t gcd(const uint64_t a, const uint64_t b) {
    if (b == 0) {
        return a;
    }

    return gcd(b, a%b);
}

std::vector<bool> erato_sieve(const uint64_t n) {
    std::vector<bool> prime(n+1, true);
    
    prime[0] = prime[1] = false;

    for (uint64_t i = 0; i <= n; ++i) {
        if (prime[i]) {
            if ((i * 1ll * i) <= n) {
                for (uint64_t j = i*i; j <= n; j += i) {
                    prime[j] = false;
                }
            }
        }
    }

    return prime;
}

std::vector<bool> atkin_sieve(const uint64_t limit) {
    // Предварительное просеивание

    std::vector<bool> sieve(limit);   

    for (unsigned long long x2 = 1ull, dx2 = 3ull; x2 < limit; x2 += dx2, dx2 += 2ull) {
        for (unsigned long long y2 = 1ull, dy2 = 3ull, n; y2 < limit; y2 += dy2, dy2 += 2ull) {
            // n = 4x² + y²
            n = (x2 << 2ull) + y2;
            if (n <= limit && (n % 12ull == 1ull || n % 12ull == 5ull))
                sieve[n].flip();
            // n = 3x² + y²
            n -= x2;
            if (n <= limit && n % 12ull == 7ull)
                sieve[n].flip();
            // n = 3x² - y² (при x > y)
            if (x2 > y2) {
                n -= y2 << 1ull;
                if (n <= limit && n % 12ull == 11ull)
                    sieve[n].flip();
            }
        }
    }

    // Все числа, кратные квадратам, помечаются как составные
    
    unsigned r = 5u;
    
    for (unsigned long long r2 = r * r, dr2 = (r << 1ull) + 1ull; r2 < limit; ++r, r2 += dr2, dr2 += 2ull)
        if (sieve[r])
            for (unsigned long long mr2 = r2; mr2 < limit; mr2 += r2)
                sieve[mr2] = false;

    // Числа 2 и 3 — заведомо простые

    if (limit > 2u)
        sieve[2u] = true;
    if (limit > 3u)
        sieve[3u] = true;
    return sieve;
}

bool is_prime(uint64_t n) {
    uint64_t i,sqrt_n = ceil(sqrt(n));

    if (n == 1) return false;

    for (i = 2; i <= sqrt_n; i++) {
        if ((n%i) == 0) return 0;
    }

    return 1;
}

int main(int argc, char * argv[]) {
    std::vector<bool> sieve;

    std::cout << "FILE - prime.cpp" << std::endl;
    
    sieve = atkin_sieve(100);
    show_sieve(sieve);

    return 0;
}