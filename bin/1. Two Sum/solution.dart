void main(List<String> args) {
  print(twoSum([3, 2, 4], 6));
}

List<int> twoSum(List<int> nums, int target) {
  var numbersMap = <int, int>{};
  for (var i = 0; i < nums.length; i++) {
    var otherNumber = target - nums[i];
    if (numbersMap.containsKey(otherNumber)) {
      return [numbersMap[otherNumber], i];
    }
    numbersMap[nums[i]] = i;
  }
  return [];
}
