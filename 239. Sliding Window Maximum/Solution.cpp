class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        priority_queue<pair<int,int>> q;
        vector<int> result;
        result.reserve(nums.size() - k + 1);
        for (size_t i = 0; i < k; i++)
        {
            q.push({nums[i],i});
        }
        result.push_back(q.top().first);
        for (size_t i = k; i < nums.size(); i++)
        {
            q.push({nums[i],i});
            while ( q.top().second <= i -k)
            {
                q.pop();
            }
            result.push_back(q.top().first);
        }
       
        return result;
    }
};