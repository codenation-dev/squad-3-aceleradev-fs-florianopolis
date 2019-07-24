import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from "react-router-dom";
import './index.css';
import App from './App';
import {createStore} from 'redux';
import {Provider} from 'react-redux';
import Logged from './redux/rootReducer';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/es/storage';
import { PersistGate } from 'redux-persist/integration/react';

const persitConfig = {key:'root',
    storage}

const pReducer = persistReducer(persitConfig, Logged)

const store = createStore(pReducer);//createStore(Logged);
const persistor = persistStore(store);

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

