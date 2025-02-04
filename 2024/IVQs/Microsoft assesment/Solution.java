
/**
 * 
 * Navigate a matrix with dimensions N*M, where each cell contains a value.
 * Starting from the cell at (1,1) and aiming for the destination at (N,M),
 * the challenge involves avoiding consecutive values with a specified
 * constraint X.
 * Movement is restricted to cells with values other than X+1.
 * Assess the feasibility of reaching the destination while adhering to these
 * constraints,
 * returning the length of the shortest path if possible, or -1 if not
 * achievable.
 * Movement options are limited to the four cardinal directions: up, down, left,
 * and right.
 * 
 * 
 * 1 2 3 4
 * 1 1 0 0
 * 1 2 2 5
 * 
 * 
 * vector<vector<int>>paths
 * vector<int>path;
 * 
 * dfs(){
 *  if(Reached){
 *      paths.push_back(path);
 *      path.pop_back();
 *      return;
 *      } 
 * 
 *   path.push_back(number);
 *   dfs();
 *   path.pop_back();
 * }
 */
public class Solution {
    public int count = 0;
    public void solve(int n, int m, int[][] matrix) {
        boolean visited[][] = new boolean[n+1][m+1];
        dfs(matrix, visited, 1, 1);
        System.out.println(count);
    }

    private void dfs(int[][] matrix, boolean[][] visited, int i, int j) {
        if (i > matrix.length || j > matrix[0].length || i < 1 || j < 1 || visited[i][j])
            return;
        visited[i][j] = true;
        int current = matrix[i][j];
        int rightNeighbor = matrix[i][j + 1];
        int downNeighbor = matrix[i + 1][j];
        int leftNeighbor = matrix[i][j - 1];
        int upNeighbor = matrix[i - 1][j];
        // out of bound
        if (current != rightNeighbor + 1) {
            dfs(matrix, visited, i, j + 1);
            count++;
        } else if (current != downNeighbor + 1) {
            dfs(matrix, visited, i + 1, j);
            count++;
        } else if (current != leftNeighbor + 1) {
            count++;
            dfs(matrix, visited, i, j - 1);
        } else if (current != upNeighbor + 1) {
            count++;
            dfs(matrix, visited, i - 1, j);
        }
    }
}

/**
 * 
 * 
 * 
 * 
 * Given a string S containing digits from 0 to 9, 
 * determine the count of occurrences of the shortest substring that 
 * encompasses all numbers from 0 to 10.
 */


 /**
  * 


  #include <iostream>
#include <vector>
#include <queue>
#include<bits/stdc++.h>
using namespace std;

struct Node{
  int x, y;
  long long dist;
};
bool operator<(const Node& a, const Node& b){
  return a.dist > b.dist;
}

// Where' X ????
int main() 
{
    int n, m;
    cin >> n >> m;
    vector<vector<int>>v(n, vector<int>(m));
    for(auto& x : v)
      for(auto& y : x)
        cin >> y;
        
    priority_queue<Node>q;
    q.push({0,0, 0});
    vector<vector<bool>>vis(n + 5, vector<bool>(m + 5));
    vis[0][0] = true;
    auto valid = [&](int i, int j){
        return i >= 0 and i < n and j >= 0 and j < m;
    };
    
    vector<int>dx{-1, 0, 1, 0};
    vector<int>dy{0, 1, 0, -1};
    
    while(!q.empty()){
      int i = q.top().x, j = q.top().y;
      int dist = q.top().dist;
      q.pop();
      if(i == n - 1 and j == m - 1){
        cout << dist;
        return 0;
      }
      for(int d = 0; d < 4; d++){
        int newI = i + dx[d], newJ = j + dy[d];
        if(valid(newI, newJ)){
              if(v[newI][newJ] == v[i][j] + 1 or vis[newI][newJ])
                continue;
              q.push({newI, newJ, dist + 1});
              vis[newI][newJ] = true;
        }
      } 
    }

    cout << -1;
    return 0;
}

  */