package personname

// renderWithSeps is a generalized function for rendering PN DCM() string with the
// correct level of null separators.
func renderWithSeps(sections []string, separator string, nullSepLevel uint) string {
	// Keep track of whether we have found any non-zero values.
	nonZeroFound := false

	// We're going to build the string on this variable. Since we are building it
	// backwards, we can't use strings.Builder.
	dcmString := ""

	// It's going to be easier to write the correct values if we iterate backwards over
	// them, since we need look-ahead to know if there are interior null components
	// from a non-zero value near the end of the sections.
	for i := len(sections) - 1; i >= 0; i-- {
		// Get the section.
		section := sections[i]

		// Add the section to the dcm string
		dcmString = section + dcmString

		// If this section is non-emtpy, we need to add all remaining separators, so
		// remember that we have encountered a non-zero value.
		if section != "" {
			nonZeroFound = true
		}

		// If this is not the (true) first section, and a non-zero section has been
		// found OR this section is under the threshold for which we are rendering null
		// separators: add a separator to the head of the string (remember, we are
		// working our way backwards).
		if i > 0 && (nonZeroFound || i <= int(nullSepLevel)) {
			dcmString = separator + dcmString
		}
	}

	return dcmString
}
