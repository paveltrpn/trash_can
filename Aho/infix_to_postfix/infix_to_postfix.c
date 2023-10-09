
#include <stdio.h>
#include <ctype.h>
#include <stdint.h>
#include <stdbool.h>

#include "expression_string.h"

expressionString_s foo;
int32_t lookahead = 0;

void expr() {
    term();
    while (true) {
        if (lookahead == '+') {
            match('+');
            term();
            printf("+");
        } else if (lookahead == '-') {
            match('-');
            term();
            printf("-");
        } else {
            return;
        }
    }
}

void term() {
        if (isdigit(lookahead)) {
            char t = lookahead;
            match(lookahead);
            printf("%c", t);
        } else {
            printf("syntax error in term()! %c - is not number!\n", (lookahead));
            exit(0);
        }
}

void match(int32_t t) {
    if (lookahead == (char)t) {
        lookahead = getNextChar(&foo);
    } else {
        printf("Syntax error in match!\n");
        exit(0);
    }
}

int main(int argc, char **argv) {
    foo = buildExpressionString("9+2-5");
    printf("%s\n", foo.string);

    lookahead = getNextChar(&foo);
    expr();
    printf("\n");
}