export function parseSponsorTextVariables(text: string): Record<string, string> {
    const start = text.indexOf("$vars:")
    if (start == -1) {
        return {}
    }

    const variables: Record<string, string> = {}
    let key = ""
    let next: "=" | ";" = "="
    let escaped = false
    for (let i = start + 6; i < text.length; i++) {
        const char = text[i]

        // Stop parsing on end marker
        if (!escaped && char == "$") {
            break
        }

        // Detect and skip escape characters (unless they are escaped themselves)
        if (!escaped && char == "\\") {
            escaped = true
            continue
        }

        // Skip (irrelevant) whitespaces
        if (!escaped && char == " " && (
            next == "=" // key whitespace (keys do not permit whitespaces)
            || next == ";" && variables[key] == "" // leading value whitespace
            || next == ";" && text.length > i + 1 && text[i + 1] == " " // consecutive value whitespace
            || next == ";" && text.length > i + 1 && text[i + 1] == ";" // trailing value whitespace
            || next == ";" && text.length > i + 1 && text[i + 1] == "$" // global trailing whitespace (with suffix)
            || next == ";" && text.length == i + 1 // global trailing whitespace (no suffix)
        )) {
            continue
        }

        const marker = !escaped && char == next
        if (marker && next == "=") {
            // Found colon after key, add key to result and switch to reading value
            variables[key] = ""
            next = ";"
        } else if (marker && next == ";") {
            // Found semicolon after value, start new key and switch to reading key
            key = ""
            next = "="
        } else if (next == "=") {
            // Currently reading key, append char to key
            key += char
        } else {
            // Currently reading value, append char to value
            variables[key] += char
        }

        // Reset escaped flag
        escaped = false
    }

    return variables
}
