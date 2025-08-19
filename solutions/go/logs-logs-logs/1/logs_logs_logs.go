package logs
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v:= range log {
        switch v {
            case '❗':
            	return "recommendation"
            case '🔍':
         	   return "search"
            case '☀':
            	return "weather"
            default:
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var out []rune
    for _, v := range log {
        switch v {
            case oldRune:
            out = append(out, newRune)
            default:
            out = append(out, v)
        }
    }
    return string(out)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
