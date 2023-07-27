package main

import(
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostname string = "localhost:27017"
	dbName string = "demo_todo"
	collectionName string = "todo"
	port string = ":9000"
)

type(
	todoModel struct {
		ID		bson.ObjectId `bson:"_id,omitempty"`
		Title	string `bson:"title"`
		Completed bool `b`
	}
)

func main() {

}
