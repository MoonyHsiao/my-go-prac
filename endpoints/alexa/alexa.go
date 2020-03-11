package alexa

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MoonyHsiao/my-go-prac/errno"
	"github.com/MoonyHsiao/my-go-prac/models"
	"github.com/MoonyHsiao/my-go-prac/viewmodels"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const defaultQueryLen = 50
const domainUrl = "https://www.alexa.com"

func GetTop(ctx *gin.Context) {

	var url = domainUrl + "/topsites"

	var viewmodel viewmodels.RankQueryParam
	if err := ctx.Bind(&viewmodel); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	emptyQueryParam := viewmodels.RankQueryParam{}
	var querylen = defaultQueryLen
	if emptyQueryParam != viewmodel {
		querylen = viewmodel.Top
	}

	res, err := GetCrawlRes(url, querylen)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	errno.Success(res, ctx)
}

func GetCountry(ctx *gin.Context) {

	var url = domainUrl + "/topsites/countries/"
	var viewmodel viewmodels.CountryQueryParam
	if err := ctx.Bind(&viewmodel); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	emptyQueryParam := viewmodels.CountryQueryParam{}
	if emptyQueryParam == viewmodel {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url = url + viewmodel.Region
	res, err := GetCrawlRes(url, 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	errno.Success(res, ctx)
}

func GetCrawlRes(url string, datalen int) ([]models.RankData, error) {

	var res []models.RankData
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("http.Get error")
		return res, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println("read error is:", err)
		err = errors.New("ReadAll error")
		return res, err
	}

	bodyString := string(body)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(bodyString))
	if err != nil {
		return res, err
	}

	var queryData []models.RankData
	dom.Find(".tr.site-listing").Each(func(i int, selection *goquery.Selection) {
		var temp models.RankData
		rank := selection.Find(".td").First().Text()
		rank = strings.Replace(rank, " ", "", -1)
		temp.Rank = rank
		url := selection.Find(".td.DescriptionCell").Text()
		url = strings.Replace(url, " ", "", -1)
		url = strings.Replace(url, "\n", "", -1)
		temp.Url = url
		queryData = append(queryData, temp)
	})

	if datalen > len(queryData) {
		datalen = len(queryData)
	}

	for i := 0; i < datalen; i++ {
		res = append(res, queryData[i])
	}
	return res, err
}
