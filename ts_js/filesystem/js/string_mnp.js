"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isDigit = exports.isAlpha = exports.splitBySpace = exports.removePunctuation = void 0;
function removePunctuation(str) {
    return str.replace(/[/.,!?;]*/g, "");
}
exports.removePunctuation = removePunctuation;
function splitBySpace(str) {
    return str.split(" ");
}
exports.splitBySpace = splitBySpace;
function isAlpha(ch) {
    return typeof ch === "string" && ch.length === 1
        && (ch >= "a" && ch <= "z" || ch >= "A" && ch <= "Z");
}
exports.isAlpha = isAlpha;
function isDigit(ch) {
    return typeof ch === "string" && ch.length === 1
        && (ch >= "0" && ch <= "9");
}
exports.isDigit = isDigit;
