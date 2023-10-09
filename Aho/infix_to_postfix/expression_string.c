
#include "string.h"
#include "expression_string.h"

/** 
  * Member function expressionString_s->buildExpressionString().
  * 
  * @param *str: String to store in new expressionString_s instance.
  * @return Return a new expressionString_s instance.
*/
expressionString_s buildExpressionString(const char* str) {
    expressionString_s rt;

    if (strlen(str) > ES_BUFLEN-1) {
        printf("buildExpressionString(): size of expression larger than buffer!\n");
        exit(0);
    };

    rt.pos = 0;
    strcpy(rt.string, str);

    return rt;
}

/** 
  * Member function expressionString_s->getNextChar().
  * 
  * @param *str - Pointer to instance of expressionString_s.
  * @return Return chracter at str->pos in str->string 
  *               and increment a str->pos.
*/
char getNextChar(expressionString_s *str) {
    if (str->pos > strlen(str->string)) {
        return '\0';
    }

    char rt = str->string[str->pos];
    str->pos += 1;
    return rt;
}

/** 
  * Member function expressionString_s->setPos()
  * 
  * @param *str: Pointer to instance of expressionString_s.
  * @param pos: New str->pos value.
*/
void setPos(expressionString_s *str, int pos) {
    if (pos > strlen(str->string) | (pos < 0)) {
        return;
    }

    str->pos = pos;
}

/** 
  * Member function expressionString_s->loadNewExpresionString()
  * 
  * @param *str: Pointer to instance of expressionString_s.
  * @param *newStr: New string to store in str->string.
*/
void loadNewExpresionString(expressionString_s *str, char *newStr) {
    if (strlen(newStr) > ES_BUFLEN-1) {
        printf("loadNewExpresionString(): size of expression larger than buffer!\n");
        return;
    }

    memset(str->string, 0, sizeof(char)*ES_BUFLEN);

    strcpy(str->string, newStr);
}
