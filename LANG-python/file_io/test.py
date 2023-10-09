
from common import checkEncoding
from text_read import textRead, textReadAsBytes

def main() -> None:
    foo = textRead("assets/raven.txt")
    print(foo[100:200])

    bar = textReadAsBytes("assets/raven.txt")
    print(bar[10:110])

if __name__ == "__main__":
    main()