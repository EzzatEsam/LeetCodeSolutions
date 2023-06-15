class Solution {
public:
    int reverse(int x) {
        string s = to_string(x);
        long int n = 0;
        bool neg = false;
        int idx =1;
        if ((int)s[0] >'9' || (int)s[0] <'0')
            {neg = true;
            idx =0;}
        else
            n = (int)s[0]  - '0';
        for (int i =1; i < s.length() ; i ++) {
            int num = (int)s[i]  - '0';
            //cout << s[i] <<"..." << num <<endl;
            
            n+= num * pow(10,idx++);
            //cout<<n <<" lll "  <<num <<endl;
        }
        
        n= neg ?  -n : n;
        return (n> INT_MAX || n < INT_MIN) ? 0 : (int)n;
    }
   
};