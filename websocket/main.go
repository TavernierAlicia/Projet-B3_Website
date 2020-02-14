package main

import (
	"fmt"
	"net/http"

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

/*
type Mail struct {
	From      string
	Name      string
	Surname   string
	Subject   string
	CmdNumber string
	Message   string
}
*/

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

	//Call to ParseForm makes form fields available.
	//err := c.Request.ParseForm()
	data := Preview{}
	//if err != nil {
	//	fmt.Println("there is an error")
	//	c.HTML(404, "404.html", data)
	//}

	fmt.Println("parsing form")
	c.Request.ParseForm()
	from := c.Request.PostForm["from"]
	name := c.Request.PostForm["name"]
	surname := c.Request.PostForm["surname"]
	subject := c.Request.PostForm["subject"]
	cmdNumber := c.Request.PostForm["cmdNumber"]
	message := c.Request.PostForm["message"]

	//mail := Mail{
	//	From:      from,
	//	Name:      name,
	//	Surname:   surname,
	//	Subject:   subject,
	//	CmdNumber: cmdNumber,
	//	Message:   message,
	//}

	fmt.Println("Hello", name, surname, from, subject, cmdNumber, message)
	fmt.Println("redirect")
	fmt.Println(data)
	c.Redirect(http.StatusMovedPermanently, "/success")
	fmt.Println("redirected")
}

//handle success page
func successForm(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "success.html", data)
}

//handle custom error page
func errorPage(c *gin.Context) {
	var cause string
	//get infos about the path
	path1 := c.Param("path1")
	path2 := c.Param("path2")
	host := "127.0.0.1"
	fullpath := host + c.FullPath() + path1 + "/" + path2

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
