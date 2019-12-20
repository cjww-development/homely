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
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func newClient(ctx context.Context) *firestore.Client {
	sa := option.WithCredentialsFile("./service-account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func newClientWithCollection(ctx context.Context, collection string) *firestore.CollectionRef {
	client := newClient(ctx)
	return client.Collection(collection)
}

func Get(collection string, documentId string) *firestore.DocumentSnapshot {
	ctx := context.Background()
	col := newClientWithCollection(ctx, collection)
	snap, err := col.Doc(documentId).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return snap
}
