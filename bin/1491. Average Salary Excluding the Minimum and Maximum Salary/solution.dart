void main() {
  print(average([1, 2, 3, 4, 5]));
}

double average(List<int> salary) {
  salary.sort();
  salary.removeAt(0);
  salary.removeAt(salary.length - 1);
  var sum = 0;
  for (var i = 0; i < salary.length; i++) {
    sum += salary[i];
  }
  return sum / salary.length;
}
