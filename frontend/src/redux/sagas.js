import { all, call, select, put, takeEvery, takeLatest } from "redux-saga/effects";
import { REQUEST_LOGIN, REQUEST_FETCH,
	 AddNotify, ActionUnauth, ActionSuccessLogin, 
	 ActionFailLogin, APIFetchResult, APIFetchFail,
	  ActionRefreshLogin, APIClear } from "./actions";
import { login, request } from "../services/api";

function* APILogin(action) {
	try {
		console.log(action);
		const data = yield call(login, action.loginFields);
		if (data.Code === 1) {
            let user = yield {token: data.token}
			yield put(ActionSuccessLogin(user));
		} else {
			yield put(AddNotify({type:"Alert",message:"Credenciais Invalidas"}))
			yield put(ActionFailLogin());
		}
	} catch (e) {
		console.log(e);
		yield put(AddNotify({type:"Alert",message:"Não foi possivel contatar o servidor"}))
	}
}

function* APICall(action) {
	try {
		console.log(action);
		let state =  yield select()
		let data 
		if (action.args === undefined) {
			data = yield call(request, state.Login.token, action.endpoint);
			console.log(data)
		} else {
		data = yield call(request, state.Login.token, action.endpoint , action.args);
			console.log(data)
	}

		if (data.status === 1) {
			yield put(ActionRefreshLogin(data.token))
			yield put(APIFetchResult(data));
		} else {
			yield put(APIFetchFail(data));
			yield put({type:"Alert",message:"Erro ao comunicar com o servidor"});
		}
	} catch (e) {
		console.log(e);
		yield put(ActionUnauth());
		yield put(APIClear());
		yield put(AddNotify({type:"Alert",message:"Token Expirado"}))
	}
}

export function* rootSaga() {
	yield all([takeLatest(REQUEST_LOGIN, APILogin),
	takeLatest(REQUEST_FETCH, APICall)]);
}
