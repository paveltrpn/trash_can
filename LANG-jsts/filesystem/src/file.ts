
import * as fs from "fs";

export class freadSync {
    public fname: string;
    public fileContent: string;

    constructor(name?: string) {
        if (name) {
            this.fname = name;
            this.fileContent = fs.readFileSync(name,"utf8"); 
        } else {
            this.fname = "unknown file";
            this.fileContent = "empty"; 
        }
        
    }

    printFileStats(name: string) {
        let stats: fs.Stats;

        try {
            stats = fs.statSync (name)
        } catch (err) {
            console.error(err)
        }

        console.log(stats)
    }
}
