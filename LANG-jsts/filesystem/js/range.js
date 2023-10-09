"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Range = void 0;
class Range extends Array {
    static range(from, to, step) {
        return Array.from({ length: Math.floor((to - from) / step) + 1 }, (v, k) => from + k * step);
    }
}
exports.Range = Range;
