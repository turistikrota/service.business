package adapters

import (
	"github.com/turistikrota/service.business/src/adapters/memory"
	"github.com/turistikrota/service.business/src/adapters/mongo"
	"github.com/turistikrota/service.business/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
