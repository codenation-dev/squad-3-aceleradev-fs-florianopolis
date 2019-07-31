import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from "react-router-dom";
import './index.css';
import App from './App';
import {createStore, applyMiddleware, compose} from 'redux';
import {Provider} from 'react-redux';
import { rootReducer } from './redux/rootReducer';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/es/storage';
import { PersistGate } from 'redux-persist/integration/react';
import  createSagaMiddleware from 'redux-saga'
import {rootSaga} from './redux/sagas';

const persitConfig = {key:'root',
    storage}

const sagaMiddleware = createSagaMiddleware()
const pReducer = persistReducer(persitConfig, rootReducer)

const store = createStore(pReducer, applyMiddleware(sagaMiddleware));//createStore(Logged);
const persistor = persistStore(store);

sagaMiddleware.run(rootSaga)

store.subscribe(() => console.log(store.getState()))

ReactDOM.render(
    <Provider store={store}>
        <Router>
        <PersistGate loading={null} persistor={persistor}>
        <App />
        </PersistGate>
        </Router>
    </Provider>
, document.getElementById('root'));

