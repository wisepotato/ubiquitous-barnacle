#include <iostream>
#include <list>
using namespace std;

namespace eulertwo {
	int main()
	{
		long long int fn2 = 0, fn1 = 2;
		long int sum = fn2 + fn1;
		long int fn;
		long int limit = 4 * 10 ^ 6;
		while (sum <= limit) {
			fn = 4 * fn1 + fn2;
			sum += fn;

			fn2 = fn1;
			fn1 = fn;
		}

		cout << sum;


		int returnvalue;
		cin >> returnvalue;
		return 0;
	}
}
