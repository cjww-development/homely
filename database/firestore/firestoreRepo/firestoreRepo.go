/*
 * Copyright 2019 CJWW Development
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package firestoreRepo

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"homely/models"
	"log"
)

func newClient(ctx context.Context) *firestore.Client {
	sa := option.WithCredentialsFile("./service-account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil { log.Fatalln(err) }
	client, err := app.Firestore(ctx)
	if err != nil { log.Fatalln(err) }
	return client
}

func Get(collection string, documentId string) *firestore.DocumentSnapshot {
	ctx := context.Background()
	client := newClient(ctx)
	col := client.Collection(collection)
	snap, err := col.Doc(documentId).Get(ctx)
	if err != nil { log.Fatalln(err) }
	return snap
}

func Query(collection string, query models.Query) []models.EnvData {
	ctx := context.Background()
	client := newClient(ctx)
	col := client.Collection(collection)
	docRefs := col.Where(query.Field, query.Operation, query.Value).Documents(ctx)
	docs, _ := docRefs.GetAll()
	var res []models.EnvData
	for _, doc := range docs {
		var model models.EnvData
		if err := doc.DataTo(&model); err != nil { log.Fatalln(err) }
		res = append(res, model)
	}
	return res
}
