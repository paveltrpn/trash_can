
export function removePunctuation(str: string): string {
    return str.replace(/[/.,!?;]*/g, "");
}

export function splitBySpace(str: string): string[] {
    return str.split(" ");
}

export function isAlpha(ch: string): boolean {
    return typeof ch === "string" && ch.length === 1
           && (ch >= "a" && ch <= "z" || ch >= "A" && ch <= "Z");
}

export function isDigit(ch: string): boolean {
    return typeof ch === "string" && ch.length === 1
           && (ch >= "0" && ch <= "9");
}