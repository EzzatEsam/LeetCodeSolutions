class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        vector<int> resVec( nums1.size() + nums2.size() ,0 );
        int idx = 0;
        int n1 =0;
        int n2 = 0;
        while (n1 < nums1.size() || n2 < nums2.size()) {
            if ( n1 == nums1.size())
                resVec[idx++] = nums2[n2++];
            else if (n2 == nums2.size())
                resVec[idx++] = nums1[n1++];
            else if( nums1[n1] < nums2[n2]  ) {
                resVec[idx++] = nums1[n1++];
            } else {
                resVec[idx++] = nums2[n2++];
            }
        }
        int md = resVec.size() /2;
        if (resVec.size() %2 == 0) {
            return (double) ((double)resVec[md-1] + (double)resVec[md ]) /2.0 ;
        } 
        return (double)resVec[md];
    }
};