
#  Даны списки:
#  a = [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89];
#  b = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13].
#  Нужно вернуть список, который состоит из элементов, общих для этих двух списков.

a = [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89]
b = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]
foo: list[int]

def naive() -> None:
    for elemA in a:
        for elemB in b:
            if elemA == elemB:
                foo.append(elemA)

    print(foo)

def better():
    result = list(set(a) & set(b))
    print(result)

def main() -> None:
    naive()
    better()
    foo2(100)

def foo2(bar: int) -> None:
    print(bar)
    # return 100

if __name__ == "__main__":
    main()