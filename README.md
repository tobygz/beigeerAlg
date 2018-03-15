# beigeerAlg
贝格尔编排法_golang version


go run beigeer.go
output:

[0 6 7 1 2 3 4 5]
[2 3 4 5 6 7 1 0]
[0 7 1 2 3 4 5 6]
[3 4 5 6 7 1 2 0]
[0 1 2 3 4 5 6 7]
[4 5 6 7 1 2 3 0]
[0 2 3 4 5 6 7 1]



round  0
0  vs  5
6  vs  4
7  vs  3
1  vs  2


round  1
2  vs  0
3  vs  1
4  vs  7
5  vs  6


round  2
0  vs  6
7  vs  5
1  vs  4
2  vs  3


round  3
3  vs  0
4  vs  2
5  vs  1
6  vs  7


round  4
0  vs  7
1  vs  6
2  vs  5
3  vs  4


round  5
4  vs  0
5  vs  3
6  vs  2
7  vs  1


round  6
0  vs  1
2  vs  7
3  vs  6
4  vs  5
