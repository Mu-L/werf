	unrecognized   parserState = "unrecognized"
	diffBegin      parserState = "diffBegin"
	diffBody       parserState = "diffBody"
	newFileDiff    parserState = "newFileDiff"
	deleteFileDiff parserState = "deleteFileDiff"
	modifyFileDiff parserState = "modifyFileDiff"
	ignoreDiff     parserState = "ignoreDiff"
			return p.handleModifyFileDiff(line)
		if strings.HasPrefix(line, "new mode ") {
			return p.writeOutLine(line)
		}
		if strings.HasPrefix(line, "diff --git ") {
			return p.handleDiffBegin(line)
		}
		if strings.HasPrefix(line, "Submodule ") {
			return p.handleSubmoduleLine(line)
		}
		if strings.HasPrefix(line, "diff --git ") {
			return p.handleDiffBegin(line)
		}
		if strings.HasPrefix(line, "Submodule ") {
			return p.handleSubmoduleLine(line)
		}