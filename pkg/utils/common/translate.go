package util_common

import (
	"log"
	"path/filepath"

	util_error "github.com/PICH-IO/admin-api/pkg/utils/errors"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var bundle *i18n.Bundle

// Initialize the i18n bundle and load locale files
func Init() *util_error.ErrorResponse {
	bundle = i18n.NewBundle(language.English)
	// fmt.Println("budle localtion:", bundle)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	localeFiles := []string{
		"pkg/translate/en.yaml",
		"pkg/translate/km.yaml",
	}

	for _, file := range localeFiles {
		_, err := bundle.LoadMessageFile(filepath.Join(file))
		if err != nil {
			log.Printf("Error loading locale file %s: %v", file, err)
			return util_error.NewError("ErrorLoadMessage", "Failed to load locale files")
		}
	}
	return nil
}

// TranslateWithError translates the given key using the i18n bundle
func TranslateWithError(c *fiber.Ctx, key string, templateData ...map[string]interface{}) (string, *util_error.ErrorResponse) {
	// fmt.Println("Hellow Error:")
	if bundle == nil {
		log.Println("Error: i18n bundle is not initialized")
		return "", util_error.NewError(key, "Translation service is unavailable")
	}
	// Set default language or fallback to English
	lang := c.Get("Accept-Language", "en")
	localizer := i18n.NewLocalizer(bundle, lang)

	// Handle optional templateData
	data := map[string]interface{}{}
	if len(templateData) > 0 && templateData[0] != nil {
		data = templateData[0]
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	})
	if err != nil {
		log.Printf("Error localizing message ID %s: %v", key, err)
		return "", util_error.NewError(key, "Translation not found")
	}
	return msg, nil
}

// translate
func Translate(c *fiber.Ctx, key string) string {
	return fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
		MessageID: key,
	})
}
