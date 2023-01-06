import 'dart:math';

void main(List<String> args) {
  print(findMinArrowShots([
    [10, 16],
    [2, 8],
    [1, 6],
    [7, 12]
  ]));
}

// @TODO
int findMinArrowShots(List<List<int>> points) {
  points.sort((a, b) => a[0].compareTo(b[0]));
  print(points);
  var count = 1;
  var endingX = points[0][1];
  for (var i = 0; i < points.length - 1; i++) {
    final generticPoint = points[i + 1];
    // Intersects
    if (endingX >= generticPoint[0]) {
      endingX = min(endingX, generticPoint[1]);
    } else {
      count += 1;
      endingX = generticPoint[1];
    }
  }
  return count;
}
