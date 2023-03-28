package main

import (
	"awesomeProject/grpc/hello/hellopb"
	"awesomeProject/grpc/products/productpb"
	"context"
	"fmt"

	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

// structure interface t omathc with mongo db
type product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Price float64            `bson:"price"`
}

func (*server) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	// parse content and save to mongo

	prod := req.GetProduct()

	data := product{
		Name:  prod.GetName(),
		Price: prod.GetPrice(),
	}

	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannont convert OID: %v", err),
		)
	}

	return &productpb.CreateProductResponse{
		Product: &productpb.Product{
			Id:    oid.Hex(),
			Name:  prod.GetName(),
			Price: prod.GetPrice(),
		},
	}, nil
}

func (*server) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	productId := req.GetProductId()

	oid, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &product{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find teh product: %v", err),
		)
	}

	return &productpb.GetProductResponse{
		Product: dbToProductPb(data),
	}, nil
}

func dbToProductPb(data *product) *productpb.Product {
	return &productpb.Product{
		Id:    data.ID.Hex(),
		Name:  data.Name,
		Price: data.Price,
	}
}

func (*server) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	fmt.Println("Update product request")
	prod := req.GetProduct()
	oid, err := primitive.ObjectIDFromHex(prod.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &product{}
	filter := bson.M{"_id": oid} // to get and to update

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find product with specified ID: %v", err),
		)
	}

	// we update our internal struct
	data.Name = prod.GetName()
	data.Price = prod.GetPrice()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
		)
	}

	return &productpb.UpdateProductResponse{
		Product: dbToProductPb(data),
	}, nil

}

func (*server) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	fmt.Println("Delete product")

	oid, err := primitive.ObjectIDFromHex(req.GetProductId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}
	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete product in MongoDB: %v", err),
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find product in MongoDB: %v", err),
		)
	}
	return &productpb.DeleteProductResponse{
		ProductId: req.GetProductId(),
	}, nil
}

func (*server) ListProduct(req *productpb.ListProductRequest, stream productpb.ProductService_ListProductServer) error {
	fmt.Println("List products")

	// Gets the cursor, throught it we could advance
	// each document json in mongo until ends
	// We use an empty filter just to accomplish with the interface contract
	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// @doc: The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding
	// function returns.
	defer cur.Close(context.Background())

	// looping the cursos until the EOF appear to end the stream
	for cur.Next(context.Background()) {
		data := &product{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error decoding data from MongoDB: %v", err),
			)
		}
		stream.Send(&productpb.ListProductResponse{Product: dbToProductPb(data)})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return nil
}

// this server has to implment whole the interfaces generated by the protobuf
type server struct{}

func (s *server) HelloManyLanguages(request *hellopb.HelloManyLanguagesRequest, languagesServer hellopb.HelloService_HelloManyLanguagesServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) HellosGoodbye(goodbyeServer hellopb.HelloService_HellosGoodbyeServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) Goodbye(goodbyeServer hellopb.HelloService_GoodbyeServer) error {
	//TODO implement me
	panic("implement me")
}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("hello world %v \n", req)
	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()
	custommerHello := "welcome " + firstName + " " + prefix
	res := &hellopb.HelloResponse{
		CustomHello: custommerHello,
	}
	return res, nil
}

func main() {
	fmt.Println("Inicializando Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to load serve %v", err)
	}
}

func Prueba() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to Mongo DB")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error creating the client DB %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error conecting the client DB %v", err)
	}

	collection = client.Database("productsdb").Collection("products")

	fmt.Println("Product service")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	productpb.RegisterProductServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	// Wait for crtl + x to exit

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until we get the signal

	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing mongo client connection")
	client.Disconnect(context.TODO())
	fmt.Println("Goodbye :D ...")
}
