/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* res = new ListNode();
        ListNode* nxt1 = l1->next;
        ListNode* nxt2 = l2->next;
        res->val = (l1->val + l2->val)%10;
        ListNode* resFirst = res;
        int carry = ((l1->val + l2->val) >9 )? 1 :0; 
        while (nxt1   || nxt2   || carry ==1) {
            int vl1 = (!nxt1 ) ? 0 : nxt1->val;
            int vl2 = (!nxt2 ) ? 0 : nxt2->val;    
            int sm = carry + vl1 + vl2;
            carry = sm >9 ? 1 :0; 
            sm = sm%10;
            res->next = new ListNode(sm ,nullptr);
            res = res->next;
            nxt1 = (!nxt1 ) ?  nullptr : nxt1->next;
            nxt2 = (!nxt2 ) ?  nullptr : nxt2->next;
        }
        return resFirst; 
    }
};