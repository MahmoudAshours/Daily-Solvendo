import 'dart:math';

void main(List<String> args) {
  print(categorizeBox(92487, 6200, 58423, 40));
}

String categorizeBox(int length, int width, int height, int mass) {
  /**
   *
   *  The box is "Bulky" if:
  Any of the dimensions of the box is greater or equal to 104.
  Or, the volume of the box is greater or equal to 109.
   */
  var isOverThanOrEqual =
      (width.toString() + length.toString() + height.toString()).length - 3 >=
          8;

  if ((length >= pow(10, 4) ||
          width >= pow(10, 4) ||
          height >= pow(10, 4) ||
          isOverThanOrEqual) &&
      mass >= 100) {
    return 'Both';
  } else if (length >= pow(10, 4) ||
      width >= pow(10, 4) ||
      height >= pow(10, 4) ||
      isOverThanOrEqual) {
    return 'Bulky';
  } else if (mass >= 100) {
    return 'Heavy';
  } else {
    return 'Neither';
  }
}
