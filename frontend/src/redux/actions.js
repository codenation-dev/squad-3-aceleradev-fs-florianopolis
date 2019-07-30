export const REQUEST_LOGIN = "REQUEST_LOGIN";
export const SUCCESS_LOGIN = "SUCCESS_LOGIN";
export const REQUEST_LOGOUT = "REQUEST_LOGOUT";
export const REFRESH_LOGIN = "REFRESH_LOGIN";
export const FAIL_LOGIN = "FAIL_LOGIN";
export const UNAUTH = "UNAUTH"

export const REQUEST_FETCH = "REQUEST_FETCH";
export const REQUEST_RESULT = "REQUEST_RESULT";
export const REQUEST_FAIL = "REQUEST_FAIL";
export const REQUEST_CLEAR = "REQUEST_CLEAR";

export const ActionRequestLogin = () => ({type:REQUEST_LOGIN});
export const ActionFailLogin = () => ({type:FAIL_LOGIN})
export const ActionSuccessLogin = user => ({type:SUCCESS_LOGIN,user})
export const ActionUnauth = () => ({type:UNAUTH})

export const AddNotify = (payload) => ({type:"ADD_NOTIFY",payload:payload})
export const ActionLoading = () => ({type:"REQUEST_LOADING",Loading:true})

export const ActionRefreshLogin = token => ({type:REFRESH_LOGIN,token})
export const ActionLogout = () => ({type:REQUEST_LOGOUT});

export const APIFetchRequest = (endpoint,...args) => ({type:REQUEST_FETCH,endpoint:endpoint,...args})
export const APIFetchResult = (data) => ({type:REQUEST_RESULT,data,Loading:false}) 
export const APIFetchFail = (data) => ({type:REQUEST_FAIL,data,Loading:false})
export const APIClear = () => ({type:REQUEST_CLEAR})