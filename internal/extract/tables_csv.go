package extract

// CSVtablesBuild
// func (e *Extractor) CSVtablesBuild(
// 	header bool, delim string, separateTables bool) {
// 	if !separateTables {
// 		e.CSVwriterGlobal = new(strings.Builder)
// 	}
// 	// Header write global
// 	if !separateTables && header {
// 		e.CSVwriterGlobal.WriteString(e.CSVheaderInternal)
// 		e.CSVwriterGlobal.WriteString(e.CSVheaderExternal)
// 	}

// 	for i, table := range e.TablesXML.Tables {
// 		if separateTables && header {
// 			table.CSVtableBuild(header, delim)
// 		}
// 		table.CSVtableBuild(header, delim)
// 		slog.Debug(
// 			"casting table to CSV", "current", i, "count", len(e.Tables))
// 	}
// }

// SaveTablesToFile
// func (e *Extractor) SaveTablesToFile(
// 	separateTables bool, dstFilePath string) error {
// 	if !separateTables {
// 		outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
// 		if err != nil {
// 			return err
// 		}
// 		defer outputFile.Close()
// 		n, err := outputFile.WriteString(e.CSVwriterGlobal.String())
// 		if err != nil {
// 			return err
// 		}
// 		slog.Debug("writen bytes to one file", "filename", dstFilePath, "bytesCount", n)
// 		return nil
// 	}
// 	current := 0
// 	bytesCountCumulative := 0
// 	for i, table := range e.Tables {
// 		current++
// 		dstFilePath := ConstructDstFilePath(table.SrcFilePath)
// 		outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
// 		if err != nil {
// 			return err
// 		}
// 		n, err := outputFile.WriteString(table.CSVwriterLocal.String())
// 		if err != nil {
// 			return err
// 		}
// 		sequnece := fmt.Sprintf("%d/%d", current, len(e.Tables))
// 		slog.Debug(
// 			"writen bytes to file in sequence", "sequence", sequnece,
// 			"filename", dstFilePath,
// 			"srcFile", i, "bytesCount", n,
// 		)

// 	}
// 	slog.Debug("finished writing files in sequence",
// 		"bytesCount", bytesCountCumulative,
// 		"filesCount", len(e.Tables))
// 	return nil
// }
