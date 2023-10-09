
// Взял вот с этой страницы - https://stackoverflow.com/questions/36947847/how-to-generate-range-of-numbers-from-0-to-n-in-es2015-only
// Range.range(0, 9, 1); - выведет - [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// Использование - Range.range( "от какого значения", 
//                              "до какого значения",
//                              "шаг" )

interface _Iterable extends Iterable<{}> {
  length: number;
}

export class Range<T> extends Array<T> {
  static range(from: number, to: number, step: number): number[] {
    return Array.from(
      <_Iterable>{ length: Math.floor((to - from) / step) + 1 },
      (v, k) => from + k * step
    );
  }
}
