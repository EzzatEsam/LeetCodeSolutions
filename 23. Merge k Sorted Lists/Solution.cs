/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode MergeKLists(ListNode[] lists) {
        if ( lists == null || lists.Length ==0  ) return null;
        var verified = 0; 
        var tempList = lists.ToList();
        while (verified != tempList.Count ) {
            verified = 0;
            for( int i =0; i < tempList.Count ; i++) { 
                var lst = tempList[i];
                if ( lst == null)  {
                    //Console.WriteLine(tempList.Count);
                    tempList.RemoveAt(i);
                    //Console.WriteLine(tempList.Count);
                    verified = 0;
                    break;
                }
                else
                    verified ++;
            }
        }
        
        if (tempList.Count == 0) return null;
        
        var myList = new List<int>();
         
        for( int i =0; i < tempList.Count ; i++) { 
            
            while (tempList[i] != null)
            {
                var lst = tempList[i];
              myList.Add(lst.val);
                tempList[i] = lst.next;  
            }
        } 
        
        var ot = new ListNode();
        var crnt = ot;
        myList.Sort();
        for ( int i = 0; i < myList.Count ; i ++) {
            crnt.val = myList[i];
            if ( i != myList.Count -1 ) {
                crnt.next = new ListNode();
                crnt = crnt.next;
                    
            }
        }
        return  ot ;
    }
    
}