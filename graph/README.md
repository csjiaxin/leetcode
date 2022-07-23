https://leetcode.com/discuss/study-guide/1326900/graph-algorithms-problems-to-practice
https://leetcode.com/discuss/interview-question/753236/List-of-graph-algorithms-for-coding-interview


Floyd-Warshall
1334 https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
399 https://leetcode.com/problems/evaluate-division/
787 https://leetcode.com/problems/cheapest-flights-within-k-stops/
1462 https://leetcode.com/problems/course-schedule-iv/
1617 https://leetcode.com/problems/count-subtrees-with-max-distance-between-cities/




Shortest Path Problem
1 BFS(unweighted)
2 Dijstra( Directed, Single source shortest path/ Non-Negative edge O(E*lg V))

Connectivity
1 union find
2 BFS/DFS

Negative cycle(weighted digraph)
1 Bellman-Ford
2 Floyd-Warshall

Topological sort
1 Picking cherry from the leaves to root, O(V+E)
1.1 DFS
1.2 indgree 

DAG(Direct Acyclic graph)
1 shortest path
  Single source - topological sort O(V+E)
2 longest path
    weight * -1 -> shortest path


Dijstra
1 Lazy 
2 Eager using Index Priority Queue

SSSP
1 Bellman-Ford O(EV) (detect negative cycles)
2 Disjstra O(E+V)*lgV ( No negative edge)

APSP( All pair shortest path)
1 Floyd-Warshall O(V3) DP[k][i][j]


Bridge/ Cut edges 
Articulation point/ Cut vertex

