
class expressionString_c:
    def __init__(self, line: str) -> None:
        self.pos: int = 0
        self.string: str = line

    
    def getNextChar(self) -> tuple[int, str]:
        """
        @return: Return chracter at self.pos in self.string 
        and increment a self.pos.
        """
        if self.pos+1 > len(self.string):
            return -1, ""

        rt = self.string[self.pos]
        self.pos = self.pos + 1
        return self.pos, rt

    def setPos(self, pos: int) -> int:
        """
        @param pos: New self.pos value.
        """
        if pos > len(self.string) or pos < 0:
            return -1
        self.pos = pos
        return 0

    def loadNewString(self, new: str) -> None:
        self.string = new
        self.pos = 0