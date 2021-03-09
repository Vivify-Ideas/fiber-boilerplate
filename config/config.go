package config

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/ace"
	"github.com/gofiber/template/amber"
	"github.com/gofiber/template/django"
	"github.com/gofiber/template/handlebars"
	"github.com/gofiber/template/html"
	"github.com/gofiber/template/jet"
	"github.com/gofiber/template/mustache"
	"github.com/gofiber/template/pug"
	"github.com/joho/godotenv"
)

var config = map[string]string{
	"APP_URL":         "http://localhost:3000",
	"APP_LISTEN_PORT": ":3000",
	"APP_ENV":         "local",

	"DB_DRIVER":   "sqlite",
	"DB_HOST":     "localhost",
	"DB_USERNAME": "root",
	"DB_PASSWORD": "root",
	"DB_PORT":     "3306",
	"DB_DATABASE": "db",

	"EMAIL_HOST":     "mailhog",
	"EMAIL_PORT":     "1025",
	"EMAIL_USERNAME": "",
	"EMAIL_PASSWORD": "",
	"EMAIL_FROM":     "Vivify <contact@vivifyideas.com>",

	"CORS_ALLOW_ORIGINS":     "*",
	"CORS_ALLOW_METHODS":     "GET,POST,HEAD,PUT,DELETE,PATCH",
	"CORS_ALLOW_HEADERS":     "",
	"CORS_ALLOW_CREDENTIALS": "false",

	"SENTRY_DSN": "",

	"FIBER_PREFORK":                   "false",
	"FIBER_SERVERHEADER":              "",
	"FIBER_STRICTROUTING":             "false",
	"FIBER_CASESENSITIVE":             "false",
	"FIBER_IMMUTABLE":                 "false",
	"FIBER_UNESCAPEPATH":              "false",
	"FIBER_ETAG":                      "false",
	"FIBER_BODYLIMIT":                 "4194304",
	"FIBER_CONCURRENCY":               "262144",
	"FIBER_VIEWS":                     "html",
	"FIBER_VIEWS_DIRECTORY":           "resources/views",
	"FIBER_VIEWS_RELOAD":              "false",
	"FIBER_VIEWS_DEBUG":               "false",
	"FIBER_VIEWS_LAYOUT":              "embed",
	"FIBER_VIEWS_DELIMS_L":            "{{",
	"FIBER_VIEWS_DELIMS_R":            "}}",
	"FIBER_READTIMEOUT":               "0",
	"FIBER_WRITETIMEOUT":              "0",
	"FIBER_IDLETIMEOUT":               "0",
	"FIBER_READBUFFERSIZE":            "4096",
	"FIBER_WRITEBUFFERSIZE":           "4096",
	"FIBER_COMPRESSEDFILESUFFIX":      ".fiber.gz",
	"FIBER_PROXYHEADER":               "",
	"FIBER_GETONLY":                   "false",
	"FIBER_DISABLEKEEPALIVE":          "false",
	"FIBER_DISABLEDEFAULTDATE":        "false",
	"FIBER_DISABLEDEFAULTCONTENTTYPE": "false",
	"FIBER_DISABLEHEADERNORMALIZING":  "false",
	"FIBER_DISABLESTARTUPMESSAGE":     "false",
	"FIBER_REDUCEMEMORYUSAGE":         "false",
}

// AppConfig - application configuration
type AppConfig struct {
	Env          map[string]string
	Views        fiber.Views
	ErrorHandler func(ctx *fiber.Ctx, err error) error
	LoadStatic   func(app *fiber.App)
}

// App - application configuration
var App = AppConfig{
	Env:          readConfig(),
	Views:        getFiberViewsEngineConfig(),
	ErrorHandler: getSentryHandler(),
	LoadStatic:   getLoadStatic(),
}

// Read configuration
func readConfig() (myConfig map[string]string) {
	var env, err = godotenv.Read()

	if err != nil {
		log.Println(err)
	}

	myConfig = make(map[string]string)
	for key, value := range config {
		if envValue, ok := env[key]; ok {
			myConfig[key] = envValue
		} else {
			myConfig[key] = value
		}
	}
	return
}

func getFiberViewsEngineConfig() fiber.Views {
	var viewsEngine fiber.Views
	config = readConfig()
	fiberViewsReload, _ := strconv.ParseBool(config["FIBER_VIEWS_RELOAD"])
	fiberViewsDebug, _ := strconv.ParseBool(config["FIBER_VIEWS_DEBUG"])

	log.Println("View Engine:", strings.ToLower(config["FIBER_VIEWS"]))

	switch strings.ToLower(config["FIBER_VIEWS"]) {
	case "ace":
		engine := ace.New(config["FIBER_VIEWS_DIRECTORY"], ".ace")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	case "amber":
		engine := amber.New(config["FIBER_VIEWS_DIRECTORY"], ".amber")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	case "django":
		engine := django.New(config["FIBER_VIEWS_DIRECTORY"], ".django")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"])
		viewsEngine = engine
		break
	case "handlebars":
		engine := handlebars.New(config["FIBER_VIEWS_DIRECTORY"], ".hbs")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	case "jet":
		engine := jet.New(config["FIBER_VIEWS_DIRECTORY"], ".jet")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	case "mustache":
		engine := mustache.New(config["FIBER_VIEWS_DIRECTORY"], ".mustache")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	case "pug":
		engine := pug.New(config["FIBER_VIEWS_DIRECTORY"], ".pug")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	default:
		engine := html.New(config["FIBER_VIEWS_DIRECTORY"], ".html")
		engine.Reload(fiberViewsReload).
			Debug(fiberViewsDebug).
			Layout(config["FIBER_VIEWS_LAYOUT"]).
			Delims(config["FIBER_VIEWS_DELIMS_L"], config["FIBER_VIEWS_DELIMS_R"])
		viewsEngine = engine
		break
	}
	return viewsEngine
}

func getSentryHandler() func(ctx *fiber.Ctx, err error) error {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	env := readConfig()
	dsn := env["SENTRY_DSN"]
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}

	return func(ctx *fiber.Ctx, err error) error {
		// Statuscode defaults to 500
		code := fiber.StatusInternalServerError

		// Retreive the custom statuscode if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		sentry.CaptureException(err)

		// Return from handler
		return ctx.Status(code).SendString(err.Error())
	}
}

func getLoadStatic() func(app *fiber.App) {
	return func(app *fiber.App) {
		app.Static("/", "./public", fiber.Static{
			Compress:      true,
			ByteRange:     true,
			CacheDuration: 24 * time.Hour,
		})
	}

}
