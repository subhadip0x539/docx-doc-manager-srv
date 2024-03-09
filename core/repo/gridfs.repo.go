package repo

import (
	"bytes"
	"context"
	E "docx-doc-manager-srv/core/entity"
	"io"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GridFSUploadMetadata struct {
	Type string `bson:"type"`
}

type GridFSDownloadFilter struct {
	DocumentId primitive.ObjectID `bson:"_id"`
}

type GridFSDownloadObject struct {
	Id       primitive.ObjectID   `bson:"_id"`
	Name     string               `bson:"filename"`
	Length   int64                `bson:"length"`
	Metadata GridFSUploadMetadata `bson:"metadata"`
}

type IGridFSRepo interface {
	GridFSFileUpload(file multipart.File, header *multipart.FileHeader) (string, error)
	GridFSFileDownload(documentId string) (E.FSFileObject, error)
}

type GridFSRepo struct {
	client   *mongo.Client
	database string
}

func (g *GridFSRepo) GridFSFileUpload(file multipart.File, header *multipart.FileHeader) (string, error) {
	db := g.client.Database(g.database)
	tags := GridFSUploadMetadata{Type: header.Header.Get("Content-Type")}
	opts := options.GridFSUpload().SetMetadata(tags)

	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}

	objectID, err := bucket.UploadFromStream(header.Filename, io.Reader(file), opts)

	return objectID.Hex(), err
}

func (g *GridFSRepo) GridFSFileDownload(documentId string) (E.FSFileObject, error) {
	db := g.client.Database(g.database)
	buffer := bytes.NewBuffer(nil)
	var objects []GridFSDownloadObject

	objectId, err := primitive.ObjectIDFromHex(documentId)
	if err != nil {
		return E.FSFileObject{}, err
	}

	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return E.FSFileObject{}, err
	}

	filter := GridFSDownloadFilter{
		DocumentId: objectId,
	}

	cursor, _ := bucket.Find(filter)
	if err = cursor.All(context.TODO(), &objects); err != nil {
		return E.FSFileObject{}, err
	}

	if _, err := bucket.DownloadToStream(objectId, buffer); err != nil {
		return E.FSFileObject{}, err
	}

	return E.FSFileObject{
		DocumentId:   objects[0].Id.Hex(),
		DocumentName: objects[0].Name,
		Type:         objects[0].Metadata.Type,
		Content:      buffer.Bytes(),
	}, nil
}

func NewGridFS(client *mongo.Client, database string) *GridFSRepo {
	return &GridFSRepo{client: client, database: database}
}
