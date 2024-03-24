void main(List<String> args) {
  print(subtractProductAndSum(10));
}

int subtractProductAndSum(int n) {
  var product = 1;
  var sum = 0;

  do {
    sum += (n % 10);
    product *= (n % 10);
    n = n ~/ 10;
  } while (n != 0);

  return product - sum;
}
