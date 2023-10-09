
//   Кругом ограничения!
//   
//   Напишите программу на языке C, которая будет выводить числа от 1 до 1000. 
//   В программе нельзя использовать ключевые слова if, switch, do, while, for, goto 
//   и подобные, логические операции, побитовую логику. Естественно, программа не 
//   должна содержать тысячу вызовов printf, нужно компактное решение. 
//   Постарайтесь не применять макросы.

#include <stdio.h>

void naive_solution() {
    int i = 0;

    // for (i = 0; i < 1001; i++) {
    //     printf("i = %i\n", i);
    // }

    // while (i < 1001) {
    //     printf("i = %i\n", i);
    //     i++;
    // }

    do {
        printf("i = %i\n", i);
        i++;
    } while (i < 1001);
}

void goto_solution() {
    int i = 0;

    inc: printf("i = %i\n", i);
    i++;

    // if (i < 1001) goto inc;

    switch (i) {
        case 1001:
            return;

        default:
            goto inc;
    }
}

int main(int argc, char** argv) {
    printf("naive_solution()\n");
    naive_solution();

    printf("goto_solution()\n");
    goto_solution();

    return 0;
}