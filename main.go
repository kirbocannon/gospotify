package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()  // parse arguments, you have to call this by yourself
//	fmt.Println(r.Form)  // print form information in server side
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
//}

//var jokes = []Joke{
//	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
//	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
//	Joke{3, 0, "How many apples grow on a tree? All of them."},
//	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
//	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
//	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
//	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
//}


//type Joke struct {
//	ID int `json:"id" binding:"required"`
//	Likes int `json:"likes"`
//	Joke string `json:"joke" binding:"required"`
//}


//// JokeHandler retrieves a list of available jokes
//func JokeHandler(c *gin.Context) {
//	c.Header("Content-Type", "application/json")
//	c.JSON(http.StatusOK, jokes)
//}

var genreCounts map[string]int

func GetGenreCounts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, genreCounts)
}


//func LikeJoke(c *gin.Context) {
//	// confirm Joke ID sent is valid
//	// remember to import the `strconv` package
//	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
//		// find joke, and increment likes
//		for i := 0; i < len(jokes); i++ {
//			if jokes[i].ID == jokeid {
//				jokes[i].Likes += 1
//			}
//		}
//
//		// return a pointer to the updated jokes list
//		c.JSON(http.StatusOK, &jokes)
//	} else {
//		// Joke ID is invalid
//		c.AbortWithStatus(http.StatusNotFound)
//	}
//}

