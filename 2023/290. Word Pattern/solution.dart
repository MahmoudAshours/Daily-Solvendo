void main(List<String> args) {
  print(wordPattern('abba', 'dog cat cat dog'));
}

bool wordPattern(String pattern, String s) {
  // ignore: omit_local_variable_types
  Map patternStringMap = {};
  final wordString = s.split(' ');
  if (wordString.length != pattern.length) return false;
  for (var i = 0; i < pattern.length; i++) {
    if (patternStringMap.containsKey(wordString[i])) {
      if (patternStringMap[wordString[i]] != pattern[i]) {
        return false;
      }
    } else {
      if (!patternStringMap.containsValue(pattern[i])) {
        patternStringMap[wordString[i]] = pattern[i];
      } else {
        return false;
      }
    }
  }
  return true;
}
