void main(List<String> args) {
  print(detectCapitalUse('TeEsT'));
}

bool detectCapitalUse(String word) {
  // I have 3 cases , Either it's all capital , all small , or first one is capital
  final allCapitalLettersLength =
      word.codeUnits.where((element) => element < 97);
  if (allCapitalLettersLength.length == word.length ||
      allCapitalLettersLength.isEmpty) {
    return true;
  } else {
    if (allCapitalLettersLength.length == 1 && word.codeUnitAt(0) < 97) {
      return true;
    }
    return false;
  }
}
