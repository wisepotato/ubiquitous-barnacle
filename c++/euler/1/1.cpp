#include <iostream>
#include <list>
using namespace std;

// Function for binary_predicate 
bool compare(int a, int b)
{
	return (a == b);
}
// get maximum amount of iterations possible < maxValue
int getMaxIterations(int multiple, int maxValue) {
	return floor((maxValue - 1) / multiple);
}
// summate, the easy way! does sum_{i=1}^n i for you!
int sumAlgebraic(int n) {
	return (n * (n + 1)) / 2;
}
int main()
{
	int sum = 0;
	list<int> results;
	list<int> multiples = { 3, 5 };
	int maxValue = 1000;
	int maxIterations;
	for (int multiple : multiples) {
		for (int multipleNested : multiples) {
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
	cout << sum;
	return 0;
}