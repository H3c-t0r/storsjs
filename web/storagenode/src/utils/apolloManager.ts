// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { HttpLink } from 'apollo-link-http';
import ApolloClient from 'apollo-client/ApolloClient';
import { InMemoryCache } from 'apollo-cache-inmemory';

// Bootstrap server url
const bootstrapUrl = new HttpLink({
    uri: 'http://localhost:10002/api/graphql/v0',

});

// Creating apollo client
export default new ApolloClient({
    link: bootstrapUrl,
    cache: new InMemoryCache(),
    connectToDevTools: true,
});
