#include <iostream>
#include <vector>
using namespace std;

vector<string> getCombinations(int stackLeft, int stackRight) {
    vector<string> combs;

    if (stackRight == 1 && stackLeft == 0) {
        return {")"};
    }

    if (stackLeft > 0) {
        auto combsIfLeft = getCombinations(stackLeft-1, stackRight+1);
        for (auto comb : combsIfLeft) {
            combs.push_back("(" + comb);
        }
    }

    if (stackRight > 0) {
        auto combsIfRight = getCombinations(stackLeft, stackRight-1);
        for (auto comb : combsIfRight) {
            combs.push_back(")" + comb);
        }
    }

    return combs;
}

vector<string> generateParenthesis(int n) {
    return getCombinations(n, 0);
}

int main() {
    vector<int> tests = {1, 2, 3};
    for (auto test : tests) {
        cout << test << endl;
        auto res = generateParenthesis(test);
        for (auto s : res) {
            cout << s << " ";
        }
        cout << endl;
    }
    return 0;
}