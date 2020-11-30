package main

import (
	"strings"
)

/**
 * <p>The string <code>&quot;PAYPALISHIRING&quot;</code> is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)</p>

<pre>
P   A   H   N
A P L S I I G
Y   I   R
</pre>

<p>And then read line by line: <code>&quot;PAHNAPLSIIGYIR&quot;</code></p>

<p>Write the code that will take a string and make this conversion given a number of rows:</p>

<pre>
string convert(string s, int numRows);
</pre>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 3
<strong>Output:</strong> &quot;PAHNAPLSIIGYIR&quot;
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 4
<strong>Output:</strong> &quot;PINALSIGYAHRPI&quot;
<strong>Explanation:</strong>
P     I    N
A   L S  I G
Y A   H R
P     I
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;A&quot;, numRows = 1
<strong>Output:</strong> &quot;A&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consists of English letters (lower-case and upper-case), <code>&#39;,&#39;</code> and <code>&#39;.&#39;</code>.</li>
	<li><code>1 &lt;= numRows &lt;= 1000</code></li>
</ul>

**/
/**
 * "PAYPALISHIRING"
3
**/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	mod1 := numRows - 1
	mod2 := mod1 * 2
	resMap := map[int]string{}
	for i, r := range s {
		c1 := i % mod1
		c2 := i % mod2
		var index int

		// 斜めの線の部分
		if c2 >= mod1 {
			index = mod1 - c1
		} else {
			index = c1
		}

		resMap[index] = strings.Join([]string{resMap[index], string(r)}, "")
	}
	//log.Printf("%v,%v", s, resMap)
	result := ""
	for i := 0; i <= mod1; i++ {
		result = strings.Join([]string{result, resMap[i]}, "")
	}
	return result
}
