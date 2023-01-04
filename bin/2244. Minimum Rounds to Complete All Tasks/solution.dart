import 'dart:collection';

void main(List<String> args) {
  print(minimumRounds([2, 3, 3]));
}

int minimumRounds(List<int> tasks) {
  final tasksMap = generateTaskMap(tasks);
  var count = 0;
  for (var element in tasksMap.keys) {
    final value = tasksMap[element];
    if (value == 1) {
      return -1;
    } else {
      if (value % 3 == 0) {
        count += value ~/ 3; 
      } else if (value - (value % 3) >= 1) {
        count += value ~/ 3 + 1; 
      } else if (value % 2 == 0) {
        count += value ~/ 2; 
      }
    }
  }
  return count;
}

HashMap generateTaskMap(List<int> tasks) {
  final tasksMap = HashMap();
  for (var task in tasks) {
    if (tasksMap.containsKey(task)) {
      tasksMap[task] = tasksMap[task] + 1;
    } else {
      tasksMap[task] = 1;
    }
  }
  return tasksMap;
}
