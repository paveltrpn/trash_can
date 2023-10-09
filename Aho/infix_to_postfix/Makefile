CC       = gcc
CFLAGS   = -Wall -Wconversion -Wextra -Wpedantic -Wno-unused-parameter -std=c17
LFLAGS 	 = 

UNAME_S := $(shell uname -s)

INCDIR	 = 	
LIBDIR	 = 

VPATH = .

SRCS = infix_to_postfix.c expression_string.c

OBJS = $(patsubst %.c, %.o, $(SRCS))

#OBJS := $(SRCS:.cpp=.o)

.PHONY: clean all

all: $(OBJS) main

infix_to_postfix.o: infix_to_postfix.c
	$(CC) -w $(addprefix -I,$(INCDIR)) -c $< -o $@

expression_string.o: expression_string.c
	$(CC) -w $(addprefix -I,$(INCDIR)) -c $< -o $@

main: $(OBJS)
	$(CC) $(OBJS) $(addprefix -L,$(LIBDIR)) -o main $(LFLAGS)

clean:
	rm -f *.o main