
# Есть список a = [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89].
# Выведите все элементы, которые меньше 5.

foo: list = [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89]

def naive() -> None:
    for elem in foo:
        if elem < 5:
            print(elem)

def main() -> None:
    naive()

    print(list(filter(lambda elem: elem < 5, foo)))

if __name__ == "__main__":
    main()