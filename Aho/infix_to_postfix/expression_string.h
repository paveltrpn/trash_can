
#ifndef __expression_string_h__
#define __expression_string_h__

#define ES_BUFLEN 128

typedef struct {
    int pos;
    char string[ES_BUFLEN];
} expressionString_s;

expressionString_s      buildExpressionString(const char* str);
char                    getNextChar(expressionString_s *str);
void                    setPos(expressionString_s *str, int pos);
void                    loadNewExpresionString(expressionString_s *str, char *newStr);


#endif