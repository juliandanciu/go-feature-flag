package flags

import (
	"github.com/nikunjy/rules/parser"
	"hash/fnv"

	"github.com/thomaspoignant/go-feature-flag/ffuser"
)

// Flag describe the fields of a flag.
type Flag struct {
	Disable    bool
	Rule       string
	Percentage int
	True       interface{} // Value if Rule applied, and in percentage
	False      interface{} // Value if Rule applied and not in percentage
	Default    interface{} // Value if Rule does not applied
}

// Value is returning the Value associate to the flag (True or False) based
// if the toggle apply to the user or not.
func (f *Flag) Value(flagName string, user ffuser.User) interface{} {
	inRule := f.evaluateRule(user)
	inPercentage := f.isInPercentage(flagName, user)

	if inRule {
		if inPercentage {
			// Rule applied and user in the cohort.
			return f.True
		}
		// Rule applied and user not in the cohort.
		return f.False
	}

	// Default Value is used if the rule does not applied to the user.
	return f.Default
}

// isInPercentage check if the user is in the cohort for the toggle.
func (f *Flag) isInPercentage(flagName string, user ffuser.User) bool {
	// 100%
	if f.Percentage == 100 {
		return true
	}

	// 0%
	if f.Percentage == 0 {
		return false
	}

	hashID := Hash(flagName+user.GetKey()) % 100
	return hashID < uint32(f.Percentage)
}

// evaluateRule is checking if the rule can apply to a specific user.
func (f *Flag) evaluateRule(user ffuser.User) bool {
	// Flag disable we cannot apply it.
	if f.Disable {
		return false
	}

	// No rule means that all user can be impacted.
	if f.Rule == "" {
		return true
	}

	// Evaluate the rule on the user.
	return parser.Evaluate(f.Rule, UserToJSON(user))
}

// Hash is taking a string and convert.
func Hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	// if we have a problem to get the hash we return 0
	if err != nil {
		return 0
	}
	return h.Sum32()
}

// UserToJSON convert the user to a MAP to use the query on it.
func UserToJSON(u ffuser.User) map[string]interface{} {
	// We don't have a json copy of the user.
	userCopy := make(map[string]interface{})

	// Duplicate the map to keep User un-mutable
	for key, value := range u.GetCustom() {
		userCopy[key] = value
	}
	userCopy["anonymous"] = u.IsAnonymous()
	userCopy["key"] = u.GetKey()
	return userCopy
}