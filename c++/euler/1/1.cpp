#include <iostream>
#include <list>
using namespace std;

// Function for binary_predicate 
bool compare(long long a, long long b)
{
	return (a == b);
}
// get maximum amount of iterations possible < maxValue
int getMaxIterations(long long multiple, long long maxValue) {
	return (long long)((maxValue - 1) / multiple);
}

/*
	Summate the easy way, in a hard way, does ∑_{i=1}^n i  for you
	Uses a bit-shift right to do division, because otherwise
	we run into precision errors. 
*/
long sumAlgebraic(long n) {
	return (n * (n + 1)) >> 1;
}
/*
	Gets the sum of the multiples of a certain maxvalues, where the multiples are 3 and 5
*/
long long getSum(long maxValue){
	long long sum = 0;
	list<long long> results;
	list<long long> multiples = { 3, 5 };
	long long maxIterations;
	for (long long multiple : multiples) {
		for (long long multipleNested : multiples) {
			if (multiple != multipleNested) {
				results.push_front(multipleNested * multiple);
			}
		}
		maxIterations = getMaxIterations(multiple, maxValue);
		sum += multiple * sumAlgebraic(maxIterations);
	}
	results.unique(compare);
	for (int result : results) {
		maxIterations = getMaxIterations(result, maxValue);
		sum -= result * sumAlgebraic(maxIterations);
	}
	return sum;
}
/*
 usage of Long here because we need precision.
*/
int main()
{
    long t;
    cin >> t;
    for(int a0 = 0; a0 < t; a0++){
        long n;
        cin >> n;
        cout << getSum(n) << endl;
    }
	return 0;
}