package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ajtroup1/DocMate/internal/types"
)

// Path to the settings file
const settingsFilePath = "settings.json"

// GetOrRetrieveSettings reads the settings file or creates a default settings object
func GetOrRetrieveSettings() (*types.Settings, error) {
	var settings types.Settings

	// Check if the settings file exists
	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		// If the file does not exist, create default settings
		fmt.Println("\033[33mSettings file not found, creating default settings...\nDO NOT MOVE THIS SETTINGS FILE\033[0m")
		settings = types.Settings{
			ProjectName:  "Include project name here...",
			ProjectPath:  "./",
			ProjectDesc:  "Include project description here...",
			ImgLink:      "",
			OutputPath:   "./",
			IncludeTests: false,
		}

		// Save the default settings to file
		err = saveSettingsToFile(&settings)
		if err != nil {
			return nil, fmt.Errorf("failed to save default settings: %v", err)
		}

		return &settings, nil
	}

	// If the settings file exists, read and unmarshal the content
	fileContent, err := os.ReadFile(settingsFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read settings file: %v", err)
	}

	err = json.Unmarshal(fileContent, &settings)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal settings file: %v", err)
	}

	return &settings, nil
}

// Helper function to save the settings to a file
func saveSettingsToFile(settings *types.Settings) error {
	fileContent, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(settingsFilePath, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
