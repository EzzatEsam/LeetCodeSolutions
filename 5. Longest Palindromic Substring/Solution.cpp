class Solution
{
public:
    string longestPalindrome(string s)
    {
        int lower = 0, upper = 0, max = 0, nxt = 0;
        for (int i = 0; i < s.length(); i++)
        {
            for (int j = s.length()-1; j > i; j--)
            {
                //cout << i << " " << j << endl;
                if (s[i] == s[j])
                {
                    bool fnd = 1;
                    int st = i;
                    int ed = j;
                    while (ed > st )
                    {
                        //cout << st << " " << ed << endl;
                        if (s[st] == s[ed] )
                        { 
                                st++;
                                ed--;
                        }
                        else
                        {
                            fnd = 0;
                            break;
                        }
                    }
                    if (fnd && j - i > max)
                    {
                        lower = i;
                        upper = j;
                        max = j - i;
                        break;
                    }
                }
            }
        }
        //cout << upper << " " << lower << endl;
        return s.substr(lower, upper - lower + 1);
    }
};