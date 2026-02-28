package namevalue

// StringAny is the backward-compatible alias for Instance[string, any].
// Use this where the old non-generic Instance was used.
type StringAny = Instance[string, any]

// StringString represents a name-value pair where both are strings.
type StringString = Instance[string, string]

// StringInt represents a name-value pair with a string key and int value.
type StringInt = Instance[string, int]

// StringMapAny represents a name-value pair with a string key and map[string]any value.
type StringMapAny = Instance[string, map[string]any]

// StringMapString represents a name-value pair with a string key and map[string]string value.
type StringMapString = Instance[string, map[string]string]
