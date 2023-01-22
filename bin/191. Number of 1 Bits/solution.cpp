#include <iostream>

using namespace std;
int hammingWeight(uint32_t n)
{
    return __builtin_popcount(n);
}
int main()
{
    cout << hammingWeight(00000000000000000000000000001011);
}
