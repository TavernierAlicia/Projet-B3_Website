package main

import (
	"fmt"
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
	data := Preview{}
	c.HTML(200, "index.html", data)
}

//handle contact page
func contactPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "contact.html", data)
}

//handle professional page
func proPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "professionnal.html", data)
}

//handle customer page
func custPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "customer.html", data)
}

//handle form pro
func formProPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "form-pro.html", data)
}

//handle form cust
func formCustPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "form-customer.html", data)
}

//handle about page
func aboutPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "about.html", data)
}

func receptForm(c *gin.Context) {

	data := Preview{}

	fmt.Println(data)
	c.Request.ParseForm()

	mail := strings.Join(c.Request.PostForm["from"], " ")
	name := strings.Join(c.Request.PostForm["name"], " ")
	surname := strings.Join(c.Request.PostForm["surname"], " ")
	subjectNum := strings.Join(c.Request.PostForm["subject"], " ")
	cmdNumber := strings.Join(c.Request.PostForm["cmdNumber"], " ")
	message := strings.Join(c.Request.PostForm["message"], " ")
	pro := false

	path := c.FullPath()

	//Define is the client is a professionnal or not
	if path == "/contact/form-pro" || path == "/professionnal/form-pro" || path == "/form-pro" {
		pro = true
	} else {
		pro = false
	}

	//choose subject and send mail
	subject := SelectSubj(pro, subjectNum)
	response := SendMail(mail, name, surname, subject, cmdNumber, message, pro)

	if response == true {
		c.Redirect(http.StatusMovedPermanently, "/success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error")
	}
}

//handle success page
func successForm(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "success.html", data)

	//TODO: find a solution for after sendmail
	//time.Sleep(2 * time.Second)
	//c.Redirect(http.StatusMovedPermanently, "/index")
}

//handle custom error page
func errorPage(c *gin.Context) {
	var cause string
	//get infos about the path
	host := "127.0.0.1"
	fullpath := host + c.FullPath()

	cause = "une erreur s'est produite"

	//display error page
	data := Preview{
		Text:       fullpath,
		ErrorCause: cause,
	}
	c.HTML(404, "404.html", data)
}

func main() {

	//zap stuff
	log, _ = zap.NewProduction()
	defer log.Sync()

	//Define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles("templates/index.html", "templates/404.html", "templates/about.html", "templates/customer.html", "templates/form-pro.html", "templates/form-customer.html", "templates/professionnal.html", "templates/success.html", "templates/contact.html")

	//to include js
	router.Static("/js", "./js")

	//to include css
	router.Static("/css", "./css")

	//to include images
	router.Static("/pictures", "./pictures")

	//GET requests
	//index routes
	router.NoRoute(errorPage)
	router.GET("/error", errorPage)
	router.GET("/", indexPage)
	router.GET("/index", indexPage)
	router.GET("/success", successForm)

	//navbar routes
	router.GET("/about", aboutPage)
	router.GET("/professionnal", proPage)
	router.GET("/customer", custPage)
	router.GET("/contact", contactPage)

	//mailing routes
	//from contact
	router.GET("/contact/form-pro", formProPage)
	router.GET("/contact/form-customer", formCustPage)

	//from infos pages
	router.GET("/professionnal/form-pro", formProPage)
	router.GET("/customer/form-customer", formCustPage)

	//from nothing
	router.GET("/form-pro", formProPage)
	router.GET("/form-customer", formCustPage)

	//POST requests
	router.POST("/form-pro", receptForm)
	router.POST("/form-customer", receptForm)
	router.POST("/professionnal/form-pro", receptForm)
	router.POST("/customer/form-customer", receptForm)
	router.POST("/contact/form-pro", receptForm)
	router.POST("/contact/form-customer", receptForm)

	//launch
	//router.Run(":3000")
	router.Run()

}
