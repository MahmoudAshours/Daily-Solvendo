void main(List<String> args) {
  print(minDeletionSize(['rrjk', 'furt', 'guzm']));
}
int minDeletionSize(List<String> strs) {
  var deletedColumnsCount = 0;
  for (var i = 0; i < strs[0].length; i++) {
    for (var j = 0; j < strs.length - 1; j++) {
      if (strs[j][i].codeUnits[0] > strs[j + 1][i].codeUnits[0]) {
        deletedColumnsCount++;
        break;
      }
    }
  }
  return deletedColumnsCount;
}
