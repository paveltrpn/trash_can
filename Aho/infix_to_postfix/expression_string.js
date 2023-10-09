
export class exprString_c {
    /**
     * 
     * @param {string} expr
     */
    constructor(expr) {
        this.expression = expr
        this.pos = 0
    }

    /**
     * 
     * @returns {(any | string)[]}
     */
    getNextChar() {
        if ((this.pos + 1) > (this.expression.length)) {
            return [-1, ""]
        }

        let rt = this.expression[this.pos]
        this.pos = this.pos + 1
        return [this.pos, rt]
    }

    /**
     * 
     * @param {number} npos 
     * @returns {number}
     */
    setPos(npos) {
        if ((npos > this.expression.length) || (npos < 0)) {
            return -1
        }

        this.pos = npos
        return 0
    }

    /**
     * 
     * @param {string} nexpr 
     */
    loadNewExpression(nexpr) {
        this.expression = nexpr
        this.pos = 0
    }

    getExprLen() {
        return this.expression.length
    }

    /**
     * 
     * @param {string} ch 
     */
    appendChar(ch) {
        this.expression = this.expression + ch
    }
}
