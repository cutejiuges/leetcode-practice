package com.github.leetcode.LC838;

public class Solution {
    public String pushDominoes(String dominoes) {
        //处理下原始输入的边界，在最左边放一个向左推的，最右边放一个向右推的，不影响结果
        String myDominoes = "L" + dominoes + "R";
        StringBuilder ans = new StringBuilder();

        // 1. 初始的L、R是一定不变的
        // 2. 两个R之间的，全部改成R，两个L之间的，全部改成L
        // 3. 左L和右R之间的不变，左R和右L之间的往中间靠
        int left = 0, right = 1;
        while (right < myDominoes.length()) {
            while (right < myDominoes.length() && myDominoes.charAt(right) == '.') {
                right++;
            }
            char leftChar = myDominoes.charAt(left), rightChar = myDominoes.charAt(right);
            if (leftChar == rightChar) {
                for (int i = left+1; i < right; i++) {
                    ans.append(leftChar);
                }
            } else if (leftChar == 'R') {
                int halfLen = (right - left - 1) >> 1;
                for (int i = 0; i < halfLen; i++) {
                    ans.append('R');
                }
                if (((right - left) & 1) == 0) {
                    ans.append('.');
                }
                for (int i = 0; i < halfLen; i++) {
                    ans.append('L');
                }
            } else {
                for (int i = left+1; i < right; i++) {
                    ans.append('.');
                }
            }
            ans.append(rightChar);
            left = right;
            right++;
        }
        return ans.substring(0, ans.length()-1);
    }
}
