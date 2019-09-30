package documentation_examples_test

import (
	"QATrader_Go/mongo/testutil"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/examples/documentation_examples"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func TestDocumentationExamples(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cs := testutil.ConnString(t)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cs.String()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	db := client.Database("documentation_examples")

	documentation_examples.InsertExamples(t, db)
	documentation_examples.QueryToplevelFieldsExamples(t, db)
	documentation_examples.QueryEmbeddedDocumentsExamples(t, db)
	documentation_examples.QueryArraysExamples(t, db)
	documentation_examples.QueryArrayEmbeddedDocumentsExamples(t, db)
	documentation_examples.QueryNullMissingFieldsExamples(t, db)
	documentation_examples.ProjectionExamples(t, db)
	documentation_examples.UpdateExamples(t, db)
	documentation_examples.DeleteExamples(t, db)
}

func TestTransactionExamples(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cs := testutil.ConnString(t)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cs.String()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	ver, err := getServerVersion(ctx, client)
	if err != nil || testutil.CompareVersions(t, ver, "4.0") < 0 || os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip("server does not support transactions")
	}
	err = documentation_examples.TransactionsExamples(ctx, client)
	require.NoError(t, err)
}

func TestChangeStreamExamples(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cs := testutil.ConnString(t)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cs.String()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	db := client.Database("changestream_examples")
	ver, err := getServerVersion(ctx, client)
	if err != nil || testutil.CompareVersions(t, ver, "3.6") < 0 || os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip("server does not support changestreams")
	}
	documentation_examples.ChangeStreamExamples(t, db)
}

func getServerVersion(ctx context.Context, client *mongo.Client) (string, error) {
	serverStatus, err := client.Database("admin").RunCommand(
		ctx,
		bsonx.Doc{{"serverStatus", bsonx.Int32(1)}},
	).DecodeBytes()
	if err != nil {
		return "", err
	}

	version, err := serverStatus.LookupErr("version")
	if err != nil {
		return "", err
	}

	return version.StringValue(), nil
}
