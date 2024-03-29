import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import CardBoard from './components/ProfileBoard';
import registerServiceWorker from './registerServiceWorker';

import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/css/bootstrap-theme.css';

import {Provider} from 'react-redux'
import {applyMiddleware, createStore} from 'redux'
import reducers from './reducers'

import thunk from 'redux-thunk';
import {createLogger} from 'redux-logger'

import {BrowserRouter} from 'react-router-dom'

import * as log from 'loglevel';

const middleware = [thunk];

if (process.env.NODE_ENV === 'development') {
    log.enableAll();
    middleware.push(createLogger())
} else {
    log.disableAll();
}

let store = createStore(
    reducers,
    applyMiddleware(...middleware)
);

ReactDOM.render(
    <Provider store={store}>
        <BrowserRouter>
            <CardBoard/>
        </BrowserRouter>
    </Provider>,
    document.getElementById('root'));

registerServiceWorker();
