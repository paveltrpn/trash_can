import * as process from "process"
import {exprString_c} from "./expression_string.js"

/**@type {string} */
let lookahead
let expressionLine = new exprString_c("9+5-2")
let expressionOut = new exprString_c("")

function expr() {
    term()
    while (true) {
        if (lookahead == "+") {
            match("+")
            term()
            expressionOut.appendChar("+")
        } else if (lookahead == "-") {
            match("-")
            term()
            expressionOut.appendChar("-")
        } else {
            return
        }
    }
}

function term() {
    if (!isNaN(Number(lookahead))) {
        let t = lookahead
        match(lookahead)
        expressionOut.appendChar(t)
    } else {
        console.log("syntax error in term()! " + lookahead + "- is not number!")
        process.exit(0)
    }
}

/**
 * @param {string} t
 */
function match(t) {
    if (lookahead == t) {
        [, lookahead] = expressionLine.getNextChar()
    } else {
        console.log("syntax error in match()!")
        process.exit(0)
    }
}

(function main() {
    [, lookahead] = expressionLine.getNextChar()

    expr()

    console.log(expressionOut.expression)
}) ()
