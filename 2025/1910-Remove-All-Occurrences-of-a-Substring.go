func removeOccurrences(s string, part string) string {
    for strings.Contains(s,part){
        start := strings.Index(s,part)
        s = s[:start] + s[start+len(part):]
    }
    return s
}