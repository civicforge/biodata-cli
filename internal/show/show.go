package show

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/civicforge/biodata-cli/internal/index"
	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/spf13/cobra"
)

func Show(cmd *cobra.Command, args []string) {
	makePretty, _ := cmd.Flags().GetBool("pretty")
	raw, _ := cmd.Flags().GetBool("raw")
	jsonOut, _ := cmd.Flags().GetBool("json")

	idxIdString := args[0]
	idxId, err := strconv.Atoi(idxIdString)
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}

	store, err := os.ReadFile(".temp_store.json")
	if err != nil {
		logging.Error(err.Error())
	}

	var indexes index.Index
	err = json.Unmarshal(store, &indexes)
	if err != nil {
		logging.Error(err.Error())
	}

	for _, idx := range indexes.IndexedFiles {
		if idx.ID == idxId {
			idxMI, _ := json.MarshalIndent(idx, "", "\t")

			if makePretty {
				showPretty(idx)
				os.Exit(1)
			}

			if raw {
				fmt.Println(idx)
				os.Exit(1)
			}

			if jsonOut {
				marshalledJson, _ := json.MarshalIndent(idx, "", "\t")
				fileName := fmt.Sprintf("%d-%s-INDEX.json", idx.ID, strings.TrimSuffix(idx.Filename, "."+idx.Extension))
				err := os.WriteFile(fileName, marshalledJson, 0666)
				if err != nil {
					logging.Error(err.Error())
					return
				}
				logging.Info("Writing to file: " + fileName)

				os.Exit(1)
			}

			fmt.Println(string(idxMI))
			os.Exit(1)
		}
	}

	logmsg := fmt.Sprintf("No index ID matching %s", idxIdString)
	logging.Info(logmsg)
}
