CC       := clang++-10
CFLAGS   := 
LFLAGS 	 := -lglut -lGL

SRCS := marchingsource.cpp
		
OBJS := marchingsource.o
		
#OBJS := $(SRCS:.cpp=.o)

.PHONY: clean all

all: $(OBJS) main

marchingsource.o: marchingsource.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c marchingsource.cpp

main: $(OBJS)
	$(CC) $(OBJS) $(LIBDIR) -o main $(LFLAGS)

clean:
	rm -f *.o main