#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class CustomQueue {
public:
    vector<pair<int, int>> q1, q2;
    vector<pair<int, int>> qm1, qm2;

    void add(int v, int i) {
        q1.push_back({v, i});
        if (qm1.empty() || v > qm1.back().first) {
            qm1.push_back({v, i});
        } else {
            qm1.push_back(qm1.back());
        }
    }

    void remove() {
        if (q2.empty()) {
            while (!q1.empty()) {
                q2.push_back({q1.back().first, q1.back().second});
                if (qm2.empty() || qm2.back().first > q1.back().first) {
                    qm2.push_back({q1.back().first, q1.back().second});
                } else {
                    qm2.push_back(qm2.back());
                }
                q1.pop_back();
                qm1.pop_back();
            }
        }
        q2.pop_back();
        qm2.pop_back();
    }

    pair<int, int> get_min() {
        if (q1.empty()) {
            return qm2.back();
        }
        if (q2.empty()) {
            return qm1.back();
        }
        return (qm1.back().first < qm2.back().first) ? qm2.back() : qm1.back();
    }
};

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(0);

    int n, k;
    cin >> n >> k;

    vector<int> a(n, 0);
    for (int i = 1; i < n - 1; i++) {
        cin >> a[i];
    }

    vector<pair<int, int>> dp_result(n);
    CustomQueue dp;
    dp.add(0, 0);
    dp_result[0] = {0, -1};

    for (int i = 1; i < min(n, k); i++) {
        dp_result[i] = {dp.get_min().first + a[i], dp.get_min().second};
        dp.add(dp_result[i].first, i);
    }

    for (int i = k; i < n; i++) {
        dp_result[i] = {dp.get_min().first + a[i], dp.get_min().second};
        dp.remove();
        dp.add(dp_result[i].first, i);
    }

    cout << dp_result.back().first << "\n";

    int idx = n - 1;
    vector<int> path;
    while (idx != -1) {
        path.push_back(idx + 1);
        idx = dp_result[idx].second;
    }
    reverse(path.begin(), path.end());

    cout << path.size() - 1 << "\n";
    for (int i = 0; i < path.size(); i++) {
        cout << path[i] << " ";
    }

    return 0;
}