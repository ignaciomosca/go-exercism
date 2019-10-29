// Package twofer contains ShareWith method that implements two fer behavior
package twofer

// ShareWith returns "One for X, one for me" where X is the name received
// if name is empty it will use 'you' as default name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
