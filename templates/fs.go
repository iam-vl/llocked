package templates

import "embed"

// The directive tells the embed package that we want to embed some files at compile time
// and store those in a var. * is a glob pattern.
// Alternatives: '*.gohtml', 'images/*.{jpg,png}'
// Can access the embedded files via the FS variable.

//go:embed *
var FS embed.FS
