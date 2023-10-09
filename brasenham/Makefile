
CC       := g++-10
CFLAGS   := --std=c++17
LIBS 	 := 

SRCS := brasenham.cpp canvas.cpp
OBJS := $(SRCS:.cpp=.o)

.PHONY: clean all

all: compile link

compile: brasenham.o canvas.o

link: 
	$(CC) $(OBJS) -o main -lstdc++ -ljpeg
	
brasenham.o: brasenham.cpp
	$(CC) $(CFLAGS) -c brasenham.cpp

canvas.o: canvas.cpp
	$(CC) $(CFLAGS) -c canvas.cpp

clean:
	rm -f *.o main