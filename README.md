# beigeerAlg
贝格尔编排法_golang version
<br>
<br>
go run beigeer.go<br>
output:<br><br>

[0 6 7 1 2 3 4 5]<br>
[2 3 4 5 6 7 1 0]<br>
[0 7 1 2 3 4 5 6]<br>
[3 4 5 6 7 1 2 0]<br>
[0 1 2 3 4 5 6 7]<br>
[4 5 6 7 1 2 3 0]<br>
[0 2 3 4 5 6 7 1]<br>



round  0<br>
0  vs  5<br>
6  vs  4<br>
7  vs  3<br>
1  vs  2<br>
<br>
<br>
round  1<br>
2  vs  0<br>
3  vs  1<br>
4  vs  7<br>
5  vs  6<br>
<br>
<br>
round  2<br>
0  vs  6<br>
7  vs  5<br>
1  vs  4<br>
2  vs  3<br>
<br>
<br>
round  3<br>
3  vs  0<br>
4  vs  2<br>
5  vs  1<br>
6  vs  7<br>
<br>
<br>
round  4<br>
0  vs  7<br>
1  vs  6<br>
2  vs  5<br>
3  vs  4<br>
<br>
<br>
round  5<br>
4  vs  0<br>
5  vs  3<br>
6  vs  2<br>
7  vs  1<br>
<br>
<br>
round  6<br>
0  vs  1<br>
2  vs  7<br>
3  vs  6<br>
4  vs  5<br>
