package adapters

import (
	"github.com/turistikrota/service.owner/src/adapters/memory"
	"github.com/turistikrota/service.owner/src/adapters/mongo"
	"github.com/turistikrota/service.owner/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
