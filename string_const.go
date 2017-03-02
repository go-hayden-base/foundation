package foundation

// --- public ---

const (
	OptionRandomStringNum   = 0 // nubmer only
	OptionRandomStringLower = 1 // lower char only
	OptionRandomStringUpper = 2 // upper char only
	OptionRandomStringAll   = 3 // number and all en char only
)

// --- private ---
const (
	cRegStrHttpURL = `^(?i:(http)|(https))://\S+$`
)