func main() {

	//c := gospotify.SpotifyClient{}
	//c.Init()
	//if c.ValidToken {
	//	genreCounts = c.GetGenreCounts()
 	//}

	genreCounts = map[string]int{"afro house":1,"alabama indie":1,"alabama rap":1,"album rock":1,"alternative dance":24,"alternative emo":1,"alternative hip hop":10,"alternative metal":8,"alternative r\u0026b":8,"alternative rock":8,"ambeat":3,"ambient":1,"anthem emo":1,"argentine hip hop":1,"art pop":11,"art rock":1,"atl hip hop":18,"atl trap":5,"aussietronica":3,"australian dance":9,"australian electropop":4,"australian indie":2,"australian pop":6,"australian psych":1,"baile pop":1,"baroque pop":2,"bass house":10,"bass music":1,"bass trap":19,"bassline":1,"belgian dance":1,"belgian edm":2,"big beat":4,"big room":90,"bmore":2,"boogaloo":1,"boston hip hop":1,"boy band":1,"brazilian edm":3,"brazilian house":1,"breakbeat":2,"brega funk":1,"brisbane indie":1,"british soul":2,"britpop":1,"brooklyn indie":1,"brostep":33,"cali rap":2,"canadian contemporary r\u0026b":2,"canadian electronic":10,"canadian electropop":2,"canadian folk":1,"canadian hip hop":6,"canadian indie":1,"canadian pop":9,"candy pop":2,"catstep":13,"chamber pop":5,"chicago house":1,"chicago indie":1,"chicago rap":4,"chicano rap":1,"chillhop":1,"chillstep":7,"chillwave":12,"christian alternative rock":1,"christian hardcore":1,"christian metal":1,"christian rock":2,"circuit":2,"classic rock":1,"colombian indie":1,"colombian pop":1,"complextro":16,"conscious hip hop":11,"contemporary country":1,"country":1,"country dawn":1,"country pop":1,"country road":1,"country rock":1,"cuban rumba":1,"cyberpunk":3,"dance pop":107,"dance-punk":12,"dancehall":3,"danish electronic":1,"danish pop":2,"dark trap":2,"deathcore":2,"deep acoustic pop":1,"deep big room":72,"deep chiptune":2,"deep disco house":1,"deep groove house":17,"deep house":10,"deep melodic euro house":1,"deep pop r\u0026b":9,"deep talent show":1,"deep tropical house":11,"deep underground hip hop":4,"deep uplifting trance":3,"destroy techno":2,"detroit hip hop":3,"detroit house":1,"devon indie":1,"dfw rap":1,"dirty south rap":6,"disco":3,"disco house":7,"diva house":1,"djent":1,"dmv rap":1,"downtempo":4,"dream pop":2,"drill":2,"dubstep":1,"dutch house":12,"dutch pop":1,"east coast hip hop":13,"edm":148,"el paso indie":1,"electra":2,"electro":1,"electro house":135,"electro latino":1,"electro trash":1,"electrofox":1,"electronic":13,"electronic rock":1,"electronic trap":32,"electropop":45,"electropowerpop":2,"emo":5,"emo rap":4,"epicore":2,"escape room":5,"estonian pop":1,"etherpop":3,"eurodance":1,"europop":4,"experimental hip hop":1,"fidget house":4,"filter house":8,"filthstep":1,"float house":1,"florida rap":1,"folk-pop":8,"folktronica":1,"fourth world":1,"freak folk":1,"french indie pop":1,"french indietronica":1,"french rock":1,"funk":4,"funk carioca":3,"funk das antigas":2,"funk metal":1,"funk ostentacao":2,"funk rock":1,"funky tech house":4,"future garage":1,"future house":11,"g funk":4,"gangster rap":20,"garage rock":5,"gauze pop":5,"german dance":3,"german house":1,"german metal":1,"german techno":4,"girl group":1,"glitch":1,"glitch hop":1,"greek indie":1,"groove metal":1,"groove room":1,"grunge":1,"hard rock":5,"hardcore hip hop":17,"hardcore techno":1,"hardstyle":1,"hip hop":75,"hip house":1,"hip pop":15,"hollywood":1,"hopebeat":1,"house":52,"hyperpop":1,"indian edm":1,"indie anthem-folk":1,"indie electro-pop":10,"indie folk":3,"indie pop":18,"indie pop rap":3,"indie poptimism":22,"indie punk":1,"indie r\u0026b":6,"indie rock":14,"indietronica":34,"indonesian hip hop":1,"industrial":1,"industrial metal":1,"industrial rock":1,"intelligent dance music":2,"j-idol":1,"j-pop girl group":2,"j-rap":2,"jazz funk":1,"jazz rap":3,"jazztronica":1,"k-hop":2,"k-pop":1,"k-pop girl group":1,"korean pop":1,"korean trap":1,"kwaito house":1,"la indie":2,"la pop":5,"latin":14,"latin arena pop":1,"latin hip hop":5,"latin jazz":1,"latin pop":3,"lgbtq+ hip hop":2,"lithuanian pop":1,"lo-fi beats":1,"lo-fi house":1,"malaysian pop":1,"mambo":1,"mashup":2,"mathcore":1,"melancholia":1,"melbourne bounce":4,"melbourne bounce international":4,"melodic metalcore":3,"memphis hip hop":1,"merengue":1,"metal":3,"metalcore":5,"metropopolis":10,"mexican pop":2,"miami hip hop":7,"microhouse":1,"minimal tech house":1,"minimal techno":1,"minneapolis sound":2,"modern alternative rock":1,"modern reggae":1,"modern rock":28,"moombahton":5,"motown":2,"nc hip hop":1,"neo classical metal":1,"neo mellow":6,"neo soul":8,"neo-psychedelic":3,"neo-synthpop":4,"neon pop punk":2,"new french touch":1,"new jack swing":7,"new jersey rap":2,"new orleans rap":1,"new rave":16,"nightrun":1,"ninja":3,"norwegian pop":4,"nu disco":7,"nu gaze":2,"nu jazz":2,"nu metal":8,"nyc pop":1,"nyc rap":1,"old school hip hop":1,"old school thrash":1,"organic electronic":1,"outlaw country":1,"oxford indie":1,"permanent wave":3,"perth indie":1,"philly indie":1,"philly rap":2,"piano rock":2,"pittsburgh rap":2,"pixie":1,"pop":169,"pop edm":90,"pop house":2,"pop nacional":1,"pop punk":7,"pop rap":91,"pop reggaeton":1,"pop rock":7,"popping":1,"post-disco":2,"post-grunge":10,"post-hardcore":1,"post-screamo":1,"post-teen pop":42,"power metal":1,"preverb":3,"progressive breaks":1,"progressive electro house":68,"progressive house":65,"progressive metal":1,"progressive metalcore":1,"progressive post-hardcore":1,"progressive trance":15,"progressive trance house":3,"progressive uplifting trance":3,"psychedelic rock":1,"quebec indie":1,"quiet storm":5,"r\u0026b":21,"r\u0026b en espanol":1,"rap":90,"rap kreyol":1,"rap latina":2,"rap metal":4,"rap rock":4,"red dirt":1,"reggaeton":7,"reggaeton flow":7,"retro electro":2,"rochester mn indie":1,"rock":17,"romanian hip hop":1,"romanian pop":5,"romanian rock":1,"romanian trap":1,"roots americana":1,"roots rock":1,"russian electronic":1,"russian hip hop":1,"russian pop":1,"russiavision":1,"sacramento indie":2,"salsa":1,"salsa international":1,"scandipop":1,"scorecore":2,"screamo":5,"serbian electronic":1,"shimmer pop":9,"shiver pop":4,"shoegaze":1,"singer-songwriter":1,"sky room":26,"slow core":1,"soul":2,"soundtrack":2,"southern hip hop":48,"spanish pop":1,"speed metal":3,"stomp and holler":2,"stoner rock":1,"swedish electropop":4,"swedish indie pop":3,"swedish pop":6,"swedish synthpop":2,"talent show":1,"tech house":5,"teen pop":1,"texas country":1,"thrash metal":2,"toronto indie":2,"toronto rap":5,"tracestep":16,"traditional country":1,"trance":14,"trancecore":2,"transpop":2,"trap argentino":1,"trap latino":6,"trap music":51,"trap queen":1,"trap soul":7,"traprun":10,"trip hop":4,"tropical":9,"tropical house":94,"turntablism":2,"uk alternative pop":3,"uk contemporary r\u0026b":3,"uk dance":15,"uk dancehall":1,"uk experimental electronic":1,"uk funky":1,"uk hip hop":1,"uk house":1,"uk metalcore":1,"uk pop":41,"underground hip hop":17,"uplifting trance":13,"urban contemporary":20,"utah indie":1,"vapor pop":1,"vapor soul":13,"vapor trap":4,"vapor twitch":13,"video game music":2,"viking metal":1,"viral pop":3,"viral trap":1,"vocal house":2,"vocal trance":6,"wave":1,"welsh rock":1,"west coast rap":4,"west coast trap":1,"western mass indie":1,"wonky":3,"zapstep":6}
	genreCounts, err := json.Marshal(genreCounts)
	if err != nil {
		// log error

	} else {
		fmt.Println("-", genreCounts)
	}

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend views files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	//router.Use(static.Serve("/scripts", static.LocalFile("./node_modules", true)))

	// Setup route group for the API
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})

	api.GET("/genre/all/counts", GetGenreCounts)
	//api.POST("/jokes/like/:jokeID", LikeJoke)

	// Start and run the server
	router.Run(":3000")



}