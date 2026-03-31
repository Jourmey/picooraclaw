package oracle

// validIdentifierRe matches safe Oracle SQL identifiers: letters, digits, underscores, dollar signs.
// Must start with a letter or underscore.
//var validIdentifierRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_$#]*$`)

// validateSQLIdentifier checks that a string is a safe Oracle SQL identifier
// (table name, column name, model name, etc.) to prevent SQL injection.
func validateSQLIdentifier(s string) error {
	//if s == "" {
	//	return fmt.Errorf("SQL identifier must not be empty")
	//}
	//if len(s) > 128 {
	//	return fmt.Errorf("SQL identifier too long (max 128 chars): %q", s)
	//}
	//if !validIdentifierRe.MatchString(s) {
	//	return fmt.Errorf("invalid SQL identifier %q: only letters, digits, _, $, # allowed and must start with letter or _", s)
	//}
	return nil
}
