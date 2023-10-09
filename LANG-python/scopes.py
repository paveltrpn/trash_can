
global_var = "global"
global_lst = ["first", "second"]

def foo() -> None:
    global global_var
    print(global_var)
    global_var = "global local"

    local_var = "local var"

    global global_lst
    global_lst = ["rediclarated first", "rediclarated second"]

    def inner() -> None:
        global global_var
        nonlocal local_var
        local_var = "rediclared"
        global_var = global_var + " concatanated"
        print(local_var)

    inner()
    print(local_var)
    print(global_var)
    print(global_lst)


def main() -> None:
    global global_lst
    foo()
    print(global_lst)


if __name__ == "__main__":
    main()