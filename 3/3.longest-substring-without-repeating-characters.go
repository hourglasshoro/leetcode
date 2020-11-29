package main

/**
 * <p>Given a string <code>s</code>, find the length of the <b>longest substring</b> without repeating characters.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;abcabcbb&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> The answer is &quot;abc&quot;, with the length of 3.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;bbbbb&quot;
<strong>Output:</strong> 1
<strong>Explanation:</strong> The answer is &quot;b&quot;, with the length of 1.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;pwwkew&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> The answer is &quot;wke&quot;, with the length of 3.
Notice that the answer must be a substring, &quot;pwke&quot; is a subsequence and not a substring.
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;&quot;
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= s.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>s</code> consists of English letters, digits, symbols and spaces.</li>
</ul>

**/
/**
 * "abcabcbb"
**/
func lengthOfLongestSubstring(s string) int {

	// My Solution
	// NOTE: test6に対応できない
	//strMap := map[rune]bool{}
	//var count, max int
	//for _, r := range s {
	//	_, ok := strMap[r]
	//	if ok {
	//		count = 1
	//		strMap = map[rune]bool{}
	//	} else {
	//		count++
	//	}
	//	if max < count {
	//		max = count
	//	}
	//	strMap[r] = true
	//}
	//return max

	// Model answer
	// Approach 2: Sliding Window
	//strMap := map[rune]bool{}
	//var i, j, ans int
	//for i < len(s) && j < len(s) {
	//	_, ok := strMap[rune(s[j])]
	//	if !ok {
	//		strMap[rune(s[j])] = true
	//		j++
	//		if ans < j-i {
	//			ans = j - i
	//		}
	//	} else {
	//		delete(strMap, rune(s[i]))
	//		i++
	//	}
	//	log.Printf("%v,%v,%v,%v", s, i, j, ans)
	//}
	//return ans

	// Approach 3: Sliding Window Optimized
	strMap := map[rune]int{}
	var i, ans int
	for j, r := range s {
		index, ok := strMap[r]
		if ok && i < index+1 {
			// NOTE: s[j]と同じ文字で前回でできたindexの次までwindowの左端を移動させる
			i = index + 1
		}
		diff := j - i + 1
		if ans < diff {
			ans = diff
		}
		strMap[r] = j
		//log.Printf("%v,%v,%v,%v", s, i, j, ans)
	}
	return ans
}
