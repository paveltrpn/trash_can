
from expression_string import *


foo: expressionString_c = expressionString_c("9+5-2")
lookahead: str = ""


def expr() -> None:
    term()
    while True:
        if lookahead == "+":
            match("+")
            term()
            print("+", end="")
        elif lookahead == "-":
            match("-")
            term()
            print("-", end="")
        else:
            return

def term() -> None:
    if lookahead.isnumeric():
        t = lookahead
        match(lookahead)
        print(t, end="")
    else:
        print("syntax error in term()! {} - is not number!".format(lookahead))
        exit()

def match(t: str) -> None:
    global lookahead
    if lookahead == t:
        lookahead = foo.getNextChar()[1]
    else:
        print("syntax error in match!")
        exit(0)


def main() -> None:
    print("read: ", end="")
    inputStr = input()

    if len(inputStr) > 0:
        foo.loadNewString(inputStr)

    global lookahead
    lookahead = foo.getNextChar()[1]

    expr()
    print()

if __name__ == "__main__":
    main()