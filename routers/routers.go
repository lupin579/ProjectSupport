package routers

import (
	"eee/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(Cors)
	api := r.Group("/api")
	{
		api.POST("/login", controller.Register)
		//api.POST("/login", controller.Login)
	}
	withToken := r.Group("/withToken" /*, middlewares.ValidateToken*/)
	{
		withToken.GET("/changePassword/:uname", controller.ChangePassword) //yes
		withToken.POST("/validateEmailCode", controller.ValidateEmailCode)
		withToken.GET("/mission/:uname", controller.Mission)      //yes
		withToken.PUT("/updateMission", controller.UpdateMission) //yes
		withToken.POST("/addMission", controller.AddMission)      //yes
		withToken.PUT("/setF", controller.UpdateMFinish)          //yes
		withToken.POST("/postWorktime", controller.PostWorktime)  //yes
		withToken.GET("/memberPercent", controller.MemberPercent) //yes
		withToken.GET("/proList/:uname", controller.ProList)      //yes
		withToken.POST("/addPro/:uname", controller.AddProject)   //
		withToken.POST("/updatePro", controller.UpdatePro)        //
		withToken.PUT("/setPF", controller.UpdateOPro)
		withToken.GET("/getMP/:pro", controller.GetMP)              //yes
		withToken.GET("/peopleList", controller.PeopleList)         //项目支撑人员信息
		withToken.GET("/allPerson", controller.PersonList)          //yes
		withToken.GET("/freePersonList", controller.FreePersonList) //yes
		withToken.POST("/distribute", controller.Distribute)        //yes
		withToken.DELETE("/cancelUP", controller.Cancel)            //yes
		withToken.POST("/addUser", controller.AddUser)              //yes
		withToken.GET("/proDetails/:pro", controller.ProDetails)    //four percent----------
		withToken.GET("/recentOutput", controller.RecentOutput)     //yes
		withToken.GET("/recentPros", controller.RecentPros)         //yes
		withToken.GET("/proM/:pro", controller.ProMoney)            //yes
		withToken.GET("/timeChart", controller.TimeChart)
		withToken.GET("/mtPro/:uid", controller.MyPro)
		withToken.GET("/mtProOne/:leader", controller.MyProOne)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "40412323156",
		})
	})
	return r
}

func Cors(context *gin.Context) {
	method := context.Request.Method
	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
	//context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Origin", "*")
	//fmt.Println(context.GetHeader("Origin"))
	// 必须，设置服务器支持的所有跨域请求的方法
	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Authorization")
	// 可选，是否允许后续请求携带认证信息Cookie，该值只能是true，不需要则不设置
	context.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}
