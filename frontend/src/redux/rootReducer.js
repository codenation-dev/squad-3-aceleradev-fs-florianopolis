import { combineReducers } from 'redux'

const Logged = (state = {value:false,firstLogin:false,token:""}, action) => {
    switch (action.type) {
        case 'REQUEST_LOGOUT':
           return {value:false,firstLogin:false,token:""};
        case 'FAIL_LOGIN':
            return {value:false,firstLogin:false,token:""};
        case 'SUCCESS_LOGIN':
            return {value:true,firstLogin:true,token:action.user.token}
        case 'REFRESH_LOGIN':
            let temp = state.token
            if (action.token !== undefined) {
                temp = action.token
            }
            return {value:true,firstLogin:false,token:temp}
        case 'UNAUTH':
            return {}
        default:
            return state
    }
}

const APIData = (state = {}, action) => {
    switch (action.type) {
    case "REQUEST_RESULT":
        return {...action.data,Loading:false}
    case"REQUEST_FAIL":
        return {...action.data,Loading:false}
    case "REQUEST_FETCH":
        return {Loading:true}
    default:
        return state
    }
}

const User = (state = {}, action) => {
    switch (action.type) {
        case "LOAD_USER":
            return action.data 
        default:
            return state
    }
}

const Notifier = (state = [], action) => {
    switch (action.type) {
        case "ADD_NOTIFY":
            var d = new Date();
            action.payload.uid = d.getTime();
            return [...state,action.payload]
        case "REMOVE_NOTIFY":
            return state.filter((val) => {
                    return (val.uid!==action.index)
            })
        default:
            return state
    }
}

export const rootReducer = combineReducers({Login:Logged,API:APIData,User: User,Notifier: Notifier})
    