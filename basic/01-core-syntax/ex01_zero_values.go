package core

// Title: Zero Values vs Missing Values
// Context: You are designing an update function for a user settings datastore.
// Users can patch their settings. In typical JSON/API payloads, an omitted basic field 
// means "do not change", while an explicitly set value (even the zero value, like `false` or `0`) 
// means "update to this value".
// Standard primitives in Go don't allow distinguishing between "missing" and "zero value" naturally.
// 
// Why this matters: In production, accidental overrides to zero values (e.g., setting age=0, 
// or isActive=false) when a partial update was intended is a ubiquitous bug.
// 
// Requirements:
// 1. `UserSettingsPatch` should be able to represent the presence or absence of fields,
//    including `int` and `bool`, without using pointers. (Pointers come with GC overhead 
//    and escape analysis implications which we want to avoid for this critical path).
//    We provide `OptionalBool` and `OptionalInt` structs for this purpose.
// 2. Implement `ApplyPatch` to merge the patch into existing settings correctly.
// 3. The `ApplyPatch` function should leave fields untouched if they weren't explicitly set.
// 4. Do not change the `UserSettings` base struct.

type UserSettings struct {
	ID            string
	Notifications bool
	Retries       int
}

type OptionalBool struct {
	Value bool
	Valid bool // True if explicitly set
}

type OptionalInt struct {
	Value int
	Valid bool // True if explicitly set
}

// UserSettingsPatch uses value wrappers capable of distinguishing zero values from unset fields.
type UserSettingsPatch struct {
	Notifications OptionalBool
	Retries       OptionalInt
}

// ApplyPatch applies valid fields from patch to base settings.
// It returns a newly constructed UserSettings safely.
func ApplyPatch(base UserSettings, patch UserSettingsPatch) UserSettings {
	// TODO: Implement the merging logic safely.
	// Only override base fields if the patch fields are marked as Valid.
	
	// Boilerplate currently returns the base unchanged.
	res := base
	
	return res
}
