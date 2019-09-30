package main
go main(){
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

// get collection
collection := client.Database("testing").Collection("numbers")

//// insert
//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
//id := res.InsertedID

// query

ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
cur, err := collection.Find(ctx, bson.D{})
if err != nil { log.Fatal(err) }
defer cur.Close(ctx)
for cur.Next(ctx) {
var result bson.M
err := cur.Decode(&result)
if err != nil { log.Fatal(err) }
// do something with result....
}
if err := cur.Err(); err != nil {
log.Fatal(err)
}

}
