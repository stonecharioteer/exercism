package parsinglogfiles
import "regexp"

func IsValidLine(text string) bool {
	var validPrefix = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return validPrefix.MatchString(text)
}

func SplitLogLine(text string) []string {
	var sep = regexp.MustCompile(`<[~*=-]*>`)
    return sep.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var quotedPassword = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    n := 0
    for _, line := range lines {
        if quotedPassword.MatchString(line) {
            n++
        }
    }
    return n
}

func RemoveEndOfLineText(text string) string {
	var eol = regexp.MustCompile(`end-of-line\d+`)
    return eol.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var userRe = regexp.MustCompile(`User\s+(\S+)`)
    out := make([]string, len(lines))
    for i, line := range lines {
        if m:= userRe.FindStringSubmatch(line); m!=nil {
            username := m[1]
            out[i] = "[USR] " + username + " " + line
        } else {
            out[i] = line
        }
    }
    return out
}
