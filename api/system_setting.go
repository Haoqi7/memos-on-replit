package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

type SystemSettingName string

const (
	// SystemSettingServerID is the key type of server id.
	SystemSettingServerID SystemSettingName = "serverId"
	// SystemSettingSecretSessionName is the key type of secret session name.
	SystemSettingSecretSessionName SystemSettingName = "secretSessionName"
	// SystemSettingAllowSignUpName is the key type of allow signup setting.
	SystemSettingAllowSignUpName SystemSettingName = "allowSignUp"
	// SystemSettingAdditionalStyleName is the key type of additional style.
	SystemSettingAdditionalStyleName SystemSettingName = "additionalStyle"
	// SystemSettingAdditionalScriptName is the key type of additional script.
	SystemSettingAdditionalScriptName SystemSettingName = "additionalScript"
	// SystemSettingCustomizedProfileName is the key type of customized server profile.
	SystemSettingCustomizedProfileName SystemSettingName = "customizedProfile"
)

// CustomizedProfile is the struct definition for SystemSettingCustomizedProfileName system setting item.
type CustomizedProfile struct {
	// Name is the server name, default is `memos`
	Name string `json:"name"`
	// LogoURL is the url of logo image.
	LogoURL string `json:"logoUrl"`
	// Description is the server description.
	Description string `json:"description"`
	// Locale is the server default locale.
	Locale string `json:"locale"`
	// Appearance is the server default appearance.
	Appearance string `json:"appearance"`
	// ExternalURL is the external url of server. e.g. https://usermemos.com
	ExternalURL string `json:"externalUrl"`
}

func (key SystemSettingName) String() string {
	switch key {
	case SystemSettingServerID:
		return "serverId"
	case SystemSettingSecretSessionName:
		return "secretSessionName"
	case SystemSettingAllowSignUpName:
		return "allowSignUp"
	case SystemSettingAdditionalStyleName:
		return "additionalStyle"
	case SystemSettingAdditionalScriptName:
		return "additionalScript"
	case SystemSettingCustomizedProfileName:
		return "customizedProfile"
	}
	return ""
}

var (
	SystemSettingAllowSignUpValue = []bool{true, false}
)

type SystemSetting struct {
	Name SystemSettingName
	// Value is a JSON string with basic value.
	Value       string
	Description string
}

type SystemSettingUpsert struct {
	Name        SystemSettingName `json:"name"`
	Value       string            `json:"value"`
	Description string            `json:"description"`
}

func (upsert SystemSettingUpsert) Validate() error {
	if upsert.Name == SystemSettingServerID {
		return errors.New("update server id is not allowed")
	} else if upsert.Name == SystemSettingAllowSignUpName {
		value := false
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting allow signup value")
		}

		invalid := true
		for _, v := range SystemSettingAllowSignUpValue {
			if value == v {
				invalid = false
				break
			}
		}
		if invalid {
			return fmt.Errorf("invalid system setting allow signup value")
		}
	} else if upsert.Name == SystemSettingAdditionalStyleName {
		value := ""
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting additional style value")
		}
	} else if upsert.Name == SystemSettingAdditionalScriptName {
		value := ""
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting additional script value")
		}
	} else if upsert.Name == SystemSettingCustomizedProfileName {
		customizedProfile := CustomizedProfile{
			Name:        "memos",
			LogoURL:     "",
			Description: "",
			Locale:      "en",
			Appearance:  "system",
			ExternalURL: "",
		}
		err := json.Unmarshal([]byte(upsert.Value), &customizedProfile)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting customized profile value")
		}
		if !slices.Contains(UserSettingLocaleValue, customizedProfile.Locale) {
			return fmt.Errorf("invalid locale value")
		}
		if !slices.Contains(UserSettingAppearanceValue, customizedProfile.Appearance) {
			return fmt.Errorf("invalid appearance value")
		}
	} else {
		return fmt.Errorf("invalid system setting name")
	}

	return nil
}

type SystemSettingFind struct {
	Name *SystemSettingName `json:"name"`
}
