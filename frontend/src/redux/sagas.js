import { all, call, select, put, takeLatest } from "redux-saga/effects";
import {
	REQUEST_LOGIN,
	REQUEST_FETCH,
	UPLOAD_FILE,
	AddNotify,
	ActionUnauth,
	ActionSuccessLogin,
	ActionFailLogin,
	APIFetchResult,
	APIFetchFail,
	ActionRefreshLogin,
	APIClear
} from "./actions";
import { login, request, uploadfile } from "../services/api";

function* APILogin(action) {
	try {
		console.log(action);
		const data = yield call(login, action.loginFields);
		if (data.Code === 1) {
			let user = yield { token: data.token };
			yield put(ActionSuccessLogin(user));
		} else {
			yield put(AddNotify({ type: "Alert", message: "Credenciais Invalidas" }));
			yield put(ActionFailLogin());
		}
	} catch (e) {
		console.log(e);
		yield put(
			AddNotify({
				type: "Alert",
				message: "Não foi possivel contatar o servidor"
			})
		);
	}
}

function* APICall(action) {
	try {
		console.log(action);
		let state = yield select();
		let data;
		if (action.args === undefined) {
			data = yield call(request, state.Login.token, action.endpoint);
			console.log(data);
		} else {
			data = yield call(
				request,
				state.Login.token,
				action.endpoint,
				action.args
			);
			console.log(data);
		}

		if (data.Code === 1) {
			yield put(ActionRefreshLogin(data.token));
			yield put(APIFetchResult(data));
		} else {
			switch (data.Code) {
				case 12:
					yield put(ActionRefreshLogin(data.token));
					yield put(
						AddNotify({ type: "Info", message: "Mudanças Salvas Com Sucesso" })
					);
					break;
				case 14:
					yield put(ActionRefreshLogin(data.token));
					yield put(
						AddNotify({ type: "Info", message: "Usuario removido com Sucesso" })
					);

					break;
				default:
					yield put(APIFetchFail(data));
					yield put(
						AddNotify({
							type: "Alert",
							message: "Erro ao comunicar com o servidor"
						})
					);
					break;
			}
		}
	} catch (e) {
		console.log(e);
		yield put(ActionUnauth());
		yield put(APIClear());
		yield put(AddNotify({ type: "Alert", message: "Token Expirado" }));
	}
}

function* UploadCSV(action) {
	try {
		console.log(action);
		let BFT = action.buffer;
		let list = atob(BFT.substring(29, BFT.length));
		let state = yield select();
		let data = yield call(uploadfile, state.Login.token, list);
		if (data.Code === 1) {
			yield put(
				AddNotify({ type: "Info", message: "Arquivo enviado com sucesso" })
			);
			yield put({ type: "CLOSE_UPLOADER" });
			yield put(ActionRefreshLogin(data.token));
			yield put(APIFetchResult(data));
		} else {
			yield put(
				AddNotify({ type: "Alert", message: "Erro ao enviar o arquivo" })
			);
		}
	} catch (e) {
		console.log(e);
		yield put(ActionUnauth());
		yield put(APIClear());
		yield put(AddNotify({ type: "Alert", message: "Token Expirado" }));
	}
}

export function* rootSaga() {
	yield all([
		takeLatest(REQUEST_LOGIN, APILogin),
		takeLatest(REQUEST_FETCH, APICall),
		takeLatest(UPLOAD_FILE, UploadCSV)
	]);
}
