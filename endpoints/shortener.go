package endpoints

import (
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/chrisenytc/gouid"
	"gopkg.in/kataras/iris.v6"

	"github.com/chrisenytc/ullli/adapters"
	"github.com/chrisenytc/ullli/config"
)

type Url struct {
	Url string `form:"url"`
}

type UrlItem struct {
	Url       string
	Clicks    string
	CreatedAt string
}

func PostShorten(ctx *iris.Context) {
	urlData := Url{}
	err := ctx.ReadForm(&urlData)

	if err != nil {
		log.Errorf("An error has occurred on ReadForm: %s", err)
		ctx.EmitError(iris.StatusBadRequest)
		return
	}

	if urlData.Url == "" {
		ctx.SetStatusCode(iris.StatusBadRequest)
		ctx.MustRender("error.html", struct{ Message string }{Message: "You need to a enter a URL."})
		return
	}

	_, err = url.ParseRequestURI(urlData.Url)

	if err != nil {
		ctx.SetStatusCode(iris.StatusBadRequest)
		ctx.MustRender("error.html", struct{ Message string }{Message: "Invalid URL."})
		return
	}

	uid := gouid.UId{6}

	uid.SetSeed()

	shortCode := uid.NewUId()

	exists, exists_err := adapters.CheckShortCode(shortCode)

	if exists_err != nil {
		log.Errorf("An error has occurred on CheckShortCode: %s", exists_err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	if exists {
		shortCode = uid.NewUId()
	}

	_, save_err := adapters.SaveUrl(shortCode, urlData.Url)

	if save_err != nil {
		log.Errorf("An error has occurred on SaveUrl: %s", save_err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	ctx.MustRender("success.html", Url{Url: config.Get().HostUrl + "/" + shortCode})
}

func GetShortCode(ctx *iris.Context) {
	shortCode := ctx.Param("shortCode")
	redirect := ctx.URLParam("redirect")

	exists, exists_err := adapters.CheckShortCode(shortCode)

	if exists_err != nil {
		log.Errorf("An error has occurred on CheckShortCode: %s", exists_err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	if !exists {
		log.Warnf("ShortCode %s not found.", shortCode)
		ctx.EmitError(iris.StatusNotFound)
		return
	}

	url, err := adapters.GetUrl(shortCode)

	if err != nil {
		log.Errorf("An error has occurred on GetUrl: %s", err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	_, count_err := adapters.CountClick(shortCode)

	if count_err != nil {
		log.Errorf("An error has occurred on CountClick: %s", count_err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	if redirect == "false" {
		ctx.MustRender("redirect.html", Url{Url: url})
		return
	}

	ctx.Redirect(url, iris.StatusTemporaryRedirect)
}

func GetShortCodeStats(ctx *iris.Context) {
	shortCode := ctx.Param("shortCode")

	exists, exists_err := adapters.CheckShortCode(shortCode)

	if exists_err != nil {
		log.Errorf("An error has occurred on CheckShortCode: %s", exists_err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	if !exists {
		log.Warnf("ShortCode %s not found.", shortCode)
		ctx.EmitError(iris.StatusNotFound)
		return
	}

	url, err := adapters.GetUrlData(shortCode)

	if err != nil {
		log.Errorf("An error has occurred on GetUrlData: %s", err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	createdAt, err := time.Parse(time.RFC3339, url[5])

	if err != nil {
		log.Errorf("An error has occurred on Time.Parse: %s", err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	ctx.MustRender("stats.html", UrlItem{Url: url[1], Clicks: url[3], CreatedAt: createdAt.UTC().Format("Jan 2, 2006 at 3:04 pm")})
}
