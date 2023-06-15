class Solution {
public:
    int myAtoi(string s) {
        long  int n =0 ;
        int idx = 0;
        bool sgn = false;
        bool spc = false;
        bool inf = false;
        
       for ( int i = s.length() -1; i >=0 ; i --) 
       {
        //cout <<((int)s[i] - '0') *(pow(10,s.length() - i -1)) << endl; 
           int res = (int)s[i] - '0';
           //cout <<i << endl;
           //cout << res <<"--" << n <<endl;
           if (res >=0 && res <=9 )
               { 
               if (spc || sgn)
                {n=0;idx = 0; sgn = false; spc = false;}
               if ( idx >= 18 ) {
                   if (res ==0 ) idx ++;
                   else n= n<0 ? INT_MIN : INT_MAX;
                   inf = true;
               }
               else {
                   n +=res *(pow(10,idx)); idx ++;
               spc = false;
               }
               
               }
           else if (res == -16 ) {
               spc = true;
           }
            else  if(s[i] =='-' && !sgn  && !spc) 
               {sgn = true;
                
                if ( n == INT_MAX && inf)
                    n = INT_MIN ;
                else 
                    n=-n;
               }
            else if ( res ==-5 && !sgn && !spc ) {sgn = true;}
           else if (res ==-2) {
               idx = 0;
               n =0;
               
           }
           else 
               {n =0;
               idx = 0;
                sgn = false ;
                inf = false;
               }
           //cout <<n <<endl;
       }
        if ( n >INT_MAX  ) return (int) INT_MAX;
        if ( n < INT_MIN  ) return (int) INT_MIN;
        return (int)n;
    }
};