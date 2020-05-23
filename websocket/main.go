package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//define variables
var (
	log *zap.Logger
)

//Preview error struct
type Preview struct {
	Text       string
	ErrorCause string
}

//handle index page
func indexPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//handle professional page
func proPage(c *gin.Context) {
	c.HTML(200, "professionnal.html", nil)
}

//handle form pro
func form(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

//handle about page
func aboutPage(c *gin.Context) {
	c.HTML(200, "about.html", nil)
}

//handle legal
func legal(c *gin.Context) {
	c.HTML(200, "legal.html", nil)
}

//handle FAQ
func faq(c *gin.Context) {
	c.HTML(200, "faq.html", nil)
}

func receptForm(c *gin.Context) {
	c.Request.ParseForm()

	ispro := strings.Join(c.Request.PostForm["ispro"], " ")
	mail := strings.Join(c.Request.PostForm["from"], " ")
	name := strings.Join(c.Request.PostForm["name"], " ")
	surname := strings.Join(c.Request.PostForm["surname"], " ")
	subjectNum := strings.Join(c.Request.PostForm["subject"], " ")
	message := strings.Join(c.Request.PostForm["message"], " ")
	tel := strings.Join(c.Request.PostForm["tel"], " ")
	entname := strings.Join(c.Request.PostForm["entname"], " ")

	var pro bool

	//Define is the client is a professionnal or not
	if ispro == "0" {
		pro = true
	} else {
		pro = false
	}

	//choose subject and send mail
	subject := SelectSubj(subjectNum)
	response := SendMail(mail, name, surname, subject, message, pro, tel, entname)

	if response == true {
		c.Redirect(http.StatusMovedPermanently, "/success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error")
	}
}

//handle success page
func successForm(c *gin.Context) {
	c.HTML(200, "success.html", nil)
}

//handle custom error 404
func errorPage(c *gin.Context) {
	var cause string
	//get infos about the path
	host := "127.0.0.1"
	fullpath := host + c.Request.URL.Path

	cause = "cette page n'est pas accessible pour le moment ou n'existe pas"

	//display error page
	data := Preview{
		Text:       fullpath,
		ErrorCause: cause,
	}
	c.HTML(404, "404.html", data)
}

//handle send form error
func errorForm(c *gin.Context) {
	c.HTML(404, "error.html", nil)
}

func main() {

	//zap stuff
	log, _ = zap.NewProduction()
	defer log.Sync()

	//Define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles("templates/index.html", "templates/404.html", "templates/error.html", "templates/about.html", "templates/legal.html", "templates/faq.html", "templates/form.html", "templates/professionnal.html", "templates/success.html")

	//to include js
	router.Static("/js", "./js")

	//to include css
	router.Static("/css", "./css")

	//to include images
	router.Static("/pictures", "./pictures")

	//GET requests
	//index routes
	router.NoRoute(errorPage)

	router.GET("/error", errorForm)
	router.GET("/", indexPage)
	router.GET("/index", indexPage)
	router.GET("/success", successForm)

	//navbar routes
	router.GET("/about", aboutPage)
	router.GET("/professionnal", proPage)

	//from infos pages
	router.GET("/professionnal/form", form)
	router.GET("/about/faq", faq)
	router.GET("/about/legal", legal)

	//from nothing
	router.GET("/form", form)
	router.GET("/faq", faq)
	router.GET("/FAQ", faq)
	router.GET("/legal", legal)

	//POST requests
	router.POST("/form", receptForm)
	router.POST("/professionnal/form", receptForm)

	//launch
	//router.Run(":3000")
	//router.Run()
	//router.Run("www.orderndrink.com:3000")
	router.Run("127.0.0.1:3000")
}
