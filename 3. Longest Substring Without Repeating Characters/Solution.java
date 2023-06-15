class Solution {
    public int lengthOfLongestSubstring(String s) {
                
var letters = new int[100];
        int lower = 0;
        int upper = 0;
        int max =0;     
        for ( var i =0; i < s.length() ; i ++) { 
            upper = i;
            int idx = s.charAt(i) -'0' +18 ;
            if ( letters[idx] > 0) {
                if(i - lower > max) max = i - lower; 
                lower =letters[idx];
                for ( var j = 0; j < letters.length ; j++) {
                    if( letters[j] < lower)
                        letters[j] =0;
                }
            } else if (i - lower +1 > max)  max = i - lower +1; 
            
            letters[idx] = i +1;     
        }
        
        return max;
    }
}