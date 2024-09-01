package mongo

import (
	"context"
	"reflect"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	"github.com/kelseyhightower/envconfig"
	"github.com/ongyoo/roomkub-api/pkg/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	ConnectionString string `envconfig:"MONGO_CONNECTION_STRING"`
	DatabaseName     string `envconfig:"MONGO_DATABASE_NAME"`
}

func NewDB() *mongo.Database {
	var config Config
	envconfig.MustProcess("", &config)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	c, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(config.ConnectionString),
		options.Client().SetMonitor(otelmongo.NewMonitor(otelmongo.WithCommandAttributeDisabled(true))),
	)
	if err != nil {
		panic(errors.Wrap(err, "Cannot connect to mongo server"))
	}

	if err := c.Ping(ctx, readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "Cannot ping mongo server"))
	}

	logrus.Info("successfully connected to mongodb")

	return c.Database(config.DatabaseName, options.Database().SetRegistry(DefaultRegistry))
}

type Encrypted string

func (e Encrypted) String() string {
	return string(e)
}

var DefaultRegistry = createRegistry()

func CreateRegistry() *bsoncodec.Registry {
	return DefaultRegistry
}

func createRegistry() *bsoncodec.Registry {
	rb := bson.NewRegistry()

	rb.RegisterTypeEncoder(
		reflect.TypeOf(Encrypted("")),
		bsoncodec.ValueEncoderFunc(func(ctx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
			encrypted, err := crypto.EncryptAes256(val.String())
			if err != nil {
				return err
			}
			return vw.WriteString(encrypted)
		}))

	rb.RegisterTypeDecoder(reflect.TypeOf(Encrypted("")), bsoncodec.ValueDecoderFunc(func(decodeContext bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value) error {
		val, err := reader.ReadString()
		if err != nil {
			return err
		}

		v, err := crypto.DecryptAes256(val)
		if err != nil {
			//// NOTE: support previous EncryptedSIV methods
			if err.Error() == crypto.ErrCipherMessageAuthenticationFailed.Error() {
				vSIV, err := crypto.DecryptAes256StaticIV(val)
				if err != nil {
					return err
				}
				value.SetString(vSIV)

				return nil
			}

			// NOTE: support non encrypted value
			if errors.Is(err, crypto.ErrInvalidEncryptedValueFormat) {
				value.SetString(val)
				return nil
			}
			return err
		}
		value.SetString(v)
		return nil
	}))
	return rb
}
