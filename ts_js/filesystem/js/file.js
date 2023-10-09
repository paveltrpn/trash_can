"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.freadSync = void 0;
const fs = require("fs");
class freadSync {
    constructor(name) {
        if (name) {
            this.fname = name;
            this.fileContent = fs.readFileSync(name, "utf8");
        }
        else {
            this.fname = "unknown file";
            this.fileContent = "empty";
        }
    }
    printFileStats(name) {
        let stats;
        try {
            stats = fs.statSync(name);
        }
        catch (err) {
            console.error(err);
        }
        console.log(stats);
    }
}
exports.freadSync = freadSync;
